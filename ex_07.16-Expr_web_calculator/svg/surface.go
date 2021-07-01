package svg

import (
	"fmt"
	"io"
	"math"
)

type Plotter interface {
	Surface(io.Writer, func(x, y float64) float64)
}

const (
	width, height = 600, 320 // canvas size in pixels
	cells         = 120      // number of grid cells
	xyrange       = 33       // axis ranges (-xyrange..+xyrange)
	//xyscale       = width / 2 / xyrange // pixels per x or y unit
	xyscale = 9 // pixels per x or y unit
	//zscale  = xyscale * 10 // pixels per z unit
	zscale = height * 0.4 // pixels per z unit
	angle  = math.Pi / 6  // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type isometric struct {
	width, height int
	cells         int
	xyrange       float64
	xyscale       float64
	zscale        float64
	angle         float64
	sin30, cos30  float64
}

func NewIsoSurface() Plotter {
	return isometric{
		width:   width,
		height:  height,
		cells:   cells,
		xyrange: xyrange,
		xyscale: xyscale,
		zscale:  zscale,
		angle:   angle,
		sin30:   math.Sin(angle),
		cos30:   math.Cos(angle),
	}
}

func (iso isometric) Surface(w io.Writer, fn func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", iso.width, iso.height)
	for i := 0; i < iso.cells; i++ {
		for j := 0; j < iso.cells; j++ {
			ax, ay := iso.corner(fn, i+1, j)
			bx, by := iso.corner(fn, i, j)
			cx, cy := iso.corner(fn, i, j+1)
			dx, dy := iso.corner(fn, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func (iso isometric) corner(fn func(x, y float64) float64, i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := iso.xyrange * (float64(i)/float64(iso.cells) - 0.5)
	y := iso.xyrange * (float64(j)/float64(iso.cells) - 0.5)

	// Compute surface height z.
	z := fn(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := float64(iso.width/2) + (x-y)*iso.cos30*iso.xyscale
	sy := float64(iso.height/2) + (x+y)*iso.sin30*iso.xyscale - z*iso.zscale
	return sx, sy
}
