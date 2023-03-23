package main

import "fmt"

func main() {
	var y, x, z int
	fmt.Println("Digite vários números inteiros (digite 0 para encerrar):")
	for {
		fmt.Scan(&y)
		if y == 0 {
			break
		}
		x += y
		z++
	}
	if z == 0 {
		fmt.Println("Nenhum numero foi digitado")
	} else {
		Media_Aritmetica := x / z
		fmt.Printf("A media aritmetica dos %d eh: %.2d\n", z, Media_Aritmetica)
	}
}
