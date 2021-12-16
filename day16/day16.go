package day16

func Part1(bs BitString) (int, error) {
	p, err := bs.ReadPacket()
	if err != nil {
		return 0, err
	}

	return p.TotalVersion(), nil
}
