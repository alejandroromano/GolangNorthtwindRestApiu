package main

import (
	"GolangNorthtwindRestApiu/database"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	databaseConnection := database.InitDB()

	//Logica

	defer databaseConnection.Close()
	fmt.Println(databaseConnection)
}
