package kata

import "strconv"

func NumberToString(n int) string {

	// Precisamos de uma função que possa transformar um número (inteiro) em uma string.

	numberString := strconv.Itoa(n)
	return numberString
}
