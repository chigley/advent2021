package advent2021

import "math"

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

func MinInts(ns []int) int {
	if len(ns) == 0 {
		panic("min of empty slice")
	}

	min := math.MaxInt
	for _, n := range ns {
		min = Min(min, n)
	}
	return min
}

func MaxInts(ns []int) int {
	if len(ns) == 0 {
		panic("max of empty slice")
	}

	max := math.MinInt
	for _, n := range ns {
		max = Max(max, n)
	}
	return max
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

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
