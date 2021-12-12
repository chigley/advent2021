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
}{
	{"exampleA", 10},
	{"exampleB", 19},
	{"exampleC", 226},
	{"input", 4495},
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

			assert.Equal(t, tt.part1, part1)
		})
	}
}
