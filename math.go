package advent2021

func Sum(ns []int) int {
	var sum int

	for _, n := range ns {
		sum += n
	}

	return sum
}
