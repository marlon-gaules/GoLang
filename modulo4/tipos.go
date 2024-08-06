package main

import "fmt"

func main() {
	fmt.Printf("Type: %T - Value: %v\n", true, false)
	fmt.Printf("Type: %T - Value: %v\n", "Marlon", "Sousa")
	fmt.Printf("Type: %T - Value: %v\n", 1, 1)
	fmt.Printf("Type: %T - Value: %v\n", "1", "1")
	fmt.Printf("Type: %T - Value: %v\n", 1.233, 1.255)
}

// Tipos:
// bool (true/false)
// string - sequência de bytes
// int
// float (float64/float32) - decimal
