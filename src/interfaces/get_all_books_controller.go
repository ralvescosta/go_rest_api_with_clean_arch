package interfaces

import (
	"encoding/json"
	"net/http"

	"restapi/src/applications"
)

// GetAllBooksController ...
type GetAllBooksController struct {
	Usecase applications.IGetAllBooksUsecase
}

// Handler ...
func (c *GetAllBooksController) Handler(res http.ResponseWriter, req *http.Request) {

	result := c.Usecase.GetAll()

	res.WriteHeader(result.StatusCode)
	if result.Body != nil {
		json.NewEncoder(res).Encode(result.Body)
	}
}
