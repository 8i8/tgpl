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
		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 12, "y": 1}, "1729"},
		{"pow(x, 3) + pow(y, 3)", eval.Env{"x": 9, "y": 10}, "1729"},
		{"min(x, y)", eval.Env{"x": 9, "y": 10}, "9"},
		{"5 / 9 * (F - 32)", eval.Env{"F": -40}, "-40"},
		{"5 / 9 * (F - 32)", eval.Env{"F": 32}, "0"},
		{"5 / 9 * (F - 32)", eval.Env{"F": 212}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		inter, err := eval.Parse(test.expr)
		if err != nil {
			fmt.Print(err) // parse error
			continue
		}
		tmp := fmt.Sprint(inter)
		expr, err := eval.Parse(tmp)
		if err != nil {
			fmt.Print("error:", err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env))
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			fmt.Printf("error: %s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
