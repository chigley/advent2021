package day16

import (
	"fmt"

	"github.com/chigley/advent2021"
)

type TypeID int

const (
	Sum TypeID = iota
	Product
	Min
	Max
	Literal
	Gt
	Lt
	Eq
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

func (p *Packet) Value() (int, error) {
	if p.Type == Literal {
		return p.LiteralValue, nil
	}

	subPacketValues := make([]int, len(p.SubPackets))
	for i, pp := range p.SubPackets {
		var err error
		subPacketValues[i], err = pp.Value()
		if err != nil {
			return 0, err
		}
	}

	switch p.Type {
	case Sum:
		return advent2021.Sum(subPacketValues), nil
	case Product:
		return advent2021.Product(subPacketValues), nil
	case Min:
		return advent2021.MinInts(subPacketValues), nil
	case Max:
		return advent2021.MaxInts(subPacketValues), nil
	case Gt:
		if subPacketValues[0] > subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}
	case Lt:
		if subPacketValues[0] < subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}
	case Eq:
		if subPacketValues[0] == subPacketValues[1] {
			return 1, nil
		} else {
			return 0, nil
		}
	default:
		return 0, fmt.Errorf("day16: invalid packet type %d", p.Type)
	}
}
