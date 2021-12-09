package day09

func Part1(m Heightmap) int {
	var riskLevel int
	for _, pos := range m.LowPoints() {
		riskLevel += m[pos] + 1
	}
	return riskLevel
}
