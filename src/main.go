package main

import (
	adapters "restapi/frameworks/adapters"
	controllers "restapi/interfaces"
)

func main() {
	httpServer := HTTPServer{}
	httpServer.Init()

	healthController := controllers.HealthController()
	httpServer.RegisterRouteHandler("/", adapters.RouteAdapt(healthController), "GET")

	httpServer.StartHTPServer()
}
