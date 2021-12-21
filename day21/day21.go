package day21

import (
	"fmt"

	"github.com/chigley/advent2021"
)

const winningScore = 1000

func Part1(p1, p2 int) int {
	var (
		pos           = [2]int{p1, p2}
		score         [2]int
		d             DeterministicDie
		currentPlayer = 0
	)

	var turns int
	for ; score[0] < winningScore && score[1] < winningScore; turns++ {
		roll := d.RollN(3)
		pos[currentPlayer] = ((pos[currentPlayer] + roll - 1) % 10) + 1

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
