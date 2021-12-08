package day08

import (
	"fmt"
)

func Part1(entries []Entry) int {
	var n int
	for _, e := range entries {
		for _, d := range e.OutputDigits {
			numSegments := len(d)
			if (2 <= numSegments && numSegments <= 4) || numSegments == 7 {
				n++
			}
		}
	}
	return n
}

func Part2(entries []Entry) (int, error) {
	var n int
	for i, e := range entries {
		output, err := e.Solve()
		if err != nil {
			return 0, fmt.Errorf("day08: solving entry %d: %w", i, err)
		}

		n += output
	}
	return n, nil
}
