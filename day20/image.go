package day20

import (
	"fmt"
	"strings"

	"github.com/chigley/advent2021"
)

var neighbourDirections = []advent2021.XY{
	advent2021.NorthEast,
	advent2021.North,
	advent2021.NorthWest,
	advent2021.East,
	{0, 0},
	advent2021.West,
	advent2021.SouthEast,
	advent2021.South,
	advent2021.SouthWest,
}

type Pixel bool

const (
	Dark  Pixel = false
	Light       = true
)

func NewPixel(c rune) (Pixel, error) {
	switch c {
	case '#':
		return Light, nil
	case '.':
		return Dark, nil
	default:
		return Dark, fmt.Errorf("day20: invalid pixel %q", c)
	}
}

func (p Pixel) String() string {
	switch p {
	case Light:
		return "#"
	default:
		return "."
	}
}

type Image struct {
	OrigW, OrigH int
	Growths      int
	Pixels       map[advent2021.XY]Pixel
}

func NewImage(in []string) (*Image, error) {
	pixels := make(map[advent2021.XY]Pixel, len(in)*len(in[0]))
	for y, l := range in {
		for x, c := range l {
			p, err := NewPixel(c)
			if err != nil {
				return nil, err
			}

			pixels[advent2021.XY{X: x, Y: y}] = p
		}
	}

	return &Image{
		OrigW:  len(in[0]),
		OrigH:  len(in),
		Pixels: pixels,
	}, nil
}

func (i *Image) Enhance(algo []Pixel) {
	i.Growths++

	// If the 0th element of the algorithm is Light and the 511th is dark then
	// we need out-of-bounds pixels to alternate state on each enhancement.
	alternateInfinite := algo[0] == Light && algo[511] == Dark

	newPixels := make(map[advent2021.XY]Pixel)
	for y := -i.Growths; y < i.OrigH+i.Growths; y++ {
		for x := -i.Growths; x < i.OrigW+i.Growths; x++ {
			pos := advent2021.XY{X: x, Y: y}
			newPixels[pos] = algo[i.Index(pos, alternateInfinite)]
		}
	}

	i.Pixels = newPixels
}

func (i *Image) Index(pos advent2021.XY, alternateInfinite bool) int {
	var algorithmIndex int

	for j, dir := range neighbourDirections {
		pixel, ok := i.Pixels[pos.Add(dir)]
		if !ok && alternateInfinite {
			pixel = i.Growths%2 == 0
		}

		if pixel == Light {
			algorithmIndex |= 1 << j
		}
	}

	return algorithmIndex
}

func (i *Image) LitPixels() int {
	var lit int
	for _, p := range i.Pixels {
		if p == Light {
			lit++
		}
	}
	return lit
}

func (i *Image) String() string {
	var b strings.Builder
	for y := -i.Growths; y < i.OrigH+i.Growths; y++ {
		for x := -i.Growths; x < i.OrigW+i.Growths; x++ {
			b.WriteString(i.Pixels[advent2021.XY{X: x, Y: y}].String())
		}
		b.WriteRune('\n')
	}
	return b.String()
}
