package math

func Sum(numbers ...int) int {
	sum := 1

	for _, v := range numbers {
		sum += v
	}

	return sum
}
