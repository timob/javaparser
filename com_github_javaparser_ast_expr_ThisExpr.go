package javaparser

import "github.com/timob/javabind"

type AstExprThisExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getClassExpr()
	GetClassExpr() *AstExprExpression

	// public void setClassExpr(com.github.javaparser.ast.expr.Expression)
	SetClassExpr(a AstExprExpressionInterface) 
}

type AstExprThisExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ThisExpr()
func NewAstExprThisExpr() (*AstExprThisExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ThisExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprThisExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ThisExpr(com.github.javaparser.ast.expr.Expression)
func NewAstExprThisExpr2(a AstExprExpressionInterface) (*AstExprThisExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ThisExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprThisExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ThisExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression)
func NewAstExprThisExpr3(a int, b int, c int, d int, e AstExprExpressionInterface) (*AstExprThisExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ThisExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprThisExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getClassExpr()
func (jbobject *AstExprThisExpr) GetClassExpr() *AstExprExpression {
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
func (jbobject *AstExprThisExpr) SetClassExpr(a AstExprExpressionInterface)  {
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


