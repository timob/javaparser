package javaparser

import "github.com/timob/javabind"

type AstExprBinaryExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getLeft()
	GetLeft() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getRight()
	GetRight() *AstExprExpression

	// public void setLeft(com.github.javaparser.ast.expr.Expression)
	SetLeft(a AstExprExpressionInterface) 

	// public void setRight(com.github.javaparser.ast.expr.Expression)
	SetRight(a AstExprExpressionInterface) 
}

type AstExprBinaryExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.BinaryExpr()
func NewAstExprBinaryExpr() (*AstExprBinaryExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/BinaryExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprBinaryExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getLeft()
func (jbobject *AstExprBinaryExpr) GetLeft() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLeft", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.Expression getRight()
func (jbobject *AstExprBinaryExpr) GetRight() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getRight", "com/github/javaparser/ast/expr/Expression")
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

// public void setLeft(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprBinaryExpr) SetLeft(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLeft", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setRight(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprBinaryExpr) SetRight(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setRight", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


