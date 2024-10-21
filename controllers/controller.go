package controllers

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/gin-gonic/gin"
)

// Get
func GetAmigos(c *gin.Context) {
	var amigoGet []models.Amigo //PEga todos os alunos
	//Procura no banco de dados
	database.DB.Find(&amigoGet)
	c.JSON(200, amigoGet)
}

//Get per id

func GetperIdAmigos(c *gin.Context) {
	var amigo models.Amigo
	id := c.Params.ByName("id")
	//Busca primeiro
	database.DB.First(&amigo, id)
	//Depoism mostra o erro
	if amigo.ID == 0 {
		c.JSON(404, gin.H{
			"err": "Amigo não encontrado",
		})
		return
	}
	c.JSON(200, amigo)
}

// Post
func PostAmigos(c *gin.Context) {
	var amigoPost models.Amigo
	if err := c.ShouldBindJSON(&amigoPost); err != nil {
		c.JSON(400, gin.H{
			"error": err,
		})
		return
	}
	database.DB.Create(&amigoPost)
	c.JSON(200, amigoPost)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(404, gin.H{
		"Mensagem de saudação": "Olá " + nome + "!" + " A rota que você deseja não existe.",
	})
}
