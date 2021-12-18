package day17

import (
	"github.com/chigley/advent2021"
)

type Probe struct {
	Pos, Vel advent2021.XY
}

func (p *Probe) Step() {
	p.Pos = p.Pos.Add(p.Vel)

	p.Vel.X -= advent2021.Sign(p.Vel.X)
	p.Vel.Y--
}
