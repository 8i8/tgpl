package fractal

import (
	"8i8/cmpbig"
	"8i8/colour"
	"fmt"
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"math/big"
	"math/cmplx"
	"os"
	"strings"
	"testing"
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *  output on/off, check these settings!
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

var (
	// Simple Mandelbrot RGBA data. Smp_r set to 1, width and height 10x10
	// both at 3.
	s1_32  = "&{[255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 255 48 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255] 40 (0,0)-(10,10)}"
	s1_64  = "&{[255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 255 48 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255] 40 (0,0)-(10,10)}"
	s1_big = "&{[255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 255 48 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255] 40 (0,0)-(10,10)}"
	s1_rat = "&{[255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 0 0 0 255 255 48 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 48 0 255 255 48 0 255 0 0 0 255 0 0 0 255 255 64 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 16 0 255 255 16 0 255 255 24 0 255 255 40 0 255 255 136 0 255 255 24 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 16 0 255 255 16 0 255 255 16 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255 255 0 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 8 0 255 255 0 0 255 255 0 0 255] 40 (0,0)-(10,10)}"
	// Set to true to print out RGBA image data when test fails.
	e_32  = bool(false)
	e_64  = bool(false)
	e_big = bool(false)
	e_rat = bool(false)
	// Set to true to show z pixel position before Mandelbrot calculation.
	e1 = bool(false)
	// Set to true to print out calculation during Mandelbrot testing.
	e2_32  = bool(false)
	e2_64  = bool(true)
	e2_big = bool(false)
	e2_rat = bool(true)

	c0 = Config{
		Width:  3,
		Height: 3,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}

	c1 = Config{
		Width:  10,
		Height: 10,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}

	c2 = Config{
		Width:  120,
		Height: 120,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}

	r1 = ConfigRAT{
		Width:  120,
		Height: 120,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Denom:  4,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}
)

/* ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 * ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~ */

// Create a fractal using 32 bit floating point precision according to the
// specifications of the Config struct coded above.
func TestDraw32(t *testing.T) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Create file for outputting debugging info.
	w, err := os.Create("t32.out")
	Check(err)
	defer w.Close()

	// Generate a scaled image from both the test suit and the library.
	img1 := t_generate32(t_mandelbrot32, c1, w)
	img2 := generate32(Mandelbrot32, c1)

	// If required, down sample.
	if c1.Smp_r > 1 {
		img1 = t_subSample(img1, smp[:], c1)
		img2 = subSample(img2, smp[:], c1)
	}

	// Compare both images with the hard coded images above.
	// Test suit generated.
	s := fmt.Sprintf("%v", img1)
	if strings.Compare(s, s1_32) != 0 {
		if e_32 {
			t.Errorf(`result test Draw32 %v, wanted %v`, s, s1_32)
		} else {
			t.Errorf(`error: test Draw32 set e1 to view colour values.`)
		}
	}
	// Library generated.
	s = fmt.Sprintf("%v", img2)
	if strings.Compare(s, s1_32) != 0 {
		if e_32 {
			t.Errorf(`result lib Draw32 %v, wanted %v`, s, s1_32)
		} else {
			t.Errorf(`error: lib Draw32 set e1 to view colour values.`)
		}
	}
}

// Iterate over the image matracies.
func t_generate32(fractal func(complex64, Config, io.Writer) color.Color, c Config, w io.Writer) *image.RGBA {

	// Types
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
			// Set e1 to true to output debugging information to
			// the terminal.
			if e1 {
				fmt.Printf("32 : i%d,j%d %v\n", py, px, z)
			}
			img.Set(px, py, fractal(z, c, w))
		}
	}
	return img
}

