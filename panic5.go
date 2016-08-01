package main

import "fmt"

func soma(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("RECOVERED! ", err)
		}
	}()
	if a > 10 {
		panic("nao e possivel soma de numeros maiores que 10")
	}
	return a + b
}

func main() {
	fmt.Printf("A soma de 10 + 10 é: %d\n", soma(10, 10))
	fmt.Printf("A soma de 1 + 4 é: %d\n", soma(1, 4))
	fmt.Printf("A soma de 11 + 4 é: %d\n", soma(11, 4))
	fmt.Printf("A soma de 2 + 4 é: %d\n", soma(2, 4))
}
