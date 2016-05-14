package javaparser

import "github.com/timob/javabind"

type AstExprIntegerLiteralMinValueExprInterface interface {
	AstExprIntegerLiteralExprInterface
}

type AstExprIntegerLiteralMinValueExpr struct {
	AstExprIntegerLiteralExpr
}

// public com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr()
func NewAstExprIntegerLiteralMinValueExpr() (*AstExprIntegerLiteralMinValueExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprIntegerLiteralMinValueExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr(int, int, int, int)
func NewAstExprIntegerLiteralMinValueExpr2(a int, b int, c int, d int) (*AstExprIntegerLiteralMinValueExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprIntegerLiteralMinValueExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


