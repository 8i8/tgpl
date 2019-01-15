// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package hue

import (
	"image/color"
	"math"
)

// ~~~ My humble addition to the program ~~~
func Hue(value float64, max float64) color.RGBA {

	// Return red at extremities instead of black.
	if value < 0 || value > max || value == max {
		return color.RGBA{255, 0, 0, 255}
	}

	var c color.RGBA
	u := 1531.0 / max                  // Size of a 'unit' for scaling output.
	v := uint16(math.Round(value * u)) // Scaled value.

	c.A = 255 // Set alpha transparency, globally as not used here.

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
