// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package colour

import (
	"image/color"
	"math"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ~~~ My humble addition to the program ~~~
func Hue(value float64, max float64) color.RGBA {

	// Return red at extremities instead of black.
	if value < 0 || value > max || value == max {
		return color.RGBA{255, 0, 0, 255}
	}

	var c color.RGBA
	u := 1531.0 / max                  // Size of a 'unit' for scaling output.
	v := uint16(math.Round(value * u)) // Scaled value.

	c.A = 255 // Set alpha transparency globally, not used here.

	switch {
	case v >= 0 && v < 255:
		c.R = 255
		c.G = uint8(v)
		c.B = 0
	case v > 254 && v < 510:
		c.R = 255 - uint8(v-255)
		c.G = 255
		c.B = 0
	case v > 509 && v < 765:
		c.R = 0
		c.G = 255
		c.B = uint8(v - 510)
	case v > 764 && v < 1020:
		c.R = 0
		c.G = 255 - uint8(v-765)
		c.B = 255
	case v > 1019 && v < 1275:
		c.R = uint8(v - 1020)
		c.G = 0
		c.B = 255
	case v > 1274 && v < 1531:
		c.R = 255
		c.G = 0
		c.B = 255 - uint8(v-1275)
	}
	return c
}

// Returns the average color.RGBA from the given array of color.RGBA objects.
func AvgRGBA(c []color.RGBA) color.RGBA {
	// If there are no items in the slice, return early
	// to avoid divided by zero at the end.
	l := uint(len(c))
	if l == 0 {
		return color.RGBA{}
	}

	// Sum up all RGBA values
	var r, g, b, a uint
	for i := uint(0); i < l; i++ {
		r += uint(c[i].R)
		g += uint(c[i].G)
		b += uint(c[i].B)
		a += uint(c[i].A)
	}

	// Return the color with averaged components.
	return color.RGBA{
		R: uint8(r / l),
		G: uint8(g / l),
		B: uint8(b / l),
		A: uint8(a / l),
	}
}