// Create a fractal using 64 bit floating point precision according to the
// specifications of the Config struct coded above.
func TestDraw64(t *testing.T) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Create file for outputting debugging info.
	w, err := os.Create("t64.out")
	Check(err)
	defer w.Close()

	// Generate a scaled image from both the test suit and the library.
	img1 := t_generate64(t_mandelbrot64, c1, w)
	img2 := generate64(Mandelbrot64, c1)

	// If required, down sample.
	if c1.Smp_r > 1 {
		img1 = t_subSample(img1, smp[:], c1)
		img2 = t_subSample(img2, smp[:], c1)
	}

	// Compare both images with the hard coded images above.
	// Test suit generated.
	s := fmt.Sprintf("%v", img1)
	if strings.Compare(s, s1_64) != 0 {
		if e_64 {
			t.Errorf(`result test Draw64 %v, wanted %v`, s, s1_64)
		} else {
			t.Errorf(`error: test Draw64 set e1 to view colour values.`)
		}
	}
	// Library generated.
	s = fmt.Sprintf("%v", img2)
	if strings.Compare(s, s1_64) != 0 {
		if e_64 {
			t.Errorf(`result lib Draw64 %v, wanted %v`, s, s1_64)
		} else {
			t.Errorf(`error: lib Draw64 set e1 to view colour values.`)
		}
	}
}

// Generate the fractal.
func t_generate64(fractal func(complex128, Config, io.Writer) color.Color, c Config, w io.Writer) *image.RGBA {

	// Type
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

			// Set the current location in a complex number and
			// call the given fractal function with that number.
			z := complex(x, y)
			// Set e1 to true to output debugging information to
			// the terminal.
			if e1 {
				fmt.Printf("64 : i%d,j%d %v\n", py, px, z)
			}
			img.Set(px, py, fractal(z, c, w))
		}
	}
	return img
}

// Create a fractal using BIG bit floating point precision according to the
// specifications of the Config struct coded above.
func TestDrawBIG(t *testing.T) {

	// Buffer for sampling pixel colour.
	var smp [BUF * BUF]color.RGBA

	// Create file for outputting debugging info.
	w, err := os.Create("tBIG.out")
	Check(err)
	defer w.Close()

	// Generate a scaled image from both the test suit and the library.
	img1 := t_generateBIG(t_mandelbrotBIG, c1, w)
	img2 := generateBIG(MandelbrotBIG, c1)

	// If required, down sample.
	if c1.Smp_r > 1 {
		img1 = t_subSample(img1, smp[:], c1)
		img2 = subSample(img2, smp[:], c1)
	}

	// Compare both images with the hard coded images above.
	// Test suit generated.
	s := fmt.Sprintf("%v", img1)
	if strings.Compare(s, s1_big) != 0 {
		if e_big {
			t.Errorf(`result test DrawBIG %v, wanted %v`, s, s1_big)
		} else {
			t.Errorf(`error: test DrawBIG set e1 to view colour values.`)
		}
	}
	// Library generated.
	s = fmt.Sprintf("%v", img2)
	if strings.Compare(s, s1_big) != 0 {
		if e_big {
			t.Errorf(`result lib DrawBIG %v, wanted %v`, s, s1_big)
		} else {
			t.Errorf(`error: lib DrawBIG set e1 to view colour values.`)
		}
	}
}

// Generate the fractal.
func t_generateBIG(fractal func(*cmpbig.ComplexBIG, Config, io.Writer) color.Color, c Config, w io.Writer) *image.RGBA {

	// Type
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

			// Set the current location in a complex number and
			// call the given fractal function with that number.
			// Set e1 to true to output debugging information to
			// the terminal.
			z.X.Set(x)
			z.Y.Set(y)
			if e1 {
				fmt.Printf("BIG: i%d,j%d %v\n", py, px, z)
			}
			img.Set(px, py, fractal(z, c, w))
		}
	}
	return img
}

