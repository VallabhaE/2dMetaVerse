package dbase

import (
	"database/sql"
	"fmt"
)

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
