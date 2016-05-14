package javaparser

import "github.com/timob/javabind"

type AstCommentsBlockCommentInterface interface {
	AstCommentsCommentInterface
}

type AstCommentsBlockComment struct {
	AstCommentsComment
}

// public com.github.javaparser.ast.comments.BlockComment()
func NewAstCommentsBlockComment() (*AstCommentsBlockComment) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/BlockComment")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsBlockComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.BlockComment(java.lang.String)
func NewAstCommentsBlockComment2(a string) (*AstCommentsBlockComment) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/BlockComment", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstCommentsBlockComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.BlockComment(int, int, int, int, java.lang.String)
func NewAstCommentsBlockComment3(a int, b int, c int, d int, e string) (*AstCommentsBlockComment) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/BlockComment", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstCommentsBlockComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}


