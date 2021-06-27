package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"prettyPrint/eval"
	"prettyPrint/plot"
	"text/template"
)

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

var errEmpty = errors.New("empty string")

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, errEmpty
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func draw(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if errors.Is(err, errEmpty) {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	plot.Surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := templ.ExecuteTemplate(w, "main", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/plot", draw)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
