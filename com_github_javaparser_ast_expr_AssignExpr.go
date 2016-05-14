package javaparser

import "github.com/timob/javabind"

type AstExprAssignExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getTarget()
	GetTarget() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getValue()
	GetValue() *AstExprExpression

	// public void setTarget(com.github.javaparser.ast.expr.Expression)
	SetTarget(a AstExprExpressionInterface) 

	// public void setValue(com.github.javaparser.ast.expr.Expression)
	SetValue(a AstExprExpressionInterface) 
}

type AstExprAssignExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.AssignExpr()
func NewAstExprAssignExpr() (*AstExprAssignExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/AssignExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprAssignExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getTarget()
func (jbobject *AstExprAssignExpr) GetTarget() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTarget", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.Expression getValue()
func (jbobject *AstExprAssignExpr) GetValue() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "com/github/javaparser/ast/expr/Expression")
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

// public void setTarget(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprAssignExpr) SetTarget(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTarget", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setValue(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprAssignExpr) SetValue(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


