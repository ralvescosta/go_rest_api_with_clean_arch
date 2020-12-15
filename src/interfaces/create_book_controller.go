package interfaces

import (
	"encoding/json"
	"net/http"

	"restapi/src/applications"
	"restapi/src/entities"
	"restapi/src/shared"
)

// CreateBookController ...
type CreateBookController struct {
	Usecase applications.ICreateBookUsecase
}

// Handler ...
func (c *CreateBookController) Handler(res http.ResponseWriter, req *http.Request) {
	body := &entities.BookDTO{}
	err := json.NewDecoder(req.Body).Decode(body)

	if err != nil {
		response := shared.HTTPBadRequest("Wrong Format")
		res.WriteHeader(response.StatusCode)
		json.NewEncoder(res).Encode(response.Body)
	}

	result := c.Usecase.Create(body)

	res.WriteHeader(result.StatusCode)
	if result.Body != nil {
		json.NewEncoder(res).Encode(result.Body)
	}
}
