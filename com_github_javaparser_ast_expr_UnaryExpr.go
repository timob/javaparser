package javaparser

import "github.com/timob/javabind"

type AstExprUnaryExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getExpr()
	GetExpr() *AstExprExpression

	// public void setExpr(com.github.javaparser.ast.expr.Expression)
	SetExpr(a AstExprExpressionInterface) 
}

type AstExprUnaryExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.UnaryExpr()
func NewAstExprUnaryExpr() (*AstExprUnaryExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/UnaryExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprUnaryExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getExpr()
func (jbobject *AstExprUnaryExpr) GetExpr() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpr", "com/github/javaparser/ast/expr/Expression")
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

// public void setExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprUnaryExpr) SetExpr(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


