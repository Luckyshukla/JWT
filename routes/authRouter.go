package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/Controller"
)
func AuthRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.POST("user/Signup",Controller.Signup())
	incomingRoutes.POST("user/login",Controller.Login())
}