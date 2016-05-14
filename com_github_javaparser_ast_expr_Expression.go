package javaparser

import "github.com/timob/javabind"

type AstExprExpressionInterface interface {
	AstNodeInterface
}

type AstExprExpression struct {
	AstNode
}

// public com.github.javaparser.ast.expr.Expression()
func NewAstExprExpression() (*AstExprExpression) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/Expression")
	if err != nil {
		panic(err)
	}
	x := &AstExprExpression{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression(int, int, int, int)
func NewAstExprExpression2(a int, b int, c int, d int) (*AstExprExpression) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/Expression", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprExpression{}
	x.Callable = &javabind.Callable{obj}
	return x
}


