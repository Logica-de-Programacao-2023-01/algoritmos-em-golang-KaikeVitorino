package main

import "fmt"

func main() {
	for x := 0; x <= 100; x++ {
		if x%3 != 0 && x%5 != 0 {
			fmt.Printf("Contando ate 100: %d\n", x)
		} else if x%3 == 0 && x%5 != 0 {
			fmt.Printf("Multiplo de 3: Fizz\n")
		} else if x%5 == 0 && x%3 != 0 {
			fmt.Printf("Multiplo de 5: Buzz\n")
		} else if x%3 == 0 && x%5 == 0 {
			fmt.Printf("Multiplo de 3 e 5: FizzBuzz\n")
		}
	}
}
