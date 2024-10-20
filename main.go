package main

import (
	"github.com/gin-gonic/gin"
)

func ExibeAmigos(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   1,
		"nome": "Vitor",
	})
}

func main() {
	request := gin.Default()
	request.GET("/amigos", ExibeAmigos)
	request.Run(":8000")
}
