package day13

import (
	"fmt"
)

func Part1(g Grid, folds []string) (int, error) {
	if err := processFolds(g, folds[:1]); err != nil {
		return 0, err
	}
	return len(g), nil
}

// Part2 assumes Part1 was already called on the same Grid.
func Part2(g Grid, folds []string) (string, error) {
	if err := processFolds(g, folds[1:]); err != nil {
		return "", err
	}
	return g.String(), nil
}

func processFolds(g Grid, folds []string) error {
	for _, f := range folds {
		var (
			axis rune
			val  int
		)
		if _, err := fmt.Sscanf(f, "fold along %c=%d", &axis, &val); err != nil {
			return fmt.Errorf("day13: failed to parse %q: %w", f, err)
		}

		switch axis {
		case 'x':
			g.FoldX(val)
		case 'y':
			g.FoldY(val)
		default:
			return fmt.Errorf("day13: invalid axis %c", axis)
		}
	}
	return nil
}
