package day14

func Part1(p *Polymer, r Rules) (int, error) {
	return scoreAfter(p, r, 10)
}

func Part2(p *Polymer, r Rules) (int, error) {
	return scoreAfter(p, r, 40)
}

func scoreAfter(p *Polymer, r Rules, n int) (int, error) {
	if err := p.StepN(r, n); err != nil {
		return 0, err
	}

	return p.Score(), nil
}
