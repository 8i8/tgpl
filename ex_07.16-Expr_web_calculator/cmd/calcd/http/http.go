package http

import (
	"Expr_web_calculator/eval"
	"errors"
	"fmt"
	"html"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

type Plotter interface {
	Surface(io.Writer, func(x, y float64) float64)
}

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

var errVar = errors.New("undefined variable")

type data struct {
	Expr    string
	Val     interface{}
	X, Y, R string
	env     eval.Env
}

func (d data) SetExpr() string {
	d.Expr = html.UnescapeString(d.Expr)
	if len(d.X) > 0 {
		d.Expr += "&x=" + d.X
	}
	if len(d.Y) > 0 {
		d.Expr += "&y=" + d.Y
	}
	if len(d.R) > 0 {
		d.Expr += "&r=" + d.R
	}
	return d.Expr
}

func getData(req *http.Request) data {
	return data{
		Expr: req.Form.Get("expr"),
		X:    req.Form.Get("x"),
		Y:    req.Form.Get("y"),
		R:    req.Form.Get("r"),
	}
}

func parseExprssion(res http.ResponseWriter, req *http.Request) (
	strs []string, expr eval.Expr, c *eval.Check, done bool) {

	// Retrieve the expression if there is one, open the
	// basic page if there is not.
	req.ParseForm()
	str := req.Form.Get("expr")
	if len(str) == 0 {
		res.Header().Set("Content-Type", "text/html")
		templ.ExecuteTemplate(res, "screen", getData(req))
		done = true
		return
	}
	strs = strings.Split(str, "&")

	// Parser expression.
	expr, err := eval.Parse(strs[0])
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotAcceptable)
		done = true
		return
	}

	// Check which variables have been used.
	c = eval.NewCheckList()
	if err := expr.Check(c); err != nil {
		http.Error(res, err.Error(), http.StatusNotAcceptable)
		done = true
		return
	}
	return
}

func plot(res http.ResponseWriter, c *eval.Check, expr eval.Expr, p Plotter) {

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

func setVariables(res http.ResponseWriter, req *http.Request, strs []string) {
	for i := 1; i < len(strs); i++ {
		str := strings.Split(strs[i], "=")
		if len(str) < 2 {
			http.Error(res, "not a valid variable: "+
				str[0], http.StatusNotAcceptable)
			return
		}
		req.Form.Set(str[0], str[1])
	}
}

func singleCalculation(res http.ResponseWriter, req *http.Request,
	c *eval.Check) (env eval.Env, done bool) {
	env = make(eval.Env)
	for v := range c.Map() {
		var x string
		if x = req.Form.Get(v.String()); len(x) == 0 {
			http.Error(res, errVar.Error()+": "+v.String(),
				http.StatusNotAcceptable)
			done = true
			return
		}
		if x != "" {
			f, err := strconv.ParseFloat(x, 10)
			if err != nil {
				exp, err := eval.Parse(x)
				if err != nil {
					http.Error(res, err.Error(),
						http.StatusNotAcceptable)
					done = true
					return
				}
				f = exp.Eval(make(eval.Env)).Float()
			}
			env[v] = f
		}
	}
	return
}

func screen(p Plotter) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Prepare data.
		strs, expr, c, done := parseExprssion(res, req)
		if done {
			return
		}

		// Are we making a plot?
		mode, err := c.Mode()
		if err != nil {
			http.Error(res, err.Error(),
				http.StatusNotAcceptable)
			return
		}
		if mode == eval.Plot {
			plot(res, c, expr, p)
			return
		}

		// Set variables from get request if present.
		if len(strs) > 1 {
			setVariables(res, req, strs)
		}

		// Compute single value.
		env, done := singleCalculation(res, req, c)
		if done {
			return
		}

		// Load page.
		data := getData(req)
		data.Val = expr.Eval(env)
		// Are we making a plot?
		mode, err = c.Mode()
		if err != nil {
			http.Error(res, err.Error(), http.StatusNotAcceptable)
			return
		}
		if mode == eval.Help || mode == eval.Helpful {
			res.Header().Set("Content-Type", "text/plain")
			res.Write([]byte(data.Val.(eval.Response).String()))
		} else {
			res.Header().Set("Content-Type", "text/html")
			templ.ExecuteTemplate(res, "screen", data)
		}
	}
}

func index(res http.ResponseWriter, req *http.Request) {

	// Prepare data.
	strs, _, _, done := parseExprssion(res, req)
	if done {
		return
	}

	// Set variables from get request if present.
	if len(strs) > 1 {
		setVariables(res, req, strs)
	}

	// Load page.
	res.Header().Set("Content-Type", "text/html")
	err := templ.ExecuteTemplate(res, "main", getData(req))
	if err != nil {
		log.Fatal(err)
	}
}

func Serve(plot Plotter) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/calc", index)
	mux.HandleFunc("/screen", screen(plot))
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
