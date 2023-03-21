package main

import "fmt"

func main() {
	var x float64
	fmt.Print("Qual seu salario: ")
	fmt.Scan(&x)
	if x < 1000 {
		fmt.Println("Vc tera um aumento de 10%, seu novo salario sera: ", x*1.1)
	} else {
		fmt.Println("Vc tera um aumento de 5%, seu novo salrio sera: ", x*1.05)
	}
}
