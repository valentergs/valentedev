package servidor

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/valentergs/valentedev/controllers"
	"github.com/valentergs/valentedev/driver"
)

//Servidor é um struct que contem todos os métodos como Banco de Dados, Mux, etc.
type Servidor struct {
	Roteador *mux.Router
	BD       *sql.DB
}

//Rotas é uma engloba todas os URLs e HandleFuncs a serem utilizados
func (s *Servidor) Rotas() {

	usuarioCtl := controllers.ControllerUser{}
	bd := driver.ConnectarBD()

	s.Roteador.HandleFunc("/", controllers.Home)
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
