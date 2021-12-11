package day11

import (
	"github.com/chigley/advent2021"
)

type Grid map[advent2021.XY]Octopus

type Octopus struct {
	Energy     int
	HasFlashed bool
}

func NewGrid(in []string) Grid {
	g := make(Grid, 100)
	for y, l := range in {
		for x, e := range l {
			g[advent2021.XY{X: x, Y: y}] = Octopus{Energy: int(e - '0')}
		}
	}
	return g
}

func (g Grid) Step() int {
	var flashes int

	for pos := range g {
		flashes += g.IncreaseEnergy(pos)
	}

	for pos, o := range g {
		if o.HasFlashed {
			o.Energy = 0
			o.HasFlashed = false
			g[pos] = o
		}
	}

	return flashes
}

func (g Grid) IncreaseEnergy(pos advent2021.XY) int {
	o := g[pos]
	o.Energy++
	g[pos] = o

	if o.Energy > 9 && !o.HasFlashed {
		return g.Flash(pos)
	}
	return 0
}

func (g Grid) Flash(pos advent2021.XY) int {
	flashes := 1

	o := g[pos]
	o.HasFlashed = true
	g[pos] = o

	for _, nPos := range pos.Neighbours(true) {
		if _, ok := g[nPos]; ok {
			flashes += g.IncreaseEnergy(nPos)
		}
	}

	return flashes
}
