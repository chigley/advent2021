package day17_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day17"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 45},
	{"input", 3655},
}

func TestDay17(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			xRange, yRange, err := day17.ParseInput(in[0])
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day17.Part1(xRange, yRange))
		})
	}
}
