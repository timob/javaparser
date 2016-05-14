package javaparser

import "github.com/timob/javabind"

type AstBodyAnnotableNodeInterface interface {

	// public abstract java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr
}

type AstBodyAnnotableNode struct {
	JavaLangObject
}

// public abstract java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstBodyAnnotableNode) GetAnnotations() []*AstExprAnnotationExpr {
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


