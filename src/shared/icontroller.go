package shared

// IController ...
type IController interface {
	Handler(httpRequest *HTTPRequest) *HTTPResponse
}
