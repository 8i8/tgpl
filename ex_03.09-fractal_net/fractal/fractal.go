package fractal

import (
	"8i8/cmpbig"
	"8i8/colour"
	//"encoding/base64"
	//"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math/big"
	"math/cmplx"
)

type Config struct {
	Width  int
	Height int
	Smp_r  int
	Frame  float64
	Long   float64
	Lati   float64
	Iter   int
	Cont   int
	Cspa   float64
}

type ConfigRAT struct {
	Width  int
	Height int
	Smp_r  int
	Frame  int64
	Long   int64
	Lati   int64
	Denom  int64
	Iter   int
	Cont   int
	Cspa   float64
}

const BUF = 3

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

// Set the dimensions and sub sampling for the programs fractal generation.
func Draw32(w io.Writer, fractal func(complex64, Config) color.Color, c Config) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Generate scaled image.
	img := generate32(fractal, c)

	// If required down sample.
	if c.Smp_r > 1 {
		img = subSample(img, smp[:], c)
	}

	// Encode.
	err := png.Encode(w, img)
	Check(err)
	//str := base64.StdEncoding.EncodeToString([]byte(img))
	//fmt.Fprintf(w, str)
}

// Iterate the image matrices.
func generate32(fractal func(complex64, Config) color.Color, c Config) *image.RGBA {

	// Types.
	frame := float32(c.Frame)
	height := float32(c.Height * c.Smp_r)
	width := float32(c.Width * c.Smp_r)
	xmin := float32(c.Long) - frame/2
	ymin := float32(c.Lati) - frame/2

	// Scale for sampling and iterate.
	i := int(c.Height * c.Smp_r)
	j := int(c.Width * c.Smp_r)
	img := image.NewRGBA(image.Rect(0, 0, j, i))
	for py := 0; py < i; py++ {

		// Scale y to frame and position.
		y := float32(py)/height*frame + ymin
		for px := 0; px < j; px++ {

			// Scale x to frame and position.
			x := float32(px)/width*frame + xmin

			// Set the current location in a complex number and
			// call the given fractal function with that number.
			z := complex(x, y)
			img.Set(px, py, fractal(z, c))
		}
	}
	return img
}

// Set the dimensions and sub sampling for the programs fractal generation.
func Draw64(w io.Writer, fractal func(complex128, Config) color.Color, c Config) {
	var smp [BUF * BUF]color.RGBA

	// Generate scaled image.
	img := generate64(fractal, c)

	// Down sample.
	if c.Smp_r > 1 {
		img = subSample(img, smp[:], c)
	}

	// Encode.
	err := png.Encode(w, img)
	Check(err)
}

// Generate the fractal.
func generate64(fractal func(complex128, Config) color.Color, c Config) *image.RGBA {

	// Types.
	frame := c.Frame
	height := float64(c.Height * c.Smp_r)
	width := float64(c.Width * c.Smp_r)
	xmin := float64(c.Long) - frame/2
	ymin := float64(c.Lati) - frame/2

	// Scale for sampling and iterate.
	i := int(c.Height * c.Smp_r)
	j := int(c.Width * c.Smp_r)
	img := image.NewRGBA(image.Rect(0, 0, j, i))
	for py := 0; py < i; py++ {

		// Scale y to frame and position.
		y := float64(py)/height*frame + ymin
		for px := 0; px < j; px++ {

			// Scale x to frame and position.
			x := float64(px)/width*frame + xmin

			// Set the current location as a complex number and
			// call fractal function with that number.
			z := complex(x, y)
			img.Set(px, py, fractal(z, c))
		}
	}
	return img
}

// Set the dimensions and sub sampling for the programs fractal generation.
func DrawBIG(w io.Writer, fractal func(*cmpbig.ComplexBIG, Config) color.Color, c Config) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Generate scaled image.
	img := generateBIG(fractal, c)

	// Down sample.
	if c.Smp_r > 1 {
		img = subSample(img, smp[:], c)
	}

	// Encode.
	err := png.Encode(w, img)
	Check(err)
}

// Generate the fractal.
func generateBIG(fractal func(*cmpbig.ComplexBIG, Config) color.Color, c Config) *image.RGBA {

	// Types.
	fframe := c.Frame
	rframe := new(big.Float).SetFloat64(fframe)
	xmin := big.NewFloat(float64(c.Long) - fframe/2)
	ymin := big.NewFloat(float64(c.Lati) - fframe/2)

	x, y := new(big.Float), new(big.Float)
	div := new(big.Rat)
	z := &cmpbig.ComplexBIG{}
	z.Init()

	// Scale for sampling and iterate.
	i := int(c.Height * c.Smp_r)
	j := int(c.Width * c.Smp_r)
	img := image.NewRGBA(image.Rect(0, 0, j, i))
	for py := 0; py < i; py++ {

		// Scale y to frame and position.
		div.SetFrac64(int64(py), int64(i))
		y.SetRat(div)
		y.Mul(y, rframe)
		y.Add(y, ymin)
		for px := 0; px < j; px++ {

			// Scale x to frame and position.
			div.SetFrac64(int64(px), int64(j))
			x.SetRat(div)
			x.Mul(x, rframe)
			x.Add(x, xmin)

			// Set the current location as a complex number and
			// call fractal function with that number.
			z.X.Set(x)
			z.Y.Set(y)
			img.Set(px, py, fractal(z, c))

		}
	}
	return img
}

