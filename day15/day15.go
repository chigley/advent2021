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
