package kata

func CountPositivesSumNegatives(numbers []int) []int {

	// Dado um array de inteiros.
	// Retorna uma matriz, onde o primeiro elemento é a contagem de números positivos e o segundo elemento é a soma de números negativos.
	// 0 não é positivo nem negativo.
	// Se a entrada for uma matriz vazia ou for nula, retorne uma matriz vazia.

	var countPositives int
	var sumNegatives int

	for _, number := range numbers {
		if number > 0 {
			countPositives += 1
		}
		if number < 0 {
			sumNegatives += number
		}
	}
	return []int{countPositives, sumNegatives}
}
