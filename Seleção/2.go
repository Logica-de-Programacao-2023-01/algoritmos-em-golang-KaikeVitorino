package main

import "fmt"

func main() {
	var x int
	var y int
	var z int
	fmt.Print("Informe três números inteiros ")
	fmt.Scan(&x, &y, &z)

	if x < y && x < z {
		fmt.Println(x, "é o menor entre eles")
	} else if y < x && y < z {
		fmt.Println(y, "é o menor entre eles")
	} else if z < y && z < x {
		fmt.Println(z, "é o menor entre eles")
	} else if x == y && x == z {
		fmt.Println("Todos são iguais")
	}
}
