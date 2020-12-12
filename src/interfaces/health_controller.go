package interfaces

import (
	"restapi/shared"

	applications "restapi/applications"
)

type healthController struct {
	usecase *applications.HealthUsecase
}

func (c *healthController) Handler(httpRequest *shared.HTTPRequest) *shared.HTTPResponse {

	return c.usecase.Health()
}

// NewHealthController ...
func NewHealthController(usecase *applications.HealthUsecase) shared.IController {
	return &healthController{usecase: usecase}
}