// Set the dimensions and sub sampling for the programs fractal generation.
// This test should not be run, although the code works it is inordinately slow,
// getting stuck in its own loop of infinite precision.
func testDrawRAT(t *testing.T) {

	var smp [BUF * BUF]color.RGBA

	w, err := os.Create("tRAT.out")
	Check(err)
	defer w.Close()

	img1 := t_generateRAT(t_mandelbrotRAT, r1, w)
	img2 := generateRAT(MandelbrotRAT, r1)

	if r1.Smp_r > 1 {
		img1 = t_subSampleRAT(img1, smp[:], r1)
		img2 = subSampleRAT(img2, smp[:], r1)
	}
	s := fmt.Sprintf("%v", img1)
	if strings.Compare(s, s1_rat) != 0 {
		if e_rat {
			t.Errorf(`result DrawRAT %v, wanted %v`, s, s1_rat)
		} else {
			t.Errorf(`error: DrawRAT set e1 in test file for more details`)
		}
	}
	s = fmt.Sprintf("%v", img2)
	if strings.Compare(s, s1_rat) != 0 {
		if e_rat {
			t.Errorf(`result DrawRAT %v, wanted %v`, s, s1_rat)
		} else {
			t.Errorf(`error: DrawRAT set e1 in test file for more details`)
		}
	}
}

// Generate the fractal.
func t_generateRAT(fractal func(*cmpbig.ComplexRAT, ConfigRAT, io.Writer) color.Color, c ConfigRAT, w io.Writer) *image.RGBA {

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
			if e1 {
				fmt.Fprintf(w, "RAT: i%d,j%d %v\n", py, px, z)
			}
			img.Set(int(px), int(py), fractal(z, c, w))
		}
	}
	return img
}

// Generate a sub image from the sampled data.
func t_subSample(img *image.RGBA, smp []color.RGBA, c Config) *image.RGBA {
	sub := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	x, y := 0, 0
	for py := 0; py < c.Height*c.Smp_r; py += c.Smp_r {
		for px := 0; px < c.Width*c.Smp_r; px += c.Smp_r {
			sub.Set(x, y, t_sample(img, smp[:], px, py, c.Smp_r))
			x++
		}
		x = 0
		y++
	}
	return sub
}

// Generate a sub image from the sampled data.
func t_subSampleRAT(img *image.RGBA, smp []color.RGBA, c ConfigRAT) *image.RGBA {
	sub := image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
	x, y := 0, 0
	for py := 0; py < c.Height*c.Smp_r; py += c.Smp_r {
		for px := 0; px < c.Width*c.Smp_r; px += c.Smp_r {
			sub.Set(x, y, t_sample(img, smp[:], px, py, c.Smp_r))
			x++
		}
		x = 0
		y++
	}
	return sub
}

