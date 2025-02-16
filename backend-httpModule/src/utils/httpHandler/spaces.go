package httpHandler

import (
	"encoding/json"
	"fmt"
	"main/src/utils/databaseHelper"
	"net/http"
	"strconv"
)

func GetAllSpaces(w http.ResponseWriter, r *http.Request) {

	rows, err := databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllSpace)
	if err != nil {
		fmt.Println("Error Getting Data from DB:", err)
		http.Error(w, "Error Getting Data from DB", http.StatusBadRequest)
		return

	}
	var MapData []databaseHelper.Space

	for rows.Next() {
		var tempMap databaseHelper.Space
		rows.Scan(&tempMap.Id, &tempMap.Thumbnail, &tempMap.UserId)

		MapData = append(MapData, tempMap)
	}

	data, err := json.Marshal(MapData)
	if err != nil {
		fmt.Println("Error With Data Unmarshaling :", err)
		http.Error(w, "Error With Data Unmarshaling", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Success(string(data))))
}

// Common Data will be send to both users who create Space by using map id
// or attempts to start existing space
// Some Random Structre we can expect from user to decide to copy and create new space to user from map
// Or just Gater all required data from admin of that space

// 1st Requirement AdminId and SpaceId should be send using quarrys localhost:8080/adminId/spaceId

func GetFullSpaceDetails(w http.ResponseWriter, r *http.Request) {
	// data, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	fmt.Println("Error Getting Data from DB:", err)
	// 	http.Error(w, "Error Getting Data from DB", http.StatusBadRequest)
	// 	return

	// }

	//Extra
	// var SpaceIdCreaedByMap int64
	// -----

	//
	SpaceId := r.URL.Query().Get("spaceID")
	// fmt.Println(SpaceId)
	// newSpaceId := r.URL.Query().Get("newSpaceId")
	//
	// if err != nil {
	// 	fmt.Println("Invalid Map Key:", err)
	// 	http.Error(w, "Invalid Map Key", http.StatusBadRequest)
	// 	return
	// }
	// if newSpaceId != "" {

	// 	for _, SingleMapElement := range mapElemets {
	// 		res, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoSpaceElement, SingleMapElement.X, SingleMapElement.Y, SingleMapElement.ElementId)
	// 		if err != nil {
	// 			fmt.Println("Invalid Map Key:", err)
	// 			http.Error(w, "Invalid Map Key", http.StatusBadRequest)
	// 			return
	// 		}
	// 		fmt.Println(res)
	// 		spaceELementId,_ := res.LastInsertId()
	// 		SingleMapElement.Id = int(spaceELementId)
	// 		// res, _ = databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoAllSpaceElementsTable,SpaceIdCreaedByMap,spaceELementId)
	// 	}
	// }
	// Keep Common Way to give data to user
	// Data sent to user
	// 1.Space Details All belongs to Id provided
	// 2.SpaceElement All Details which delongs to that spaceId
	// 3.All Elements Which Belongs to that SpaceElement or Send All Space ELements to Client and Let Him Choose accordingly to manipulate data
	// Note : later code should be updated to only send data from here insted of all, move to only needed

	var Space databaseHelper.Space
	var SpaceElements []databaseHelper.SpaceElement
	var AllSpaceElements []databaseHelper.AllSpaceElements
	var Elements []databaseHelper.Element

	ConvertedID, err := strconv.Atoi(SpaceId)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	row := databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.GetSpaceById, ConvertedID)

	err = row.Scan(&Space.Id, &Space.Thumbnail, &Space.UserId)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	rows, err := databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllFromAllSpaceElementsTableWithID, ConvertedID)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	for rows.Next() {
		var CurrentAllSpaceElement databaseHelper.AllSpaceElements

		err := rows.Scan(&CurrentAllSpaceElement.Id, &CurrentAllSpaceElement.SpaceId, &CurrentAllSpaceElement.SpaceElementId)
		if err != nil {
			fmt.Println("Invalid Map Key:", err)
			http.Error(w, "Invalid Map Key", http.StatusBadRequest)
			return
		}
		AllSpaceElements = append(AllSpaceElements, CurrentAllSpaceElement)
	}

	// Get All SpaceElements by using spaceelementid we got from all space element id datailss table
	for _, AllSpaceElement := range AllSpaceElements {
		fmt.Println(AllSpaceElement)
		var CurrentSpaceElement databaseHelper.SpaceElement
		rows := databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.GetAllSpaceElementsBySpaceElementId, AllSpaceElement.SpaceElementId)
		err := rows.Scan(&CurrentSpaceElement.Id, &CurrentSpaceElement.X, &CurrentSpaceElement.Y, &CurrentSpaceElement.ElementId)
		if err != nil {
			fmt.Println("Invalid Map Key:", err)
			http.Error(w, "Invalid Map Key", http.StatusBadRequest)
			return
		}
		SpaceElements = append(SpaceElements, CurrentSpaceElement)
	}

	// Attempt to get all elemnts and store them

	rows, err = databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllElements)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}

	for rows.Next() {
		var currentElement databaseHelper.Element
		err := rows.Scan(&currentElement.Id, &currentElement.Witdth, &currentElement.Height, &currentElement.ImageUrl)
		if err != nil {
			fmt.Println("Invalid Map Key:", err)
			http.Error(w, "Invalid Map Key", http.StatusBadRequest)
			return
		}
		Elements = append(Elements, currentElement)
	}

	var spaceString []byte
	var spaceELementString []byte
	var AllElementsString []byte
	spaceString, _ = json.Marshal(&Space)
	spaceELementString, _ = json.Marshal(&SpaceElements)
	AllElementsString, err = json.Marshal(&Elements)

	if err != nil {
		fmt.Println("Invalid Data:", err)
		http.Error(w, "Invalid Data", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(SendMapDetailsToUser(string(spaceString), string(spaceELementString), string(AllElementsString))))

}

