package routes

import (
	"github.com/Babiel09/Gin-rest-api/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	request := gin.Default()
	request.GET("/amigos", controllers.ExibeAmigos)
	request.GET("/:nome", controllers.Saudacao)
	request.Run(":8000")
}
