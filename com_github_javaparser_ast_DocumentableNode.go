package javaparser

import "github.com/timob/javabind"

type AstDocumentableNodeInterface interface {

	// public abstract com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment

	// public abstract void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 
}

type AstDocumentableNode struct {
	JavaLangObject
}

// public abstract com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
func (jbobject *AstDocumentableNode) GetJavaDoc() *AstCommentsJavadocComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJavaDoc", "com/github/javaparser/ast/comments/JavadocComment")
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
	unique_x := &AstCommentsJavadocComment{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstDocumentableNode) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJavaDoc", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


