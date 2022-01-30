package routes

import (
	"cloudAuthentication/controller"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {

	router := gin.Default()
	router.POST("/signUp/", controller.SignUp)             //POST
	router.POST("/validateUser/", controller.ValidateUser) //get
	return router
}
