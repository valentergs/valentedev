package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	//Abrir o arquivo a ser trabalhado
	fileName := "./templates/index.html"

	//Ler o arquivo e armazenar na variável "content"
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}

	//Criar o arquivo final, que neste caso se chama "generated.go"
	out, _ := os.Create("./controllers/generated.go")

	//Usar a func os.Write para compor o arquivo de transformação de STRING para BINARIO.
	out.Write([]byte("package controllers\n\n"))
	out.Write([]byte("var EmbedFile = map[string][]byte{\n"))
	out.Write([]byte("\"" + fileName + "\": []byte(`"))
	out.Write(content)
	out.Write([]byte("`),\n}"))

	//Mandar para o stdout mensagem e nome do arquivo criado.
	fmt.Println("Generated: ", fileName)
}
