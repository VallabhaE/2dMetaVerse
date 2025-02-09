package v1

import (
	"github.com/gin-gonic/gin"
	"main/dbase"
	"main/v1/controllers"
	"net/http"
)

func GetAllSpaceInfo(ctx *gin.Context) {
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

	quarry := "select * from spaces where userid = ?"

	rows, err := dbase.GLOBAL_DB_CONNECTION.Query(quarry, userInfo.Id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with Database Please contact to admins"})
		return
	}
	var Spaces []controllers.Space
	for rows.Next() {
		var space controllers.Space
		if err := rows.Scan(&space.Id, &space.Thumbnail, &space.UserId); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning row"})
			return
		}
		Spaces = append(Spaces, space)

	}


	ctx.JSON(http.StatusOK, gin.H{"message": "Found Spaces","Spaces":Spaces})

}


func CreateSpaceElement(ctx *gin.Context) {
    // Validate the user session using the "auth" cookie
    data, err := ctx.Request.Cookie("auth")
    if err != nil {
        ctx.JSON(http.StatusOK, gin.H{"error": "Issue with cookie please login"})
        return
    }

    // Retrieve the user information from the session map
    userInfo, exist := controllers.SESSION_MAP[data.Value]
    if !exist {
        ctx.SetCookie("auth", "", -1, "/", "", true, true)
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Issue with cookie please login"})
        return
    }

    // Struct to bind the request data (space info)
    var req struct {
        Thumbnail  string `json:"thumbnail"`
        MapId      int    `json:"mapId"`
        ElementIds []int  `json:"elementIds"` // List of space element IDs
    }

    // Bind the request JSON to the struct
    if err := ctx.ShouldBindJSON(&req); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
        return
    }

    // Create a new space entry
    quarryCreateSpace := "INSERT INTO Space (thumbnail, userId) VALUES (?, ?)"
    result, err := dbase.GLOBAL_DB_CONNECTION.Exec(quarryCreateSpace, req.Thumbnail, userInfo.Id)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating space"})
        return
    }

    // Get the ID of the newly created space
    spaceID, err := result.LastInsertId()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving space ID"})
        return
    }

    // Insert space elements for this new space
    for _, elementID := range req.ElementIds {
        // Add the element to the space (from spaceElement table)
        quarryInsertElement := "INSERT INTO allSpaceElements (spaceId, spaceElementId) VALUES (?, ?)"
        _, err := dbase.GLOBAL_DB_CONNECTION.Exec(quarryInsertElement, spaceID, elementID)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error associating element to space"})
            return
        }
    }

    // If everything is successful, return the created space ID and success message
    ctx.JSON(http.StatusOK, gin.H{"message": "Space created successfully", "spaceId": spaceID})
}
