package day09

import (
	"container/list"

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

	for _, nPos := range pos.Neighbours(false) {
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

func (m Heightmap) BasinSize(lowPoint advent2021.XY) int {
	var size int

	q := list.New()
	q.PushBack(lowPoint)

	visited := map[advent2021.XY]struct{}{
		lowPoint: {},
	}

	for e := q.Front(); e != nil; e = e.Next() {
		pos := e.Value.(advent2021.XY)

		height, ok := m[pos]
		if !ok || height == 9 {
			continue
		}

		size++

		for _, nPos := range pos.Neighbours(false) {
			if _, ok := visited[nPos]; ok {
				continue
			}
			visited[nPos] = struct{}{}
			q.PushBack(nPos)
		}
	}

	return size
}
