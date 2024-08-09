package main

import "fmt"

func main() {

	//Array
	fmt.Println("Arrays")
	var array [2]string
	array[0] = "Hello"
	array[1] = "World"
	fmt.Println(array[0], array[1])
	fmt.Println(array[1])
	fmt.Println(array[1])
	fmt.Println(array)

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(numPrimos)
	fmt.Println(numPrimos[2])
	fmt.Println(numPrimos[0], numPrimos[5])
	fmt.Println(numPrimos[0:3])
	fmt.Println(numPrimos[:4])
	fmt.Println("##########################")
	//Slice
	fmt.Println("Slices")
	slice := make([]string, 5)
	slice[0] = "Hello"
	slice[1] = "World"
	slice[4] = "Marlin"
	fmt.Println(slice)
	fmt.Println(slice[3])
	fmt.Println(slice[0], slice[1])
	fmt.Println(len(slice))

	numPares := []int{2, 4, 6, 8, 10, 12}
	fmt.Println(numPares)

	numPares = append(numPares, 14, 16, 18, 20)
	fmt.Println(numPares)
}
