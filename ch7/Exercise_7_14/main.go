package main

import (
	"fmt"
	"math"
	"strings"
)

type Env map[Var]float64

type Expr interface {
	Eval(env Env) float64
	Check(vars map[Var]bool) error
	String() string
}

type Var string

type literal float64

type unary struct {
	op rune
	x  Expr
}

type binary struct {
	op   rune
	x, y Expr
}

type call struct {
	fn   string
	args []Expr
}

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

func (l literal) Eval(env Env) float64 {
	return float64(l)
}

func (l literal) Check(vars map[Var]bool) error {
	return nil
}

func (l literal) String() string {
	return fmt.Sprintf("%g", float64(l))
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unsupported unary operator: %q", u.op))
}

func (u unary) Check(vars map[Var]bool) error {
	if u.op != '+' && u.op != '-' {
		return fmt.Errorf("unsupported unary operator: %q", u.op)
	}
	return u.x.Check(vars)
}

func (u unary) String() string {
	return fmt.Sprintf("(%c%s)", u.op, u.x)
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
	panic(fmt.Sprintf("unsupported binary operator: %q", b.op))
}

func (b binary) Check(vars map[Var]bool) error {
	if b.op != '+' && b.op != '-' && b.op != '*' && b.op != '/' {
		return fmt.Errorf("unsupported binary operator: %q", b.op)
	}
	if err := b.x.Check(vars); err != nil {
		return err
	}
	return b.y.Check(vars)
}

func (b binary) String() string {
	return fmt.Sprintf("(%s %c %s)", b.x, b.op, b.y)
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.args[0].Eval(env), c.args[1].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unsupported function call: %s", c.fn))
}

func (c call) Check(vars map[Var]bool) error {
	var argStrs []string
	for _, arg := range c.args {
		if err := arg.Check(vars); err != nil {
			return err
		}
		argStrs = append(argStrs, arg.String())
	}
	return nil
}

func (c call) String() string {
	var argStrs []string
	for _, arg := range c.args {
		argStrs = append(argStrs, arg.String())
	}
	return fmt.Sprintf("%s(%s)", c.fn, strings.Join(argStrs, ", "))
}

type Min struct {
	x, y Expr
}

func (m Min) Eval(env Env) float64 {
	return math.Min(m.x.Eval(env), m.y.Eval(env))
}

func (m Min) Check(vars map[Var]bool) error {
	if err := m.x.Check(vars); err != nil {
		return err
	}
	if err := m.y.Check(vars); err != nil {
		return err
	}
	return nil
}

func (m Min) String() string {
	return fmt.Sprintf("min(%s, %s)", m.x, m.y)
}

func main() {
	expr := Min{
		x: literal(5),
		y: binary{
			op: '+',
			x:  Var("x"),
			y:  Var("y"),
		},
	}

	fmt.Println(expr.String())
}
