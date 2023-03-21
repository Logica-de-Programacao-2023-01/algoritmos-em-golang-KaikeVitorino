package main

import "fmt"

func main() {
	var x int
	var y int
	fmt.Print("Quantos dias vc trabalha?")
	fmt.Scan(&x)
	fmt.Print("Qual seu salario mensal?")
	fmt.Scan(&y)
	fmt.Println("Valor do seu dia trabalho:", y/x)
}
