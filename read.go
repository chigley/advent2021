package advent2021

import (
	"bufio"
	"encoding/csv"
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

func ReadStrings(name string) ([]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2021: opening %s: %w", name, err)
	}
	defer f.Close()

	return readStrings(f)
}

func readStrings(r io.Reader) ([]string, error) {
	var ret []string

	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("advent2021: scanner: %w", err)
	}

	return ret, nil
}

// ReadStringGroups reads a file of string data formed of groups of lines, where
// groups are separated by an empty line.
func ReadStringGroups(name string) ([][]string, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, fmt.Errorf("advent2021: opening %s: %w", name, err)
	}
	defer f.Close()

	return readStringGroups(f)
}

func readStringGroups(r io.Reader) ([][]string, error) {
	lines, err := readStrings(r)
	if err != nil {
		return nil, fmt.Errorf("advent2021: reading strings: %w", err)
	}

	var groups [][]string

	var group []string
	for i, l := range lines {
		if l == "" {
			groups = append(groups, group)
			group = nil
			continue
		}

		group = append(group, l)

		if i == len(lines)-1 {
			groups = append(groups, group)
		}
	}

	return groups, nil
}

func ReadIntLines(r io.Reader) ([][]int, error) {
	var ret [][]int

	reader := csv.NewReader(r)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			return ret, nil
		}
		if err != nil {
			return nil, fmt.Errorf("advent2021: read error: %w", err)
		}

		line := make([]int, len(record))
		for i, nStr := range record {
			n, err := strconv.Atoi(nStr)
			if err != nil {
				return nil, fmt.Errorf("advent2021: atoi: %w", err)
			}

			line[i] = n
		}

		ret = append(ret, line)
	}
}
