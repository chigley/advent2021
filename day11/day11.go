package day11

func Part1(g Grid) int {
	var flashes int
	for i := 0; i < 100; i++ {
		flashes += g.Step()
	}

	return flashes
}

func Part2(g Grid) int {
	steps := 1
	for g.Step() != 100 {
		steps++
	}

	return steps
}
