package response

// ErrorResponse represents the error response
type ErrorResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Errors interface{} `json:"errors"`
}