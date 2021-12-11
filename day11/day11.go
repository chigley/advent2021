package day11

func Part1(in []string) int {
	g := NewGrid(in)

	var flashes int
	for i := 0; i < 100; i++ {
		flashes += g.Step()
	}

	return flashes
}

func Part2(in []string) int {
	g := NewGrid(in)

	steps := 1
	for g.Step() != 100 {
		steps++
	}

	return steps
}
