package main

import (
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"main/dbase"
	"main/v1"
)

func initDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/metaverse")

	if err != nil {
		panic(err.Error())
	}

	quarry, err := db.Query("show tables;")
	if err != nil {
		panic(err.Error())
	}

	dbase.InitDataBase()
	for quarry.Next() {
		var tableName string
		err := quarry.Scan(&tableName)
		if err != nil {
			panic(err.Error())
		}
		// Print the table name
		dbase.TablesMap[tableName] = true

	}

	return db
}

func main() {

	// Database logic is being written here
	db := initDb()

	defer db.Close()
	fmt.Println("Database Connected")

	dbase.CheckAndCreateTables(db)
	//Database logic ended here

	//further we use db to call database from API
	router := gin.Default()

	// Simple group: v1
	v1.V1Group(router)

	router.Run(":8080")

}
