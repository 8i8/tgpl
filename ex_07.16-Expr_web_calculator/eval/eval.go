package eval

import (
	"fmt"
	"math"
	"strings"
)

type T int

const (
	tString T = iota
	tFloat64
)

type Response struct {
	Type  T
	value interface{}
}

func (r Response) String() string {
	switch r.Type {
	case tString:
		return r.value.(string)
	case tFloat64:
		return fmt.Sprintf("%f", r.value.(float64))
	}
	return ""
}

func (r Response) Float() float64 {
	return r.value.(float64)
}

type Expr interface {
	Eval(env Env) Response
	Check(vars map[Var]bool) error
	String() string
}

// A Var indentifies a variable, e.g., x.
type Var string

// Env maps envirnoment variable names to values.
type Env map[Var]float64

// Vars keeps track of which variables have been checked.
type Vars map[Var]bool

func (v Var) Eval(env Env) Response {
	return Response{tFloat64, env[v]}
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

func (l literal) Eval(env Env) Response {
	return Response{tFloat64, float64(l)}
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

func (u unary) Eval(env Env) Response {
	val := u.x.Eval(env)
	switch u.op {
	case '+':
		return Response{tFloat64, +val.value.(float64)}
	case '-':
		return Response{tFloat64, -val.value.(float64)}
	}
	panic(fmt.Sprintf("unsupparted unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-", u.op) {
		return fmt.Errorf("unexpected unary op %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("%c%s", u.op, u.x)
}

// A binary represents a binary operator expression, e.g., x+y.
type binary struct {
	op   rune // one of '+', '-', '*', '/', '%'.
	x, y Expr
}

func (b binary) Eval(env Env) Response {
	x := b.x.Eval(env).Float()
	y := b.y.Eval(env).Float()
	switch b.op {
	case '+':
		return Response{tFloat64, x + y}
	case '-':
		return Response{tFloat64, x - y}
	case '*':
		return Response{tFloat64, x * y}
	case '/':
		return Response{tFloat64, x / y}
	case '^':
		return Response{tFloat64, math.Pow(x, y)}
	}
	panic(fmt.Sprintf("unsupparted binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if !strings.ContainsRune("+-*/^", b.op) {
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

func (c call) Eval(env Env) Response {
	args := make([]float64, len(c.args))
	for i := range c.args {
		args[i] = c.args[i].Eval(env).value.(float64)
	}
	switch c.fn {
	case "cos":
		return Response{tFloat64, math.Cos(args[0])}
	case "asin":
		return Response{tFloat64, math.Asin(args[0])}
	case "acos":
		return Response{tFloat64, math.Acos(args[0])}
	case "atan":
		return Response{tFloat64, math.Atan(args[0])}
	case "asinh":
		return Response{tFloat64, math.Asinh(args[0])}
	case "acosh":
		return Response{tFloat64, math.Acosh(args[0])}
	case "atanh":
		return Response{tFloat64, math.Atanh(args[0])}
	case "atan2":
		return Response{tFloat64, math.Atan2(args[0], args[1])}
	case "abs":
		return Response{tFloat64, math.Abs(args[0])}
	case "ceil":
		return Response{tFloat64, math.Ceil(args[0])}
	case "cbrt":
		return Response{tFloat64, math.Cbrt(args[0])}
	case "copysign":
		return Response{tFloat64, math.Copysign(args[0], args[1])}
	case "dim":
		return Response{tFloat64, math.Dim(args[0], args[1])}
	case "exp":
		return Response{tFloat64, math.Exp(args[0])}
	case "exp2":
		return Response{tFloat64, math.Exp2(args[0])}
	case "expm1":
		return Response{tFloat64, math.Expm1(args[0])}
	case "FMA":
		return Response{tFloat64, math.FMA(args[0], args[1], args[2])}
	case "floor":
		return Response{tFloat64, math.Floor(args[0])}
	case "gamma":
		return Response{tFloat64, math.Gamma(args[0])}
	case "hypot":
		return Response{tFloat64, math.Hypot(args[0], args[1])}
	case "inf":
		return Response{tFloat64, math.Inf(int(args[0]))}
	case "J0":
		return Response{tFloat64, math.J0(args[0])}
	case "J1":
		return Response{tFloat64, math.J1(args[0])}
	case "Jn":
		return Response{tFloat64, math.Jn(int(args[0]), args[1])}
	case "ldexp":
		return Response{tFloat64, math.Ldexp(args[0], int(args[0]))}
	case "log":
		return Response{tFloat64, math.Log(args[0])}
	case "log10":
		return Response{tFloat64, math.Log10(args[0])}
	case "log1p":
		return Response{tFloat64, math.Log1p(args[0])}
	case "log2":
		return Response{tFloat64, math.Log2(args[0])}
	case "logb":
		return Response{tFloat64, math.Logb(args[0])}
	case "max":
		return Response{tFloat64, math.Max(args[0], args[1])}
	case "min":
		return Response{tFloat64, math.Min(args[0], args[1])}
	case "mod":
		return Response{tFloat64, math.Mod(args[0], args[1])}
	case "nextafter":
		return Response{tFloat64, math.Nextafter(args[0], args[1])}
	case "pow":
		return Response{tFloat64, math.Pow(args[0], args[1])}
	case "pow10":
		return Response{tFloat64, math.Pow10(int(args[0]))}
	case "remainder":
		return Response{tFloat64, math.Remainder(args[0], args[1])}
	case "round":
		return Response{tFloat64, math.Round(args[0])}
	case "roundtoeven":
		return Response{tFloat64, math.RoundToEven(args[0])}
	case "sin":
		return Response{tFloat64, math.Sin(args[0])}
	case "sinh":
		return Response{tFloat64, math.Sinh(args[0])}
	case "sqrt":
		return Response{tFloat64, math.Sqrt(args[0])}
	case "tan":
		return Response{tFloat64, math.Tan(args[0])}
	case "tanh":
		return Response{tFloat64, math.Tanh(args[0])}
	case "trunc":
		return Response{tFloat64, math.Trunc(args[0])}
	case "Y0":
		return Response{tFloat64, math.Y0(args[0])}
	case "Y1":
		return Response{tFloat64, math.Y1(args[0])}
	case "Yn":
		return Response{tFloat64, math.Yn(int(args[0]), args[1])}
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	arity, ok := fnData[c.fn]
	if !ok {
		return fmt.Errorf("unknown function %q", c.fn)
	}
	if len(c.args) != arity.params {
		return fmt.Errorf("call to %s has %d args, want %d",
			c.fn, len(c.args), arity.params)
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
	case "FMA":
		return fmt.Sprintf("%s(%s, %s, %s)", c.fn, c.args[0].String(),
			c.args[1].String(), c.args[2].String())
	case "atan2", "copysign", "dim", "Jn", "ldexp", "max", "min",
		"mod", "nextafter", "pow", "remainder", "Yn":
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

func (b bracket) Eval(env Env) (r Response) {
	for i := range b.args {
		r = b.args[i].Eval(env)
	}
	return
}

func (b bracket) Check(vars map[Var]bool) error {
	for i := range b.args {
		b.args[i].Check(vars)
	}
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

// mode sets a different running mode on the calculator, used primarily
// to redirect output type.
type mode struct {
	id   rmode
	args []Expr
}

func (m mode) Eval(env Env) (r Response) {
	for i := range m.args {
		r = m.args[i].Eval(env)
	}
	if m.id == plot {
	}
	return
}

func (m mode) Check(vars map[Var]bool) error {
	for i := range m.args {
		m.args[i].Check(vars)
	}
	return nil
}

func (m mode) String() string {
	buf := strings.Builder{}
	buf.WriteString(m.id.String())
	buf.WriteByte('(')
	for i := range m.args {
		buf.WriteString(m.args[i].String())
	}
	buf.WriteByte(')')
	return buf.String()
}

// helpout sets help output data for a given calulator function.
type helpout struct {
	mode  rmode
	usage string
}

func (h helpout) Eval(env Env) Response {
	return Response{tString, h.String()}
}

func (h helpout) Check(vars map[Var]bool) error {
	return nil
}

func (h helpout) String() string {
	args := strings.Split(h.usage, ",")
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	buf := strings.Builder{}
	for _, arg := range args {
		for key, val := range fnData {
			if h.mode == help && key == arg {
				printhelp(&buf, key, val)
			}
			if h.mode == helpful && strings.Contains(key, arg) {
				printhelpful(&buf, key, val)
			}
		}
	}
	return buf.String()
}

func printhelpful(buf *strings.Builder, key string, val fn) {
	v := []string{"x", "y", "z"}
	buf.WriteString(key + "(")
	buf.WriteString(v[0])
	for i := 1; i < val.params; i++ {
		buf.WriteString(", ")
		buf.WriteString(v[i])
	}
	buf.WriteString(") " + val.usage + "\n")
}

func printhelp(buf *strings.Builder, key string, val fn) {
	v := []string{"x", "y", "z"}
	buf.WriteString(key + "(")
	buf.WriteString(v[0])
	for i := 1; i < val.params; i++ {
		buf.WriteString(", " + v[i])
	}
	buf.WriteString(") " + val.usage + val.special + "\n")
}
