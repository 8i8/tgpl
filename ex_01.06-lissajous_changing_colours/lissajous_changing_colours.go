package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

//!-main
// Packages not needed by version in book.
// TODO this program is not working, possible du to web serving functionality
import (
	"log"
	"net/http"
	"time"
)

//!+main

var palette = []color.Color{
	color.RGBA{0, 0, 0, 255},
	color.RGBA{0, 100, 0, 255},
	color.RGBA{0, 125, 0, 255},
	color.RGBA{0, 150, 0, 255},
	color.RGBA{0, 175, 0, 255},
	color.RGBA{0, 200, 0, 255},
	color.RGBA{0, 225, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 225, 0, 255},
	color.RGBA{0, 200, 0, 255},
	color.RGBA{0, 175, 0, 255},
	color.RGBA{0, 150, 0, 255},
	color.RGBA{0, 125, 0, 255},
	color.RGBA{0, 100, 0, 255},
}

//var palette = []color.Color{color.White, color.Black}

const (
	backgroundIndex = 0 // first color in palette
)

func main() {
	/*
	 * !-main
	 * The sequence of images is deterministic unless we seed
	 * the pseudo-random number generator using the current time.
	 * Thanks to Randall McPherson for pointing out the omission.
	 */
	rand.Seed(time.Now().UTC().UnixNano())

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			lissajous(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//!+main
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5      // number of complete x oscillator revolutions
		res     = 0.0001 // angular resolution
		size    = 100    // image canvas covers [-size..+size]
		nframes = 64     // number of animation frames
		delay   = 8      // delay between frames in 10ms units
	)
	freq := rand.Float64() * 3.0 // relative frequency of y oscillator
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0 // phase difference
	var colorIndex uint8 = 1
	var l uint8 = uint8(len(palette))
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				colorIndex)
		}
		phase += 0.1

		colorIndex %= l - 1
		colorIndex++
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}
