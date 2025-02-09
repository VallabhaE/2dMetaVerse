package v1

import (
	"github.com/gin-gonic/gin"
	"main/dbase"
	"main/v1/controllers"
	"net/http"
)

func GetAllMapInfo(ctx *gin.Context) {
	data, err := ctx.Request.Cookie("auth")
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
		return
	}

	userInfo, exist := controllers.SESSION_MAP[data.Value]
	if !exist {
		ctx.SetCookie("auth", "", -1, "/", "", true, true)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with cookie please login"})
		return
	}

	quarry := "select * from map where userid = ?"

	rows, err := dbase.GLOBAL_DB_CONNECTION.Query(quarry, userInfo.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with Database Please contact to admins"})
		return
	}
	var maps []controllers.Map
	for rows.Next() {
		var Map controllers.Map
		if err := rows.Scan(&Map.Id, &Map.Thumbnail, &Map.AdminId); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		maps = append(maps, Map)

	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Found Spaces","Spaces":maps})

}


// controlers should be created which 
// 1.Expects element data and loc details of it and tables will be created accouringly for both
// MapElemet and SpaceElement accoring to it

// Assuming that ElementObject has fields such as width, height, and imageURL
func CreateElement(ctx *gin.Context) {
    var req controllers.ElementObject

    // Bind the incoming JSON data into the struct
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // Insert the element into the Element table
    quarry1forElement := "INSERT INTO Element (width, height, imageURL) VALUES (?, ?, ?)"
    result, err := dbase.GLOBAL_DB_CONNECTION.Exec(quarry1forElement, req.Width, req.Height, req.ImageURL)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert element"})
        return
    }

    // Get the ID of the inserted element
    elementID, err := result.LastInsertId()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve element ID"})
        return
    }

    // Now insert into spaceElement or mapElement based on your context (assuming SpaceID or MapID is provided in the request)
    // Assuming spaceID is provided in the request for this example
    quarry2forSpaceElement := "INSERT INTO spaceElement (x, y, ElementId) VALUES (?, ?, ?)"
    _, err = dbase.GLOBAL_DB_CONNECTION.Exec(quarry2forSpaceElement, req.X, req.Y, elementID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into spaceElement"})
        return
    }

    // Optionally, you can insert into mapElement if needed (depends on your logic)
    quarry3forMapElement := "INSERT INTO mapElement (x, y, ElementId) VALUES (?, ?, ?)"
    _, err = dbase.GLOBAL_DB_CONNECTION.Exec(quarry3forMapElement, req.X, req.Y, elementID)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into mapElement"})
        return
    }

    // If everything is successful, return the created element ID as a response
    ctx.JSON(http.StatusOK, gin.H{"message": "Element created successfully", "elementId": elementID})
}



func CreateMapElementInAllMapElements(ctx *gin.Context) {
    var req struct {
        MapId        int `json:"mapId"`         // The Map ID
        MapElementId int `json:"mapElementId"`   // The ID of the map element
    }

    // Bind the incoming JSON data into the struct
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // Insert the map element into the allMapElements table
    quarry := "INSERT INTO allMapElements (mapId, mapElementId) VALUES (?, ?)"
    result, err := dbase.GLOBAL_DB_CONNECTION.Exec(quarry, req.MapId, req.MapElementId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert into allMapElements"})
        return
    }

    // Get the ID of the inserted allMapElements entry (you might want to return this if needed)
    allMapElementID, err := result.LastInsertId()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve allMapElement ID"})
        return
    }

    // Optionally, you could store this entry's ID in the Map (depending on your logic)
    quarryUpdateMap := "UPDATE Map SET allMapElementId = ? WHERE id = ?"
    _, err = dbase.GLOBAL_DB_CONNECTION.Exec(quarryUpdateMap, allMapElementID, req.MapId)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update Map with allMapElement ID"})
        return
    }

    // If everything is successful, return the success response
    ctx.JSON(http.StatusOK, gin.H{"message": "Map element added to allMapElements successfully", "allMapElementId": allMapElementID})
}
