package main

import (
	"time"

	applications "restapi/src/applications"
	frameworks "restapi/src/frameworks"
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

	healthUsecase := applications.NewHealthUsecase(startup)
	healthController := controllers.HealthController{Usecase: healthUsecase}
	httpServer.RegisterRouteHandler("/", healthController.Handler, "GET")

	booksRepository := repositories.NewBooksRepository(DbCon)
	createBookUsecase := applications.NewCreateBookUsecase(booksRepository)
	createBookController := controllers.CreateBookController{Usecase: createBookUsecase}
	httpServer.RegisterRouteHandler("/books", createBookController.Handler, "POST")

	getAllBooksUsecase := applications.NewGetAllBooksUsecase(booksRepository)
	getAllBooksController := controllers.GetAllBooksController{Usecase: getAllBooksUsecase}
	httpServer.RegisterRouteHandler("/books", getAllBooksController.Handler, "GET")

	getOneBookUsecase := applications.NewGetOneBookUsecase(booksRepository)
	getOneBookController := controllers.GetOneBookController{Usecase: getOneBookUsecase}
	httpServer.RegisterRouteHandler("/books/{id}", getOneBookController.Handler, "GET")

	httpServer.StartHTTPServer()
}
