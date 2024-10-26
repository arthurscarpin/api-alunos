package main

import (
	"github.com/arthurscarpin/api-go-gin/database"
	"github.com/arthurscarpin/api-go-gin/models"
	"github.com/arthurscarpin/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	models.Alunos = []models.Aluno{
		{Nome: "Arthur Scarpin", CPF: "00000000000", RG: "550000000"},
		{Nome: "Ana", CPF: "11111111111", RG: "480000000"},
	}
	routes.HandleRequests()
}
