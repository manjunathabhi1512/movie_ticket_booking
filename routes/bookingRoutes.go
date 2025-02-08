package routes

import (
	"movie_booking/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(controller *controller.BookingController) *gin.Engine {
	r := gin.Default()
	//crate a group pf routes for name booking

	r.POST("bookTicket", controller.BookTicket)
	r.GET("getBookingDetails", controller.GetBookingDetails)
	r.GET("getAllBookingDetailsForMovie", controller.GetAllBookingDetailsForMovie)
	r.PUT("modifyBooking", controller.ModifyBooking)
	return r
}
