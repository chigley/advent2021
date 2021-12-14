package day14

import (
	"fmt"
)

func ParseInput(in [][]string) (*Polymer, Rules, error) {
	if len(in) != 2 {
		return nil, nil, fmt.Errorf("day14: got %d input groups, expected 2", len(in))
	}

	if len(in[0]) != 1 {
		return nil, nil, fmt.Errorf("day14: got %d polymer(s), expected 1", len(in[0]))
	}

	polymer, err := NewPolymer(in[0][0])
	if err != nil {
		return nil, nil, fmt.Errorf("day14: initialising polymer: %w", err)
	}

	rules, err := NewRules(in[1])
	if err != nil {
		return nil, nil, fmt.Errorf("day14: initialising rules: %w", err)
	}

	return polymer, rules, nil
}
