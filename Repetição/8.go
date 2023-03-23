package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite um numero inteiro positivo a ser mostrado seus divisores: ")
	fmt.Scan(&x)
	for y := 1; y <= x; y++ {
		if x%y == 0 {
			fmt.Printf("Divisor de %d: %d\n", x, y)
		}
	}
}
