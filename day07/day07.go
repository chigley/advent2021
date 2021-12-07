package day07

import (
	"math"

	"github.com/chigley/advent2021"
)

func Part1(crabs []int) int {
	return minimiseCost(crabs, false)
}

func Part2(crabs []int) int {
	return minimiseCost(crabs, true)
}

func minimiseCost(crabs []int, isBinomial bool) int {
	minCost := math.MaxInt32
	for i := advent2021.MinInts(crabs); i <= advent2021.MaxInts(crabs); i++ {
		minCost = advent2021.Min(minCost, cost(crabs, i, isBinomial))
	}
	return minCost
}

func cost(crabs []int, target int, isBinomial bool) int {
	var cost int

	for _, pos := range crabs {
		steps := advent2021.Abs(pos - target)

		if isBinomial {
			cost += steps * (steps + 1) / 2
		} else {
			cost += steps
		}
	}

	return cost
}
