package main

import (
	"fmt"
	"sync"
)

func soma(a, b int, wg sync.WaitGroup) {
	defer wg.Done()
	if a > 10 {
		panic("nao e possivel soma de numeros maiores que 10")
	}
	fmt.Println(a + b)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(4)
	go soma(10, 10, wg)
	go soma(1, 4, wg)
	go soma(11, 4, wg)
	go soma(2, 4, wg)
	wg.Wait()
}
