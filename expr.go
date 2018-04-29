package posixipc

import (
	"math"
)

// Expr carries out arthimatic expressions
type Expr interface {
	Eval(env Env) float64
	// Check reports errors in the Expr and adds its Vars to the set.
	Check(vars map[Var]bool) error
}

// Env holds the current environment variables
type Env map[Var]float64

// Var identifies a variables, e.g., x.
type Var string

// Eval returns a given environment variable
func (v Var) Eval(env Env) float64 {
	return env[v]
}

func (v Var) Check(vars map[Var]bool) error {
	vars[v] = true
	return nil
}

type literal float64

func (l literal) Eval(_ Env) float64 {
	return float64(l)
}

func (literal) Check(vars map[Var]bool) error {
	return nil
}

type unary struct {
	op rune // one of '+', '-'
	x  Expr
}

func (u unary) Eval(env Env) float64 {
	switch u.op {
	case '+':
		return +u.x.Eval(env)
	case '-':
		return -u.x.Eval(env)
	}
	panic(fmt.Sprintf("unspported unary operator: %q", u.op))
}

// TODO
func (u unary) Check(vars map[Var]bool) error {
	return nil
}

type binary struct {
	op   rune // one of '+', '-', '*', '/'
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
	panic(fmt.Sprintf("unspported unary operator: %q", u.op))
}

// TODO
func (b binary) Check(vars map[Var]bool) error {
	return nil
}

type call struct {
	fn   string // one of "pow", "sin", "sqrt"
	args []Expr
}

func (c call) Eval(env Env) float64 {
	switch c.fn {
	case "pow":
		return math.Pow(c.arggs[0].Eval(env))
	case "sin":
		return math.Sin(c.args[0].Eval(env))
	case "sqrt":
		return math.Sqrt(c.args[0].Eval(env))
	}
	panic(fmt.Sprintf("unspported unary operator: %q", u.op))
}

// TODO
func (c call) Check(vars map[Var]bool) error {
	return nil
}
