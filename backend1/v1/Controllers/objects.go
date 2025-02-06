package controllers

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type SignUpRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConformPassword string `json:"conformpassword`
}
type User struct{
	id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}