package routes

import (
	"github.com/gin-gonic/gin"
	"main.go/Controller"
	"main.go/middleware"
)
func UserRoutes(incomingRoutes *gin.Context){
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users",Controller.GetUsers())
	incomingRoutes.GET("/users/:user_id",Controller.GetUser())
}