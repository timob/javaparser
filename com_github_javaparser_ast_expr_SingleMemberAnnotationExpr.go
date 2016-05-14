package javaparser

import "github.com/timob/javabind"

type AstExprSingleMemberAnnotationExprInterface interface {
	AstExprAnnotationExprInterface

	// public com.github.javaparser.ast.expr.Expression getMemberValue()
	GetMemberValue() *AstExprExpression

	// public void setMemberValue(com.github.javaparser.ast.expr.Expression)
	SetMemberValue(a AstExprExpressionInterface) 
}

type AstExprSingleMemberAnnotationExpr struct {
	AstExprAnnotationExpr
}

// public com.github.javaparser.ast.expr.SingleMemberAnnotationExpr()
func NewAstExprSingleMemberAnnotationExpr() (*AstExprSingleMemberAnnotationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprSingleMemberAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.SingleMemberAnnotationExpr(com.github.javaparser.ast.expr.NameExpr, com.github.javaparser.ast.expr.Expression)
func NewAstExprSingleMemberAnnotationExpr2(a AstExprNameExprInterface, b AstExprExpressionInterface) (*AstExprSingleMemberAnnotationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprSingleMemberAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.SingleMemberAnnotationExpr(int, int, int, int, com.github.javaparser.ast.expr.NameExpr, com.github.javaparser.ast.expr.Expression)
func NewAstExprSingleMemberAnnotationExpr3(a int, b int, c int, d int, e AstExprNameExprInterface, f AstExprExpressionInterface) (*AstExprSingleMemberAnnotationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprSingleMemberAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getMemberValue()
func (jbobject *AstExprSingleMemberAnnotationExpr) GetMemberValue() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMemberValue", "com/github/javaparser/ast/expr/Expression")
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
	unique_x := &AstExprExpression{}
	unique_x.Callable = dst
	return unique_x
}

// public void setMemberValue(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprSingleMemberAnnotationExpr) SetMemberValue(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMemberValue", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


