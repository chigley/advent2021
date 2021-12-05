package day05

import (
	"fmt"

	"github.com/chigley/advent2021"
)

func ParseInput(in []string) ([][2]advent2021.XY, error) {
	ret := make([][2]advent2021.XY, len(in))

	for i, l := range in {
		var x1, y1, x2, y2 int
		if _, err := fmt.Sscanf(l, "%d,%d -> %d,%d", &x1, &y1, &x2, &y2); err != nil {
			return nil, fmt.Errorf("day05: parsing %q: %w", l, err)
		}

		ret[i] = [2]advent2021.XY{{x1, y1}, {x2, y2}}
	}

	return ret, nil
}
