package day16

func Part1(in string) (int, error) {
	bs, err := NewBitString(in)
	if err != nil {
		return 0, err
	}

	p, err := bs.ReadPacket()
	if err != nil {
		return 0, err
	}

	return p.TotalVersion(), nil
}

func Part2(in string) (int, error) {
	bs, err := NewBitString(in)
	if err != nil {
		return 0, err
	}

	p, err := bs.ReadPacket()
	if err != nil {
		return 0, err
	}

	return p.Value()
}
