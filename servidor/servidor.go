package servidor

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
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

//Inicializar estabelece um novo Roteador e rotas a serem utilizadas
func (s *Servidor) Inicializar(endereco string) {

	s.Roteador = mux.NewRouter()
	s.Rotas()

	srv := &http.Server{
		Handler:      s.Roteador,
		Addr:         os.Getenv("BD_HOST") + endereco,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Printf("Servindo na porta: %s", endereco)
	log.Fatal(srv.ListenAndServe())
}

//Rotas é uma engloba todas os URLs e HandleFuncs a serem utilizados
func (s *Servidor) Rotas() {

	usuarioCtl := controllers.ControllerUser{}
	bd := driver.ConnectarBD()

	s.Roteador.HandleFunc("/", controllers.Home).Methods("GET")
	s.Roteador.HandleFunc("/usuarios", usuarioCtl.ChamarUsuarios(bd)).Methods("GET")
}
