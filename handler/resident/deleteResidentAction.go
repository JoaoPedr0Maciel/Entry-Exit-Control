package handler

import (
	"entry-exit-api/database"
	"entry-exit-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)


func DeleteResidentAction(ctx *gin.Context) {
	db := database.GetMysql()
	id := ctx.Param("id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "ID is required",
		})
    return
	}

	resident := schemas.Resident{}

	if err := db.First(&resident, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
      "error": "Resident not found",
      "details": err.Error(),
    })
    return
	}

	if err := db.Delete(&resident).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to delete resident",
      "details": err.Error(),
    })
    return
	}

	ctx.JSON(http.StatusOK, http.StatusOK)
}