package day14_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day14"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 1588},
	{"input", 2321},
}

func TestDay14(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			polymer, rules, err := day14.ParseInput(in)
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day14.Part1(polymer, rules)
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
		})
	}
}
