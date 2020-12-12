package shared

import "io"

// HTTPRequest ....
type HTTPRequest struct {
	Body   io.Reader
	Params map[string]string
}

// HTTPResponse ....
type HTTPResponse struct {
	StatusCode int
	Body       interface{}
	Headers    interface{}
}

// HTTPSuccess ...
func HTTPSuccess(body interface{}) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 200,
		Body:       body,
	}
}

// HTTPCreated ...
func HTTPCreated() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 201,
	}
}

// HTTPNoContent ...
func HTTPNoContent() *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 204,
	}
}

type httpErrorBody struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

// HTTPBadRequest ...
func HTTPBadRequest(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 400,
		Body: httpErrorBody{
			StatusCode: 400,
			Message:    message,
		},
	}
}

// HTTPUnauthorized ...
func HTTPUnauthorized(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 401,
		Body: httpErrorBody{
			StatusCode: 401,
			Message:    message,
		},
	}
}

// HTTPForbbiden ...
func HTTPForbbiden(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 403,
		Body: httpErrorBody{
			StatusCode: 403,
			Message:    message,
		},
	}
}

// HTTPInternalServerError ...
func HTTPInternalServerError(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 500,
		Body: httpErrorBody{
			StatusCode: 500,
			Message:    message,
		},
	}
}
