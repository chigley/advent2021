package day04

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/chigley/advent2021"
)

func ParseInput(groups [][]string) (draw []int, boards []Board, err error) {
	drawLines, err := advent2021.ReadIntLines(strings.NewReader(groups[0][0]))
	if err != nil {
		return nil, nil, fmt.Errorf("day04: parsing draw: %w", err)
	}
	draw = drawLines[0]

	boards = make([]Board, len(groups)-1)

	for i, b := range groups[1:] {
		boards[i] = make(Board, BoardDimension)

		for y, l := range b {
			boards[i][y] = make([]int, BoardDimension)

			for x, nAsc := range strings.Fields(l) {
				n, err := strconv.Atoi(nAsc)
				if err != nil {
					return nil, nil, fmt.Errorf("day04: atoi: %w", err)
				}

				boards[i][y][x] = n
			}
		}
	}

	_ = boards[0].IsWon(nil)

	return
}
