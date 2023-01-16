package main

import (
    "example/web-service-gin/configs"
    "example/web-service-gin/routes"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    //run database
    configs.ConnectDB()

    //routes
    routes.UserRoute(router)

    router.Run("localhost:8080")
}