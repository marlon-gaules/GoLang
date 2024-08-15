package main

import "fmt"

func main() {

	posicao := 2
	switch posicao {
	case 1:
		fmt.Println("Primeiro lugar")
	case 2:
		fmt.Println("Segundo lugar")
	case 3:
		fmt.Println("Terceiro lugar")
	}

	nome := "Bob"
	switch nome {
	case "Marlon":
		fmt.Println("É uma pessoa")
	case "Billy":
		fmt.Println("É um cachorro")
	case "Bob":
		fmt.Println("É um personagem")
	}
}
