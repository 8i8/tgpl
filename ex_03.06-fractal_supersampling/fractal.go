// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 61.
//!+

// Mandelbrot emits a PNG image of the Mandelbrot fractal.
package main

import (
	"../ex_03.06-fractal_supersampling/colour"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

const (
	width, height = 3, 3
	smp_r         = 1
)

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Set the dimentions and sub sampeling for the programs fractal generation.
func main() {
	var smp [smp_r * smp_r]color.RGBA
	img := generate(width*smp_r, height*smp_r)
	sub := subSample(img, smp[:], width, height, smp_r)

	f, err := os.Create("fractal.png")
	Check(err)
	defer f.Close()

	err = png.Encode(f, sub)
	Check(err)
}

// Generate a large image for colour sampeling.
func generate(width, height int) *image.RGBA {

	var xmin, ymax, ymin, xmax float64
	xmin, ymin = -2, -2
	xmax = xmin + 4
	ymax = ymin + (xmax - xmin)
	sub := image.NewRGBA(image.Rect(0, 0, width, height))
	// Fill Sub sampling array with color.RGBA structs for down sampling.
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			sub.Set(py, px, mandelbrot(z))
			//sub.Set(py, px, newton(z))
		}
	}
	return sub
}

// Generate a sub image from the sampled data.
func subSample(img *image.RGBA, smp []color.RGBA, width, height, smp_r int) *image.RGBA {

	sub := image.NewRGBA(image.Rect(0, 0, width, height))
	x, y := 0, 0
	for py := 0; py < height*smp_r; py += smp_r {
		for px := 0; px < width*smp_r; px += smp_r {
			sub.Set(y, x, sample(img, smp[:], px, py, smp_r))
			x++
		}
		x = 0
		y++
	}
	return sub
}

// Sample a submarix of size fac^2.
func sample(img *image.RGBA, smp []color.RGBA, px, py, fac int) color.RGBA {

	count := 0
	for x := 0; x < fac; x++ {
		for y := 0; y < fac; y++ {
			smp[count] = img.RGBAAt(px+x, py+y)
			count++
		}
	}
	return colour.AvgRGBA(smp[:])
}

// The mandelbrot set.
func mandelbrot(z complex128) color.RGBA {
	const iterations = 8
	const contrast = 200
	var v complex128
	for n := uint64(0); n < iterations; n++ {
		// 1 fmt.Printf("%d\nbef %1.15f %1.15f\n", n, real(v), imag(v))
		v = v * v
		// 1 fmt.Printf("mul %1.15f %1.15f\n", real(v), imag(v))
		v = v + z
		// 1 fmt.Printf("add %1.15f %1.15f\n", real(v), imag(v))
		// 1 fmt.Printf("abs %1.15f\n", cmplx.Abs(v))
		if cmplx.Abs(v) > 2 {
			return colour.Hue(float64(contrast*n), 1531)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

//!-

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
func newton(z complex128) color.RGBA {
	const iterations = 37
	const contrast = 7
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		if cmplx.Abs(z*z*z*z-1) < 1e-6 {
			return colour.Hue(255-float64(contrast*i), 255)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
