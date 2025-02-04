//DataBase Related Functions only should add in this folder
package dbase

var TablesMap map[string]bool


const CreateUserTable = `create table users(
	id int PRIMARY KEY auto_increment,
    username varchar(255) UNIQUE,
    email varchar(255) UNIQUE,
    password varchar(255)
);
`

const CreateAdminTable = `CREATE TABLE admins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);`

const ElementTable = `CREATE TABLE Element (
    id INT PRIMARY KEY AUTO_INCREMENT,
    width INT,
    height INT,
    imageURL VARCHAR(255)
);`

const SpaceElement = `CREATE TABLE spaceElement (
    id INT PRIMARY KEY AUTO_INCREMENT,
    x INT,
    y INT,
    ElementId int,
    FOREIGN KEY (ElementId) REFERENCES Element(id)
);`

const MapElementTable = `CREATE TABLE mapElement (
    id INT PRIMARY KEY AUTO_INCREMENT,
    x INT,
    y INT,
    ElementId int,
    FOREIGN KEY (ElementId) REFERENCES Element(id)
);`

const SpaceTable = `create table Space(
	    id INT PRIMARY KEY AUTO_INCREMENT,
        thumbnail varchar(255),
        userId int,
        FOREIGN KEY (userId) REFERENCES users(id)
        -- A single space can have Multiple SpaceElements so lets create table
);`


const AllSpaceElementsTable = `create table allSpaceElements(
 id INT PRIMARY KEY AUTO_INCREMENT, 
 spaceId int,
 FOREIGN KEY (spaceId) REFERENCES Space(id),
 spaceElementId int,
 FOREIGN KEY (spaceElementId) REFERENCES spaceElement(id)
);
`


const MapTable = `create table Map(
	    id INT PRIMARY KEY AUTO_INCREMENT,
        thumbnail varchar(255),
        adminId int,
        FOREIGN KEY (adminId) REFERENCES admins(id)
        -- A single space can have Multiple SpaceElements so lets create table
);`

const AllMapElementTable = `create table allMapElements(
 id INT PRIMARY KEY AUTO_INCREMENT, 
 mapId int,
 FOREIGN KEY (mapId) REFERENCES Map(id),
 mapElementId int,
 FOREIGN KEY (mapElementId) REFERENCES mapElement(id)
);
`


func InitDataBase(){
	TablesMap = make(map[string]bool)
}
