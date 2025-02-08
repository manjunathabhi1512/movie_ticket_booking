package models

type Bookingrequest struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Movie     string `json:"movie"`
	Showtime  string `json:"showtime"`
	NoOfSeats int    `json:"noOfSeats"`
	ScreenId  string `json:"screenId"`
}

type BookingResponse struct {
	MovieName   string   `json:"moviename"`
	Showtime    string   `json:"showtime"`
	UserName    string   `json:"username"`
	UserEmail   string   `json:"useremail"`
	TotalTicket int      `json:"totalticket"`
	SeatNums    []string `json:"seatno"`
}

type ModifyBookingRequest struct {
	Email          string   `json:"email"`
	NewSeatNumbers []string `json:"newSeatNumbers"`
}
