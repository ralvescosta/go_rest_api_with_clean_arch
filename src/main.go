package main

import (
	"time"

	applications "restapi/applications"
	frameworks "restapi/frameworks"
	adapters "restapi/frameworks/adapters"
	database "restapi/frameworks/database"
	controllers "restapi/interfaces"
)

func main() {
	startup := time.Now()

	httpServer := frameworks.HTTPServer{}
	httpServer.Init()

	_, err := database.DbConnection()
	if err != nil {
		panic(err)
	}

	healthUsecase := &applications.HealthUsecase{StartupTime: startup}
	healthController := controllers.NewHealthController(healthUsecase)
	httpServer.RegisterRouteHandler("/", adapters.RouteAdapt(healthController), "GET")

	createBookUsecase := &applications.CreateBookUsecase{}
	createBookController := controllers.NewCreatBookController(createBookUsecase)
	httpServer.RegisterRouteHandler("/books", adapters.RouteAdapt(createBookController), "POST")

	httpServer.StartHTTPServer()
}
