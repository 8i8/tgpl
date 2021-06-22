package eval

import (
	"fmt"
	"math"
	"strings"
)

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

// A Var indentifies a variable, e.g., x.
type Var string

// Env maps envirnoment variable names to values.
type Env map[Var]float64

func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

func (v Var) String() string {
	return string(v)
}

// A literal is a numeric constant, e.g., 3.141.
type literal float64

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%.2f", float64(l))
}

// A unary represents a unary operator expression, e.g., -x.
type unary struct {
	op rune // one of '+', '-'.
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupparted unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected enary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/'.
	x, y Expr
}

func (b binary) Eval(env Env) float64 {
	switch b.op {
	case '+':
		return b.x.Eval(env) + b.y.Eval(env)
	case '-':
		return b.x.Eval(env) - b.y.Eval(env)
	case '*':
		return b.x.Eval(env) * b.y.Eval(env)
	case '/':
		return b.x.Eval(env) / b.y.Eval(env)
	}
	panic(fmt.Sprintf("unsupparted binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/", b.op) {
		return fmt.Errorf("unsupported binary op %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) String() string {
	return fmt.Sprintf("%s%c%s", b.x.String(), b.op, b.y.String())
}

// A call represents a function call expression, e.g., sin(x).
type call struct {
	fn   string // one of "pow", "sin", "sqrt".
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "cos":
		return math.Cos(c.args[0].Eval(env))
	case "tan":
		return math.Tan(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

var numParams = map[string]int{"pow": 2, "sin": 1, "cos": 1, "tan": 1, "sqrt": 1}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := numParams[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity)
	}
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
	}
	return nil
}

func (c call) String() string {
	switch c.fn {
	case "pow":
		return fmt.Sprintf("%s(%s, %s)", c.fn, c.args[0].String(),
			c.args[1].String())
	default:
		return fmt.Sprintf("%s(%s)", c.fn, c.args[0].String())
	}
	return ""
}

// A bracket is an empty place holder when parsing an expression, it is
// required to maintain information about the position of brackets in an
// input expression, enabling the correct functioning of the Stringer
// interface for pretty printing an expression.
type bracket struct {
	args []Expr
}

func (b bracket) Eval(env Env) (f float64) {
	for i := range b.args {
		f = b.args[i].Eval(env)
	}
	return
}

func (b bracket) Check(vars map[Var]bool) error {
	return nil
}

func (b bracket) String() string {
	buf := strings.Builder{}
	buf.WriteByte('(')
	for i := range b.args {
		buf.WriteString(b.args[i].String())
	}
	buf.WriteByte(')')
	return buf.String()
}
