package driver

import (
	"database/sql"
	"fmt"
	"os"
)

//ConnectarBD conecta o programa ao Banco de Dados
func ConnectarBD() *sql.DB {

	var err error

	BDInfo := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable", os.Getenv("BD_USUARIO"), os.Getenv("BD_SENHA"), os.Getenv("BD_HOST"), os.Getenv("BD_PORTA"), os.Getenv("BD_NOME"))
	db, err := sql.Open("postgres", BDInfo)
	if err != nil {
		panic(err)
	}

	return db
}
