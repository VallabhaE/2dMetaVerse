create database metaverse;
use metaverse;


-- Table for Creation of Users and validation 
create table users(
	id int PRIMARY KEY auto_increment,
    username varchar(255) UNIQUE,
    email varchar(255) UNIQUE,
    password varchar(255)
);


-- Table for Creation of Admins and validation 
CREATE TABLE admins (
    id INT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(255) UNIQUE,
    email VARCHAR(255) UNIQUE,
    password VARCHAR(255)
);





--  Space Creation by User

-- User attempts to create space by selection space or map 
-- So key is Elements(Contains Size and imageURL), 
-- Lets Attempt to store images in public folder so we can just send image names
CREATE TABLE Element (
    id INT PRIMARY KEY AUTO_INCREMENT,
    width INT,
    height INT,
    imageURL VARCHAR(255)
);

-- Create Element Holders for admin and users so it contains ref to Element and pos(x,y) in FE
CREATE TABLE spaceElement (
    id INT PRIMARY KEY AUTO_INCREMENT,
    x INT,
    y INT,
    ElementId int,
    FOREIGN KEY (ElementId) REFERENCES Element(id)
);

CREATE TABLE mapElement (
    id INT PRIMARY KEY AUTO_INCREMENT,
    x INT,
    y INT,
    ElementId int,
    FOREIGN KEY (ElementId) REFERENCES Element(id)
);


--  Lets Create Space and Map for indidual user

create table Space(
	    id INT PRIMARY KEY AUTO_INCREMENT,
        thumbnail varchar(255),
        userId int,
        FOREIGN KEY (userId) REFERENCES users(id)
        -- A single space can have Multiple SpaceElements so lets create table
);

create table allSpaceElements(
 id INT PRIMARY KEY AUTO_INCREMENT, 
 spaceId int,
 FOREIGN KEY (spaceId) REFERENCES Space(id),
 spaceElementId int,
 FOREIGN KEY (spaceElementId) REFERENCES spaceElement(id),
);



-- do same for Admin Table


create table Map(
	    id INT PRIMARY KEY AUTO_INCREMENT,
        thumbnail varchar(255),
        adminId int,
        FOREIGN KEY (adminId) REFERENCES admins(id)
        -- A single space can have Multiple SpaceElements so lets create table
);

create table allMapElements(
 id INT PRIMARY KEY AUTO_INCREMENT, 
 mapId int,
 FOREIGN KEY (mapId) REFERENCES Map(id),
 mapElementId int,
 FOREIGN KEY (mapElementId) REFERENCES mapElement(id),
);

