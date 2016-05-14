package javaparser

import "github.com/timob/javabind"

type AstExprLongLiteralMinValueExprInterface interface {
	AstExprLongLiteralExprInterface
}

type AstExprLongLiteralMinValueExpr struct {
	AstExprLongLiteralExpr
}

// public com.github.javaparser.ast.expr.LongLiteralMinValueExpr()
func NewAstExprLongLiteralMinValueExpr() (*AstExprLongLiteralMinValueExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LongLiteralMinValueExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprLongLiteralMinValueExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.LongLiteralMinValueExpr(int, int, int, int)
func NewAstExprLongLiteralMinValueExpr2(a int, b int, c int, d int) (*AstExprLongLiteralMinValueExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LongLiteralMinValueExpr", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprLongLiteralMinValueExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


