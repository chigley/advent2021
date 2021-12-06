package day06

func Part1(in []int) int {
	for i := 0; i < 80; i++ {
		in = step(in)
	}
	return len(in)
}

func step(fish []int) []int {
	var newFishCount int
	for i, n := range fish {
		if n == 0 {
			fish[i] = 6
			newFishCount++
		} else {
			fish[i] = n - 1
		}
	}

	newFish := make([]int, newFishCount)
	for i := 0; i < newFishCount; i++ {
		newFish[i] = 8
	}

	return append(fish, newFish...)
}
