package entities

type GetPassengerByIdRequest struct {
	NationalId string `form:"national_id" binding:"required,len=11,numeric"`
}

type GetPassengerByPnrRequest struct {
	PNR     string `form:"pnr" binding:"required,len=6,alphanum"`
	Surname string `form:"surname" binding:"required,alpha,min=2,max=50"`
}

type OnlineCheckInRequest struct {
	PNR     string `json:"pnr" binding:"len=6,alphanum"`
	Surname string `json:"surname" binding:"alpha,min=2,max=50"`
}

type GetEmployeeByIdRequest struct {
	ID uint `json:"id" binding:"required"`
}

type RegisterEmployeeRequest struct {
	Employee Employee
}

type AddPlaneRequest struct {
	Plane Plane
}

type SetPlaneStatusRequest struct {
	PlaneRegistration string `json:"plane_registration" binding:"required"`
	IsAvailable       bool   `json:"is_available" binding:"required"`
}

type GetPlaneByRegistrationRequest struct {
	PlaneRegistration string `json:"plane_registration" binding:"required"`
}

type GetPlaneByFlightNumberRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
}

type GetPlaneByLocationRequest struct {
	Location string `json:"location" binding:"required"`
}

type GetSpecificFlightRequest struct {
	FlightNumber      string `json:"flight_number" binding:"required"`
	DepartureDateTime string `json:"departure_datetime" binding:"required"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
