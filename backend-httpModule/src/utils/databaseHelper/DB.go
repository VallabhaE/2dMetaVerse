package databaseHelper
// package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var DATABASE_GLOBAL *sql.DB



func ConnectAndPingDataBase() {
    cfg := mysql.Config{
        User:   "root",
        Passwd: "root",
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "metaverse",
    }
    var err error
    DATABASE_GLOBAL, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    pingErr := DATABASE_GLOBAL.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
    }
    fmt.Println("Connected!")
}

func GetGlobalDatabase() *sql.DB{
	return DATABASE_GLOBAL
}



