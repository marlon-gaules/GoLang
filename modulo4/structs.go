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

	alunos := map[string][]Pessoa{}
	alunos["programação"] = pessoas
	fmt.Println(alunos)

	var alunos2 = map[string][]Pessoa{
		"Programação": {{Nome: "Marlon", Idade: 33}, {Nome: "Billy", Idade: 2}},
		"Engenharia":  {{Nome: "Marlindo", Idade: 19}, {Nome: "Biloso", Idade: 1}},
	}
	fmt.Println(alunos2)

	type Profissao struct {
		Pessoa
		Tipo string
	}

	prof := Profissao{p2, "dev"}
	fmt.Println(prof)
	fmt.Println(prof.Pessoa.Sobrenome)
	fmt.Println(prof.Pessoa.AnoNasc)
	fmt.Println(prof.Tipo)
}
