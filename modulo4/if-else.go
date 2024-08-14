package main

import "fmt"

func main() {

	valor := 4
	if valor == 1 {
		fmt.Println("Valor é igual a 1")
	} else {
		fmt.Println("Valor não é igual a 1")
	}

	numero := 9
	if numero == 1 {
		fmt.Println("Numero é igual a 1")
	} else if numero == 5 {
		fmt.Println("Numero é igual a 5")
	} else {
		fmt.Println("Valor é diferente de 1 e 5")
	}

	fmt.Println(1 + 1)
	fmt.Println(3 - 1)
	fmt.Println(3 * 4)
	fmt.Println(1 / 1)
	fmt.Println(10 % 2)

	number := 8
	if number%2 == 0 {
		fmt.Printf("%d é par", number)
	} else {
		fmt.Printf("%d é impar", number)
	}
}
