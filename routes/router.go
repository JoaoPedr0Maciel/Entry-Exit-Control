package routes

import "github.com/gin-gonic/gin"

func Init() {
	router := gin.Default()

	InitializeRoutes(router)

	router.Run()
}
