package controllers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	"github.com/valentergs/valentedev/models"
)

//ControllerUser inicializa um struct
type ControllerUser struct{}

// Home é a função que será executada na Rota "/"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem-vindo à Home Page!")
}

//ChamarUsuarios faz uma Query no Banco de Dados
func (c ControllerUser) ChamarUsuarios(bd *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		// É feita o Query para trazer todos os registros da tabela usuários
		sqlQuery := `SELECT user_id, first_name, last_name, email, profile FROM users;`

		rows, _ := bd.Query(sqlQuery)
		defer rows.Close()

		//Criamos uma slice vazia do tipo models.Usuario que será vinculada a uma variavel chamada resultados.
		resultados := make([]models.Usuario, 0)
		for rows.Next() {
			resultado := models.Usuario{}
			err := rows.Scan(&resultado.ID, &resultado.FirstName, &resultado.LastName, &resultado.Email, &resultado.Profile)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			resultados = append(resultados, resultado)
		}

		//Com html/template injetamos dinamicamente a variavel "resultados" dentro do HTML templates/index.html
		// var tpl *template.Template
		// tpl = template.Must(template.ParseGlob("./templates/*.html"))
		// err := tpl.ExecuteTemplate(w, "index.html", resultados)
		// if err != nil {
		// 	log.Fatal(err)
		// }

		t := template.Must(template.ParseFiles("templates/index.html"))
		err := t.ExecuteTemplate(w, "index.html", resultados)
		if err != nil {
			panic(err)
		}

	}

}
