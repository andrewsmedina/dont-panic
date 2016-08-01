package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		if i == 7 {
			panic("algo errado que nao esta certo =(")
		}
		fmt.Println(i)
	}
}
