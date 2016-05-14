package javaparser

import "github.com/timob/javabind"

type AstExprEnclosedExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getInner()
	GetInner() *AstExprExpression

	// public void setInner(com.github.javaparser.ast.expr.Expression)
	SetInner(a AstExprExpressionInterface) 
}

type AstExprEnclosedExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.EnclosedExpr()
func NewAstExprEnclosedExpr() (*AstExprEnclosedExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/EnclosedExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprEnclosedExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.EnclosedExpr(com.github.javaparser.ast.expr.Expression)
func NewAstExprEnclosedExpr2(a AstExprExpressionInterface) (*AstExprEnclosedExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/EnclosedExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprEnclosedExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.EnclosedExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression)
func NewAstExprEnclosedExpr3(a int, b int, c int, d int, e AstExprExpressionInterface) (*AstExprEnclosedExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/EnclosedExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprEnclosedExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getInner()
func (jbobject *AstExprEnclosedExpr) GetInner() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInner", "com/github/javaparser/ast/expr/Expression")
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

// public void setInner(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprEnclosedExpr) SetInner(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInner", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


