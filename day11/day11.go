package day11

func Part1(in []string) int {
	g := NewGrid(in)

	var flashes int
	for i := 0; i < 100; i++ {
		flashes += g.Step()
	}

	return flashes
}
