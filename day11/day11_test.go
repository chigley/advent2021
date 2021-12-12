package day11_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day11"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 1656, 195},
	{"input", 1667, 488},
}

func TestDay11(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			g1 := day11.NewGrid(in)
			g2 := g1.Clone()

			assert.Equal(t, tt.part1, day11.Part1(g1))
			assert.Equal(t, tt.part2, day11.Part2(g2))
		})
	}
}
