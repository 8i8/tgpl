package http

import (
	"Expr_web_calculator/eval"
	"Expr_web_calculator/svg"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func calc(e eval.Expr, p svg.Plotter) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		// Retrieve data from request.
		req.ParseForm()
		expr := req.Form.Get("expr")

		// Set up expression parser.
		expr, err := eval.Parse(expr)
		env := eval.NewEnv()

		vars := map[eval.Var]bool
		err := e.Check(vars)
		if err != nil {
			http.Error(res, err, http.StatusNotAcceptable)
		}

		io.WriteString(res, "hello world")
	}
}

func Serve(e eval.Expr, p svg.Plotter) {
	mux := http.NewServeMux()
	mux.HandleFunc("/calc", calc(e, p))
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		log.Fatal(err)
	}

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Kill, os.Interrupt)
	<-sig
	fmt.Println("\nServer shutting down ...")
}
