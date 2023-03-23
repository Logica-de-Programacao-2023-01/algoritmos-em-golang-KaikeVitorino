package main

import "fmt"

func main() {
	for x := 0; x <= 30; x++ {
		if x%3 == 0 {
			fmt.Printf("Numero multiplo de 30: %d\n", x)
		}
	}
}
