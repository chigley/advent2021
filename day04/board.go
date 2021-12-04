package day04

const BoardDimension = 5

type Board [][]int

func (b Board) IsWon(draw map[int]struct{}) bool {
row:
	for _, row := range b {
		for _, n := range row {
			if _, ok := draw[n]; !ok {
				continue row
			}
		}
		return true
	}

column:
	for i := 0; i < BoardDimension; i++ {
		for _, row := range b {
			if _, ok := draw[row[i]]; !ok {
				continue column
			}
		}
		return true
	}

	return false
}

func (b Board) SumUnmarked(draw map[int]struct{}) int {
	var sumUnmarked int

	for _, row := range b {
		for _, n := range row {
			if _, ok := draw[n]; !ok {
				sumUnmarked += n
			}
		}
	}

	return sumUnmarked
}
