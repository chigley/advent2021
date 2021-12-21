package day21_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day21"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 739785, 444356092776315},
	{"input", 605070, 218433063958910},
}

func TestDay21(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			p1, p2, err := day21.ParseStartingPositions(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day21.Part1(p1, p2))
			assert.Equal(t, tt.part2, day21.Part2(p1, p2))
		})
	}
}
