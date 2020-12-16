package interfaces

import (
	"encoding/json"
	"net/http"
	application "restapi/src/applications/health"
)

// HealthController ...
type HealthController struct {
	Usecase application.IHealthUsecase
}

// Handler ...
func (c *HealthController) Handler(res http.ResponseWriter, req *http.Request) {

	result := c.Usecase.Health()

	res.WriteHeader(result.StatusCode)
	json.NewEncoder(res).Encode(result.Body)

}
