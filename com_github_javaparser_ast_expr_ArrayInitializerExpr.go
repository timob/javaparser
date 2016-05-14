package javaparser

import "github.com/timob/javabind"

type AstExprArrayInitializerExprInterface interface {
	AstExprExpressionInterface

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getValues()
	GetValues() []*AstExprExpression

	// public void setValues(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetValues(a []*AstExprExpression) 
}

type AstExprArrayInitializerExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ArrayInitializerExpr()
func NewAstExprArrayInitializerExpr() (*AstExprArrayInitializerExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayInitializerExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprArrayInitializerExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayInitializerExpr(java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstExprArrayInitializerExpr2(a []*AstExprExpression) (*AstExprArrayInitializerExpr) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayInitializerExpr", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprArrayInitializerExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayInitializerExpr(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstExprArrayInitializerExpr3(a int, b int, c int, d int, e []*AstExprExpression) (*AstExprArrayInitializerExpr) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayInitializerExpr", a, b, c, d, conv_e.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprArrayInitializerExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getValues()
func (jbobject *AstExprArrayInitializerExpr) GetValues() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValues", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprExpression)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setValues(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstExprArrayInitializerExpr) SetValues(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValues", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


