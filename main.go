package main

import (
	"time"

	applications "restapi/src/applications"
	frameworks "restapi/src/frameworks"
	adapters "restapi/src/frameworks/adapters"
	database "restapi/src/frameworks/database"
	controllers "restapi/src/interfaces"
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
