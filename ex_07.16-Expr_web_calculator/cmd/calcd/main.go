package main

import (
	"Expr_web_calculator/cmd/calcd/http"
	"Expr_web_calculator/eval"
	"Expr_web_calculator/svg"
)

func main() {
	var e eval.Expr
	var p svg.Plotter
	http.Serve(e, p)
}
