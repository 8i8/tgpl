package main

import (
	"../ex_03.07-fractal_newtons_colours/fractal"
	"os"
)

func main() {
	fractal.Draw(os.Stdout, fractal.Newton)
}
