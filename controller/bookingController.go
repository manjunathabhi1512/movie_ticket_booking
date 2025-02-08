package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"movie_booking/models"
	"movie_booking/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BookingController struct {
	BookingService *service.BookingService
}

func NewBookingController(service *service.BookingService) *BookingController {
	return &BookingController{
		BookingService: service,
	}
}

func (bc *BookingController) BookTicket(r *gin.Context) {
	var bookingrequest models.Bookingrequest
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		fmt.Println("Unable to Read the data", err)
		return
	}
	if err := json.Unmarshal(body, &bookingrequest); err != nil {
		fmt.Println("Unable to Unmarshal the data", err)
		return
	}
	bookingDetails, err := bc.BookingService.BookTicket(bookingrequest)
	if err != nil {
		fmt.Println("Failed to Book the ticket", err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message":  "Booking successful",
		"response": bookingDetails,
	})
}

func (bc *BookingController) GetBookingDetails(r *gin.Context) {
	userEmail := r.Param("email")
	bookingDetails, err := bc.BookingService.GetBookingDetails(userEmail)
	if err != nil {
		fmt.Println("Failed to Fetch the ticket details", err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message":  "Fetched data successfully",
		"response": bookingDetails,
	})
}

func (bc *BookingController) GetAllBookingDetailsForMovie(r *gin.Context) {
	movieName := r.Param("movie_name")
	showtime := r.Param("showtime")
	fmt.Println("movieName", movieName, "showtime", showtime)
	bookingDetails, err := bc.BookingService.GetAllBookingDetailsForMovie(movieName, showtime)
	if err != nil {
		fmt.Println("Failed to Fetch the ticket details", err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message":  "Fetched data successfully",
		"response": bookingDetails,
	})
	if bookingDetails.UserName == "" {
		r.JSON(http.StatusOK, gin.H{
			"message": "No bookings found for the given movie and showtime",
		})
	}
}

func (bc *BookingController) ModifyBooking(r *gin.Context) {
	var modifyRequest models.ModifyBookingRequest
	body, err := io.ReadAll(r.Request.Body)
	if err != nil {
		fmt.Println("Unable to Read the data", err)
		return
	}
	if err := json.Unmarshal(body, &modifyRequest); err != nil {
		fmt.Println("Unable to Unmarshal the data", err)
		return
	}
	modifiedDetails, err := bc.BookingService.ModifyBooking(modifyRequest)
	if err != nil {
		fmt.Println("Failed to Modify the Seats", err)
		return
	}
	r.JSON(http.StatusOK, gin.H{
		"message":  "Seats Modified successfully",
		"response": modifiedDetails,
	})
}