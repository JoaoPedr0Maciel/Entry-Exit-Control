package routes

import (
	residentHandler "entry-exit-api/handler/resident"
	visitorHandler "entry-exit-api/handler/visitor"

	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	{
		// Residents routes
		v1.GET("/resident/:cpf", residentHandler.GetResidentEntrysAndExits)
		v1.POST("/resident-entry", residentHandler.ResidentEntry)	
		v1.PUT("/resident-exit/:cpf", residentHandler.ResidentExit)
		v1.DELETE("/resident/:id", residentHandler.DeleteResidentAction)

		// Visitors routes
		v1.POST("/visitor-entry", visitorHandler.VisitorEntry)
	}
}
