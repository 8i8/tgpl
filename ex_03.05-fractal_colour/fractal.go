// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"../ex_03.05-fractal_colour/hue"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	err := png.Encode(os.Stdout, img)
	check(err)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return hue.Hue(float64(contrast*n), iterations)
			//return color.Gray{255 - contrast*n}
		}
	}
	return color.Black
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

//!-

func myfunc(z complex128, width int) color.Color {
	return Hue(float64(real(z)+1), float64(width))
}

// Some other interesting functions:

func acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func sqrt(z complex128) color.Color {
	v := cmplx.Sqrt(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{128, blue, red}
}

// f(x) = x^4 - 1
//
// z' = z - f(z)/f'(z)
//    = z - (z^4 - 1) / (4 * z^3)
//    = z - (z - 1/z^3) / 4
func newton(z complex128) color.Color {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return Hue(255-float64(contrast*i), 255)
			//return color.Gray{255 - contrast*i}
		}
	}
	return color.Black
}
