package shared

// HTTPRequest ....
type HTTPRequest struct {
	Body   interface{}
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
	statusCode int
	message    string
}

// HTTPBadRequest ...
func HTTPBadRequest(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 400,
		Body: httpErrorBody{
			statusCode: 400,
			message:    message,
		},
	}
}

// HTTPUnauthorized ...
func HTTPUnauthorized(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 401,
		Body: httpErrorBody{
			statusCode: 401,
			message:    message,
		},
	}
}

// HTTPForbbiden ...
func HTTPForbbiden(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 403,
		Body: httpErrorBody{
			statusCode: 403,
			message:    message,
		},
	}
}

// HTTPInternalServerError ...
func HTTPInternalServerError(message string) *HTTPResponse {
	return &HTTPResponse{
		StatusCode: 500,
		Body: httpErrorBody{
			statusCode: 500,
			message:    message,
		},
	}
}
