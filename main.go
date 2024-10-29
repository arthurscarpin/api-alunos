package main

import (
	"github.com/arthurscarpin/gin-api-rest/database"
	"github.com/arthurscarpin/gin-api-rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
