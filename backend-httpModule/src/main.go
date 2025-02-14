package main

import (
	"main/src/utils/databaseHelper"
	"main/src/utils/httpHandler"
)

func main() {
	databaseHelper.ConnectAndPingDataBase()
	

	httpHandler.ListenAndServeRoutes()


}