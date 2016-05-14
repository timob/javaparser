package javaparser

import "github.com/timob/javabind"

type AstExprAnnotationExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.NameExpr getName()
	GetName() *AstExprNameExpr

	// public void setName(com.github.javaparser.ast.expr.NameExpr)
	SetName(a AstExprNameExprInterface) 
}

type AstExprAnnotationExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.AnnotationExpr()
func NewAstExprAnnotationExpr() (*AstExprAnnotationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/AnnotationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.AnnotationExpr(int, int, int, int)
func NewAstExprAnnotationExpr2(a int, b int, c int, d int) (*AstExprAnnotationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/AnnotationExpr", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstExprAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NameExpr getName()
func (jbobject *AstExprAnnotationExpr) GetName() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "com/github/javaparser/ast/expr/NameExpr")
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
	unique_x := &AstExprNameExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstExprAnnotationExpr) SetName(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


