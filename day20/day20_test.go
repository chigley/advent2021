package day20_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day20"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 35},
	{"input", 5339},
}

func TestDay20(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			algo, img, err := day20.ParseInput(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day20.Part1(algo, img))
		})
	}
}
