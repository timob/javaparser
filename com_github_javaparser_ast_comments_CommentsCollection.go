package javaparser

import "github.com/timob/javabind"

type AstCommentsCommentsCollectionInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.github.javaparser.ast.comments.LineComment> getLineComments()
	GetLineComments() []*AstCommentsLineComment

	// public java.util.List<com.github.javaparser.ast.comments.BlockComment> getBlockComments()
	GetBlockComments() []*AstCommentsBlockComment

	// public java.util.List<com.github.javaparser.ast.comments.JavadocComment> getJavadocComments()
	GetJavadocComments() []*AstCommentsJavadocComment

	// public void addComment(com.github.javaparser.ast.comments.LineComment)
	AddComment(a AstCommentsLineCommentInterface) 

	// public void addComment(com.github.javaparser.ast.comments.BlockComment)
	AddComment2(a AstCommentsBlockCommentInterface) 

	// public void addComment(com.github.javaparser.ast.comments.JavadocComment)
	AddComment3(a AstCommentsJavadocCommentInterface) 

	// public boolean contains(com.github.javaparser.ast.comments.Comment)
	Contains(a AstCommentsCommentInterface) bool

	// public java.util.List<com.github.javaparser.ast.comments.Comment> getAll()
	GetAll() []*AstCommentsComment

	// public int size()
	Size() int

	// public com.github.javaparser.ast.comments.CommentsCollection minus(com.github.javaparser.ast.comments.CommentsCollection)
	Minus(a AstCommentsCommentsCollectionInterface) *AstCommentsCommentsCollection
}

type AstCommentsCommentsCollection struct {
	JavaLangObject
}

// public com.github.javaparser.ast.comments.CommentsCollection()
func NewAstCommentsCommentsCollection() (*AstCommentsCommentsCollection) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/CommentsCollection")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsCommentsCollection{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.comments.LineComment> getLineComments()
func (jbobject *AstCommentsCommentsCollection) GetLineComments() []*AstCommentsLineComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsLineComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.comments.BlockComment> getBlockComments()
func (jbobject *AstCommentsCommentsCollection) GetBlockComments() []*AstCommentsBlockComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlockComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsBlockComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.comments.JavadocComment> getJavadocComments()
func (jbobject *AstCommentsCommentsCollection) GetJavadocComments() []*AstCommentsJavadocComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJavadocComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsJavadocComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void addComment(com.github.javaparser.ast.comments.LineComment)
func (jbobject *AstCommentsCommentsCollection) AddComment(a AstCommentsLineCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addComment", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/LineComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addComment(com.github.javaparser.ast.comments.BlockComment)
func (jbobject *AstCommentsCommentsCollection) AddComment2(a AstCommentsBlockCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addComment", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/BlockComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void addComment(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstCommentsCommentsCollection) AddComment3(a AstCommentsJavadocCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addComment", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean contains(com.github.javaparser.ast.comments.Comment)
func (jbobject *AstCommentsCommentsCollection) Contains(a AstCommentsCommentInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/comments/Comment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public java.util.List<com.github.javaparser.ast.comments.Comment> getAll()
func (jbobject *AstCommentsCommentsCollection) GetAll() []*AstCommentsComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAll", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public int size()
func (jbobject *AstCommentsCommentsCollection) Size() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "size", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.github.javaparser.ast.comments.CommentsCollection minus(com.github.javaparser.ast.comments.CommentsCollection)
func (jbobject *AstCommentsCommentsCollection) Minus(a AstCommentsCommentsCollectionInterface) *AstCommentsCommentsCollection {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "minus", "com/github/javaparser/ast/comments/CommentsCollection", conv_a.Value().Cast("com/github/javaparser/ast/comments/CommentsCollection"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstCommentsCommentsCollection{}
	unique_x.Callable = dst
	return unique_x
}


