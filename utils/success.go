package utils

// NewOkReponse return a server error response
func NewOkReponse(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Created",
		Message: message,
	}
}
