package main

import "fmt"

func main() {
	idade := map[string]int{}
	idade["Marlon"] = 33
	idade["Billy"] = 1
	fmt.Println(idade)
	fmt.Println(idade["Marlon"])
	fmt.Println(idade["Billy"])

	anoNasc := map[string]int{
		"Marlon": 1991,
		"Billy":  2023,
	}
	fmt.Println(anoNasc)
	fmt.Println(anoNasc["Marlon"])
	fmt.Println(anoNasc["Billy"])
	anoNasc["golangDoZero"] = 2024
	fmt.Println(anoNasc)
}
