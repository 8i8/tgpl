package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"prettyPrint/eval"
	"strconv"
)

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

func expression(sc *bufio.Scanner, txt string) (eval.Expr, map[eval.Var]bool) {
	var expr eval.Expr
	var err error
	var vars map[eval.Var]bool
	fmt.Print(txt)
	for {
		sc.Scan()
		str := sc.Text()
		if str == "q" || str == "quit" {
			os.Exit(1)
		}
		expr, err = eval.Parse(str)
		if err != nil {
			fmt.Println(err)
			continue
		}
		vars = make(map[eval.Var]bool)
		err := expr.Check(vars)
		if err != nil {
			fmt.Println(err)
			continue
		}
		break
	}

	return expr, vars
}

func variables(sc *bufio.Scanner, vars map[eval.Var]bool) eval.Env {
	env := make(eval.Env)
	var val float64
	var err error
	for key, _ := range vars {
		fmt.Printf("%s = ", key)
		for {
			sc.Scan()
			val, err = strconv.ParseFloat(sc.Text(), 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			break
		}
		env[key] = val
	}
	return env
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	for {
		expr, vars := expression(sc, "~E: ")
		env := variables(sc, vars)
		x := expr.Eval(env)
		fmt.Printf("\t%f\n", x)
		fmt.Println(expr)
	}
}
