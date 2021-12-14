package day14

func Part1(p *Polymer, r Rules) (int, error) {
	if err := p.StepN(r, 10); err != nil {
		return 0, err
	}

	return p.Score(), nil
}
