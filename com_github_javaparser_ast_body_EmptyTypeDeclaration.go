package javaparser

import "github.com/timob/javabind"

type AstBodyEmptyTypeDeclarationInterface interface {
	AstBodyTypeDeclarationInterface

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyEmptyTypeDeclaration struct {
	AstBodyTypeDeclaration
}

// public com.github.javaparser.ast.body.EmptyTypeDeclaration()
func NewAstBodyEmptyTypeDeclaration() (*AstBodyEmptyTypeDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EmptyTypeDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyEmptyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EmptyTypeDeclaration(int, int, int, int)
func NewAstBodyEmptyTypeDeclaration2(a int, b int, c int, d int) (*AstBodyEmptyTypeDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EmptyTypeDeclaration", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstBodyEmptyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyEmptyTypeDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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

// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
func (jbobject *AstBodyEmptyTypeDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


