package main

import (
	"Expr_cli_calculator/eval"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

var errEmpty = errors.New("empty string")

// expression takes input from the command line until an exceptible
// expressoin has been parsed.
func expression(sc *bufio.Scanner, prompt string) (
	eval.Expr, map[eval.Var]bool) {

	var expr eval.Expr
	var err error
	var vars map[eval.Var]bool
	for {
		fmt.Print(prompt)
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
		break
	}
	vars = make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return expr, vars
}

// cliArgExpression parses the given string as a mathmatical expression.
func cliArgExpression(exp string) (eval.Expr, map[eval.Var]bool) {
	expr, err := eval.Parse(exp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	vars := make(map[eval.Var]bool)
	err = expr.Check(vars)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return expr, vars
}

// variables sets any variables that do not already have a value by
// asking the user for input.
func variables(sc *bufio.Scanner, vars map[eval.Var]bool) eval.Env {
	env := eval.NewEnv()
	var val float64
	var err error
	for key, _ := range vars {
		for {
			fmt.Printf("%s = ", key)
			sc.Scan()
			val, err = strconv.ParseFloat(sc.Text(), 64)
			if err != nil {
				fmt.Println(err)
				continue
			}
			break
		}
		env.Add(key, val)
	}
	return env
}

// activeInput sets up an active command line data imput.
func activeInput(sc *bufio.Scanner) {
	expr, vars := expression(sc, "~E: ")
	env := variables(sc, vars)
	res := expr.Eval(env)
	fmt.Println(res)
}

// cliArgInput takes the first argument given on the command line after
// the program command and parses that as an expession.
func cliArgInput(sc *bufio.Scanner) {
	expr, vars := cliArgExpression(os.Args[1])
	env := variables(sc, vars)
	res := expr.Eval(env)
	fmt.Println(res)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	if len(os.Args) > 1 {
		cliArgInput(sc)
		return
	}
	for {
		activeInput(sc)
	}
}
