package day04

import (
	"github.com/chigley/advent2021"
)

func Part1(drawOrder []int, boards []Board) (int, error) {
	draw := make(map[int]struct{})

	for _, n := range drawOrder {
		draw[n] = struct{}{}

		for _, b := range boards {
			if b.IsWon(draw) {
				return b.SumUnmarked(draw) * n, nil
			}
		}

	}

	return 0, advent2021.ErrNoSolution
}
