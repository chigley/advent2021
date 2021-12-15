package day15_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day15"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 40, 315},
	{"input", 592, 2897},
}

func TestDay15(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			g := day15.NewGrid(in)

			assert.Equal(t, tt.part1, day15.Part1(g))
			assert.Equal(t, tt.part2, day15.Part2(g))
		})
	}
}
