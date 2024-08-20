package kata

func CountPositivesSumNegatives(numbers []int) []int {
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
