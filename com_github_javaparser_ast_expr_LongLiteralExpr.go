package javaparser

import "github.com/timob/javabind"

type AstExprLongLiteralExprInterface interface {
	AstExprStringLiteralExprInterface

	// public final boolean isMinValue()
	IsMinValue() bool
}

type AstExprLongLiteralExpr struct {
	AstExprStringLiteralExpr
}

// public com.github.javaparser.ast.expr.LongLiteralExpr()
func NewAstExprLongLiteralExpr() (*AstExprLongLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LongLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprLongLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.LongLiteralExpr(java.lang.String)
func NewAstExprLongLiteralExpr2(a string) (*AstExprLongLiteralExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LongLiteralExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprLongLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.LongLiteralExpr(int, int, int, int, java.lang.String)
func NewAstExprLongLiteralExpr3(a int, b int, c int, d int, e string) (*AstExprLongLiteralExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LongLiteralExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprLongLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final boolean isMinValue()
func (jbobject *AstExprLongLiteralExpr) IsMinValue() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isMinValue", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


