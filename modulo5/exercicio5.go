package kata

import "strconv"

func StringToNumber(str string) int {

	// Precisamos de uma função que possa transformar uma string em um número.
	// Quais maneiras de conseguir isso você conhece? Observação:
	// não se preocupe, todas as entradas serão strings e cada string é uma representação perfeitamente válida de um número integral.

	stringNumber, _ := strconv.Atoi(str)
	return stringNumber
}
