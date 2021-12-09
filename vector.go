package advent2021

var (
	North = XY{0, 1}
	East  = XY{1, 0}
	South = XY{0, -1}
	West  = XY{-1, 0}

	Dirs = [4]XY{North, East, South, West}
)

type XY struct {
	X, Y int
}

func (p1 XY) Add(p2 XY) XY {
	return XY{
		X: p1.X + p2.X,
		Y: p1.Y + p2.Y,
	}
}

func (p1 XY) Subtract(p2 XY) XY {
	return p1.Add(p2.Multiply(-1))
}

func (p XY) Multiply(n int) XY {
	return XY{
		X: n * p.X,
		Y: n * p.Y,
	}
}

func (p XY) Sign() XY {
	return XY{
		X: Sign(p.X),
		Y: Sign(p.Y),
	}
}

func (p XY) Neighbours() [4]XY {
	var ns [4]XY
	for i, dir := range Dirs {
		ns[i] = p.Add(dir)
	}
	return ns
}
