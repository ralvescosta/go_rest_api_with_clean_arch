package interfaces

import (
	"encoding/json"
	"net/http"
	"restapi/src/applications"
	"restapi/src/shared"
	"strconv"

	"github.com/gorilla/mux"
)

// GetOneBookController ...
type GetOneBookController struct {
	Usecase applications.IGetOneBookUsecase
}

// Handler ...
func (c *GetOneBookController) Handler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)

	u64, err := strconv.ParseUint(vars["id"], 10, 64)
	if err != nil {
		response := shared.HTTPBadRequest("Wrong Format")
		res.WriteHeader(response.StatusCode)
		json.NewEncoder(res).Encode(response.Body)
	}
	result := c.Usecase.GetOne(u64)

	res.WriteHeader(result.StatusCode)
	if result.Body != nil {
		json.NewEncoder(res).Encode(result.Body)
	}
}
