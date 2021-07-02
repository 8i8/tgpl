package http

import (
	"Expr_web_calculator/eval"
	"Expr_web_calculator/svg"
	"errors"
	"fmt"
	"html/template"
	"io"
	"log"
	"math"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"strconv"
	"strings"
)

// Plotter is passed into the plot function/s and used buy the
// calculator to generate graphs.
type Plotter interface {
	Surface(io.Writer, func(x, y float64) float64)
}

var templ = template.Must(template.ParseFiles("./assets/templ.gohtml"))

var errVar = errors.New("undefined variable")

// data contains the data for html template ouput.
type data struct {
	Expr    string
	Val     interface{}
	X, Y, R string
	env     eval.Env
	List    []string
}

// URLEscExpr outputs the expression with any variables in an unescaped URL
// GET format, to postfix href links in anchor tag.
func (d data) URLEscExpr() string {
	buf := strings.Builder{}
	buf.WriteString("?expr=" + url.QueryEscape(d.Expr))
	if len(d.X) > 0 {
		buf.WriteString("&x=" + d.X)
	}
	if len(d.Y) > 0 {
		buf.WriteString("&y=" + d.Y)
	}
	if len(d.R) > 0 {
		buf.WriteString("&r=" + d.R)
	}
	return buf.String()
}

// prepareData fills a data struct with the required data for a plot.
func prepareData(req *http.Request, list ...string) data {
	return data{
		Expr: req.Form.Get("expr"),
		X:    req.Form.Get("x"),
		Y:    req.Form.Get("y"),
		R:    req.Form.Get("r"),
		List: list,
	}
}

// parseForm retrieves any expression in the request form,
func parseForm(res http.ResponseWriter, req *http.Request, tmpl string) (
	str string, done bool) {

	// Retrieve the expression if there is one, open the
	// basic page if there is not.
	req.ParseForm()
	str = req.Form.Get("expr")
	if len(str) == 0 {
		res.Header().Set("Content-Type", "text/html")
		templ.ExecuteTemplate(res, tmpl, nil)
		done = true
		return
	}
	if x := req.Form.Get("x"); len(x) != 0 {
		str += "&x=" + x
	}
	if y := req.Form.Get("y"); len(y) != 0 {
		str += "&y=" + y
	}
	if r := req.Form.Get("r"); len(r) != 0 {
		str += "&r=" + r
	}
	return
}

// divideExpression splis the give expression, seperating out any & key
// value pairs from the GET requests, then parsing the exprssion. It
// theb generates and runs an eval.CheckList to asecetain the existace
// of variables and the programs required running mode.
func divideExpression(res http.ResponseWriter, str string) (
	strs []string, expr eval.Expr, c *eval.CheckList, done bool) {

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
	// Remove the expression to leave only any key value pairs.
	strs = strs[1:]
	return
}

// setVariables sets any variables in the array of strings into the Form
// value map.
func setVariables(res http.ResponseWriter, req *http.Request, strs []string) {
	for i := 0; i < len(strs); i++ {
		str := strings.Split(strs[i], "=")
		if len(str) < 2 {
			http.Error(res, "not a valid variable: "+
				str[0], http.StatusNotAcceptable)
			return
		}
		req.Form.Set(str[0], str[1])
	}
}

// singleCalculation evaluate a simple calculator expression.
func singleCalculation(res http.ResponseWriter, req *http.Request,
	c *eval.CheckList) (env eval.Env, done bool) {
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

// isometricPlot evaluates the given expression, running it with an
// environment that consists of x, y and r variables. Plotting the
// result to a 3D surface.
func isometricPlot(res http.ResponseWriter,
	c *eval.CheckList, expr eval.Expr, p Plotter) {

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

// plot handles the iframe that is used by the calculator.
func plot(res http.ResponseWriter, req *http.Request) {
	const fname = "plot"
	p := svg.NewIsoSurface()

	// Prepare data.
	str, done := parseForm(res, req, fname)
	if done {
		return
	}

	strs, expr, c, done := divideExpression(res, str)
	if done {
		return
	}

	// Are we making a plot?
	mode, err := c.Mode()
	if err != nil {
		http.Error(res, err.Error(), http.StatusNotAcceptable)
		return
	}
	if mode == eval.Plot {
		isometricPlot(res, c, expr, p)
		return
	}

	// Set variables from get request if present.
	if len(strs) > 0 {
		setVariables(res, req, strs)
	}

	// Compute single value.
	env, done := singleCalculation(res, req, c)
	if done {
		return
	}

	// Load page.
	data := prepareData(req)
	data.Val = expr.Eval(env)
	mode, err = c.Mode()
	if err != nil {
		http.Error(res, err.Error(),
			http.StatusNotAcceptable)
		return
	}
	if mode == eval.Help || mode == eval.Helpful {
		res.Header().Set("Content-Type", "text/plain")
		res.Write([]byte(data.Val.(eval.Response).String()))
	} else {
		res.Header().Set("Content-Type", "text/html")
		templ.ExecuteTemplate(res, fname, data)
	}
}

func clear(res http.ResponseWriter, req *http.Request) {
	http.SetCookie(res, &http.Cookie{Name: "buffer", MaxAge: -1})
	res.Header().Set("Content-Type", "text/html")
	templ.ExecuteTemplate(res, "screen", nil)
}

// screen handles the intermediary calculator screen that wraps either a
// plot or a calculation, nested in the main page.
func screen(res http.ResponseWriter, req *http.Request) {
	const fname = "screen"

	str, done := parseForm(res, req, fname)
	if done {
		return
	}

	buf := eval.NewBuffer(15)
	cookie, err := req.Cookie("buffer")
	if err != nil && err != http.ErrNoCookie {
		http.Error(res, err.Error(),
			http.StatusInternalServerError)
		panic(err)
		return
	} else if err == http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:     "buffer",
			Path:     "/",
			Value:    str, // There is only one expression.
			HttpOnly: true,
			SameSite: http.SameSiteStrictMode,
		}
		if len(str) > 0 {
			buf.Add(str)
		}
	} else {
		if len(cookie.Value) > 0 {
			buf.Load(cookie.Value)
		}
		buf.MoveUp(str)
	}
	cookie.Value = buf.Unload()
	http.SetCookie(res, cookie)

	// Prepare data.
	strs, _, _, done := divideExpression(res, str)
	if done {
		return
	}

	// Set variables from get request if present.
	if len(strs) > 0 {
		setVariables(res, req, strs)
	}

	// Set data into struct for output.
	data := prepareData(req, buf.List()...)

	// Load page.
	res.Header().Set("Content-Type", "text/html")
	err = templ.ExecuteTemplate(res, fname, data)
	if err != nil {
		log.Fatal(err)
	}
}

// index handles the main calculator page.
func index(res http.ResponseWriter, req *http.Request) {
	const fname = "index"

	str, done := parseForm(res, req, fname)
	if done {
		return
	}

	// Prepare data.
	strs, _, _, done := divideExpression(res, str)
	if done {
		return
	}

	// Set variables from get request if present.
	if len(strs) > 1 {
		setVariables(res, req, strs)
	}

	// Load page.
	res.Header().Set("Content-Type", "text/html")
	err := templ.ExecuteTemplate(res, fname, prepareData(req))
	if err != nil {
		log.Fatal(err)
	}
}

func Serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/screen", screen)
	mux.HandleFunc("/screen/plot", plot)
	mux.HandleFunc("/screen/clear", clear)
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
