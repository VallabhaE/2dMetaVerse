package main

import (
	"database/sql"
	"fmt"

	"main/dbase"
	"main/v1"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
)






func main() {

	// Database logic is being written here
	db := dbase.InitDb()
	GLOBAL_DB_CONNECTION = db

	defer db.Close()
	fmt.Println("Database Connected")

	dbase.CheckAndCreateTables(db)
	//Database logic ended here

	//further we use db to call database from API
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
  

	// Simple group: v1
	v1.V1Group(router,db)

	router.Run(":8080")

}
