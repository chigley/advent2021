package day03

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
		if ones[i] > len(in)/2 {
			gamma |= 1 << (bitLength - i - 1)
		}
	}

	epsilon := ^gamma & ((1 << bitLength) - 1)

	return gamma * epsilon
}
