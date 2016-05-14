package javaparser

import "github.com/timob/javabind"

type AstCommentsLineCommentInterface interface {
	AstCommentsCommentInterface
}

type AstCommentsLineComment struct {
	AstCommentsComment
}

// public com.github.javaparser.ast.comments.LineComment()
func NewAstCommentsLineComment() (*AstCommentsLineComment) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/LineComment")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsLineComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.LineComment(java.lang.String)
func NewAstCommentsLineComment2(a string) (*AstCommentsLineComment) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/LineComment", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstCommentsLineComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.LineComment(int, int, int, int, java.lang.String)
func NewAstCommentsLineComment3(a int, b int, c int, d int, e string) (*AstCommentsLineComment) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/LineComment", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstCommentsLineComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public boolean isLineComment()
func (jbobject *AstCommentsLineComment) IsLineComment() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLineComment", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


