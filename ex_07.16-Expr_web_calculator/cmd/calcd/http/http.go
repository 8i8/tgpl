package http

import (
	"Expr_web_calculator/eval"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"text/template"
)

type Plotter interface {
	Surface(io.Writer, func(x, y float64) float64)
}

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

var errVar = errors.New("undefined variable")

func plot(res http.ResponseWriter, req *http.Request, plot Plotter, str string) {
	expr, err := eval.Parse(str)
	if err != nil {
		http.Error(res, err.Error(),
			http.StatusNotAcceptable)
		return
	}

	// Check which variables have been used.
	vars := make(eval.Vars)
	if err := expr.Check(vars); err != nil {
		http.Error(res, err.Error(),
			http.StatusNotAcceptable)
		return
	}

	// Set environment with given variables.
	for v := range vars {
		if v.String() != "x" && v.String() != "y" &&
			v.String() != "r" {
			http.Error(res, errVar.Error()+": "+v.String(),
				http.StatusNotAcceptable)
			return
		}
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	plot.Surface(res, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r}).Float()
	})
}

func calc(p Plotter) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Set up expression parser.
		req.ParseForm()
		if str := req.Form.Get("expr"); len(str) > 0 {
			plot(res, req, p, str)
			return
		}
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := templ.ExecuteTemplate(w, "main", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Serve(plot Plotter) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/calc", calc(plot))
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
