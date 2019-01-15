package orchid

import (
	"fmt"
	"io"
	"math"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 50                  // number of cells in grid
	xyrange       = 10.0                // axis range
	xyscale       = width / 2 / xyrange // pixels per x y unit
	zscale        = height * 0.2        // pixels per z unit
	angle         = math.Pi / 6         // angle of x y axes = 30 deg
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func Orchid(out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			e, ax, ay := corner(i+1, j)
			f, bx, by := corner(i, j)
			g, cx, cy := corner(i, j+1)
			h, dx, dy := corner(i+1, j+1)
			colour := zDepthColour(e+f+g+h/4, zscale)
			fmt.Fprintf(out, "<polygon "+
				"points='%g, %g, %g, %g, %g, %g, %g, %g' "+
				"style='stroke: grey; fill: %s; "+
				"stroke-width: 0.4' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, colour)
		}
	}
	fmt.Fprintf(out, "</svg>")
}

func zDepthColour(i float64, scale float64) string {
	i = (i + 5) / 10
	i *= 255
	out := fmt.Sprintf("#%2x00%2x", uint8(i), uint8(255-i))
	return out
}

func corner(i, j int) (float64, float64, float64) {
	//find point (x,y) at the corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y) / 2

	// Project (x,y,z) isometricaly onto 2D svg canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return z, sx, sy
}

func f(x, y float64) float64 {
	x += 2
	x = math.Sin(x)
	y += -1
	y = math.Sin(y)
	return x - y
}
