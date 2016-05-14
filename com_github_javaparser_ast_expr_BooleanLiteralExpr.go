package javaparser

import "github.com/timob/javabind"

type AstExprBooleanLiteralExprInterface interface {
	AstExprLiteralExprInterface

	// public boolean getValue()
	GetValue() bool

	// public void setValue(boolean)
	SetValue(a bool) 
}

type AstExprBooleanLiteralExpr struct {
	AstExprLiteralExpr
}

// public com.github.javaparser.ast.expr.BooleanLiteralExpr()
func NewAstExprBooleanLiteralExpr() (*AstExprBooleanLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/BooleanLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprBooleanLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.BooleanLiteralExpr(boolean)
func NewAstExprBooleanLiteralExpr2(a bool) (*AstExprBooleanLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/BooleanLiteralExpr", a)
	if err != nil {
		panic(err)
	}
	x := &AstExprBooleanLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.BooleanLiteralExpr(int, int, int, int, boolean)
func NewAstExprBooleanLiteralExpr3(a int, b int, c int, d int, e bool) (*AstExprBooleanLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/BooleanLiteralExpr", a, b, c, d, e)
	if err != nil {
		panic(err)
	}
	x := &AstExprBooleanLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean getValue()
func (jbobject *AstExprBooleanLiteralExpr) GetValue() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setValue(boolean)
func (jbobject *AstExprBooleanLiteralExpr) SetValue(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


