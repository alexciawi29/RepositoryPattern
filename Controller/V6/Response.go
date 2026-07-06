package V6

type ErrorResponse struct {
	Error string `json:"Error" example:"error message"`
}

type MessageResponse struct {
	Message string `json:"Message" example:"operation successful"`
}
