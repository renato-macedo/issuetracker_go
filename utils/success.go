package utils

// NewSuccess return a server error response
func NewSuccess(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Created",
		Message: message,
	}
}
