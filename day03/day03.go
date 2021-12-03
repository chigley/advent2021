package day03

import (
	"fmt"

	"github.com/chigley/advent2021"
)

func Part1(in []string) uint {
	if len(in) == 0 {
		return 0
	}

	bitLength := len(in[0])

	ones := make([]int, bitLength)
	for _, n := range in {
		for i, c := range n {
			if c == '1' {
				ones[i] += 1
			}
		}
	}

	var gamma uint
	for i := 0; i < bitLength; i++ {
		if ones[i] >= (len(in)+1)/2 {
			gamma |= 1 << (bitLength - i - 1)
		}
	}

	epsilon := ^gamma & ((1 << bitLength) - 1)

	return gamma * epsilon
}

func Part2(in []string) (uint, error) {
	oxyGen, err := evaluateCriteria(in, true)
	if err != nil {
		return 0, fmt.Errorf("advent2021: computing oxygen generator rating: %w", err)
	}

	co2Scrubber, err := evaluateCriteria(in, false)
	if err != nil {
		return 0, fmt.Errorf("advent2021: computing CO2 scrubber rating: %w", err)
	}

	return oxyGen * co2Scrubber, nil
}

func evaluateCriteria(in []string, wantMajority bool) (uint, error) {
	if len(in) == 0 {
		return 0, advent2021.ErrMissingInput
	}

	// Clone the input since we modify it in-place
	in = append([]string(nil), in...)

	for i := 0; i < len(in[0]); i++ {
		if len(in) == 1 {
			break
		}

		var ones int
		for _, n := range in {
			if n[i] == '1' {
				ones++
			}
		}

		p := ones >= (len(in)+1)/2
		if !wantMajority {
			p = !p
		}

		desiredBit := byte('0')
		if p {
			desiredBit = '1'
		}

		// Filter in to those with the desired bit
		var j int
		for _, n := range in {
			if n[i] == desiredBit {
				in[j] = n
				j++
			}
		}
		in = in[:j]
	}

	var ret uint
	if _, err := fmt.Sscanf(in[0], "%b", &ret); err != nil {
		return 0, fmt.Errorf("advent2021: scanning %s: %w", in[0], err)
	}

	return ret, nil
}
