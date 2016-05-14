package javaparser

import "github.com/timob/javabind"

type AstExprIntegerLiteralExprInterface interface {
	AstExprStringLiteralExprInterface

	// public final boolean isMinValue()
	IsMinValue() bool
}

type AstExprIntegerLiteralExpr struct {
	AstExprStringLiteralExpr
}

// public com.github.javaparser.ast.expr.IntegerLiteralExpr()
func NewAstExprIntegerLiteralExpr() (*AstExprIntegerLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/IntegerLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprIntegerLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.IntegerLiteralExpr(java.lang.String)
func NewAstExprIntegerLiteralExpr2(a string) (*AstExprIntegerLiteralExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/IntegerLiteralExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprIntegerLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.IntegerLiteralExpr(int, int, int, int, java.lang.String)
func NewAstExprIntegerLiteralExpr3(a int, b int, c int, d int, e string) (*AstExprIntegerLiteralExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/IntegerLiteralExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprIntegerLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final boolean isMinValue()
func (jbobject *AstExprIntegerLiteralExpr) IsMinValue() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMinValue", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