// Set the dimensions and sub sampling for the programs fractal generation.
func DrawRAT(w io.Writer, fractal func(*cmpbig.ComplexRAT, ConfigRAT) color.Color, c ConfigRAT) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Generate scaled image.
	img := generateRAT(fractal, c)

	// Down sample.
	if c.Smp_r > 1 {
		img = subSampleRAT(img, smp[:], c)
	}

	// Encode.
	err := png.Encode(w, img)
	Check(err)
}

func ratSetTypes(long, lati, frame, denom int64) (*big.Rat, *big.Rat, *big.Rat) {
	width := big.NewRat(frame*denom, denom)
	half := big.NewRat(int64(1), int64(2))
	xmin := big.NewRat(long, denom)
	ymin := big.NewRat(lati, denom)
	half.Mul(half, width)
	xmin.Sub(xmin, half)
	ymin.Sub(ymin, half)
	return xmin, ymin, width
}

// Generate the fractal.
func generateRAT(fractal func(*cmpbig.ComplexRAT, ConfigRAT) color.Color, c ConfigRAT) *image.RGBA {

	// Types.
	xmin, ymin, frame := ratSetTypes(c.Long, c.Lati, c.Frame, c.Denom)

	z := &cmpbig.ComplexRAT{}
	z.Init()

	// Scale for sampling and iterate.
	i := int64(c.Height * c.Smp_r)
	j := int64(c.Width * c.Smp_r)
	img := image.NewRGBA(image.Rect(0, 0, int(j), int(i)))
	for py := int64(0); py < i; py++ {

		z.Y.SetFrac64(py, i)
		z.Y.Mul(z.Y, frame)
		z.Y.Add(z.Y, ymin)
		for px := int64(0); px < j; px++ {

			z.X.SetFrac64(px, j)
			z.X.Mul(z.X, frame)
			z.X.Add(z.X, xmin)
			img.Set(int(px), int(py), fractal(z, c))
		}
	}
	return img
}

// Generate a sub image from the sampled data.
func subSample(img *image.RGBA, smp []color.RGBA, c Config) *image.RGBA {
	sub := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	x, y := 0, 0
	for py := 0; py < c.Height*c.Smp_r; py += c.Smp_r {
		for px := 0; px < c.Width*c.Smp_r; px += c.Smp_r {
			sub.Set(x, y, sample(img, smp[:], px, py, c.Smp_r))
			x++
		}
		x = 0
		y++
	}
	return sub
}

// Generate a sub image from the sampled data.
func subSampleRAT(img *image.RGBA, smp []color.RGBA, c ConfigRAT) *image.RGBA {
	sub := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	x, y := 0, 0
	for py := 0; py < c.Height*c.Smp_r; py += c.Smp_r {
		for px := 0; px < c.Width*c.Smp_r; px += c.Smp_r {
			sub.Set(x, y, sample(img, smp[:], px, py, c.Smp_r))
			x++
		}
		x = 0
		y++
	}
	return sub
}

// Sample a submarix of size fac^2.
func sample(img *image.RGBA, smp []color.RGBA, px, py, smp_r int) color.RGBA {
	count := 0
	for x := 0; x < smp_r; x++ {
		for y := 0; y < smp_r; y++ {
			smp[count] = img.RGBAAt(px+x, py+y)
			count++
		}
	}
	return colour.AvgRGBA(smp[:], uint(smp_r))
}

// The mandelbrot set.
func Mandelbrot32(z complex64, c Config) color.Color {
	var v complex64
	for i := int(0); i < c.Iter; i++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandelbrot set.
func Mandelbrot64(z complex128, c Config) color.Color {
	var v complex128
	for i := int(0); i < c.Iter; i++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandlebrot for big.Floats
func MandelbrotBIG(z *cmpbig.ComplexBIG, c Config) color.Color {
	v := &cmpbig.ComplexBIG{}
	v.Init()
	two := big.NewFloat(2.0)
	for i := int(0); i < c.Iter; i++ {
		v.Mul(v, v)
		v.Add(v, z)
		r := v.Abs()
		if r.Cmp(two) > 0 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandlebrot for big.Floats using Sqr()
func MandelbrotBIGmul(z *cmpbig.ComplexBIG, c Config) color.Color {
	v := &cmpbig.ComplexBIG{}
	v.Init()
	two := big.NewFloat(2.0)
	for i := int(0); i < c.Iter; i++ {
		v.Sqr()
		v.Add(v, z)
		r := v.Abs()
		if r.Cmp(two) > 0 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandlebrot for big.Rat
func MandelbrotRAT(z *cmpbig.ComplexRAT, c ConfigRAT) color.Color {
	v := &cmpbig.ComplexRAT{}
	v.Init()
	for i := int(0); i < c.Iter; i++ {
		v.Mul(v, v)
		v.Add(v, z)
		r := v.Abs()
		if r > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

//!-

// Some other interesting functions:

func Acos(z complex128, c Config) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func Sqrt(z complex128, c Config) color.Color {
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
func Newton(z complex128, c Config) color.Color {
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
