package day21

import (
	"fmt"

	"github.com/chigley/advent2021"
)

const (
	part1WinningScore = 1000
	part2WinningScore = 21
)

func Part1(p1, p2 int) int {
	var (
		pos           = [2]int{p1, p2}
		score         [2]int
		d             DeterministicDie
		currentPlayer = 0
	)

	var turns int
	for ; score[0] < part1WinningScore && score[1] < part1WinningScore; turns++ {
		roll := d.RollN(3)
		pos[currentPlayer] = (pos[currentPlayer]+roll-1)%10 + 1

		score[currentPlayer] += pos[currentPlayer]

		currentPlayer = (currentPlayer + 1) % 2
	}

	return turns * 3 * advent2021.MinInts(score[:])
}

type DeterministicDie struct {
	Last int
}

func (d *DeterministicDie) Roll() int {
	d.Last = d.Last%100 + 1
	return d.Last
}

func (d *DeterministicDie) RollN(n int) int {
	var sum int
	for i := 0; i < n; i++ {
		sum += d.Roll()
	}
	return sum
}

type gameState struct {
	pos, score [2]int
}

func Part2(p1, p2 int) int {
	universesWon := wins(
		make(map[gameState][2]int),
		gameState{pos: [2]int{p1, p2}},
	)
	return advent2021.MaxInts(universesWon[:])
}

func wins(cache map[gameState][2]int, state gameState) [2]int {
	universesWon, ok := cache[state]
	if ok {
		return universesWon
	}

	pos, score := state.pos, state.score

	if score[0] >= part2WinningScore {
		return [2]int{1, 0}
	}
	if score[1] >= part2WinningScore {
		return [2]int{0, 1}
	}

	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			for k := 1; k <= 3; k++ {
				// Work out player 0's new position and score. Don't modify
				// pos[0] and score[0] in place since they're shared by other
				// loop iterations (you absolute melon).
				newPos := (pos[0]+i+j+k-1)%10 + 1
				newScore := score[0] + newPos

				// Having updated player 0's position and score, swap the
				// players around and recurse. Player 1 now becomes player 0.
				nextWins := wins(
					cache,
					gameState{
						pos:   [2]int{pos[1], newPos},
						score: [2]int{score[1], newScore},
					},
				)

				// In this universe, we need to swap the results back again,
				// since we swapped the players in our recursive call to wins().
				universesWon[0] += nextWins[1]
				universesWon[1] += nextWins[0]
			}
		}
	}

	cache[state] = universesWon
	return universesWon
}

func ParseStartingPositions(in []string) (int, int, error) {
	if len(in) != 2 {
		return 0, 0, fmt.Errorf("day21: got %d input lines, expected 2", len(in))
	}

	var p1 int
	if _, err := fmt.Sscanf(in[0], "Player 1 starting position: %d", &p1); err != nil {
		return 0, 0, err
	}

	var p2 int
	if _, err := fmt.Sscanf(in[1], "Player 2 starting position: %d", &p2); err != nil {
		return 0, 0, err
	}

	return p1, p2, nil
}
