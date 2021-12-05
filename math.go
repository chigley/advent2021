package advent2021

func Sign(x int) int {
	if x < 0 {
		return -1
	}
	if x > 0 {
		return 1
	}
	return 0
}

func Sum(ns []int) int {
	var sum int

	for _, n := range ns {
		sum += n
	}

	return sum
}
