package main

import "fmt"

func main() {
	var y int
	var maior int
	fmt.Println("Digite vários números inteiros maiores que 0: ")
	for {
		fmt.Scan(&y)
		if y == 0 {
			break
		}
		if y > maior {
			maior = y
		}
	}
	if maior == 0 {
		fmt.Printf("Nenhum numero valido foi digitado")
	} else {
		fmt.Printf("O maior numero digitado foi: %d\n", maior)
	}

}
