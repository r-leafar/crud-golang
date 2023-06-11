package model

import (
	"log"

	"github.com/r-leafar/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscarTodosProdutos() []Produto {
	produtos := []Produto{}
	p := Produto{}

	var conn = db.ConnDb()
	rows, err := conn.Query("SELECT * FROM produto order by id")

	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = rows.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade

		produtos = append(produtos, p)

	}
	defer conn.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConnDb()

	insereDadosNoBanco, err := db.Prepare("INSERT INTO produto (nome,descricao,preco,quantidade)values($1,$2,$3,$4)")

	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)

	defer db.Close()
}

func DeletaProduto(idproduto string) {
	db := db.ConnDb()

	deletarProduto, err := db.Prepare("DELETE FROM produto WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	deletarProduto.Exec(idproduto)

	defer db.Close()
}

func GetProduto(idproduto string) Produto {
	db := db.ConnDb()
	p := Produto{}

	prod, err := db.Prepare("SELECT * FROM produto WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}
	row, err := prod.Query(idproduto)

	if err != nil {
		panic(err.Error())
	}
	for row.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = row.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
	}
	defer db.Close()
	return p
}
func AtualizaProduto(id, nome, descricao string, preco float64, quantidade int) {
	db := db.ConnDb()

	updateProduto, err := db.Prepare("UPDATE produto SET nome=$1,descricao=$2,preco=$3,quantidade=$4 WHERE id=$5")

	if err != nil {
		panic(err.Error())
	}
	updateProduto.Exec(nome, descricao, preco, quantidade, id)

	defer db.Close()
}
