package shared

import (
	"net/http"
)

// HTTPResponse ....
type HTTPResponse struct {
	StatusCode int
	Body       interface{}
	Headers    interface{}
}

// HTTPSuccess ...
func HTTPSuccess(body interface{}) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusOK,
		Body:       body,
	}
}

// HTTPCreated ...
func HTTPCreated() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusCreated,
	}
}

// HTTPNoContent ...
func HTTPNoContent() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusNoContent,
	}
}

// HTTPErrorMessage ...
type HTTPErrorMessage struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// HTTPBadRequest ...
func HTTPBadRequest(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusBadRequest,
		Body: HTTPErrorMessage{
			StatusCode: http.StatusBadRequest,
			Message:    message,
		},
	}
}

// HTTPUnauthorized ...
func HTTPUnauthorized(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusUnauthorized,
		Body: HTTPErrorMessage{
			StatusCode: http.StatusUnauthorized,
			Message:    message,
		},
	}
}

// HTTPForbbiden ...
func HTTPForbbiden(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusForbidden,
		Body: HTTPErrorMessage{
			StatusCode: http.StatusForbidden,
			Message:    message,
		},
	}
}

// HTTPInternalServerError ...
func HTTPInternalServerError(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: http.StatusInternalServerError,
		Body: HTTPErrorMessage{
			StatusCode: http.StatusInternalServerError,
			Message:    message,
		},
	}
}
