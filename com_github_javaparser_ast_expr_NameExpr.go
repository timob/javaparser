package javaparser

import "github.com/timob/javabind"

type AstExprNameExprInterface interface {
	AstExprExpressionInterface

	// public final java.lang.String getName()
	GetName() string

	// public final void setName(java.lang.String)
	SetName(a string) 
}

type AstExprNameExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.NameExpr()
func NewAstExprNameExpr() (*AstExprNameExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NameExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NameExpr(java.lang.String)
func NewAstExprNameExpr2(a string) (*AstExprNameExpr) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NameExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NameExpr(int, int, int, int, java.lang.String)
func NewAstExprNameExpr3(a int, b int, c int, d int, e string) (*AstExprNameExpr) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NameExpr", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final java.lang.String getName()
func (jbobject *AstExprNameExpr) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public final void setName(java.lang.String)
func (jbobject *AstExprNameExpr) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


