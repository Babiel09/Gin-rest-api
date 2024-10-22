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
	request.DELETE("/amigos/:id", controllers.DeleteAmigos)
	request.PATCH("/amigos/:id", controllers.PatchAmigosNome)
	request.PUT("/amigos/:id", controllers.PutAmigosNome)
	request.GET("/:nome", controllers.Saudacao)
	request.Run()
}
