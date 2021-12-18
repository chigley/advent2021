package day18

import (
	"fmt"
	"math"

	"github.com/chigley/advent2021"
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

func Part2(in []string) (int, error) {
	snailfish := make([]Snailfish, len(in))
	for i, l := range in {
		var err error
		snailfish[i], err = NewSnailfish(l)
		if err != nil {
			return 0, fmt.Errorf("day18: parsing snailfish %d: %w", i, err)
		}
	}

	max := math.MinInt
	for i := 0; i < len(snailfish); i++ {
		for j := 0; j < len(snailfish); j++ {
			if i == j {
				continue
			}

			// No need to do two Add() calls here, since our j loops from [0, len]
			max = advent2021.Max(max, snailfish[i].Add(snailfish[j]).Magnitude())
		}
	}

	return max, nil
}
