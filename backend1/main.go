package main

import (
	"fmt"

	"main/dbase"
	"main/v1"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)






func main() {

	// Database logic is being written here
	db := dbase.InitDb()
	dbase.GLOBAL_DB_CONNECTION = db

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
