package fractal

import (
	"8i8/colour"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/cmplx"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Set the dimentions and sub sampeling for the programs fractal generation.
func Draw(w io.Writer, frac func(complex128) color.Color) {
	const (
		width, height = 1000, 1000
		smp_f         = 3
	)
	var smp [smp_f * smp_f]color.RGBA
	img := generate(frac, width*smp_f, height*smp_f)
	sub := subSample(img, smp[:], width, height, smp_f)

	err := png.Encode(w, sub)
	check(err)
}

// Generate the fractal.
func generate(frac func(complex128) color.Color, width, height int) *image.RGBA {

	const xmin, ymin, xmax, ymax = -2, -2, +2, +2
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			z := complex(x, y)
			img.Set(py, px, frac(z))
		}
	}
	return img
}

// Generate a sub image from the sampled data.
func subSample(img *image.RGBA, smp []color.RGBA, width, height, smp_f int) *image.RGBA {

	sub := image.NewRGBA(image.Rect(0, 0, width, height))
	x, y := 0, 0
	for py := 0; py < height*smp_f; py += smp_f {
		for px := 0; px < width*smp_f; px += smp_f {
			sub.Set(y, x, sample(img, smp[:], px, py, smp_f))
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
func Mandelbrot(z complex128) color.Color {

	const iterations = 200
	const contrast = 8
	var v complex128
	for n := int(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colour.Hue(float64(contrast*n), 1531)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

//!-

// Some other interesting functions:

func Acos(z complex128) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func Sqrt(z complex128) color.Color {
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
func Newton(z complex128) color.Color {
	const iterations = 51
	const contrast = 5
	const width = 255
	for i := uint8(0); i < iterations; i++ {
		z -= (z - 1/(z*z*z)) / 4
		y := z*z*z*z - 1
		if cmplx.Abs(y) < 1e-6 {
			switch {
			case real(z) >= 0 && imag(z) >= 0:
				c := color.RGBA{255, 0, 0, 255}
				c = colour.LumRGBA(c, width-float64(contrast*i))
				return c
			case real(z) >= 0 && imag(z) <= 0:
				c := color.RGBA{0, 255, 0, 255}
				c = colour.LumRGBA(c, width-float64(contrast*i))
				return c
			case real(z) <= 0 && imag(z) >= 0:
				c := color.RGBA{0, 0, 255, 255}
				c = colour.LumRGBA(c, width-float64(contrast*i))
				return c
			case real(z) <= 0 && imag(z) <= 0:
				c := color.RGBA{0, 255, 255, 255}
				c = colour.LumRGBA(c, width-float64(contrast*i))
				return c
			}
		}
	}
	return color.RGBA{0, 0, 0, 255}
}
