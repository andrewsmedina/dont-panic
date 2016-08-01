package main

import (
	"fmt"
	"os"
)

func main() {
	if os.Getenv("MENSAGEMDOALEM") == "" {
		panic("a mensagem do alem nao foi recebida")
	}
	fmt.Println("a mensagem do alem é: ", os.Getenv("MENSAGEMDOALEM"))
}
