package day02

type Command struct {
	Direction Direction
	Units     int
}

type Direction int

const (
	Forward Direction = iota
	Down
	Up
)

func Part1(cmds []Command) int {
	var pos, depth int

	for _, cmd := range cmds {
		switch cmd.Direction {
		case Forward:
			pos += cmd.Units
		case Down:
			depth += cmd.Units
		case Up:
			depth -= cmd.Units
		}
	}

	return pos * depth
}
