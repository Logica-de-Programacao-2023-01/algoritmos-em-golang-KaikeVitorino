package main

import "fmt"

func main() {
	var peso float64
	var altura float64
	var sexo float64
	var IMC float64
	fmt.Println("Digite seu peso, altura e sexo: ")
	fmt.Scan(&peso, &altura, &sexo)
	IMC = peso / (altura * altura)
	if IMC > 18.49 && IMC < 25.5 {
		fmt.Print("vc esta no peso ideal")
	} else if IMC < 18.50 {
		fmt.Print("Vc esta abaixo o peso")
	} else {
		fmt.Print("Vc esta acima do peso")
	}
}
