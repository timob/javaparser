package javaparser

import "github.com/timob/javabind"

type AstExprNullLiteralExprInterface interface {
	AstExprLiteralExprInterface
}

type AstExprNullLiteralExpr struct {
	AstExprLiteralExpr
}

// public com.github.javaparser.ast.expr.NullLiteralExpr()
func NewAstExprNullLiteralExpr() (*AstExprNullLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NullLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprNullLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NullLiteralExpr(int, int, int, int)
func NewAstExprNullLiteralExpr2(a int, b int, c int, d int) (*AstExprNullLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NullLiteralExpr", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprNullLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


