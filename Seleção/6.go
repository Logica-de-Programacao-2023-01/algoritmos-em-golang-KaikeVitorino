package main

import "fmt"

func main() {
	var x int64
	var y int64
	fmt.Print("Digite os 2  inteiros: ")
	fmt.Scan(&x, &y)
	if x >= 0 && y >= 0 {
		fmt.Println("Dois numeros positivos, ent, aq esta a multiplicacao deles: ", x*y)
	} else {
		fmt.Println("Ja q um dos dois nn eh positivo, toma a soma deles: ", x+y)
	}
}
