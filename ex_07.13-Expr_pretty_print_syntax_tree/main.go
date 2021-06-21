package main

import (
	"fmt"
	"math"
	"prettyPrint/eval"
)

func main() {
	tests := []struct {
		expr string
		env  eval.Env
		want string
	}{
		{"sqrt(A / pi)", eval.Env{"A": 87616, "pi": math.Pi}, "167"},
		{"pow(x, 3.1) + pow(y, 3)", eval.Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 9, "y": 10}, "1729"},
		{"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", eval.Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", eval.Env{"F": 212}, "100"},
	}
	//var prevExpr string
	for _, test := range tests {
		expr, err := eval.Parse(test.expr)
		if err != nil {
			fmt.Print(err) // parse error
			continue
		}
		fmt.Println(test.expr)
		fmt.Println(expr, "\n")
	}
}
