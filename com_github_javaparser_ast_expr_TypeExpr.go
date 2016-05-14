package javaparser

import "github.com/timob/javabind"

type AstExprTypeExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 
}

type AstExprTypeExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.TypeExpr()
func NewAstExprTypeExpr() (*AstExprTypeExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/TypeExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprTypeExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.TypeExpr(int, int, int, int, com.github.javaparser.ast.type.Type)
func NewAstExprTypeExpr2(a int, b int, c int, d int, e AstTypeTypeInterface) (*AstExprTypeExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/TypeExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprTypeExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstExprTypeExpr) GetType() *AstTypeType {
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
func (jbobject *AstExprTypeExpr) SetType(a AstTypeTypeInterface)  {
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


