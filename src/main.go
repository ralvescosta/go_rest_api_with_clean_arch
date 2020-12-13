package main

import (
	"time"

	applications "restapi/src/applications"
	frameworks "restapi/src/frameworks"
	adapters "restapi/src/frameworks/adapters"
	database "restapi/src/frameworks/database"
	repositories "restapi/src/frameworks/repositories"
	controllers "restapi/src/interfaces"
)

func main() {
	startup := time.Now()

	httpServer := frameworks.HTTPServer{}
	httpServer.Init()

	DbCon, err := database.DbConnection()
	if err != nil {
		panic(err)
	}

	healthUsecase := &applications.HealthUsecase{StartupTime: startup}
	healthController := controllers.NewHealthController(healthUsecase)
	httpServer.RegisterRouteHandler("/", adapters.RouteAdapt(healthController), "GET")

	booksRepository := &repositories.BooksRepository{Db: DbCon}
	createBookUsecase := &applications.CreateBookUsecase{BooksRepository: booksRepository}
	createBookController := controllers.NewCreatBookController(createBookUsecase)
	httpServer.RegisterRouteHandler("/books", adapters.RouteAdapt(createBookController), "POST")

	httpServer.StartHTTPServer()
}
