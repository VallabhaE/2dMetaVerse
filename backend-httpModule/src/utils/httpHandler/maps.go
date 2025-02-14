package httpHandler

import (
	"encoding/json"
	"fmt"
	"io"
	"main/src/utils/databaseHelper"
	"net/http"
	"strconv"
)

func GetAllMaps(w http.ResponseWriter, r *http.Request) {
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

func AddMap(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error With Data Reading :", err)
		http.Error(w, "Error With Data Reading", http.StatusBadRequest)
	}
	var UserMapDetails databaseHelper.Map
	err = json.Unmarshal(data, &UserMapDetails)
	if err != nil {
		fmt.Println("Error With Data Unmarshaling :", err)
		http.Error(w, "Error With Data Unmarshaling", http.StatusBadRequest)
	}

	res, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoMap, UserMapDetails.Thumbnail, UserMapDetails.AdminId)
	if err != nil {
		fmt.Println("Database Error :", err)
		http.Error(w, "Error With DataBase"+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	Ints, _ := res.LastInsertId()
	w.Write([]byte(Success(strconv.FormatInt(Ints, 10))))
}

func AddMapElements(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error With Data Reading :", err)
		http.Error(w, "Error With Data Reading", http.StatusBadRequest)
	}
	var UserMapDetails databaseHelper.MapElement
	err = json.Unmarshal(data, &UserMapDetails)
	if err != nil {
		fmt.Println("Error With Data Unmarshaling :", err)
		http.Error(w, "Error With Data Unmarshaling", http.StatusBadRequest)
	}

	res, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertIntoMapELement, UserMapDetails.X, UserMapDetails.Y, UserMapDetails.ElementId)
	if err != nil {
		fmt.Println("Database Error :", err)
		http.Error(w, "Error With DataBase"+err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(http.StatusOK)
	Ints, _ := res.LastInsertId()
	w.Write([]byte(Success(strconv.FormatInt(Ints, 10))))
}
