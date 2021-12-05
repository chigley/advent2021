package day05

import (
	"github.com/chigley/advent2021"
)

func Part1(lines [][2]advent2021.XY) int {
	return countOverlaps(lines, false)
}

func Part2(lines [][2]advent2021.XY) int {
	return countOverlaps(lines, true)
}

func countOverlaps(lines [][2]advent2021.XY, includeDiagonals bool) int {
	overlaps := make(map[advent2021.XY]int)
	for _, line := range lines {
		p1, p2 := line[0], line[1]

		minX, maxX := advent2021.Min(p1.X, p2.X), advent2021.Max(p1.X, p2.X)
		minY, maxY := advent2021.Min(p1.Y, p2.Y), advent2021.Max(p1.Y, p2.Y)

		if p1.X == p2.X {
			for y := minY; y <= maxY; y++ {
				overlaps[advent2021.XY{X: p1.X, Y: y}]++
			}
		} else if p1.Y == p2.Y {
			for x := minX; x <= maxX; x++ {
				overlaps[advent2021.XY{X: x, Y: p1.Y}]++
			}
		} else if includeDiagonals {
			direction := p2.Subtract(p1).Sign()
			for p := p1; p != p2; p = p.Add(direction) {
				overlaps[p]++
			}
			overlaps[p2]++
		}
	}

	var ret int
	for _, n := range overlaps {
		if n >= 2 {
			ret++
		}
	}

	return ret
}
