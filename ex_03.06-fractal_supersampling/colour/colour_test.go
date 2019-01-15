package colour

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"os"
	"testing"
	"time"
)

// 1531 is the `width` used for to scale the hue functions input. This very
// much guides the logic of both the function and its test, as such, it need be
// grasped in order that the reader understand the values used. 1536 is chosen
// as a multiple of both 6 and 256, permitting each segment of the switch to
// either gradually increase or decrease one of the three colours in all but
// the final of its 256 values. Only the last of the six parts, having no
// successor, includes the final value. Thus the `fence post` effect does not
// create flat spots, by way of duplicate values, upon the output spectrum.
// The output has also been designed such that any over or underflow of the
// given width are colored red and effectively hidden. The image and values
// produced by this test respond to 1536 and not the `functional` range that is
// given 1531, to encompass this over and underflow.

const MAX = 1531 // The range of the rgb output spectrum

func TestHue(t *testing.T) {
	var c, e color.RGBA
	var i uint8
	j := 0

	c = Hue(0, 0)
	e = color.RGBA{255, 0, 0, 255}
	if c != e {
		t.Errorf(`hue(%d, %d) == %v, wanted %v`, 0, MAX, c, e)
	}

	// Underflow the range by using a negative value, should be violet.
	c = Hue(-1, float64(MAX))
	e = color.RGBA{255, 0, 0, 255}
	if c != e {
		t.Errorf(`hue(%d, %d) == %v, wanted %v`, 0, MAX, c, e)
	}

	// Check all values in the normal range, j is the `value` that is being
	// output as a colour in the range MAX.
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{255, i, 0, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{255 - i, 255, 0, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{0, 255, i, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{0, 255 - i, 255, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{i, 0, 255, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	for i = 0; i < 255; i++ {
		c = Hue(float64(j), float64(MAX))
		e = color.RGBA{255, 0, 255 - i, 255}
		if c != e {
			t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
		}
		j++
	}
	// Add one for the final final fence post.
	c = Hue(float64(j), float64(MAX))
	e = color.RGBA{255, 0, 0, 255}
	if c != e {
		t.Errorf(`hue(%d, %d) == %v, wanted %v`, j, MAX, c, e)
	}

	// Overflow the range by the value, should be red.
	c = Hue(float64(MAX), float64(MAX))
	e = color.RGBA{255, 0, 0, 255}
	if c != e {
		t.Errorf(`hue(%d, %d) == %v, wanted %v`, MAX, MAX, c, e)
	}

	// Overflow the range by the value, should be red.
	c = Hue(float64(MAX+1), float64(MAX))
	e = color.RGBA{255, 0, 0, 255}
	if c != e {
		t.Errorf(`hue(%d, %d) == %v, wanted %v`, MAX+1, MAX, c, e)
	}
	//huePrint(MAX)
}

// When called this function will output to the terminal the full spectrum of
// color.RGBA structs created during the test in a descriptive text format.
func huePrint() {
	for i := 0; i < MAX; i++ {
		if i%255 == 0 {
			fmt.Println()
		}
		fmt.Printf("Hue(%d, %d) : %+v\n",
			i, MAX, Hue(float64(i), float64(MAX)))
	}
}

// Outputs a png of the colour spectrum produced, the file RGB.png is created
// and overwrites the existing file every time the test is run.
func TestHueSpectrumPng(t *testing.T) {
	const (
		width, height = 1536, 128
	)
	img := image.NewRGBA(image.Rect(-5, 0, width, height))
	for py := 0; py < height; py++ {
		// Left end of image or range underflow.
		for px := 0; px < width-5; px++ {
			img.Set(px, py, Hue(float64(px), float64(width-5)))
		}
		// Normal working range.
		for px := -5; px < 0; px++ {
			img.Set(px, py, Hue(float64(px), float64(width-5)))
		}
		// Right end of image, range overflow.
		for px := 1531; px < 1537; px++ {
			img.Set(px, py, Hue(float64(px), float64(width-5)))
		}
	}

	f, err := os.Create("RGB.png")
	check(err)
	defer f.Close()

	err = png.Encode(f, img)
	check(err)
}

// Runs a benchmark of the Hue function, for testing variants of the program.
func BenchmarkHue(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	for i := MAX; i < b.N; i++ {
		Hue(float64(rand.Intn(i)), float64(i))
	}
}
