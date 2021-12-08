package day08

func Part1(entries []Entry) int {
	var n int
	for _, e := range entries {
		for _, d := range e.OutputDigits {
			numSegments := len(d)
			if (2 <= numSegments && numSegments <= 4) || numSegments == 7 {
				n++
			}
		}
	}
	return n
}
