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
	fmt.Println(Pessoa{Nome: "Marlon"})

	p1 := Pessoa{Sobrenome: "Araujo", AnoNasc: 2023}
	fmt.Println(p1.Sobrenome)
	fmt.Println(p1.AnoNasc)

	p1.Sobrenome = "Sousa"
	fmt.Println(p1.Sobrenome)

	p2 := Pessoa{Sobrenome: "Costa", AnoNasc: 2020}

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)
	fmt.Println(pessoas)
}
