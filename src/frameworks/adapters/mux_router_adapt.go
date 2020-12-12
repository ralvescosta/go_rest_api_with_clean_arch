package adapters

import (
	shared "restapi/shared"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// RouteAdapt ...
func RouteAdapt(controller shared.IController) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		httpRequest := &shared.HTTPRequest{
			Body:   req.Body,
			Params: vars,
		}

		result := controller.Handler(httpRequest)

		res.WriteHeader(result.StatusCode)
		if result.Body != nil {
			json.NewEncoder(res).Encode(result.Body)
		}
	}
}
