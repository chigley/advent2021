package day18_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day18"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 4140, 3993},
	{"input", 2501, 4935},
}

func TestDay18(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day18.Part1(in)
			if err != nil {
				t.Error(err)
			}

			part2, err := day18.Part2(in)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
			assert.Equal(t, tt.part2, part2)
		})
	}
}
