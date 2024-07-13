package handler

import (
	"entry-exit-api/database"
	"entry-exit-api/schemas"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type residentEntryRequest struct {
	Name        string    `json:"name"`
	Sexo        string    `json:"sexo"`
	Cpf         string    `json:"cpf"`
	HouseNumber int64     `json:"house_number"`
}

func ResidentEntry(ctx *gin.Context) {
	db := database.GetMysql()
	request := residentEntryRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	residentEntry := schemas.Resident{
		Name:        request.Name,
    Sexo:        request.Sexo,
    Cpf:         request.Cpf,
    HouseNumber: request.HouseNumber,
    EntryTime:   time.Now(),
    ExitTime:    nil,
	}

	if err := db.Create(&residentEntry).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
      "error": "Failed to create resident entry",
			"details": err.Error(),
    })
    return
	}

	ctx.JSON(http.StatusOK, gin.H{
    "message": "Resident entry created successfully",
    "data":    residentEntry,
  })
}