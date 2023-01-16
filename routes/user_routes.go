package routes

import (
    "example/web-service-gin/controllers" 
    "github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine)  {
    router.POST("/user", controllers.CreateUser()) 
    router.GET("/users", controllers.GetAllUsers()) 
    router.GET("/user/:userId", controllers.GetAUser()) 
    router.PUT("/user/:userId", controllers.EditAUser())
    router.DELETE("/user/:userId", controllers.DeleteAUser())
}