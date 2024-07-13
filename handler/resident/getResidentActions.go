package handler

import (
	"entry-exit-api/database"
	"entry-exit-api/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)


func GetResidentEntrysAndExits(ctx *gin.Context) {
	db := database.GetMysql()
	cpf := ctx.Param("cpf")
	
	println("CPF DO RESIDENTE", cpf)

	var resident  []schemas.Resident

	if err := db.Where("cpf = ?", cpf).Find(&resident).Error; err!= nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get resident entrys and exits",
      "details": err.Error(),
		})
		return
	}

	if len(resident) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{
      "error": "Resident not found",
    })
    return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"residente": resident[0].Name,
		"data": resident,
	})
}