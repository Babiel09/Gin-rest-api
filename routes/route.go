package routes

import (
	"github.com/Babiel09/Gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	request := gin.Default()
	request.GET("/amigos", controllers.GetAmigos)
	request.GET("/amigos/:id", controllers.GetperIdAmigos)
	request.POST("/amigos", controllers.PostAmigos)
	request.GET("/:nome", controllers.Saudacao)
	request.Run()
}
