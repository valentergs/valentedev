package controllers

import (
	"database/sql"
	"fmt"
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

		sqlQuery := `SELECT user_id, email, profile FROM users;`

		rows, _ := bd.Query(sqlQuery)
		defer rows.Close()

		resultados := make([]models.Usuario, 0)
		for rows.Next() {
			resultado := models.Usuario{}
			err := rows.Scan(&resultado.ID, &resultado.Email, &resultado.Profile)
			if err != nil {
				http.Error(w, http.StatusText(500), 500)
				fmt.Println(err)
				return
			}
			resultados = append(resultados, resultado)
		}

		fmt.Fprint(w, resultados)

	}

}