func MoveMapToSpace(w http.ResponseWriter, r *http.Request) {
	MapId := r.URL.Query().Get("MapId")
	SpaceOwnerId := r.URL.Query().Get("userId")
	SpaceOwnId, err := strconv.Atoi(SpaceOwnerId)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	fmt.Println("Creating Map to Space Attempted")

	// Add Logic to Move All Data from Map to Space
	var MapDetails databaseHelper.Map
	id, err := strconv.Atoi(MapId)
	fmt.Println(id, "This is map Id")
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	row := databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.GetMapWithId, id)

	err = row.Scan(&MapDetails.Id, &MapDetails.Thumbnail, &MapDetails.AdminId)
	if err != nil {
		fmt.Println("Error With Database:", err)
		http.Error(w, "Error With Database", http.StatusBadRequest)
		return
	}
	res, err := databaseHelper.DATABASE_GLOBAL.Exec(databaseHelper.InsertIntoSpace, MapDetails.Thumbnail, SpaceOwnId)
	fmt.Println("Space Insertion Happend")
	
	// SpaceID, err := res.LastInsertId()
	// SpaceId = strconv.FormatInt(SpaceID, 10)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	fmt.Println(res, "This is Map Retrieved", MapDetails)
	SpaceIdCreaedByMap, _ := res.LastInsertId()
	rows, err := databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllFromAllMapElementsTableUsingMapId, MapDetails.Id)
	if err != nil {
		fmt.Println("Invalid Map Key:", err)
		http.Error(w, "Invalid Map Key", http.StatusBadRequest)
		return
	}
	var mapElementsDetails []databaseHelper.AllMapElements
	for rows.Next() {
		var curr databaseHelper.AllMapElements
		rows.Scan(&curr.Id, &curr.MapId, &curr.MapElementId)
		mapElementsDetails = append(mapElementsDetails, curr)
	}
	fmt.Println("All Map Elements Retrieved",mapElementsDetails)

	var mapElemets []databaseHelper.MapElement
	for _, AllMapELementObj := range mapElementsDetails {
		var current databaseHelper.MapElement
		row := databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.GetAll_MapElementsById, AllMapELementObj.MapElementId)
		err := row.Scan(&current.Id, &current.X, &current.Y, &current.ElementId)
		if err != nil {
			fmt.Println("Invalid Db :", err)
			http.Error(w, "Invalid ", http.StatusBadRequest)
			return
		}
		mapElemets = append(mapElemets, current)
	}

	fmt.Println(" Map Elements Retrieved",mapElemets)

	for i := 0; i < len(mapElemets); i++ {

		row, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoSpaceElement, mapElemets[i].X, mapElemets[i].Y, mapElemets[i].ElementId)
		if err != nil {
			fmt.Println("Issue With DB:", err)
			http.Error(w, "Issue With DB", http.StatusBadRequest)
			return
		}
		lastRowId, _ := row.LastInsertId()
		mapElemets[i].Id = int(lastRowId)
	}

	for i := 0; i < len(mapElemets); i++ {
		row, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoAllSpaceElementsTable, SpaceIdCreaedByMap, mapElemets[i].Id)
		if err != nil {
			fmt.Println("Issue With DB:", err)
			http.Error(w, "Issue With DB", http.StatusBadRequest)
			return
		}
		data, _ := row.LastInsertId()
		fmt.Println(data)
		fmt.Println("pushing initiated,",SpaceIdCreaedByMap,mapElemets[i])
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Success("Good Map is Sucessfully converted to Space ")))

}

// the below functions are can be updated when attempting to revisit again this project
// currently flow of this project is getSpace if user already created or just select map from screen
// by using id he provided move all elements from map to space with respect to userId
func AddSpace(w http.ResponseWriter, r *http.Request) {
}

func AddSpaceElements(w http.ResponseWriter, r *http.Request) {
}
