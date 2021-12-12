package day12_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day12"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"exampleA", 10, 36},
	{"exampleB", 19, 103},
	{"exampleC", 226, 3509},
	{"input", 4495, 131254},
}

func TestDay12(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			cs, err := day12.NewCaveSystem(in)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day12.Part1(cs)
			if err != nil {
				t.Error(err)
			}

			part2, err := day12.Part2(cs)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
			assert.Equal(t, tt.part2, part2)
		})
	}
}
