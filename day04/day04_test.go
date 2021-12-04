package day04_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day04"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 int
}{
	{"example", 4512, 1924},
	{"input", 23177, 6804},
}

func TestDay04(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			draw, boards, err := day04.ParseInput(in)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day04.Part1(draw, boards)
			if err != nil {
				t.Error(err)
			}

			part2, err := day04.Part2(draw, boards)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
			assert.Equal(t, tt.part2, part2)
		})
	}
}
