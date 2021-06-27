package main

import (
	"Expr_web_calculator/cmd/calcd/http"
	"Expr_web_calculator/svg"
)

func main() {
	http.Serve(svg.NewIsoSurface())
}
