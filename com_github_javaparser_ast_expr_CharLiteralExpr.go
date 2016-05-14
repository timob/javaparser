package javaparser

import "github.com/timob/javabind"

type AstExprCharLiteralExprInterface interface {
	AstExprStringLiteralExprInterface
}

type AstExprCharLiteralExpr struct {
	AstExprStringLiteralExpr
}

// public com.github.javaparser.ast.expr.CharLiteralExpr()
func NewAstExprCharLiteralExpr() (*AstExprCharLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CharLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprCharLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.CharLiteralExpr(java.lang.String)
func NewAstExprCharLiteralExpr2(a string) (*AstExprCharLiteralExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CharLiteralExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprCharLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.CharLiteralExpr(int, int, int, int, java.lang.String)
func NewAstExprCharLiteralExpr3(a int, b int, c int, d int, e string) (*AstExprCharLiteralExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CharLiteralExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprCharLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


