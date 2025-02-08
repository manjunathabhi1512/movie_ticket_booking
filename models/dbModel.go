package models

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type DBConnection struct {
	gorm.Model
	DBName   string
	Port     string
	UserName string
	Password string
	Host     string
}

type Booking struct {
	gorm.Model
	BookingId   uuid.UUID `gorm:"column:bookingid;type:uuid;primary_key"`
	UserName    string    `gorm:"column:user_name;type:text; not null"`
	UserEmail   string    `gorm:"column:user_email;type:text; not null"`
	MovieName   string    `gorm:"column:movie_name;type:text; not null"`
	Showtime    string    `gorm:"column:showtime;type:text; not null"`
	NoOfSeats   int       `gorm:"column:no_of_seats;type:int; not null"`
	SeatNumbers []string  `gorm:"column:seat_numbers;type:text; not null"`
}

type SeatStructure struct {
	gorm.Model
	ScreenId    string    `gorm:"column:screen_id; not null"`
	Showtime    time.Time `gorm:"column:showtime; not null"`
	SeatsNumber string    `json:"seat_number"`
	IsBooked    bool      `json:"is_booked"`
}
