package day08

import (
	"fmt"
	"math"

	mapset "github.com/deckarep/golang-set"
)

type Entry struct {
	Patterns     []string
	OutputDigits []string
}

func (e *Entry) Solve() (int, error) {
	var digits [10]mapset.Set

	fiveDigitPatterns, sixDigitPatterns := make([]string, 0, 3), make([]string, 0, 3)
	for _, p := range e.Patterns {
		switch len(p) {
		case 2:
			// 1 is the only digit using two segments
			digits[1] = patternToSet(p)
		case 3:
			// 7 is the only digit using three segments
			digits[7] = patternToSet(p)
		case 4:
			// 4 is the only digit using four segments
			digits[4] = patternToSet(p)
		case 5:
			fiveDigitPatterns = append(fiveDigitPatterns, p)
		case 6:
			sixDigitPatterns = append(sixDigitPatterns, p)
		case 7:
			// 8 is the only digit using seven segments
			digits[8] = patternToSet(p)
		}
	}

	// 9 is the only six-segment superset of 4. Find it and remove it from
	// sixDigitPatterns
	var j int
	for _, p := range sixDigitPatterns {
		candidate := patternToSet(p)
		if candidate.IsSuperset(digits[4]) {
			digits[9] = candidate
		} else {
			sixDigitPatterns[j] = p
			j++
		}
	}
	sixDigitPatterns = sixDigitPatterns[:j]

	// Now that 9 has been identified, 0 is the only remaining six-segment
	// superset of 1. Once we've identified 0, the other must be 6
	p0, p1 := patternToSet(sixDigitPatterns[0]), patternToSet(sixDigitPatterns[1])
	if p0.IsSuperset(digits[1]) {
		digits[0] = p0
		digits[6] = p1
	} else {
		digits[0] = p1
		digits[6] = p0
	}

	for _, p := range fiveDigitPatterns {
		digit := patternToSet(p)
		if digit.IsSuperset(digits[1]) {
			// 3 is the only five-segment superset of 1 (and 7, in fact - either
			// would work)
			digits[3] = digit
		} else if digit.IsSubset(digits[6]) {
			// 5 is the only five-segment subset of 6
			digits[5] = digit
		} else {
			// 2 is neither 3 nor 5
			digits[2] = digit
		}
	}

	var n int

digit:
	for i, d := range e.OutputDigits {
		ourDigit := patternToSet(d)

		for candidateDigit, candidateDigitSet := range digits {
			if ourDigit.Equal(candidateDigitSet) {
				n += candidateDigit * int(math.Pow10(len(e.OutputDigits)-(i+1)))
				continue digit
			}
		}

		return 0, fmt.Errorf("day08: failed to map %s to a digit", d)
	}

	return n, nil
}

func patternToSet(pattern string) mapset.Set {
	set := mapset.NewSet()
	for _, c := range pattern {
		set.Add(c)
	}
	return set
}
