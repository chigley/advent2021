package day06

import (
	"github.com/chigley/advent2021"
)

func Part1(in []int) int {
	return stepN(in, 80)
}

func Part2(in []int) int {
	return stepN(in, 256)
}

func stepN(fish []int, steps int) int {
	// freq maps from a timer value to the number of fish with that value
	var freq [9]int
	for _, n := range fish {
		freq[n]++
	}

	for i := 0; i < steps; i++ {
		var nextFreq [9]int

		for timerValue, numFish := range freq {
			if timerValue == 0 {
				nextFreq[6] += numFish
				nextFreq[8] += numFish
			} else {
				nextFreq[timerValue-1] += numFish
			}
		}

		freq = nextFreq
	}

	return advent2021.Sum(freq[:])
}
