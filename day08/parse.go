package day08

import (
	"fmt"
	"strings"
)

func ParseEntries(in []string) ([]Entry, error) {
	const (
		numPatterns     = 10
		numOutputDigits = 4
	)

	entries := make([]Entry, len(in))
	for i, e := range in {
		toks := strings.SplitN(e, " | ", 2)
		if len(toks) != 2 {
			return nil, fmt.Errorf("day08: failed to parse %q", e)
		}

		patterns := strings.SplitN(toks[0], " ", numPatterns)
		if len(patterns) != numPatterns {
			return nil, fmt.Errorf("day08: got %d patterns in %q, expected %d", len(patterns), e, numPatterns)
		}

		outputDigits := strings.SplitN(toks[1], " ", numOutputDigits)
		if len(outputDigits) != 4 {
			return nil, fmt.Errorf("day08: got %d output digits in %q, expected %d", len(outputDigits), e, numOutputDigits)
		}

		entries[i] = Entry{
			Patterns:     patterns,
			OutputDigits: outputDigits,
		}
	}

	return entries, nil
}
