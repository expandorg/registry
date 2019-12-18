package apierror

import (
	"encoding/json"
	"fmt"
)

type APIError struct {
	statusCode int    `json:"-"`
	Message    string `json:"message"`
	error      error  `json:"-"`
}

type errorPayload struct {
	Message string `json:"message"`
}

func New(statusCode int, message string, err error) *APIError {
	return &APIError{
		statusCode: statusCode,
		Message:    message,
		error:      err,
	}
}

func (a *APIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(errorPayload{
		Message: a.Message,
	})
}

func (a *APIError) StatusCode() int {
	return a.statusCode
}

func (a *APIError) Error() string {
	return fmt.Sprintf("Status: %d, Message: [%s]:%s", a.statusCode, a.Message, a.error)
}
