package day01

import (
	"github.com/chigley/advent2021"
)

func Part1(measurements []int) int {
	var increaseCount int

	for i := 1; i < len(measurements); i++ {
		if measurements[i] > measurements[i-1] {
			increaseCount++
		}
	}

	return increaseCount
}

func Part2(measurements []int) int {
	var increaseCount int

	for i := 0; i < len(measurements)-3; i++ {
		w1 := measurements[i : i+3]
		w2 := measurements[i+1 : i+4]

		if advent2021.Sum(w2) > advent2021.Sum(w1) {
			increaseCount++
		}
	}

	return increaseCount
}
