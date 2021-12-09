package day09

import (
	"sort"

	"github.com/chigley/advent2021"
)

func Part1(m Heightmap) int {
	var riskLevel int
	for _, pos := range m.LowPoints() {
		riskLevel += m[pos] + 1
	}
	return riskLevel
}

func Part2(m Heightmap) int {
	lowPoints := m.LowPoints()

	basinSizes := make([]int, len(lowPoints))
	for i, lowPoint := range lowPoints {
		basinSizes[i] = m.BasinSize(lowPoint)
	}

	sort.Ints(basinSizes)

	return advent2021.Product(basinSizes[len(basinSizes)-3:])
}
