package javaparser

import "github.com/timob/javabind"

type AstExprMethodReferenceExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getScope()
	GetScope() *AstExprExpression

	// public void setScope(com.github.javaparser.ast.expr.Expression)
	SetScope(a AstExprExpressionInterface) 

	// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
	GetTypeParameters() []*AstTypeParameter

	// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
	SetTypeParameters(a []*AstTypeParameter) 

	// public java.lang.String getIdentifier()
	GetIdentifier() string

	// public void setIdentifier(java.lang.String)
	SetIdentifier(a string) 
}

type AstExprMethodReferenceExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.MethodReferenceExpr()
func NewAstExprMethodReferenceExpr() (*AstExprMethodReferenceExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodReferenceExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprMethodReferenceExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MethodReferenceExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.TypeParameter>, java.lang.String)
func NewAstExprMethodReferenceExpr2(a int, b int, c int, d int, e AstExprExpressionInterface, f []*AstTypeParameter, g string) (*AstExprMethodReferenceExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodReferenceExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstExprMethodReferenceExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getScope()
func (jbobject *AstExprMethodReferenceExpr) GetScope() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScope", "com/github/javaparser/ast/expr/Expression")
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

// public void setScope(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprMethodReferenceExpr) SetScope(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScope", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
func (jbobject *AstExprMethodReferenceExpr) GetTypeParameters() []*AstTypeParameter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeParameters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeParameter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
func (jbobject *AstExprMethodReferenceExpr) SetTypeParameters(a []*AstTypeParameter)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeParameters", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String getIdentifier()
func (jbobject *AstExprMethodReferenceExpr) GetIdentifier() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIdentifier", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setIdentifier(java.lang.String)
func (jbobject *AstExprMethodReferenceExpr) SetIdentifier(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIdentifier", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


