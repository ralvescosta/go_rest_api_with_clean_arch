package interfaces

import (
	"encoding/json"
	"net/http"
	applications "restapi/src/applications"
)

// HealthController ...
type HealthController struct {
	Usecase applications.IHealthUsecase
}

// Handler ...
func (c *HealthController) Handler(res http.ResponseWriter, req *http.Request) {

	result := c.Usecase.Health()

	res.WriteHeader(result.StatusCode)
	json.NewEncoder(res).Encode(result.Body)

}
