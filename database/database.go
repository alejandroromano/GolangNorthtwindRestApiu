package database

import "database/sql"

func InitDB() *sql.DB {
	connectionString := "root:admin@tcp(localhost:3310)/northwind"
	databaseConnection, err := sql.Open("mysql", connectionString)
	
	if err != nil {
		panic(err.Error)
	}
	return databaseConnection
}
