package interfaces

import (
	"restapi/shared"

	applications "restapi/applications"
)

type healthController struct {
	usecase applications.IHealthUsecase
}

func (c *healthController) Handler(httpRequest *shared.HTTPRequest) *shared.HTTPResponse {

	result, err := c.usecase.Health()
	if err != nil {
		return shared.HTTPInternalServerError(err.Error())
	}

	return result
}

// NewHealthController ...
func NewHealthController(usecase applications.IHealthUsecase) shared.IController {
	return &healthController{usecase: usecase}
}
