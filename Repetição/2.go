package main

import "fmt"

func main() {
	numeros := []string{"2", "4", "6", "8", "10", "12", "14", "16", "18", "20"}
	for index, num := range numeros {
		fmt.Printf("Numero par %d: %s\n", index, num)
	}
}
