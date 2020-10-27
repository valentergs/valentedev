package driver

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load(); err != nil {
		fmt.Println("arquivo .env náo encontrado")
	}
}

//ConnectarBD conecta o programa ao Banco de Dados
func ConnectarBD() *sql.DB {

	var err error

	// const (
	// 	user     = os.Getenv("BD_USUARIO")
	// 	password = "Gustavo2012"
	// 	// Quando rodar em Docker o HOST precisa ter o mesmo nome do container onde roda o Postgresql - nesse caso vai ficar postgres:5432
	// 	host   = "localhost"
	// 	port   = 5432
	// 	dbname = "valentedevBlog"
	// )

	psqlInfo := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable", os.Getenv("BD_USUARIO"), os.Getenv("BD_SENHA"), os.Getenv("BD_HOST"), os.Getenv("BD_PORTA"), os.Getenv("BD_NOME"))
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
