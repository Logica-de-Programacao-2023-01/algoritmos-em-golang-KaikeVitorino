package main

import "fmt"

func main() {
	numeros := []string{"1", "3", "5", "7", "9", "11", "13", "15", "17", "19"}
	for index, numero := range numeros {
		fmt.Printf("Numero impar %d: %s\n", index, numero)
	}
}
