package interfaces

import (
	"encoding/json"
	"net/http"

	application "restapi/src/applications/get_all_books"
)

// GetAllBooksController ...
type GetAllBooksController struct {
	Usecase application.IGetAllBooksUsecase
}

// Handler ...
func (c *GetAllBooksController) Handler(res http.ResponseWriter, req *http.Request) {

	result := c.Usecase.GetAll()

	res.WriteHeader(result.StatusCode)
	if result.Body != nil {
		json.NewEncoder(res).Encode(result.Body)
	}
}
