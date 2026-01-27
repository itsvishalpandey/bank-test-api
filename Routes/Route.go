package routes

import (
	controller "bank-test-api/Controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {

	router := gin.Default()

	router.GET("/ifsc", controller.GetIfscCode)
	router.POST("/ifsc", controller.PostIFSC)

	return router
}
