package day02

import (
	"fmt"
	"strconv"
	"strings"
)

func ParseCommands(in []string) ([]Command, error) {
	cmds := make([]Command, len(in))

	for i, l := range in {
		toks := strings.SplitN(l, " ", 2)
		if len(toks) != 2 {
			return nil, fmt.Errorf("day02: failed to parse %q", l)
		}

		var dir Direction
		switch toks[0] {
		case "forward":
			dir = Forward
		case "down":
			dir = Down
		case "up":
			dir = Up
		default:
			return nil, fmt.Errorf("day02: unknown direction %q", toks[0])
		}

		units, err := strconv.Atoi(toks[1])
		if err != nil {
			return nil, fmt.Errorf("day02: atoi: %w", err)
		}

		cmds[i] = Command{
			Direction: dir,
			Units:     units,
		}
	}

	return cmds, nil
}
