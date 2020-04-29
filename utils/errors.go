package utils

// ErrorResponse for json errors
type ErrorResponse struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

// NewBadRequest return a error response with Type of Bad Request
func NewBadRequest(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Bad Request",
		Message: message,
	}
}

// NewNotFound return a not found response
func NewNotFound(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Not Found",
		Message: message,
	}
}

// NewServerError return a server error response
func NewServerError(message string) *ErrorResponse {
	return &ErrorResponse{
		Type:    "Server Error",
		Message: message,
	}
}
