package visitor

import (
	"entry-exit-api/database"
	"entry-exit-api/schemas"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type visitorRequest struct {
	Name          string `json:"name"`
	Sexo          string `json:"sexo"`
	Cpf           string `json:"cpf"`
	HouseToVisit  int64  `json:"house_to_visit_number"`
	WhoAllowedCpf string `json:"who_allowed_cpf"`
}

func VisitorEntry(ctx *gin.Context) {
	db := database.GetMysql()

	request := visitorRequest{}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request body",
		})
		return
	}

	resident := schemas.Resident{}
	houseToVisit := request.HouseToVisit

	println("House to visit number", request.HouseToVisit)

	if err := db.Where("house_number = ?", houseToVisit).First(&resident).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error":   "House not found",
			"details": err.Error(),
		})
		return
	}

	visitor := schemas.Visitor{
		Name:         request.Name,
		Sexo:         request.Sexo,
		Cpf:          request.Cpf,
		HouseToVisit: request.HouseToVisit,
		WhoAllowed:   request.WhoAllowedCpf,
		EntryTime:    time.Now(),
		ExitTime:     nil,
	}

	if err := db.Create(&visitor).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create visitor",
			"details": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "Visitor entry successful",
		"data":    visitor,
	})
}
