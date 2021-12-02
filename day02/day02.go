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

func Part2(cmds []Command) int {
	var pos, depth, aim int

	for _, cmd := range cmds {
		switch cmd.Direction {
		case Forward:
			pos += cmd.Units
			depth += aim * cmd.Units
		case Down:
			aim += cmd.Units
		case Up:
			aim -= cmd.Units
		}
	}

	return pos * depth
}
