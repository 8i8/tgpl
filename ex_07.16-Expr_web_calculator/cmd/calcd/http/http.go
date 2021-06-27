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
	"strconv"
	"text/template"
)

type Plotter interface {
	Surface(io.Writer, func(x, y float64) float64)
}

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

var errVar = errors.New("undefined variable")

func calc(p Plotter) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Set up expression parser.
		req.ParseForm()
		str := req.Form.Get("expr")
		if len(str) == 0 {
			res.Header().Set("Content-Type", "text/html")
			templ.ExecuteTemplate(res, "main", nil)
			return
		}

		expr, err := eval.Parse(str)
		if err != nil {
			http.Error(res, err.Error(),
				http.StatusNotAcceptable)
			return
		}

		// Check which variables have been used.
		c := eval.NewCheckList()
		if err := expr.Check(c); err != nil {
			http.Error(res, err.Error(),
				http.StatusNotAcceptable)
			return
		}

		if c.Mode() == eval.Plot {
			// Set environment with given variables.
			for v := range c.Map() {
				if v.String() != "x" && v.String() != "y" &&
					v.String() != "r" {
					http.Error(res, errVar.Error()+": "+v.String(),
						http.StatusNotAcceptable)
					return
				}
			}

			res.Header().Set("Content-Type", "image/svg+xml")
			p.Surface(res, func(x, y float64) float64 {
				r := math.Hypot(x, y) // distance from (0,0)
				return expr.Eval(eval.Env{"x": x, "y": y, "r": r}).Float()
			})
			return
		}
		env := make(eval.Env)
		for v := range c.Map() {
			var x string
			if x = req.Form.Get(v.String()); len(x) == 0 {
				http.Error(res, errVar.Error()+": "+v.String(),
					http.StatusNotAcceptable)
				return
			}
			f, err := strconv.ParseFloat(x, 10)
			if err != nil {
				http.Error(res, err.Error(), http.StatusNotAcceptable)
			}
			env[v] = f
		}
		res.Header().Set("Content-Type", "text/html")
		templ.ExecuteTemplate(res, "main", expr.Eval(env).Float())
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
