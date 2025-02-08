package service

import (
	"errors"
	"movie_booking/models"
	"movie_booking/repository"
	"time"
)

type BookingService struct {
	Repo repository.BookingRepository
}

func NewBookingService(repo repository.BookingRepository) *BookingService {
	return &BookingService{
		Repo: repo,
	}
}

func (bs *BookingService) BookTicket(bookingrequest models.Bookingrequest) (models.BookingResponse, error) {
	movieShowtime, err := time.Parse("15:04", bookingrequest.Showtime)
	if err != nil {
		return models.BookingResponse{}, err
	}
	available, err := bs.Repo.CheckAvailability(bookingrequest.Movie, movieShowtime, bookingrequest.NoOfSeats)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if !available {
		return models.BookingResponse{}, errors.New("tickets Sold out")
	}

	response, err := bs.Repo.BookTicket(bookingrequest)
	if err != nil {
		return models.BookingResponse{}, err
	}
	return response, nil
}

func (bs *BookingService) GetBookingDetails(userEmail string) (models.BookingResponse, error) {
	bookingResponse, err := bs.Repo.GetBookingDetails(userEmail)
	if err != nil {
		return models.BookingResponse{}, err
	}
	return bookingResponse, nil
}

func (bs *BookingService) GetAllBookingDetailsForMovie(movieName string, showtime string) (models.BookingResponse, error) {
	movieShowtime, err := time.Parse("15:04", showtime)
	if err != nil {
		return models.BookingResponse{}, err
	}
	bookingResponse, err := bs.Repo.GetAllBookingDetailsForMovie(movieName, movieShowtime)
	if err != nil {
		return models.BookingResponse{}, err
	}
	return bookingResponse, nil
}

func (bs *BookingService) ModifyBooking(bookingrequest models.ModifyBookingRequest) (models.BookingResponse, error) {
	getBooking, err := bs.Repo.GetBookingDetails(bookingrequest.Email)
	if err != nil {
		return models.BookingResponse{}, err
	}
	movieShowtime, err := time.Parse("15:04", getBooking.Showtime)
	if err != nil {
		return models.BookingResponse{}, err
	}
	checkAvailability, err := bs.Repo.CheckSpecificSeatAvailability(getBooking.MovieName, movieShowtime, bookingrequest.NewSeatNumbers)
	if err != nil {
		return models.BookingResponse{}, err
	}
	if !checkAvailability {
		return models.BookingResponse{}, errors.New("tickets Sold out")
	}
	modifiedResponse, err := bs.Repo.ModifyBooking(bookingrequest)
	if err != nil {
		return models.BookingResponse{}, err
	}
	return modifiedResponse, nil
}
