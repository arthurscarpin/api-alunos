package main

import (
	"github.com/arthurscarpin/api-go-gin/database"
	"github.com/arthurscarpin/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
