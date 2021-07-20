package main

// godos.org
import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectarBD()
	defer db.Close()

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)

	// teste

}

func index(w http.ResponseWriter, r *http.Request) {
	p := Produto{}
	produtos := []Produto{}
	db := conectarBD()
	selProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	for selProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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

	temp.ExecuteTemplate(w, "index", produtos)
	defer db.Close()
}

// devolve um ponteiro do banco
func conectarBD() *sql.DB {
	conexao := "user=postgres dbname=teste password=123 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}

	return db
}
