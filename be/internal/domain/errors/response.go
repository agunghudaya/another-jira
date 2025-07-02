package errors

import (
	"encoding/json"
	"net/http"
)

// Response represents a standardized API response
type Response struct {
	Success bool           `json:"success"`
	Data    interface{}    `json:"data,omitempty"`
	Error   *ErrorResponse `json:"error,omitempty"`
	Meta    *MetaResponse  `json:"meta,omitempty"`
}

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Details interface{} `json:"details,omitempty"`
}

// MetaResponse represents metadata for the response
type MetaResponse struct {
	Page     int `json:"page,omitempty"`
	PageSize int `json:"page_size,omitempty"`
	Total    int `json:"total,omitempty"`
}

// NewSuccessResponse creates a new success response
func NewSuccessResponse(data interface{}, meta *MetaResponse) *Response {
	return &Response{
		Success: true,
		Data:    data,
		Meta:    meta,
	}
}

// NewErrorResponse creates a new error response
func NewErrorResponse(err error) *Response {
	var errorResp *ErrorResponse

	switch e := err.(type) {
	case *DomainError:
		errorResp = &ErrorResponse{
			Code:    e.Code,
			Message: e.Message,
		}
	default:
		errorResp = &ErrorResponse{
			Code:    ErrInternalServer,
			Message: "Internal server error",
		}
	}

	return &Response{
		Success: false,
		Error:   errorResp,
	}
}

// WriteJSON writes a JSON response to the http.ResponseWriter
func (r *Response) WriteJSON(w http.ResponseWriter, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(r)
}

// HTTP status codes for different error types
func GetHTTPStatus(err error) int {
	switch e := err.(type) {
	case *DomainError:
		switch e.Code {
		// System errors
		case ErrInternalServer, ErrDatabaseError:
			return http.StatusInternalServerError
		case ErrServiceUnavailable:
			return http.StatusServiceUnavailable
		case ErrExternalService:
			return http.StatusBadGateway

		// Validation errors
		case ErrInvalidInput, ErrMissingRequired, ErrInvalidFormat, ErrValidationFailed:
			return http.StatusBadRequest

		// Authentication errors
		case ErrUnauthorized, ErrInvalidToken, ErrTokenExpired:
			return http.StatusUnauthorized
		case ErrForbidden:
			return http.StatusForbidden

		// Resource errors
		case ErrNotFound:
			return http.StatusNotFound
		case ErrAlreadyExists:
			return http.StatusConflict
		case ErrResourceConflict:
			return http.StatusConflict
		case ErrResourceLocked:
			return http.StatusLocked

		// Business logic errors
		case ErrBusinessRule, ErrInvalidOperation, ErrInsufficientData, ErrInvalidState:
			return http.StatusUnprocessableEntity

		default:
			return http.StatusInternalServerError
		}
	default:
		return http.StatusInternalServerError
	}
}
