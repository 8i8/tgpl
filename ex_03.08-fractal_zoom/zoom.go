package main

import (
	"../ex_03.08-fractal_zoom/fractal"
	"os"
)

func main() {

	var a = 4

	f, err := os.Create("out.zoom.png")
	fractal.Check(err)
	defer f.Close()

	switch {
	case a == 1:
		fractal.Draw32(f, fractal.Mandelbrot32, f1)
	case a == 2:
		fractal.Draw64(f, fractal.Mandelbrot64, f1)
	case a == 3:
		fractal.DrawBIG(f, fractal.MandelbrotBIG, f1)
	case a == 4:
		fractal.DrawRAT(f, fractal.MandelbrotRAT, r1)
	}
}

var (
	f1 = fractal.Config{
		Width:  400,
		Height: 400,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}

	r1 = fractal.ConfigRAT{
		Width:  400,
		Height: 400,
		Smp_r:  1,
		Frame:  4,
		Long:   0,
		Lati:   0,
		Denom:  4,
		Iter:   200,
		Cont:   8,
		Cspa:   1531,
	}

	f2 = fractal.Config{
		Width:  600,
		Height: 600,
		Smp_r:  2,
		Frame:  1e-12,
		Long:   -56031461529402095E-17,
		Lati:   -640000039800001978E-18,
		Iter:   700,
		Cont:   4,
		Cspa:   4000,
	}
)
