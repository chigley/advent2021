package day01

func Part1(measurements []int) int {
	var increaseCount int

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increaseCount++
		}
	}

	return increaseCount
}
