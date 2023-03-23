package main

import "fmt"

func main() {
	numeros := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for index, numero := range numeros {
		fmt.Printf("O índice %d contém o número %s\n", index, numero)
	}
}
