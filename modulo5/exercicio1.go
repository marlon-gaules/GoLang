package kata

func MakeNegative(x int) int {

	//Nesta tarefa simples, você recebe um número e tem que torná-lo negativo. Mas talvez o número já seja negativo?

	if x <= 0 {
		return x
	}
	return x * -1
}
