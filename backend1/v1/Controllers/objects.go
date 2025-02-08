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
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	AvatarId string `json:"avatarid"`
}

type Avatars struct{
	Id int `json"id"`
	AvatarName string `json"avatarName"`
	AvatarImg string `json"avatarImg"`
	Height int `json"Height"`
	Width int `json"Width"`
}

// id int PRIMARY key auto_increment,
//     avatarName varchar(255),
//     avatarImg varchar(255),
//     height int
//     width int