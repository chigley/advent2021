package day13_test

import (
	"path"
	"testing"

	"github.com/chigley/advent2021"
	"github.com/chigley/advent2021/day13"
	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	in    string
	part1 int
	part2 string
}{
	{"example", 17, "#####\n#...#\n#...#\n#...#\n#####\n"},
	{"input", 755, "###..#....#..#...##.###..###...##...##.\n#..#.#....#.#.....#.#..#.#..#.#..#.#..#\n###..#....##......#.#..#.###..#..#.#...\n#..#.#....#.#.....#.###..#..#.####.#.##\n#..#.#....#.#..#..#.#.#..#..#.#..#.#..#\n###..####.#..#..##..#..#.###..#..#..###\n"},
}

func TestDay13(t *testing.T) {
	for _, tt := range tests {
		tt := tt

		t.Run(tt.in, func(t *testing.T) {
			t.Parallel()

			in, err := advent2021.ReadStringGroups(path.Join("testdata", tt.in))
			if err != nil {
				t.Fatal(err)
			}

			g, err := day13.NewGrid(in[0])
			if err != nil {
				t.Fatal(err)
			}

			part1, err := day13.Part1(g, in[1])
			if err != nil {
				t.Error(err)
			}

			part2, err := day13.Part2(g, in[1])
			if err != nil {
				t.Error(err)
			}

			assert.Equal(t, tt.part1, part1)
			assert.Equal(t, tt.part2, part2)
		})
	}
}
