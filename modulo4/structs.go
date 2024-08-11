package main

import "fmt"

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
	AnoNasc   int
}

func main() {
	fmt.Println(Pessoa{"Marlon", "Sousa", 33, 1991})
	fmt.Println(Pessoa{Nome: "Marlon", Sobrenome: "Sousa", Idade: 33, AnoNasc: 1991})
}
