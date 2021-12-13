package day13

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/chigley/advent2021"
)

type Grid map[advent2021.XY]struct{}

func NewGrid(coordinates []string) (Grid, error) {
	g := make(Grid, len(coordinates))

	for _, c := range coordinates {
		toks := strings.SplitN(c, ",", 2)
		if len(toks) != 2 {
			return nil, fmt.Errorf("day13: failed to parse %q", c)
		}

		x, err := strconv.Atoi(toks[0])
		if err != nil {
			return nil, fmt.Errorf("day13: atoi: %w", err)
		}

		y, err := strconv.Atoi(toks[1])
		if err != nil {
			return nil, fmt.Errorf("day13: atoi: %w", err)
		}

		g[advent2021.XY{X: x, Y: y}] = struct{}{}
	}

	return g, nil
}

func (g Grid) FoldX(val int) {
	for pos := range g {
		if pos.X > val {
			reflectedPos := advent2021.XY{X: val - (pos.X - val), Y: pos.Y}
			g[reflectedPos] = struct{}{}
			delete(g, pos)
		}
	}
}

func (g Grid) FoldY(val int) {
	for pos := range g {
		if pos.Y > val {
			reflectedPos := advent2021.XY{X: pos.X, Y: val - (pos.Y - val)}
			g[reflectedPos] = struct{}{}
			delete(g, pos)
		}
	}
}

func (g Grid) String() string {
	maxX, maxY := math.MinInt, math.MinInt
	for pos := range g {
		maxX = advent2021.Max(maxX, pos.X)
		maxY = advent2021.Max(maxY, pos.Y)
	}

	var b strings.Builder
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			_, ok := g[advent2021.XY{X: x, Y: y}]
			if ok {
				b.WriteRune('#')
			} else {
				b.WriteRune('.')
			}
		}
		b.WriteRune('\n')
	}
	return b.String()
}
