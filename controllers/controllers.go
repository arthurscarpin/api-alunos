package controllers

import (
	"net/http"

	"github.com/arthurscarpin/gin-api-rest/database"
	"github.com/arthurscarpin/gin-api-rest/models"
	"github.com/gin-gonic/gin"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// ExibeTodosAlunos godoc
// @Summary Exibe todos os alunos
// @Description Rota para listar todos os alunos cadastrados
// @Tags alunos
// @Accept json
// @Produce json
// @Success 200 {array} models.Aluno
// @Router /alunos [get]
func ExibeTodosAlunos(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

// Saudacao godoc
// @Summary Exibe uma saudação personalizada
// @Description Rota para exibir uma saudação com o nome fornecido
// @Tags saudacao
// @Accept json
// @Produce json
// @Param nome path string true "Nome do Aluno"
// @Success 200 {object} map[string]string
// @Router /saudacao/{nome} [get]
func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo beleza?",
	})
}

// CriaAluno godoc
// @Summary Cria um aluno
// @Description Rota para cadastrar um novo aluno
// @Tags alunos
// @Accept json
// @Produce json
// @Param aluno body models.Aluno true "Dados do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} map[string]string
// @Router /alunos [post]
func CriaAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidadDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorId godoc
// @Summary Busca aluno por ID
// @Description Rota para buscar um aluno específico pelo seu ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} map[string]string
// @Router /alunos/{id} [get]
func BuscaAlunoPorId(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

// DeletaAluno godoc
// @Summary Deleta aluno
// @Description Rota para deletar um aluno pelo seu ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do Aluno"
// @Success 200 {object} map[string]string
// @Router /alunos/{id} [delete]
func DeletaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.Delete(&aluno, id)
	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})
}

// EditaAluno godoc
// @Summary Edita aluno
// @Description Rota para editar as informações de um aluno pelo seu ID
// @Tags alunos
// @Accept json
// @Produce json
// @Param id path string true "ID do Aluno"
// @Param aluno body models.Aluno true "Dados do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 400 {object} map[string]string
// @Router /alunos/{id} [put]
func EditaAluno(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	if err := models.ValidadDadosDeAluno(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

// BuscaAlunoPorCPF godoc
// @Summary Busca aluno por CPF
// @Description Rota para buscar um aluno específico pelo seu CPF
// @Tags alunos
// @Accept json
// @Produce json
// @Param cpf path string true "CPF do Aluno"
// @Success 200 {object} models.Aluno
// @Failure 404 {object} map[string]string
// @Router /alunos/cpf/{cpf} [get]
func BuscaAlunoPorCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado"})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func ExibePaginaIndex(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)

	c.HTML(http.StatusOK, "index.html", gin.H{
		"alunos": alunos,
	})
}

func RotaNaoEncontrada(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404.html", nil)
}
