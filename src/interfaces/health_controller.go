package interfaces

import (
	"restapi/shared"
	"time"
)

type healthController struct{}

func (*healthController) Handler(httpRequest *shared.HTTPRequest) *shared.HTTPResponse {
	var response = make(map[string]interface{})
	response["now"] = time.Now()

	return &shared.HTTPResponse{
		StatusCode: 200,
		Body:       response,
	}
}

// HealthController ...
func HealthController() shared.IController {
	return &healthController{}
}
