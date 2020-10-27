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
	s.Roteador.HandleFunc("/", Home)
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
	servidor.ChamarUsuarios()
	servidor.Rodar("8080")

}

// Home é a função que será executada na Rota "/"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bem-vindo à Home Page!")
}

//ChamarUsuarios faz uma Query no Banco de Dados
func (s *Servidor) ChamarUsuarios() *sql.Rows {
	var (
		email   string
		profile string
	)
	rows, err := s.ConnectarBD().Query("SELECT email, profile FROM users;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&email, &profile)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(email, profile)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return rows
}
