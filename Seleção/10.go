package main

import "fmt"

func main() {
	var x int
	fmt.Print("Digite sua idade: ")
	fmt.Scan(&x)
	if x <= 9 {
		fmt.Println("sua classificação: Mirim")
	} else if x <= 13 && x >= 10 {
		fmt.Println("sua classificação: Infantil")
	} else if x <= 17 && x >= 14 {
		fmt.Println("sua classificação: Juvenil")
	} else {
		fmt.Println("sua classificação: Adulto")
	}
}
