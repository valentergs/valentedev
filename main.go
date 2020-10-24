package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

//Servidor é um struct que contem todos os métodos como Banco de Dados, Mux, etc.
type Servidor struct {
	Roteador *mux.Router
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
	servidor.Rodar("8080")

}

// Home é a função que será executada na Rota "/"
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Bemvindo à Home Page!")
}
