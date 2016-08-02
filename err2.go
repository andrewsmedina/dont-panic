package main

import "fmt"

type errorString struct{ s string }

func (e *errorString) Error() string { return e.s }

func soma(a, b int) (int, error) {
	if a > 10 {
		return 0, &errorString{"nao e possivel soma de numeros maiores que 10"}
	}
	return a + b, nil
}

func main() {
	result, err := soma(11, 1)
	if err != nil {
		fmt.Printf("Erro ao somar 11 + 1: %v\n", err)
	} else {
		fmt.Printf("A soma de 11 + 1 é: %d\n", result)
	}
}
