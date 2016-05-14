package javaparser

import "github.com/timob/javabind"

type AstBodyBodyDeclarationInterface interface {
	AstNodeInterface

	// public final java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public final void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 
}

type AstBodyBodyDeclaration struct {
	AstNode
}

// public com.github.javaparser.ast.body.BodyDeclaration()
func NewAstBodyBodyDeclaration() (*AstBodyBodyDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BodyDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyBodyDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BodyDeclaration(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func NewAstBodyBodyDeclaration2(a []*AstExprAnnotationExpr) (*AstBodyBodyDeclaration) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BodyDeclaration", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstBodyBodyDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.BodyDeclaration(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func NewAstBodyBodyDeclaration3(a int, b int, c int, d int, e []*AstExprAnnotationExpr) (*AstBodyBodyDeclaration) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/BodyDeclaration", a, b, c, d, conv_e.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstBodyBodyDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstBodyBodyDeclaration) GetAnnotations() []*AstExprAnnotationExpr {
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

// public final void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstBodyBodyDeclaration) SetAnnotations(a []*AstExprAnnotationExpr)  {
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


