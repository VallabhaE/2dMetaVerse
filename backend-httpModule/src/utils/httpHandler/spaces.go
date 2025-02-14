package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"main/src/utils/databaseHelper"
	"net/http"
)

func GetAllSpaces(w http.ResponseWriter, r *http.Request) {
	rows, err := databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllMap)
	if err != nil {
		fmt.Println("Error Getting Data from DB:", err)
		http.Error(w, "Error Getting Data from DB", http.StatusBadRequest)
		return

	}
	var MapData []databaseHelper.Map

	for rows.Next() {
		var tempMap databaseHelper.Map
		rows.Scan(&tempMap.Id, &tempMap.Thumbnail, &tempMap.AdminId)

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

func GetFullMapDetails(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error Getting Data from DB:", err)
		http.Error(w, "Error Getting Data from DB", http.StatusBadRequest)
		return

	}
	SpaceOwnerId := r.URL.Query().Get("adminId")
	SpaceId := r.URL.Query().Get("spaceID")
	newSpaceId := r.URL.Query().Get("newSpaceId")

	if newSpaceId != "" {
		// Add Logic to Move All Data from Map to Space
	}
	// Keep Common Way to give data to user
	// Data sent to user
	// 1.Space Details All belongs to Id provided
	// 2.SpaceElement All Details which delongs to that spaceId
	// 3.All Elements Which Belongs to that SpaceElement or Send All Space ELements to Client and Let Him Choose accordingly to manipulate data
	// Note : later code should be updated to only send data from here insted of all, move to only needed


	

}

// the below functions are can be updated when attempting to revisit again this project
// currently flow of this project is getSpace if user already created or just select map from screen
// by using id he provided move all elements from map to space with respect to userId
func AddSpace(w http.ResponseWriter, r *http.Request) {
}

func AddSpaceElements(w http.ResponseWriter, r *http.Request) {
}
