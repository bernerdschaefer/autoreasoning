package main

import "fmt"

type (
	Expr interface {
		Simplify() Expr
	}

	Add   struct{ a, b Expr }
	Mult  struct{ a, b Expr }
	Const int
	Var   string
)

func (v Var) Simplify() Expr { return v }

func (c Const) Simplify() Expr { return c }

func isConst(e Expr) bool {
	_, ok := e.(Const)
	return ok
}

func (add Add) Simplify() Expr {
	var (
		a = add.a.Simplify()
		b = add.b.Simplify()
	)

	switch {
	case a == Const(0):
		return b
	case b == Const(0):
		return a
	case isConst(a) && isConst(b):
		return Const(a.(Const) + b.(Const))
	default:
		return Add{a, b}
	}
}

func (mult Mult) Simplify() Expr {
	var (
		a = mult.a.Simplify()
		b = mult.b.Simplify()
	)

	switch {
	case a == Const(0):
		return Const(0)
	case b == Const(0):
		return Const(0)
	case a == Const(1):
		return b
	case b == Const(1):
		return a
	case isConst(a) && isConst(b):
		return Const(a.(Const) * b.(Const))
	default:
		return Mult{a, b}
	}
}

func main() {
	var e Expr = Add{Mult{Add{Const(1), Mult{Const(0), Var("x")}}, Const(3)}, Const(12)}
	e = e.Simplify()

	switch e.(type) {
	case Const:
		fmt.Printf("%v\n", e)
	default:
		fmt.Printf("The expression could not be simplified to a constant\n")
	}
}
