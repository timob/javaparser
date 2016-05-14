package javaparser

import "github.com/timob/javabind"

type AstExprLiteralExprInterface interface {
	AstExprExpressionInterface
}

type AstExprLiteralExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.LiteralExpr()
func NewAstExprLiteralExpr() (*AstExprLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.LiteralExpr(int, int, int, int)
func NewAstExprLiteralExpr2(a int, b int, c int, d int) (*AstExprLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LiteralExpr", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


