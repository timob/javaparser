package javaparser

import "github.com/timob/javabind"

type AstExprDoubleLiteralExprInterface interface {
	AstExprStringLiteralExprInterface
}

type AstExprDoubleLiteralExpr struct {
	AstExprStringLiteralExpr
}

// public com.github.javaparser.ast.expr.DoubleLiteralExpr()
func NewAstExprDoubleLiteralExpr() (*AstExprDoubleLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/DoubleLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprDoubleLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.DoubleLiteralExpr(java.lang.String)
func NewAstExprDoubleLiteralExpr2(a string) (*AstExprDoubleLiteralExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/DoubleLiteralExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprDoubleLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.DoubleLiteralExpr(int, int, int, int, java.lang.String)
func NewAstExprDoubleLiteralExpr3(a int, b int, c int, d int, e string) (*AstExprDoubleLiteralExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/DoubleLiteralExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprDoubleLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


