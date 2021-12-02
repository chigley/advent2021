package day02_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day02"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
}{
	{"example", 150},
	{"input", 1451208},
}

func TestDay02(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStrings(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			cmds, err := day02.ParseCommands(in)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, tt.part1, day02.Part1(cmds))
		})
	}
}
