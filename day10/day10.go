package day10

var closerToOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var characterScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

func Part1(lines []string) int {
	var score int

	for _, l := range lines {
		corrupt, firstIllegal := isCorrupt(l)
		if corrupt {
			score += characterScore[firstIllegal]
		}
	}

	return score
}

// isCorrupt returns true iff the given line is corrupt, also returning the
// first illegal character if so.
func isCorrupt(line string) (bool, rune) {
	var stack []rune

	for _, c := range line {
		expectedOpeningChar, isCloser := closerToOpener[c]
		if !isCloser {
			// We've got an opening bracket. Push it onto the stack
			stack = append(stack, c)
			continue
		}

		// We've got a closing bracket. Pop an opening bracket off the stack and
		// make sure it matches

		if len(stack) == 0 {
			return true, c
		}

		var actualOpeningChar rune
		actualOpeningChar, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if actualOpeningChar != expectedOpeningChar {
			return true, c
		}
	}

	return false, 0
}
