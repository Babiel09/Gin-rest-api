package main

import (
	"github.com/Babiel09/Gin-rest-api/database"
	"github.com/Babiel09/Gin-rest-api/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
