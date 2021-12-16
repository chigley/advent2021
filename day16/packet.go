package day16

type TypeID int

const (
	Literal TypeID = 4
)

type Packet struct {
	Version int
	Type    TypeID

	LiteralValue int
	SubPackets   []Packet
}

func (p *Packet) TotalVersion() int {
	sum := p.Version
	if p.Type != Literal {
		for _, pp := range p.SubPackets {
			sum += pp.TotalVersion()
		}
	}
	return sum
}
