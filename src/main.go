package main

import (
	"time"

	applications "restapi/applications"
	adapters "restapi/frameworks/adapters"
	controllers "restapi/interfaces"
)

func main() {
	startup := time.Now()

	httpServer := HTTPServer{}
	httpServer.Init()

	healthUsecase := applications.NewHealthUsecase(startup)
	healthController := controllers.NewHealthController(healthUsecase)
	httpServer.RegisterRouteHandler("/", adapters.RouteAdapt(healthController), "GET")

	httpServer.StartHTTPServer()
}
