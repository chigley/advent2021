package day16

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BitString struct {
	r *strings.Reader
}

func NewBitString(hexString string) (BitString, error) {
	bytes, err := hex.DecodeString(hexString)
	if err != nil {
		return BitString{}, err
	}

	var b strings.Builder
	for _, bb := range bytes {
		if _, err := fmt.Fprintf(&b, "%08b", bb); err != nil {
			return BitString{}, err
		}
	}

	return BitString{
		r: strings.NewReader(b.String()),
	}, nil
}

func (bs BitString) ReadPacket() (Packet, error) {
	version, err := bs.readInt(3)
	if err != nil {
		return Packet{}, fmt.Errorf("day16: reading version: %w", err)
	}

	typeID, err := bs.readInt(3)
	if err != nil {
		return Packet{}, fmt.Errorf("day16: reading type ID: %w", err)
	}

	p := Packet{
		Version: version,
		Type:    TypeID(typeID),
	}

	if p.Type == Literal {
		v, err := bs.readLiteral()
		if err != nil {
			return Packet{}, fmt.Errorf("day16: reading literal: %w", err)
		}
		p.LiteralValue = v
	} else {
		subPackets, err := bs.readSubPackets()
		if err != nil {
			return Packet{}, fmt.Errorf("day16: reading sub-packets: %w", err)
		}
		p.SubPackets = subPackets
	}

	return p, nil
}

func (bs BitString) readInt(len int) (int, error) {
	buf := make([]byte, len)
	if _, err := bs.r.Read(buf); err != nil {
		return 0, err
	}

	n, err := strconv.ParseInt(string(buf), 2, 0)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func (bs BitString) readLiteral() (int, error) {
	var b strings.Builder
	for {
		buf := make([]byte, 5)
		if _, err := bs.r.Read(buf); err != nil {
			return 0, err
		}

		if _, err := b.Write(buf[1:5]); err != nil {
			return 0, err
		}

		if buf[0] == '0' {
			break
		}
	}

	n, err := strconv.ParseInt(b.String(), 2, 0)
	if err != nil {
		return 0, err
	}

	return int(n), nil
}

func (bs BitString) readSubPackets() ([]Packet, error) {
	lengthTypeID, err := bs.r.ReadByte()
	if err != nil {
		return nil, err
	}

	if lengthTypeID == '0' {
		return bs.readSubPacketsByLength()
	}
	return bs.readSubPacketsByCount()
}

func (bs BitString) readSubPacketsByLength() ([]Packet, error) {
	bitLen, err := bs.readInt(15)
	if err != nil {
		return nil, err
	}

	var ps []Packet

	var read int
	for origLen := bs.r.Len(); read < bitLen; read = origLen - bs.r.Len() {
		p, err := bs.ReadPacket()
		if err != nil {
			return nil, err
		}

		ps = append(ps, p)
	}

	if read > bitLen {
		return nil, errors.New("day16: sub-packets were longer than expected")
	}
	return ps, nil
}

func (bs BitString) readSubPacketsByCount() ([]Packet, error) {
	numSubPackets, err := bs.readInt(11)
	if err != nil {
		return nil, err
	}

	ps := make([]Packet, numSubPackets)
	for i := 0; i < numSubPackets; i++ {
		ps[i], err = bs.ReadPacket()
		if err != nil {
			return nil, err
		}
	}
	return ps, nil
}
