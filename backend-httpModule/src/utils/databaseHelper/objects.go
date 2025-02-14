package databaseHelper



// User Table Struct
type User struct{
	Id int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"Password"`
	AvatarId int `josn:"avatarId"`
}


// Admin Table Struct

type Admin struct{
	Id int
	Username string
	Email string
	Password string
}

// AllMapElements Table Struct

type AllMapElements struct{
	Id int
	MapId int
	MapElementId int
}


// AllSpaceelements  Table Struct

type AllSpaceElements struct{
	Id int
	SpaceId int
	SpaceElementId int
}


// Element Table Struct
// {
//     "width":21,
//     "height":22,
//     "imageUrl":2
// }
type Element struct{
	Id int `json:"id"`
	Witdth int `json:"width"`
	Height int `json:"height"`
	ImageUrl string `json:"imageUrl"`
}

// Map Table Struct

type Map struct{
	Id int `json:"id"`
	Thumbnail string `json:"thumbnail"`
	AdminId int `json:"adminId"`
}

// MapElement table struct

type MapElement struct{
	Id int
	X int
	Y int
	ElementId int
}


// Space table Struct
type Space struct{
	Id int
	Thumbnail int
	UserId int
}


// SpaceElement table Struct
type SpaceElement struct{
	Id int 
	X int
	Y int
	ElementId int
}