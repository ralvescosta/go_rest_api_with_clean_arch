package main

import (
	"time"

	frameworks "restapi/src/frameworks"
	database "restapi/src/frameworks/database"
	repositories "restapi/src/frameworks/repositories"

	healthUsecase "restapi/src/applications/health"
	healthController "restapi/src/interfaces/health"

	createBookUsecase "restapi/src/applications/create_book"
	createBookController "restapi/src/interfaces/create_book"

	getAllBooksUsecase "restapi/src/applications/get_all_books"
	getAllBooksController "restapi/src/interfaces/get_all_books"
)

func main() {
	startup := time.Now()

	httpServer := frameworks.HTTPServer{}
	httpServer.Init()

	DbCon, err := database.DbConnection()
	if err != nil {
		panic(err)
	}

	healthUsecase := healthUsecase.NewHealthUsecase(startup)
	healthController := healthController.HealthController{Usecase: healthUsecase}
	httpServer.RegisterRouteHandler("/", healthController.Handler, "GET")

	booksRepository := repositories.NewBooksRepository(DbCon)
	createBookUsecase := createBookUsecase.NewCreateBookUsecase(booksRepository)
	createBookController := createBookController.CreateBookController{Usecase: createBookUsecase}
	httpServer.RegisterRouteHandler("/books", createBookController.Handler, "POST")

	getAllBooksUsecase := getAllBooksUsecase.NewGetAllBooksUsecase(booksRepository)
	getAllBooksController := getAllBooksController.GetAllBooksController{Usecase: getAllBooksUsecase}
	httpServer.RegisterRouteHandler("/books", getAllBooksController.Handler, "GET")

	httpServer.StartHTTPServer()
}
