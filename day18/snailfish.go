package day18

import (
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
)

type Snailfish interface {
	Reduce() Snailfish

	Explode(depth int) (didExplode bool, next Snailfish, left, right Leaf)
	AddLeftmost(x Leaf) Snailfish
	AddRightmost(x Leaf) Snailfish

	Split() (didSplit bool, sf Snailfish)

	Add(x Snailfish) Snailfish

	Magnitude() int

	fmt.Stringer
}

func NewSnailfish(in string) (Snailfish, error) {
	r := strings.NewReader(in)

	sf, err := newSnailfish(r)
	if err != nil {
		return nil, err
	}

	c, err := r.ReadByte()
	if !errors.Is(err, io.EOF) {
		return nil, fmt.Errorf("day18: expected EOF, got %q", c)
	}

	return sf, nil
}

func newSnailfish(r io.ByteReader) (Snailfish, error) {
	c, err := r.ReadByte()
	if err != nil {
		return nil, err
	}
	if unicode.IsDigit(rune(c)) {
		return Leaf(c - '0'), nil
	}
	if c != '[' {
		return nil, fmt.Errorf("day18: expected '[', got %q", c)
	}

	left, err := newSnailfish(r)
	if err != nil {
		return nil, err
	}

	c, err = r.ReadByte()
	if err != nil {
		return nil, err
	}
	if c != ',' {
		return nil, fmt.Errorf("day18: expected ',', got %q", c)
	}

	right, err := newSnailfish(r)
	if err != nil {
		return nil, err
	}

	c, err = r.ReadByte()
	if err != nil {
		return nil, err
	}
	if c != ']' {
		return nil, fmt.Errorf("day18: expected ']', got %q", c)
	}

	return Tree{
		Left:  left,
		Right: right,
	}, nil
}

type Tree struct {
	Left, Right Snailfish
}

func (t Tree) Reduce() Snailfish {
	var sf Snailfish = t

	for {
		var didExplode bool
		if didExplode, sf, _, _ = sf.Explode(0); didExplode {
			continue
		}

		var didSplit bool
		if didSplit, sf = sf.Split(); !didSplit {
			break
		}
	}

	return sf
}

func (t Tree) Explode(depth int) (bool, Snailfish, Leaf, Leaf) {
	if depth == 4 {
		return true, Leaf(0), t.Left.(Leaf), t.Right.(Leaf)
	}

	if didExplode, acc, left, right := t.Left.Explode(depth + 1); didExplode {
		return true, Tree{Left: acc, Right: t.Right.AddLeftmost(right)}, left, 0
	}

	if didExplode, acc, left, right := t.Right.Explode(depth + 1); didExplode {
		return true, Tree{Left: t.Left.AddRightmost(left), Right: acc}, 0, right
	}

	return false, t, 0, 0
}

func (t Tree) AddLeftmost(x Leaf) Snailfish {
	return Tree{
		Left:  t.Left.AddLeftmost(x),
		Right: t.Right,
	}
}

func (t Tree) AddRightmost(x Leaf) Snailfish {
	return Tree{
		Left:  t.Left,
		Right: t.Right.AddRightmost(x),
	}
}

func (t Tree) Split() (bool, Snailfish) {
	if didSplit, acc := t.Left.Split(); didSplit {
		return true, Tree{Left: acc, Right: t.Right}
	}

	if didSplit, acc := t.Right.Split(); didSplit {
		return true, Tree{Left: t.Left, Right: acc}
	}

	return false, t
}

func (t Tree) Add(x Snailfish) Snailfish {
	return Tree{
		Left:  t,
		Right: x,
	}.Reduce()
}

func (t Tree) Magnitude() int {
	return 3*t.Left.Magnitude() + 2*t.Right.Magnitude()
}

func (t Tree) String() string {
	var b strings.Builder
	b.WriteRune('[')
	b.WriteString(t.Left.String())
	b.WriteRune(',')
	b.WriteString(t.Right.String())
	b.WriteRune(']')
	return b.String()
}

type Leaf int

func (l Leaf) Reduce() Snailfish {
	panic("can't reduce a leaf")
}

func (l Leaf) Explode(_ int) (bool, Snailfish, Leaf, Leaf) {
	return false, l, 0, 0
}

func (l Leaf) AddLeftmost(x Leaf) Snailfish {
	return l + x
}

func (l Leaf) AddRightmost(x Leaf) Snailfish {
	return l + x
}

func (l Leaf) Split() (bool, Snailfish) {
	if l < 10 {
		return false, nil
	}

	return true, Tree{
		Left:  l / 2,
		Right: (l + 1) / 2,
	}
}

func (l Leaf) Add(x Snailfish) Snailfish {
	return Tree{
		Left:  l,
		Right: x,
	}.Reduce()
}

func (l Leaf) Magnitude() int {
	return int(l)
}

func (l Leaf) String() string {
	return strconv.FormatInt(int64(l), 10)
}
