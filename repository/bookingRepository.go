package repository

import (
	"fmt"
	"log"
	"movie_booking/models"
	"strings"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type BookingRepository interface {
	CheckAvailability(movie string, showtime time.Time, noofseats int) (bool, error)
	GetAvailableSeats(screenId string, showtime time.Time) ([]string, error)
	CheckSpecificSeatAvailability(movieName string, showtime time.Time, seatNumbers []string) (bool, error)
	BookTicket(bookingrequest models.Bookingrequest) (models.BookingResponse, error)
	GetAvailabeSeatNumbers(movie string, showtime time.Time, noofseats int) ([]string, error)
	GetBookingDetails(userEmail string) (models.BookingResponse, error)
	GetAllBookingDetailsForMovie(movieName string, showtime time.Time) (models.BookingResponse, error)
	ModifyBooking(bookingrequest models.ModifyBookingRequest) (models.BookingResponse, error)
}

type BookingRepo struct {
	db *gorm.DB
}

func NewBookingRepository() (*BookingRepo, error) {
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return &BookingRepo{}, err
	}

	fmt.Println("Database connection established")
	return &BookingRepo{
		db: db,
	}, nil
}

func (br *BookingRepo) GetAvailabeSeatNumbers(movie string, showtime time.Time, noofseats int) ([]string, error) {
	var seatStructure models.SeatStructure
	if err := br.db.Where("screen_id = ? AND is_booked = FALSE").Find(&seatStructure).Error; err != nil {
		return nil, err
	}
	return strings.Split(seatStructure.SeatsNumber, ","), nil
}

func (br *BookingRepo) CheckAvailability(movie string, showtime time.Time, noofseats int) (bool, error) {
	var booking models.Booking
	if err := br.db.Where("movie_name = ? AND showtime = ? AND no_of_seats >= ?", movie, showtime, noofseats).First(&booking).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (br *BookingRepo) BookTicket(bookingrequest models.Bookingrequest) (models.BookingResponse, error) {
	var booking models.Booking
	movieShowtime, err := time.Parse("15:04", bookingrequest.Showtime)
	if err != nil {
		return models.BookingResponse{}, err
	}
	seatNums, err := br.GetAvailabeSeatNumbers(bookingrequest.Movie, movieShowtime, bookingrequest.NoOfSeats)
	if err != nil {
		return models.BookingResponse{}, err
	}

	if err := br.db.Model(&models.SeatStructure{}).
		Where("screen_id = ? AND seats_number IN (?)", bookingrequest.ScreenId, seatNums).
		Updates(map[string]interface{}{"is_booked": true}).Error; err != nil {
		return models.BookingResponse{}, err
	}

	if err := br.db.Create(&models.Booking{
		UserName:    bookingrequest.Name,
		UserEmail:   bookingrequest.Email,
		MovieName:   bookingrequest.Movie,
		Showtime:    bookingrequest.Showtime,
		NoOfSeats:   bookingrequest.NoOfSeats,
		SeatNumbers: seatNums,
	}).Error; err != nil {
		return models.BookingResponse{}, err
	}
	return models.BookingResponse{
		MovieName:   booking.MovieName,
		Showtime:    booking.Showtime,
		TotalTicket: booking.NoOfSeats,
		UserName:    booking.UserName,
		UserEmail:   booking.UserEmail,
		SeatNums:    booking.SeatNumbers,
	}, nil
}

func (br *BookingRepo) GetBookingDetails(userEmail string) (models.BookingResponse, error) {
	var booking models.Booking
	if err := br.db.Where("user_email = ?", userEmail).Error; err != nil {
		return models.BookingResponse{}, err
	}
	return models.BookingResponse{
		MovieName:   booking.MovieName,
		Showtime:    booking.Showtime,
		TotalTicket: booking.NoOfSeats,
		UserName:    booking.UserName,
		UserEmail:   booking.UserEmail,
		SeatNums:    booking.SeatNumbers,
	}, nil
}

func (br *BookingRepo) GetAllBookingDetailsForMovie(movieName string, showtime time.Time) (models.BookingResponse, error) {
	var booking models.Booking
	if err := br.db.Where("movie_name = ? AND showtime = ?", movieName, showtime).Error; err != nil {
		return models.BookingResponse{}, err
	}
	return models.BookingResponse{
		UserName:    booking.UserName,
		UserEmail:   booking.UserEmail,
		SeatNums:    booking.SeatNumbers,
		TotalTicket: booking.NoOfSeats,
	}, nil
}

func (br *BookingRepo) ModifyBooking(bookingrequest models.ModifyBookingRequest) (models.BookingResponse, error) {
	var booking models.Booking
	if err := br.db.Where("user_email = ?", bookingrequest.Email).Error; err != nil {
		return models.BookingResponse{}, err
	}
	return models.BookingResponse{
		UserName:    booking.UserName,
		UserEmail:   booking.UserEmail,
		SeatNums:    booking.SeatNumbers,
		TotalTicket: booking.NoOfSeats,
	}, nil
}

func (br *BookingRepo) CheckSpecificSeatAvailability(movieName string, showtime time.Time, seatNumbers []string) (bool, error) {
	var screenId string
	if err := br.db.Where("movie_name = ? AND showtime = ?", movieName, showtime).Select("screen_id").First(&screenId).Error; err != nil {
		return false, err
	}
	if screenId == "" {
		return false, fmt.Errorf("no screen found for the given movie and showtime")
	}
	var unavailableSeats []string
	if err := br.db.Where("screen_id = ? AND seats_number IN (?) AND is_booked = FALSE", screenId, seatNumbers).Find(&unavailableSeats).Error; err != nil {
		return false, err
	}
	if len(unavailableSeats) < len(seatNumbers) {
		return false, nil
	}
	return true, nil
}


func (br *BookingRepo) GetAvailableSeats(screenId string, showtime time.Time) ([]string, error) {
	var seatStructure models.SeatStructure
	if err := br.db.Where("screen_id = ? AND showtime = ? AND is_booked = FALSE", screenId, showtime).Find(&seatStructure).Error; err != nil {
		return nil, err
	}
	return strings.Split(seatStructure.SeatsNumber, ","), nil
}
