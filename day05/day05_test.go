package day05_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day05"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 5, 12},
	{"input", 6007, 19349},
}

func TestDay05(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			lines, err := day05.ParseInput(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day05.Part1(lines))
			assert.Equal(t, tt.part2, day05.Part2(lines))
		})
	}
}
