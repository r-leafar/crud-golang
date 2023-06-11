package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/r-leafar/model"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	var produtos = model.BuscarTodosProdutos()
	templates.ExecuteTemplate(w, "Index", produtos)
}
func New(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		preco_float64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco", err)
		}
		quantidade_int, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		model.CriarNovoProduto(nome, descricao, preco_float64, quantidade_int)
	}

	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	model.DeletaProduto(idproduto)
	http.Redirect(w, r, "/", 301)
}
func MapTest(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "MapTest", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idproduto := r.URL.Query().Get("id")
	produto := model.GetProduto(idproduto)
	templates.ExecuteTemplate(w, "Edit", produto)
}
func Update(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		id := r.FormValue("id")
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		preco_float64, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro na conversão do preco", err)
		}
		quantidade_int, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro na conversão da quantidade", err)
		}

		model.AtualizaProduto(id, nome, descricao, preco_float64, quantidade_int)
		http.Redirect(w, r, "/", 301)
	}
}
