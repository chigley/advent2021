package day12

import (
	"fmt"
	"strings"
)

type (
	CaveSystem map[string]map[string]struct{}
	Caves      map[string]struct{}
)

func NewCaveSystem(in []string) (CaveSystem, error) {
	cs := make(CaveSystem)

	for _, l := range in {
		toks := strings.SplitN(l, "-", 2)
		if len(toks) != 2 {
			return nil, fmt.Errorf("day12: failed to parse %q", l)
		}

		// Initialise neighbour set if this is the first time we've seen either
		// cave
		for _, c := range toks {
			if _, ok := cs[c]; !ok {
				cs[c] = make(Caves)
			}
		}

		// Mark each cave as a neighbour of the other
		c1, c2 := toks[0], toks[1]
		cs[c1][c2] = struct{}{}
		cs[c2][c1] = struct{}{}
	}

	return cs, nil
}
