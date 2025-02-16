package httpHandler

import (
	// "main/src/httpHandler"
	"encoding/json"
	"fmt"
	"io"
	"main/src/utils/databaseHelper"
	"net/http"
	"strconv"
)

// Function Will Get All Elements Available in the database
// USECASE : if user created empty map and want to add elements in his map manually,this REST_API will be Helpfull
func GetAllElements(w http.ResponseWriter, r *http.Request) {
	rows, err := databaseHelper.GetGlobalDatabase().Query(databaseHelper.GetAllElements)

	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	var AllElement []databaseHelper.Element
	for rows.Next() {
		var Element databaseHelper.Element
		rows.Scan(&Element.Id, &Element.Witdth, &Element.Height, &Element.ImageUrl)
		AllElement = append(AllElement, Element)
	}

	data, err := json.Marshal(AllElement)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Success(string(data))))

}

// Sample Message to below Func
// {
//     "width":"21",
//     "height":"22",
//     "imageUrl":"212"
// }

func SetElement(w http.ResponseWriter, r *http.Request) {
	var Element databaseHelper.Element
	data, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	json.Unmarshal(data, &Element)
	res, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertElement, Element.Witdth, Element.Height, Element.ImageUrl)
	if err != nil {
		fmt.Println("Error Loading body:", err)
		http.Error(w, "Failed to Load body", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	w.Write([]byte(Success(strconv.FormatInt(id, 10))))

}

func SetAvatar(w http.ResponseWriter, r *http.Request) {
	var AvatarId Avatar

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}

	json.Unmarshal(data, &AvatarId)

	res, err := databaseHelper.GetGlobalDatabase().Exec(databaseHelper.UpdateAvaterIdByUsername, AvatarId.AvatarId,AvatarId.Username)

	if err != nil {
		fmt.Println("Error Moving to database:", err)
		http.Error(w, "Error Moving to database", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	w.Write([]byte(Success(strconv.FormatInt(id, 10))))
}
