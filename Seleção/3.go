package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite o numero: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Println("Numero eh par")
	} else {
		fmt.Println("Numero eh impar")
	}
}
