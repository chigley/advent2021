package day10

import (
	"sort"
)

var openerToCloser = map[rune]rune{
	'(': ')',
	'[': ']',
	'{': '}',
	'<': '>',
}

var closerToOpener = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
	'>': '<',
}

var syntaxScore = map[rune]int{
	')': 3,
	']': 57,
	'}': 1197,
	'>': 25137,
}

var autocompleteScore = map[rune]int{
	')': 1,
	']': 2,
	'}': 3,
	'>': 4,
}

func Part1(lines []string) int {
	var score int

	for _, l := range lines {
		corrupt, firstIllegal, _ := isCorrupt(l)
		if corrupt {
			score += syntaxScore[firstIllegal]
		}
	}

	return score
}

func Part2(lines []string) int {
	var autocompleteScores []int
	for _, l := range lines {
		corrupt, _, completionString := isCorrupt(l)
		if !corrupt {
			autocompleteScores = append(autocompleteScores, completionStringScore(completionString))
		}
	}

	sort.Ints(autocompleteScores)

	return autocompleteScores[len(autocompleteScores)/2]
}

// isCorrupt returns true iff the given line is corrupt. If the line is corrupt,
// the first illegal character is returned in the second parameter.
// If the line is incomplete, the required completion string is returned in the
// third parameter.
func isCorrupt(line string) (bool, rune, string) {
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
			return true, c, ""
		}

		var actualOpeningChar rune
		actualOpeningChar, stack = stack[len(stack)-1], stack[:len(stack)-1]

		if actualOpeningChar != expectedOpeningChar {
			return true, c, ""
		}
	}

	// Line isn't corrupt, so it must be incomplete. Reverse the stack and
	// invert the brackets
	completionString := make([]rune, len(stack))
	for i, c := range stack {
		completionString[len(completionString)-(i+1)] = openerToCloser[c]
	}

	return false, 0, string(completionString)
}

func completionStringScore(s string) int {
	var score int
	for _, c := range s {
		score = score*5 + autocompleteScore[c]
	}
	return score
}
