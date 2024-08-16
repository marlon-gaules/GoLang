package main

import (
	"fmt"
)

func main() {

	sum := 0

	for i := 0; i < 10; i++ {
		fmt.Println("Sum:", sum)
		fmt.Println("Indice:", i)
		sum += i
	}
	fmt.Println(sum)

	//	for {
	//	fmt.Println("Golang do zero")
	//time.Sleep(2 * time.Second)
	//}

	frutas := []string{"laranja", "maça", "banana", "uva", "kiwi"}
	for i, fruta := range frutas {
		fmt.Println("Frutas", fruta, "Indice", i)
	}
}
