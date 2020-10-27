package main

import (
	_ "github.com/lib/pq"
	"github.com/valentergs/valentedev/servidor"
)

func main() {

	servidor := servidor.Servidor{}
	servidor.Inicializar()
	servidor.Rodar("8080")

}
