package day10_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day10"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 26397},
	{"input", 462693},
}

func TestDay10(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day10.Part1(in))
		})
	}
}
