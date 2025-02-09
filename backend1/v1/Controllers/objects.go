package controllers

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}


type SignUpRequest struct {
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	ConformPassword string `json:"conformpassword"`
}
type User struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
	AvatarId string `json:"avatarid"`
}

type Avatars struct {
    Id         int    `json:"id"`
    AvatarName string `json:"avatarName"`
    AvatarImg  string `json:"avatarImg"`
    Height     int    `json:"Height"`
    Width      int    `json:"Width"`
}
type Space struct{
	Id int `json:"id"`
	Thumbnail string `json:"thumbnail"`
	UserId int `json:"userId"`
}


type Map struct{
	Id int `json:"id"`
	Thumbnail string `json:"thumbnail"`
	AdminId int `json:"adminid"`
}


type ElementObject struct {
    X         int    `json:"x"`
    Y         int    `json:"y"`
    Width     int    `json:"width"`
    Height    int    `json:"height"`
    ImageURL  string `json:"imageURL"`
}

// id INT PRIMARY KEY AUTO_INCREMENT,
// x INT,
// y INT,
// ElementId int,
// FOREIGN KEY (ElementId) REFERENCES Element(id)
// );
// CREATE TABLE Element (
//     id INT PRIMARY KEY AUTO_INCREMENT,
//     width INT,
//     height INT,
//     imageURL VARCHAR(255)
// );