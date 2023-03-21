package main

import "fmt"

func main() {
	var x int
	fmt.Print("Qual o numero a ser verificado: ")
	fmt.Scan(&x)
	if x%3 == 0 && x%5 == 0 {
		fmt.Println("Eh multiplo de 3 e 5 ao mesmos tempo")
	} else {
		fmt.Println("N eh multiplo dos dois")
	}
}
