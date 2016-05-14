package javaparser

import "github.com/timob/javabind"

type AstCommentsCommentsParserInterface interface {
	JavaLangObjectInterface
}

type AstCommentsCommentsParser struct {
	JavaLangObject
}

// public com.github.javaparser.ast.comments.CommentsParser()
func NewAstCommentsCommentsParser() (*AstCommentsCommentsParser) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/comments/CommentsParser")
	if err != nil {
		panic(err)
	}
	x := &AstCommentsCommentsParser{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.comments.CommentsCollection parse(java.lang.String) throws java.io.IOException, java.io.UnsupportedEncodingException
func (jbobject *AstCommentsCommentsParser) Parse(a string) (*AstCommentsCommentsCollection, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parse", "com/github/javaparser/ast/comments/CommentsCollection", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstCommentsCommentsCollection
		return zero, err
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
	return unique_x, nil
}

// public com.github.javaparser.ast.comments.CommentsCollection parse(java.io.InputStream, java.lang.String) throws java.io.IOException, java.io.UnsupportedEncodingException
func (jbobject *AstCommentsCommentsParser) Parse2(a JavaIoInputStreamInterface, b string) (*AstCommentsCommentsCollection, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "parse", "com/github/javaparser/ast/comments/CommentsCollection", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstCommentsCommentsCollection
		return zero, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstCommentsCommentsCollection{}
	unique_x.Callable = dst
	return unique_x, nil
}


