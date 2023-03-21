package main

import "fmt"

func main() {
	var P float64
	var A float64
	fmt.Print("Qual seu peso?")
	fmt.Scan(&P)
	fmt.Print("Qual sua altura?")
	fmt.Scan(&A)
	fmt.Println("Seu IMC:", P/(A*A))
}
