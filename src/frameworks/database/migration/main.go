package main

import (
	"fmt"
	"restapi/src/frameworks/database"
)

func main() {
	dbConn, err := database.DbConnection()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	fmt.Println("Running migrations....")

	_, err = dbConn.Exec(`
		CREATE TABLE books
		(
			id SERIAL NOT NULL PRIMARY KEY,
			title varchar(255) NOT NULL,
			author varchar(255) NOT NULL,
			publishing_company varchar(255) NOT NULL,
			edition int,
			created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
			deleted_at TIMESTAMP WITH TIME ZONE
		);
	`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Migration run successfully")

}