// Sample a submarix of size fac^2.
func t_sample(img *image.RGBA, smp []color.RGBA, px, py, smp_r int) color.RGBA {
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
func t_mandelbrot32(z complex64, c Config, w io.Writer) color.Color {
	var v complex64
	var r float64
	for i := int(0); i < c.Iter; i++ {
		if e2_32 {
			fmt.Fprintf(w, "32 :i%d\n   v:   %+v\n   z:   %+v\n", i, v, z)
			v = v * v
			fmt.Fprintf(w, "   :v*v %+v %+v\n", v, z)
			v += z
			fmt.Fprintf(w, "   :v+z %+v %+v\n", v, z)
			r = cmplx.Abs(complex128(v))
			fmt.Fprintf(w, "   :v~  %+v %+v\n", v, z)
			fmt.Fprintf(w, "   :abs %+v bool %v\n", r, r > 2)
		} else {
			v = v*v + z
			r = cmplx.Abs(complex128(v))
		}
		if r > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandelbrot set.
func t_mandelbrot64(z complex128, c Config, w io.Writer) color.Color {
	var v complex128
	var r float64
	for i := int(0); i < c.Iter; i++ {
		if e2_64 {
			fmt.Fprintf(w, "64 :i%d\n   v:   %+v\n   z:   %+v\n", i, v, z)
			v = v * v
			fmt.Fprintf(w, "   :v*v %+v %+v\n", v, z)
			v += z
			fmt.Fprintf(w, "   :v+z %+v %+v\n", v, z)
			r = cmplx.Abs(v)
			fmt.Fprintf(w, "   :v~  %+v %+v\n", v, z)
			fmt.Fprintf(w, "   :abs %+v bool %v\n", r, r > 2)
		} else {
			v = v*v + z
			r = cmplx.Abs(v)
		}
		if r > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandlebrot for big.Floats
func t_mandelbrotBIG(z *cmpbig.ComplexBIG, c Config, w io.Writer) color.Color {
	v := &cmpbig.ComplexBIG{}
	v.Init()
	r := new(big.Float)
	two := big.NewFloat(2.0)
	for i := int(0); i < c.Iter; i++ {
		if e2_big {
			fmt.Fprintf(w, "BIG:i%d\n   v:   %+v\n   z:   %+v\n", i, v, z)
			v.Mul(v, v)
			fmt.Fprintf(w, "   :v*v %+v %+v\n", v, z)
			v.Add(v, z)
			fmt.Fprintf(w, "   :v+z %+v %+v\n", v, z)
			r = v.Abs()
			fmt.Fprintf(w, "   :v~  %+v %+v\n", v, z)
			fmt.Fprintf(w, "   :abs %+v bool %v\n", r, r.Cmp(two) > 0)
		} else {
			v.Mul(v, v)
			v.Add(v, z)
			r = v.Abs()
		}
		if r.Cmp(two) > 0 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

// The mandlebrot for big.Rat
func t_mandelbrotRAT(z *cmpbig.ComplexRAT, c ConfigRAT, w io.Writer) color.Color {
	v := &cmpbig.ComplexRAT{}
	v.Init()
	var r float64
	var b bool
	for i := int(0); i < c.Iter; i++ {
		if e2_rat {
			fmt.Fprintf(w, "RAT:i%d\n   v:   %s/%s\n   z:   %s/%s\n", i, v.X.FloatString(6), v.Y.FloatString(6), z.X.FloatString(6), z.X.FloatString(6))
			v.Mul(v, v)
			fmt.Fprintf(w, "   :v*v %v/%v %v/%v\n", v.X.FloatString(6), v.Y.FloatString(6), z.X.FloatString(6), z.X.FloatString(6))
			v.Add(v, z)
			fmt.Fprintf(w, "   :v+z %v/%v %v/%v\n", v.X.FloatString(6), v.Y.FloatString(6), z.X.FloatString(6), z.X.FloatString(6))
			r = v.Abs()
			fmt.Fprintf(w, "   :v~  %v %v/%v %v/%v\n", b, v.X.FloatString(6), v.Y.FloatString(6), z.X.FloatString(6), z.X.FloatString(6))
			fmt.Fprintf(w, "   :abs %v bool %v\n", r, r > 2)
		} else {
			v.Mul(v, v)
			v.Add(v, z)
			r = v.Abs()
		}
		if r > 2 {
			return colour.Hue(float64(c.Cont*i), c.Cspa)
		}
	}
	return color.RGBA{0, 0, 0, 255}
}

//!-

// Some other interesting functions:

func t_acos(z complex128, c Config) color.Color {
	v := cmplx.Acos(z)
	blue := uint8(real(v)*128) + 127
	red := uint8(imag(v)*128) + 127
	return color.YCbCr{192, blue, red}
}

func t_sqrt(z complex128, c Config) color.Color {
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
func t_newton(z complex128, c Config) color.Color {
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

func BenchmarkFractal32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Draw32(ioutil.Discard, Mandelbrot32, c1)
	}
}

func BenchmarkFractal64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Draw64(ioutil.Discard, Mandelbrot64, c1)
	}
}

func BenchmarkFractalBIG(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawBIG(ioutil.Discard, MandelbrotBIG, c1)
	}
}

func BenchmarkFractalRAT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawRAT(ioutil.Discard, MandelbrotRAT, r1)
	}
}
