package day14

import (
	"fmt"

	"github.com/chigley/advent2021"
)

type Pair [2]rune

type Polymer struct {
	// pairFreq tracks the frequency of each pair of elements in the polymer.
	pairFreq map[Pair]int

	// lastChar stores the last element from the input string used to
	// initialsise the polymer, used to compute its score.
	lastChar rune
}

func NewPolymer(template string) (*Polymer, error) {
	if len(template) < 2 {
		return nil, fmt.Errorf("day14: got template length %d, expected 2+", len(template))
	}

	pairFreq := make(map[Pair]int)
	for i := 0; i < len(template)-1; i++ {
		p1, p2 := rune(template[i]), rune(template[i+1])
		pairFreq[Pair{p1, p2}]++
	}

	return &Polymer{
		pairFreq: pairFreq,
		lastChar: rune(template[len(template)-1]),
	}, nil
}

func (p *Polymer) StepN(r Rules, n int) error {
	for i := 0; i < n; i++ {
		if err := p.Step(r); err != nil {
			return fmt.Errorf("day14: step %d: %w", i, err)
		}
	}
	return nil
}

func (p *Polymer) Step(r Rules) error {
	nextPairFreq := make(map[Pair]int)
	for pair, freq := range p.pairFreq {
		element, ok := r[pair]
		if !ok {
			return fmt.Errorf("day14: missing rule for %c%c", pair[0], pair[1])
		}

		nextPairFreq[Pair{pair[0], element}] += freq
		nextPairFreq[Pair{element, pair[1]}] += freq
	}

	p.pairFreq = nextPairFreq
	return nil
}

func (p *Polymer) Score() int {
	// Count occurrences of the first character in every pair, then manually
	// increment the count for the last character from the initial template
	// input.
	elementFreqs := make(map[rune]int)
	for pair, freq := range p.pairFreq {
		elementFreqs[pair[0]] += freq
	}
	elementFreqs[p.lastChar]++

	var i int
	occs := make([]int, len(elementFreqs))
	for _, f := range elementFreqs {
		occs[i] = f
		i++
	}

	return advent2021.MaxInts(occs) - advent2021.MinInts(occs)
}

type Rules map[Pair]rune

// https://www.youtube.com/watch?v=k2qgadSvNyU
func NewRules(rules []string) (Rules, error) {
	r := make(Rules, len(rules))

	for _, rr := range rules {
		var p1, p2, out rune
		if _, err := fmt.Sscanf(rr, "%c%c -> %c", &p1, &p2, &out); err != nil {
			return nil, fmt.Errorf("day14: failed to parse %q", rr)
		}

		r[Pair{p1, p2}] = out
	}

	return r, nil
}
