package javaparser

import "github.com/timob/javabind"

type AstExprClassExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 
}

type AstExprClassExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ClassExpr()
func NewAstExprClassExpr() (*AstExprClassExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ClassExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprClassExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ClassExpr(com.github.javaparser.ast.type.Type)
func NewAstExprClassExpr2(a AstTypeTypeInterface) (*AstExprClassExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ClassExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprClassExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ClassExpr(int, int, int, int, com.github.javaparser.ast.type.Type)
func NewAstExprClassExpr3(a int, b int, c int, d int, e AstTypeTypeInterface) (*AstExprClassExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ClassExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprClassExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstExprClassExpr) GetType() *AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/Type")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstTypeType{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstExprClassExpr) SetType(a AstTypeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


