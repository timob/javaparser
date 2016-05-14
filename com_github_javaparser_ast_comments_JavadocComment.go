package javaparser

import "github.com/timob/javabind"

type AstCommentsJavadocCommentInterface interface {
	AstCommentsCommentInterface
}

type AstCommentsJavadocComment struct {
	AstCommentsComment
}

// public com.github.javaparser.ast.comments.JavadocComment()
func NewAstCommentsJavadocComment() (*AstCommentsJavadocComment) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/JavadocComment")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsJavadocComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.JavadocComment(java.lang.String)
func NewAstCommentsJavadocComment2(a string) (*AstCommentsJavadocComment) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/JavadocComment", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstCommentsJavadocComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.JavadocComment(int, int, int, int, java.lang.String)
func NewAstCommentsJavadocComment3(a int, b int, c int, d int, e string) (*AstCommentsJavadocComment) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/JavadocComment", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstCommentsJavadocComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}


