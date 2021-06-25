package eval

import (
	"fmt"
	"math"
	"testing"
)

func TestEval(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{m: map[Var]float64{"A": 87616, "pi": math.Pi}}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 12, "y": 1}}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 9, "y": 10}}, "1729"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": -40}}, "-40"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 32}}, "0"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 212}}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env).Float())
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestStringer(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{m: map[Var]float64{"A": 87616, "pi": math.Pi}}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 12, "y": 1}}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 9, "y": 10}}, "1729"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": -40}}, "-40"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 32}}, "0"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 212}}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		inter, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		tmp := fmt.Sprint(inter)
		expr, err := Parse(tmp)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		got := fmt.Sprintf("%.6g", expr.Eval(test.env).Float())
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}

func TestEnv(t *testing.T) {
	tests := []struct {
		expr string
		env  Env
		want string
	}{
		{"sqrt(A / pi)", Env{m: map[Var]float64{"A": 87616, "pi": math.Pi}}, "167"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 12, "y": 1}}, "1729"},
		{"pow(x, 3) + pow(y, 3)", Env{m: map[Var]float64{"x": 9, "y": 10}}, "1729"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": -40}}, "-40"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 32}}, "0"},
		{"5 / 9 * (F - 32)", Env{m: map[Var]float64{"F": 212}}, "100"},
	}
	var prevExpr string
	for _, test := range tests {
		// Print expr only when it changes.
		if test.expr != prevExpr {
			fmt.Printf("\n%s\n", test.expr)
			prevExpr = test.expr
		}
		expr, err := Parse(test.expr)
		if err != nil {
			t.Error(err) // parse error
			continue
		}

		vars := make(map[Var]bool)
		err = expr.Check(vars)
		if err != nil {
			t.Error(err) // parse error
			continue
		}
		for v, _ := range test.env.m {
			if !vars[v] {
				t.Errorf("%s not found in vars\n", v)
				continue
			}
		}

		got := fmt.Sprintf("%.6g", expr.Eval(test.env).Float())
		fmt.Printf("\t%v => %s\n", test.env, got)
		if got != test.want {
			t.Errorf("%s.Eval() in %v = %q, want %q\n",
				test.expr, test.env, got, test.want)
		}
	}
}
