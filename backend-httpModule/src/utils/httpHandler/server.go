package httpHandler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"main/src/utils/databaseHelper"
	"net/http"
)

var Mux *http.ServeMux

func MethodCheck(method string, fun func(w http.ResponseWriter, r *http.Request)) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == method {
			fun(w, r) // Call the provided handler function if method matches
		} else {
			// Handle case when method does not match
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}

func ListenAndServeRoutes() {
	Mux = http.NewServeMux()

	// User Level Routes
	Mux.Handle("/signup", MethodCheck(http.MethodPost, SignUp))
	Mux.Handle("/signin", MethodCheck(http.MethodPost, SignIn))

	// All Below Routes are Must Be protected every one are expected to send with auth barier and jwt token

	// Information Level Routes  -- Level -1
	// 1.GetAllElements,GetAllAvatars,GetALLMaps,GetAllSpaces,

	// For Elements
	Mux.Handle("/GetAllElements", Middleware(MethodCheck(http.MethodGet, GetAllElements)))
	Mux.Handle("/AddElementToDB", Middleware(MethodCheck(http.MethodPost, SetElement)))

	// For Avatars
	Mux.Handle("/SetAvatar", Middleware(MethodCheck(http.MethodPut, SetAvatar)))

	// For Maps
	Mux.Handle("/GetAllMaps", Middleware(MethodCheck(http.MethodGet, GetAllMaps)))
	Mux.Handle("/AddMap", Middleware(MethodCheck(http.MethodPost, AddMap)))
	Mux.Handle("/AddMapElements", Middleware(MethodCheck(http.MethodPost, AddMapElements)))

	//Get All Space
	Mux.Handle("/GetAllSpace", Middleware(MethodCheck(http.MethodGet, GetAllSpaces)))

	// Advanced Information Level Routes
	// 1. If user Selected Map that should be created a space for him and move all elements and detials to user space who requested
	// 2. when user attempts to join a space he can get roomId so that FE will attempt to connect that perticular room
	Mux.Handle("/GetFullMapDetails", Middleware(MethodCheck(http.MethodGet, GetFullSpaceDetails)))
	Mux.Handle("/MoveMapToSpace",Middleware(MethodCheck(http.MethodGet,MoveMapToSpace)))

	log.Println("Listning on Prot :8080")
	err := http.ListenAndServe(":8080", Mux)

	if err != nil {
		fmt.Println(err)
	}
}

// Sample Raw Json Accepted
//
//	{
//	    "type":"admin",
//	    "username":"BhaiBhai",
//	    "Email":"123455678812@gmail.com",
//	    "Password":"12345678",
//	    "ConformPassword":"12345678"
//	}
func SignUp(w http.ResponseWriter, r *http.Request) {
	var newUser = struct {
		Type            string `json:"type"`
		Username        string `json:"username"`
		Email           string `json:"Email"`
		Password        string `json:"Password"`
		ConformPassword string `json:"ConformPassword"`
	}{}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &newUser)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	json.Unmarshal(data, &newUser)
	fmt.Println("Data Reached")
	fmt.Println(newUser)

	if newUser.ConformPassword != newUser.Password || newUser.Email == "" || newUser.Username == "" || len(newUser.ConformPassword) < 5 {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(DetailsError))
		return
	}

	if newUser.Type == "admin" {
		_, err = databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertAdmin, newUser.Username, newUser.Email, newUser.Password)
	} else {
		_, err = databaseHelper.GetGlobalDatabase().Exec(databaseHelper.InsertUser, newUser.Username, newUser.Email, newUser.Password)
	}

	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
		w.Write([]byte(DetailsError))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Success("OK")))

}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var newUser = struct {
		Type     string `json:"type"`
		Username string `json:"username"`
		Password string `json:"Password"`
		AvatarId int    `josn:"avatarId"`
	}{}

	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body:", err)
		http.Error(w, "Failed to read body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(data, &newUser)

	if err != nil {
		fmt.Println("err", err)
		return
	}

	json.Unmarshal(data, &newUser)
	fmt.Println("Data Reached")

	if newUser.Password == "" || newUser.Username == "" {
		w.WriteHeader(http.StatusExpectationFailed)
		w.Write([]byte(DetailsError))
		return
	}
	var rows *sql.Row
	if newUser.Type == "admin" {
		rows = databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.CheckAdminExistence, newUser.Username, newUser.Password)
	} else {
		rows = databaseHelper.GetGlobalDatabase().QueryRow(databaseHelper.CheckUserExistence, newUser.Username, newUser.Password)
	}

	var UserCheck databaseHelper.User
	err = rows.Scan(&UserCheck.Id, &UserCheck.Username, &UserCheck.Email, &UserCheck.Password)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
		w.Write([]byte(DetailsError))
		return
	}

	token, err := createToken(newUser.Username)
	if err != nil {
		w.WriteHeader(http.StatusExpectationFailed)
		fmt.Println(err)
		w.Write([]byte(DetailsError))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(Success(token)))
	
}
