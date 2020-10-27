package driver

import (
	"database/sql"
	"fmt"
)

//ConnectarBD conecta o programa ao Banco de Dados
func ConnectarBD() *sql.DB {

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
