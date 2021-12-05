package advent2021

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Sum(ns []int) int {
	var sum int

	for _, n := range ns {
		sum += n
	}

	return sum
}
