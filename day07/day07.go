package day07

import (
	"math"

	"github.com/chigley/advent2021"
)

func Part1(crabs []int) int {
	minCost := math.MaxInt32
	for i := advent2021.MinInts(crabs); i <= advent2021.MaxInts(crabs); i++ {
		minCost = advent2021.Min(minCost, cost(crabs, i))
	}
	return minCost
}

func cost(crabs []int, target int) int {
	var cost int
	for _, pos := range crabs {
		cost += advent2021.Abs(pos - target)
	}
	return cost
}
