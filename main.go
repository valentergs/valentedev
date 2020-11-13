package main

//go:generate go run gen/generator.go

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/joho/godotenv"
	"github.com/valentergs/valentedev/servidor"
)

func init() {
	if err := godotenv.Load(); err != nil {
		godotenv.Load(".env_docker")
	}
}

func main() {
	servidor := servidor.Servidor{}
	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORTA"))
	servidor.Inicializar(appPort)
}
