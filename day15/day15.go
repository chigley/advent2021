package day15

import (
	"github.com/chigley/advent2021"
)

func Part1(g Grid) int {
	return g.ShortestPath(
		advent2021.XY{X: 0, Y: 0},
		g.BottomRight(),
	)
}

func Part2(g Grid) int {
	expandedGrid := g.Expand(5)
	return expandedGrid.ShortestPath(
		advent2021.XY{X: 0, Y: 0},
		expandedGrid.BottomRight(),
	)
}
