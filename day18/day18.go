package day18

import (
	"fmt"
)

func Part1(in []string) (int, error) {
	snailfish := make([]Snailfish, len(in))
	for i, l := range in {
		var err error
		snailfish[i], err = NewSnailfish(l)
		if err != nil {
			return 0, fmt.Errorf("day18: parsing snailfish %d: %w", i, err)
		}
	}

	sum := snailfish[0]
	for i := 1; i < len(snailfish); i++ {
		sum = sum.Add(snailfish[i])
	}

	return sum.Magnitude(), nil
}
