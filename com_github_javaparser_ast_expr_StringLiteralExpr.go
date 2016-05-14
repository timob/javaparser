package javaparser

import "github.com/timob/javabind"

type AstExprStringLiteralExprInterface interface {
	AstExprLiteralExprInterface

	// public final java.lang.String getValue()
	GetValue() string

	// public final void setValue(java.lang.String)
	SetValue(a string) 
}

type AstExprStringLiteralExpr struct {
	AstExprLiteralExpr
}

// public com.github.javaparser.ast.expr.StringLiteralExpr()
func NewAstExprStringLiteralExpr() (*AstExprStringLiteralExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/StringLiteralExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprStringLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.StringLiteralExpr(java.lang.String)
func NewAstExprStringLiteralExpr2(a string) (*AstExprStringLiteralExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/StringLiteralExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprStringLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.StringLiteralExpr(int, int, int, int, java.lang.String)
func NewAstExprStringLiteralExpr3(a int, b int, c int, d int, e string) (*AstExprStringLiteralExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/StringLiteralExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprStringLiteralExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final java.lang.String getValue()
func (jbobject *AstExprStringLiteralExpr) GetValue() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final void setValue(java.lang.String)
func (jbobject *AstExprStringLiteralExpr) SetValue(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


