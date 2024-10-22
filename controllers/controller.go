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
	c.JSON(201, amigoPost)
}

//Delete

func DeleteAmigos(c *gin.Context) {
	var amigoDelete models.Amigo
	id := c.Params.ByName("id")
	database.DB.Delete(&amigoDelete, id)
	c.JSON(200, gin.H{"user": "Amigo deletado"})
}

//Patch

func PatchAmigosNome(c *gin.Context) {
	var amigoPatch models.Amigo
	id := c.Params.ByName("id")
	//Pega as informações do DB
	database.DB.First(&amigoPatch, id)
	//Verifica se a informação existe
	if err := c.ShouldBindJSON(&amigoPatch); err != nil {
		c.JSON(500, gin.H{"err": err})
	}
	//Implmenta as atualizações caso a informação exista
	database.DB.Model(&amigoPatch).UpdateColumn(amigoPatch.Nome, amigoPatch)
	c.JSON(200, amigoPatch)
}

//Put

func PutAmigosNome(c *gin.Context) {
	var amigoPut models.Amigo
	id := c.Params.ByName("id")
	//Procura as informações no banco de dados
	database.DB.First(&amigoPut, id)
	//Caso não encontre
	if err := c.ShouldBindJSON(&amigoPut); err != nil {
		c.JSON(400, gin.H{"err": err})
	}
	//Caso ache as informações no banco de dados
	database.DB.Model(&amigoPut).Updates(amigoPut)
	c.JSON(200, amigoPut)
}

// Função para aprender sobre params
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(404, gin.H{
		"Mensagem de saudação": "Olá " + nome + "!" + " A rota que você deseja não existe.",
	})
}
