package javaparser

import "github.com/timob/javabind"

type AstExprLambdaExprInterface interface {
	AstExprExpressionInterface

	// public java.util.List<com.github.javaparser.ast.body.Parameter> getParameters()
	GetParameters() []*AstBodyParameter

	// public void setParameters(java.util.List<com.github.javaparser.ast.body.Parameter>)
	SetParameters(a []*AstBodyParameter) 

	// public com.github.javaparser.ast.stmt.Statement getBody()
	GetBody() *AstStmtStatement

	// public void setBody(com.github.javaparser.ast.stmt.Statement)
	SetBody(a AstStmtStatementInterface) 

	// public boolean isParametersEnclosed()
	IsParametersEnclosed() bool

	// public void setParametersEnclosed(boolean)
	SetParametersEnclosed(a bool) 
}

type AstExprLambdaExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.LambdaExpr()
func NewAstExprLambdaExpr() (*AstExprLambdaExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LambdaExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprLambdaExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.LambdaExpr(int, int, int, int, java.util.List<com.github.javaparser.ast.body.Parameter>, com.github.javaparser.ast.stmt.Statement, boolean)
func NewAstExprLambdaExpr2(a int, b int, c int, d int, e []*AstBodyParameter, f AstStmtStatementInterface, g bool) (*AstExprLambdaExpr) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/LambdaExpr", a, b, c, d, conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/Statement"), g)
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprLambdaExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.body.Parameter> getParameters()
func (jbobject *AstExprLambdaExpr) GetParameters() []*AstBodyParameter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParameters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyParameter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setParameters(java.util.List<com.github.javaparser.ast.body.Parameter>)
func (jbobject *AstExprLambdaExpr) SetParameters(a []*AstBodyParameter)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setParameters", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.github.javaparser.ast.stmt.Statement getBody()
func (jbobject *AstExprLambdaExpr) GetBody() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBody", "com/github/javaparser/ast/stmt/Statement")
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
	unique_x := &AstStmtStatement{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBody(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstExprLambdaExpr) SetBody(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBody", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean isParametersEnclosed()
func (jbobject *AstExprLambdaExpr) IsParametersEnclosed() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isParametersEnclosed", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setParametersEnclosed(boolean)
func (jbobject *AstExprLambdaExpr) SetParametersEnclosed(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setParametersEnclosed", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


