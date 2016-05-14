package javaparser

import "github.com/timob/javabind"

type AstCommentsCommentInterface interface {
	AstNodeInterface

	// public final java.lang.String getContent()
	GetContent() string

	// public void setContent(java.lang.String)
	SetContent(a string) 

	// public boolean isLineComment()
	IsLineComment() bool

	// public com.github.javaparser.ast.comments.LineComment asLineComment()
	AsLineComment() *AstCommentsLineComment

	// public com.github.javaparser.ast.Node getCommentedNode()
	GetCommentedNode() *AstNode

	// public void setCommentedNode(com.github.javaparser.ast.Node)
	SetCommentedNode(a AstNodeInterface) 

	// public boolean isOrphan()
	IsOrphan() bool
}

type AstCommentsComment struct {
	AstNode
}

// public com.github.javaparser.ast.comments.Comment()
func NewAstCommentsComment() (*AstCommentsComment) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/Comment")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.Comment(java.lang.String)
func NewAstCommentsComment2(a string) (*AstCommentsComment) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/Comment", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstCommentsComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.Comment(int, int, int, int, java.lang.String)
func NewAstCommentsComment3(a int, b int, c int, d int, e string) (*AstCommentsComment) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/Comment", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstCommentsComment{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final java.lang.String getContent()
func (jbobject *AstCommentsComment) GetContent() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getContent", "java/lang/String")
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

// public void setContent(java.lang.String)
func (jbobject *AstCommentsComment) SetContent(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setContent", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean isLineComment()
func (jbobject *AstCommentsComment) IsLineComment() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isLineComment", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public com.github.javaparser.ast.comments.LineComment asLineComment()
func (jbobject *AstCommentsComment) AsLineComment() *AstCommentsLineComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "asLineComment", "com/github/javaparser/ast/comments/LineComment")
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
	unique_x := &AstCommentsLineComment{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node getCommentedNode()
func (jbobject *AstCommentsComment) GetCommentedNode() *AstNode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCommentedNode", "com/github/javaparser/ast/Node")
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCommentedNode(com.github.javaparser.ast.Node)
func (jbobject *AstCommentsComment) SetCommentedNode(a AstNodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCommentedNode", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean isOrphan()
func (jbobject *AstCommentsComment) IsOrphan() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isOrphan", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


