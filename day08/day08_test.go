package day08_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day08"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 26},
	{"input", 261},
}

func TestDay08(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			entries, err := day08.ParseEntries(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day08.Part1(entries))
		})
	}
}
