package main

import "fmt"

func main() {
	var y int
	fmt.Print("Digite numero a ser tabuado: ")
	fmt.Scan(&y)
	x := 0
	for x < 10 {
		x++
		fmt.Printf("Tabuada do numero digitado: %d*%d=%d\n", y, x, y*x)
	}
}
