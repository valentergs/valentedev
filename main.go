package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//Servidor é um struct que contem todos os métodos como Banco de Dados, Mux, etc.
type Servidor struct {
	Roteador *mux.Router
	BD       *sql.DB
}

//ConnectarBD conecta o programa ao Banco de Dados
func (s *Servidor) ConnectarBD() *sql.DB {

	var err error

	const (
		user     = "rodrigovalente"
		password = "Gustavo2012"
		// Quando rodar em Docker o HOST precisa ter o mesmo nome do container onde roda o Postgresql - nesse caso vai ficar postgres:5432
		host   = "localhost"
		port   = 5432
		dbname = "valentedevBlog"
	)

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%d dbname=%s sslmode=disable", user, password, host, port, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}

//Rotas é uma engloba todas os URLs e HandleFuncs a serem utilizados
func (s *Servidor) Rotas() {

	usuarioCtl := ControllerUser{}
	bd := s.ConnectarBD()

	s.Roteador.HandleFunc("/", Home)
	s.Roteador.HandleFunc("/usuarios", usuarioCtl.ChamarUsuarios(bd)).Methods("GET")
}

//Inicializar estabelece um novo Roteador e rotas a serem utilizadas
func (s *Servidor) Inicializar() {
	s.Roteador = mux.NewRouter()
	s.Rotas()
}

//Rodar inicia o servidor e aplica configurações de Timeout
func (s *Servidor) Rodar(endereco string) {

	srv := &http.Server{
		Handler:      s.Roteador,
		Addr:         "127.0.0.1:" + endereco,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Servindo na porta: %s", endereco)
	log.Fatal(srv.ListenAndServe())

}

func main() {

	servidor := Servidor{}
	servidor.Inicializar()
	servidor.Rodar("8080")

}

// Home é a função que será executada na Rota "/"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem-vindo à Home Page!")
}

type ControllerUser struct{}

type Usuario struct {
	ID      int
	Email   string
	Profile string
}

//ChamarUsuarios faz uma Query no Banco de Dados
func (c ControllerUser) ChamarUsuarios(bd *sql.DB) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		sqlQuery := `SELECT user_id, email, profile FROM users;`

		rows, _ := bd.Query(sqlQuery)
		defer rows.Close()

		resultados := make([]Usuario, 0)
		for rows.Next() {
			resultado := Usuario{}
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
