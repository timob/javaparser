package javaparser

import "github.com/timob/javabind"

type AstExprSuperExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getClassExpr()
	GetClassExpr() *AstExprExpression

	// public void setClassExpr(com.github.javaparser.ast.expr.Expression)
	SetClassExpr(a AstExprExpressionInterface) 
}

type AstExprSuperExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.SuperExpr()
func NewAstExprSuperExpr() (*AstExprSuperExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SuperExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprSuperExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.SuperExpr(com.github.javaparser.ast.expr.Expression)
func NewAstExprSuperExpr2(a AstExprExpressionInterface) (*AstExprSuperExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SuperExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprSuperExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.SuperExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression)
func NewAstExprSuperExpr3(a int, b int, c int, d int, e AstExprExpressionInterface) (*AstExprSuperExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/SuperExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprSuperExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getClassExpr()
func (jbobject *AstExprSuperExpr) GetClassExpr() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClassExpr", "com/github/javaparser/ast/expr/Expression")
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

// public void setClassExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprSuperExpr) SetClassExpr(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClassExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


