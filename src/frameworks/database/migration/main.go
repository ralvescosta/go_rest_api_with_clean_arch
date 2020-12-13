package main

import (
	"database/sql"
	"fmt"

	"restapi/src/frameworks/database"
)

func createBookTable(dbConn *sql.DB) error {
	fmt.Println("\n****************************************")
	fmt.Println("1607860952225 - Create Book Table")

	_, err := dbConn.Exec(`
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

	return err
}

func main() {
	dbConn, err := database.DbConnection()
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	fmt.Println("Running migrations....")

	err = createBookTable(dbConn)
	if err != nil {
		if err.Error() == "pq: relation \"books\" already exists" {
			fmt.Println("Book Table already exist")
			fmt.Println("****************************************")
			fmt.Println()
		} else {
			panic(err)
		}
	} else {
		fmt.Println("Book was Table created")
		fmt.Println("****************************************")
		fmt.Println()
	}

	fmt.Println("Migration run successfully")

}
