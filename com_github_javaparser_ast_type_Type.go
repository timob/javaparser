package javaparser

import "github.com/timob/javabind"

type AstTypeTypeInterface interface {
	AstNodeInterface

	// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 
}

type AstTypeType struct {
	AstNode
}

// public com.github.javaparser.ast.type.Type()
func NewAstTypeType() (*AstTypeType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/Type")
	if err != nil {
		panic(err)
	}
	x := &AstTypeType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func NewAstTypeType2(a []*AstExprAnnotationExpr) (*AstTypeType) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/Type", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type(int, int, int, int)
func NewAstTypeType3(a int, b int, c int, d int) (*AstTypeType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/Type", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstTypeType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func NewAstTypeType4(a int, b int, c int, d int, e []*AstExprAnnotationExpr) (*AstTypeType) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/Type", a, b, c, d, conv_e.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstTypeType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstTypeType) GetAnnotations() []*AstExprAnnotationExpr {
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

// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstTypeType) SetAnnotations(a []*AstExprAnnotationExpr)  {
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


