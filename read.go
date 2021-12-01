package advent2021

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func ReadInts(name string) ([]int, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2021: opening %s: %w", name, err)
	}
	defer f.Close()

	return readInts(f)
}

func readInts(r io.Reader) ([]int, error) {
	var ret []int

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("advent2021: atoi: %w", err)
		}

		ret = append(ret, n)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2021: scanner: %w", err)
	}

	return ret, nil
}
