package day17

import (
	"fmt"
	"math"

	"github.com/chigley/advent2021"
)

// fudge represents the [-fudge..fudge) range we try for the X and Y components
// of each starting velocity attempted, and also the number of steps we take
// before concluding that a velocity doesn't hit the target area.
const fudge = 250

func Part1(xRange, yRange [2]int) int {
	maxY := math.MinInt
	for x := -fudge; x < fudge; x++ {
		for y := -fudge; y < fudge; y++ {
			if thisMaxY, ok := maxHeight(advent2021.XY{X: x, Y: y}, xRange, yRange); ok {
				maxY = advent2021.Max(maxY, thisMaxY)
			}
		}
	}
	return maxY
}

func Part2(xRange, yRange [2]int) int {
	var velocities int
	for x := -fudge; x < fudge; x++ {
		for y := -fudge; y < fudge; y++ {
			if _, ok := maxHeight(advent2021.XY{X: x, Y: y}, xRange, yRange); ok {
				velocities++
			}
		}
	}
	return velocities
}

func maxHeight(vel advent2021.XY, xRange, yRange [2]int) (int, bool) {
	p := Probe{Vel: vel}
	maxY := p.Pos.Y // 0

	for i := 0; i < fudge; i++ {
		p.Step()
		maxY = advent2021.Max(maxY, p.Pos.Y)

		if p.Pos.InRange(xRange, yRange) {
			return maxY, true
		}
	}

	return maxY, false
}

func ParseInput(in string) ([2]int, [2]int, error) {
	var x, y [2]int
	_, err := fmt.Sscanf(in, "target area: x=%d..%d, y=%d..%d", &x[0], &x[1], &y[0], &y[1])
	return x, y, err
}
