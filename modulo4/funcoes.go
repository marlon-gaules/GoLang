package main

import "fmt"

func main() {
	somaDosValores := soma(42, 13)
	sub := subtracao(10, 5)
	funcPrintaNome := "Marlon"
	funcPrintaNome2 := "Sousa"
	funcPrintaNome3, sobrenome := "Billy", "Araujo"

	fmt.Println(somaDosValores)
	fmt.Println(sub)
	fmt.Println(funcPrintaNome)
	fmt.Println(funcPrintaNome2)
	fmt.Println(funcPrintaNome3, sobrenome)

	nome1, _ := printaNome4("Marlindo")
	fmt.Println(nome1)
	//	fmt.Println(nome2)

	nome3, nome4, _ := printaNome5("Armando", "Góix")
	fmt.Println(nome3)
	fmt.Println(nome4)
}

func soma(x int, y int) int {
	return x + y
}

func subtracao(x int, y int) int {
	return x - y
}

func printaNome(nome string) string {
	return nome
}

func printaNome2(nome string) (string, string) {
	return nome, nome
}

func printaNome3(nome string) (string, string) {
	return nome, nome
}

func printaNome4(nome string) (string, string) {
	return nome, nome
}

func printaNome5(nome, sobreNome string) (string, string, string) {
	return nome, sobreNome, nome
}
