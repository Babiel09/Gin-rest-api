package controllers

import (
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

func ExibeAmigos(c *gin.Context) {
	c.JSON(200, models.Amigos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"Mensagem de saudação": "Olá " + nome + "!" + " A rota que você deseja não existe.",
	})
}
