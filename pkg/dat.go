package pkg

import (
	"image/color"
	"math"
)

const (
	numColors = 256
)

// Unmarshal loads a DAT file into a color.Palette
func Unmarshal(data []byte) (color.Palette, error) {
	const (
		// index offset helpers
		b = iota
		g
		r
		o
	)

	palette := make(color.Palette, numColors)

	for idx := range palette {
		// offsets look like idx*3+n, where n is 0,1,2
		palette[idx] = &color.RGBA{B: data[idx*o+b], G: data[idx*o+g], R: data[idx*o+r], A: math.MaxUint8}
	}

	return palette, nil
}

// Marshal encodes a color.Palette into a byte slice in DAT palette format
func Marshal(p color.Palette) []byte {
	result := make([]byte, 0)

	for idx := range p {
		r, g, b, _ := p[idx].RGBA()
		result = append(result, byte(b), byte(g), byte(r))
	}

	return result
}
