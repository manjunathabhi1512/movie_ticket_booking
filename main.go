package main

import (
	"fmt"
	"movie_booking/controller"
	"movie_booking/repository"
	"movie_booking/routes"
	"movie_booking/service"
)

func main() {

	bookingRepo, err := repository.NewBookingRepository()
	if err != nil {
		fmt.Printf("Failed to Initiate Database connetion : %s", err)
	}
	bookingService := service.NewBookingService(bookingRepo)
	bookingController := controller.NewBookingController(bookingService)
	bookingRoutes := routes.InitRouter(bookingController)
	bookingRoutes.Run(":4000")
}
