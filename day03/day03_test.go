package day03_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day03"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 uint
	part2 uint
}{
	{"example", 198, 230},
	{"input", 3969000, 4267809},
}

func TestDay03(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day03.Part1(in))

			part2, err := day03.Part2(in)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part2, part2)
		})
	}
}
