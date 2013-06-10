package script

import (
	"go/ast"
	"reflect"
)

func (w *World) compileCallExpr(n *ast.CallExpr) Expr {
	// only call idents for now
	id, ok := (n.Fun).(*ast.Ident)
	if !ok {
		panic(err(n.Pos(), "not allowed:", typ(n.Fun)))
	}
	// only call reflectFunc for now
	f, ok2 := w.resolve(n.Pos(), id.Name).(*function)
	if !ok2 {
		panic(err(n.Pos(), "can not call", id.Name))
	}
	// check args count. no strict check for varargs
	variadic := f.Type().IsVariadic()
	if !variadic && len(n.Args) != f.NumIn() {
		panic(err(n.Pos(), id.Name, "needs", f.NumIn(), "arguments, got", len(n.Args))) // TODO: varargs
	}
	// convert args
	args := make([]Expr, len(n.Args))
	for i := range args {
		if variadic {
			args[i] = w.compileExpr(n.Args[i]) // no type check or conversion
		} else {
			args[i] = typeConv(n.Args[i].Pos(), w.compileExpr(n.Args[i]), f.In(i))
		}
	}
	return &call{*f, args}
}

type call struct {
	f    function
	args []Expr
}

func (c *call) Eval() interface{} {
	argv := make([]reflect.Value, len(c.args))

	for i := range c.args {
		argv[i] = evalue(c.args[i])
	}

	ret := c.f.Call(argv)
	assert(len(ret) <= 1)
	if len(ret) == 0 {
		return nil
	} else {
		return ret[0].Interface()
	}
}

func (c *call) Type() reflect.Type { return c.f.ReturnType() }

// eval and return as value
func evalue(e Expr) reflect.Value {
	return reflect.ValueOf(e.Eval()) // later we can typeswitch on EvalValue() and avoid unbox+box
}
