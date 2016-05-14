package javaparser

import "github.com/timob/javabind"

type AstBodyBaseParameterInterface interface {
	AstNodeInterface

	// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public com.github.javaparser.ast.body.VariableDeclaratorId getId()
	GetId() *AstBodyVariableDeclaratorId

	// public int getModifiers()
	GetModifiers() int

	// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 

	// public void setId(com.github.javaparser.ast.body.VariableDeclaratorId)
	SetId(a AstBodyVariableDeclaratorIdInterface) 

	// public void setModifiers(int)
	SetModifiers(a int) 
}

type AstBodyBaseParameter struct {
	AstNode
}

// public com.github.javaparser.ast.body.BaseParameter()
func NewAstBodyBaseParameter() (*AstBodyBaseParameter) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BaseParameter")
	if err != nil {
		panic(err)
	}
	x := &AstBodyBaseParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BaseParameter(com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyBaseParameter2(a AstBodyVariableDeclaratorIdInterface) (*AstBodyBaseParameter) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BaseParameter", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstBodyBaseParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BaseParameter(int, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyBaseParameter3(a int, b AstBodyVariableDeclaratorIdInterface) (*AstBodyBaseParameter) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BaseParameter", a, conv_b.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyBaseParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BaseParameter(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyBaseParameter4(a int, b []*AstExprAnnotationExpr, c AstBodyVariableDeclaratorIdInterface) (*AstBodyBaseParameter) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BaseParameter", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstBodyBaseParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BaseParameter(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyBaseParameter5(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstBodyVariableDeclaratorIdInterface) (*AstBodyBaseParameter) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BaseParameter", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstBodyBaseParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstBodyBaseParameter) GetAnnotations() []*AstExprAnnotationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAnnotations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprAnnotationExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.body.VariableDeclaratorId getId()
func (jbobject *AstBodyBaseParameter) GetId() *AstBodyVariableDeclaratorId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getId", "com/github/javaparser/ast/body/VariableDeclaratorId")
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
	unique_x := &AstBodyVariableDeclaratorId{}
	unique_x.Callable = dst
	return unique_x
}

// public int getModifiers()
func (jbobject *AstBodyBaseParameter) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstBodyBaseParameter) SetAnnotations(a []*AstExprAnnotationExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAnnotations", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setId(com.github.javaparser.ast.body.VariableDeclaratorId)
func (jbobject *AstBodyBaseParameter) SetId(a AstBodyVariableDeclaratorIdInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setId", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setModifiers(int)
func (jbobject *AstBodyBaseParameter) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


