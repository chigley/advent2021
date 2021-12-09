package day09

import (
	"github.com/chigley/advent2021"
)

type Heightmap map[advent2021.XY]int

func NewHeightmap(in []string) Heightmap {
	m := make(map[advent2021.XY]int)
	for y, l := range in {
		for x, h := range l {
			m[advent2021.XY{X: x, Y: y}] = int(h - '0')
		}
	}
	return m
}

func (m Heightmap) LowPoints() []advent2021.XY {
	var lowPoints []advent2021.XY
	for pos := range m {
		if m.IsLowPoint(pos) {
			lowPoints = append(lowPoints, pos)
		}
	}
	return lowPoints
}

func (m Heightmap) IsLowPoint(pos advent2021.XY) bool {
	ourHeight, ok := m[pos]
	if !ok {
		return false
	}

	for _, nPos := range pos.Neighbours() {
		neightbourHeight, ok := m[nPos]
		if !ok {
			continue
		}

		if neightbourHeight <= ourHeight {
			return false
		}
	}

	return true
}
