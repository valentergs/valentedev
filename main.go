package main

import (
	_ "github.com/lib/pq"
	"github.com/valentergs/valentedev/servidor"
)

func main() {
	servidor := servidor.Servidor{}
	servidor.Inicializar("8080")
}
