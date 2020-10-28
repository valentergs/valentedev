package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/valentergs/valentedev/servidor"
)

func main() {
	servidor := servidor.Servidor{}
	appPort := fmt.Sprintf(":%s", os.Getenv("APP_PORTA"))
	servidor.Inicializar(appPort)
}
