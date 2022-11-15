package arraysSlices

func Sum(numbers []int) int {
	var bekir int;
	for  _, value := range numbers {
		bekir += value
	}
	return bekir
}

func SumAll(numbersToSum ...[]int)[]int {
	var sums []int

	for _, value := range numbersToSum {
		sums = append(sums, Sum(value))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
			continue
		}

		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}