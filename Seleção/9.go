package main

import "fmt"

func main() {
	var x float64
	var y float64
	var z float64
	fmt.Print("Digite 3 numeros Reais: ")
	fmt.Scan(&x, &y, &z)
	if x <= y && y < z {
		fmt.Println("Os numeros em ordem crescente: ", x, y, z)
	} else if z >= y && z < x {
		fmt.Println("Os numeros em ordem crescente: ", y, z, x)
	} else if x >= z && y < z {
		fmt.Println("Os numeros em ordem crescente: ", x, z, y)
	} else if x >= z && x < y {
		fmt.Println("Os numeros em ordem crescente: ", z, x, y)
	} else if x > z && y <= z {
		fmt.Println("Os numeros em ordem crescente: ", z, y, x)
	} else if x == z && y == z {
		fmt.Println("Os numeros tem o msm valor: ", z, y, x)
	}
}
