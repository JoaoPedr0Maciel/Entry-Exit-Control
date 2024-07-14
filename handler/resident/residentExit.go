package handler

import (
	"entry-exit-api/database"
	"entry-exit-api/schemas"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func ResidentExit(ctx *gin.Context) {
	db := database.GetMysql()
	cpf := ctx.Param("cpf")

	resident := schemas.Resident{}

	if err := db.Where("cpf = ? AND exit_time IS NULL", cpf).First(&resident).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
      "error": "Resident not found",
      "details": err.Error(),
    })
    return
	}

	exitTime := time.Now()
	resident.ExitTime = &exitTime

	if err := db.Save(&resident).Error; err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to update resident exit time",
      "details": err.Error(),
    })
    return
	}

	ctx.JSON(http.StatusOK, gin.H{
    "message": "Resident exit time updated successfully",
    "data":    resident,
  })
}