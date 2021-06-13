package pkg

import (
	"image/color"
	"math"
)

const (
	numColors = 256
)

type DAT color.Palette

// Decode loads a DAT file into a color.Palette
func Decode(data []byte) (DAT, error) {
	const (
		// index offset helpers
		b = iota
		g
		r
		o
	)

	palette := make(DAT, numColors)

	for idx := range palette {
		// offsets look like idx*3+n, where n is 0,1,2
		palette[idx] = &color.RGBA{B: data[idx*o+b], G: data[idx*o+g], R: data[idx*o+r], A: math.MaxUint8}
	}

	return palette, nil
}

// Encode encodes a color.Palette into a byte slice in DAT palette format
func Encode(p DAT) []byte {
	result := make([]byte, 0)

	for idx := range p {
		r, g, b, _ := p[idx].RGBA()
		result = append(result, byte(b), byte(g), byte(r))
	}

	return result
}
