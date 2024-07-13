package routes

import (
	handler "entry-exit-api/handler/resident"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		v1.GET("/resident/:cpf", handler.GetResidentEntrysAndExits)
		v1.POST("/resident-entry", handler.ResidentEntry)	
		v1.PUT("/resident-exit/:cpf", handler.ResidentExit)
	}
}
