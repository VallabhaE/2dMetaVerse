package dbase

import (
	"database/sql"
	"fmt"
)
func InitDb() *sql.DB {
	var err error
	GLOBAL_DB_CONNECTION, err = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/metaverse")

	if err != nil {
		panic(err.Error())
	}
	

	quarry, err := GLOBAL_DB_CONNECTION.Query("show tables;")
	if err != nil {
		panic(err.Error())
	}

	InitDataBase()
	for quarry.Next() {
		var tableName string
		err := quarry.Scan(&tableName)
		if err != nil {
			panic(err.Error())
		}
		// Print the table name
		TablesMap[tableName] = true
	}
	return GLOBAL_DB_CONNECTION
}

func CheckAndCreateTables(db *sql.DB) {

	var tables = map[string]string{"users":CreateUserTable,"admins":CreateAdminTable,"element":ElementTable,"spaceelement":SpaceElement,"mapelement":MapElementTable,"space":SpaceTable,"allspaceelements":AllSpaceElementsTable,"map":MapTable,"allmapelements":AllMapElementTable}
	for k,v := range tables{
		if TablesMap[k]==false{
			_,err:= db.Query(v)
			if err!=nil{
				fmt.Println(err.Error())
			}
			fmt.Println("ExecutedQuarrry")
		}else{
			fmt.Println(k,"->Available")
		}
	}

	fmt.Println("Total Tables Created",len(TablesMap))
}
