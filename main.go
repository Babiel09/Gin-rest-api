package main

import (
	"github.com/Babiel09/Gin-rest-api/models"
	"github.com/Babiel09/Gin-rest-api/routes"
)

func main() {
	models.Amigos = []models.Amigo{
		{Id: 1, Nome: "Vitor", Divertido: true},
		{Id: 2, Nome: "Clarice", Divertido: true},
	}
	routes.HandleRequests()
}
