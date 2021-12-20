package day20

import (
	"fmt"
)

func ParseInput(in [][]string) ([]Pixel, *Image, error) {
	if len(in) != 2 {
		return nil, nil, fmt.Errorf("day20: got %d input groups, expected 2", len(in))
	}

	if len(in[0]) != 1 {
		return nil, nil, fmt.Errorf("day20: got %d lines for algorithm, expected 1", len(in[0]))
	}

	algo, err := parseAlgorithm(in[0][0])
	if err != nil {
		return nil, nil, fmt.Errorf("day20: parsing algorithm: %w", err)
	}

	img, err := NewImage(in[1])
	if err != nil {
		return nil, nil, fmt.Errorf("day20: parsing image: %w", err)
	}

	return algo, img, nil
}

func parseAlgorithm(in string) ([]Pixel, error) {
	algo := make([]Pixel, len(in))
	for i, p := range in {
		var err error
		algo[i], err = NewPixel(p)
		if err != nil {
			return nil, err
		}
	}
	return algo, nil
}
