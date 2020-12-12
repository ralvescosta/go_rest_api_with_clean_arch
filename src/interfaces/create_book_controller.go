package interfaces

import (
	"encoding/json"
	"restapi/applications"
	"restapi/entities"
	"restapi/shared"
)

type createBookController struct {
	usecase *applications.CreateBookUsecase
}

func (c *createBookController) Handler(httpRequest *shared.HTTPRequest) *shared.HTTPResponse {
	body := &entities.BookDTO{}
	err := json.NewDecoder(httpRequest.Body).Decode(body)

	if err != nil {
		return shared.HTTPBadRequest("Body Wrong Format")
	}

	return c.usecase.Create(body)
}

// NewCreatBookController ...
func NewCreatBookController(usecase *applications.CreateBookUsecase) shared.IController {
	return &createBookController{usecase: usecase}
}
