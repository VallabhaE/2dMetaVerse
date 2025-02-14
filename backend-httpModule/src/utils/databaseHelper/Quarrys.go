package databaseHelper


// User Insertion and Verification
var InsertUser = `insert into users(username,email,password) values (?,?,?);`
var CheckUserExistence = `select * from users where username=? and password =?;`
var GetAllUsers = `select * from users;`
var UpdateAvaterIdByUsername = `update users set avatarid=? where username = ?;`

// Admin Insertion and Verification
var InsertAdmin = `insert into admins(username,email,password) values (?,?,?);`
var CheckAdminExistence = `select * from admins where username=? and password =?;`
var GetAllAdmins = `select * from admins;`

// AllMapElements Table

var GetAllFromAllMapElementsTable = `select * from allmapelements;`
var InsertIntoAllMapElementsTable = `insert into allmapelements(mapId,mapElementId) values(?,?);`
var DeleteFromAllMapElementsTable = `delete from allmapelements where id = ?;`


// AllSpaceElements Table

var GetAllFromAllSpaceElementsTable = `select * from allspaceelements;`
var InsertIntoAllSpaceElementsTable = `insert into allSpaceelements(spaceId,spaceElementId) value(?,?);`
var DeleteFromAllSpaceElementsTable = `delete from allSpaceelements where id = ?;`


// Element Table

var GetAllElements = `select * from element;`
var InsertElement = `insert into element(width,height,imageURL) values (?,?,?);`
var DeleteElement = `delete from element where id = ?;`

// Map Table

var GetAllMap = `select * from map;`
var InsertIntoMap = `insert into map(thumbnail,adminid) values (?,?);`
var DeleteMap = `delete from map where id = ?`

// Space Table
var GetAllSpace = `select * from space;`
var InsertIntoSpace = `insert into space(thumbnail,userid) values (?,?);`
var DeleteSpace = `delete from space where id = ?;`



// MapElement Table

var GetAllMapElements = `select * from mapElement;`
var InsertIntoMapELement = `insert into mapelement (x,y,elementId) values (?,?,?);`
var DeleteMapElement = `delete from mapElement where id = ?;`

// Space Table
var GetAllSpaceElements = `select * from spaceelement;`
var InsertIntoSpaceElement = `insert into spaceelement(x,y,elementid) values (?,?,?);`
var DeleteSpaceElement = `delete from spaceelement where id = ?;`

