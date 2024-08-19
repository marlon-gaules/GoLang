package kata

func PositiveSum(numbers []int) int {

	// Você obtém uma matriz de números e retorna a soma de todos os positivos.
	// Exemplo `[1,-4,7,12]`=>`1 + 7 + 12 = 20`
	// Observação: se não houver nada para somar, a soma padrão será `0`.

	var sum int
	for _, number := range numbers {
		if number > 0 {
			sum += number
		}
	}
	return sum
}
