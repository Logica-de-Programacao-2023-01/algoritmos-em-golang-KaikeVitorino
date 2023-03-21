package main

import "fmt"

func main() {
	var x int
		fmt.Print("Digite um numero de 1 a 7: ")
		fmt.Scan(&x)
		if x <= 0 {
			fmt.Println("Numero invalido, tente novamente...")
		} else if x >= 8 {
			fmt.Println("Numero invalido, tente novamente...")
		} else if x == 1 {
			fmt.Println("Numero digitado equivale ao primeiro dia da semana, Domingo")
		} else if x == 2 {
			fmt.Println("Numero digitado equivale ao segundo dia da semana, Segunda-Feira")
		} else if x == 3 {
			fmt.Println("Numero digitado equivale ao terceiro dia da semana, Terça-feira")
		} else if x == 4 {
			fmt.Println("Numero digitado equivale ao quarto dia da semana, Quarta-feira")
		} else if x == 5 {
			fmt.Println("Numero digitado equivale ao quinta dia da semana, Quinta-feira")
		} else if x == 6 {
			fmt.Println("Numero digitado equivale ao sexto dia da semana, Sexta-feira")
		} else {
			fmt.Println("Numero digitado equivale ao setimo e ultimo dia da semana, Sabado")
		}
}
