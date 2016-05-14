package javaparser

import "github.com/timob/javabind"

type AstInternalUtilsInterface interface {
	JavaLangObjectInterface
}

type AstInternalUtils struct {
	JavaLangObject
}

// public com.github.javaparser.ast.internal.Utils()
func NewAstInternalUtils() (*AstInternalUtils) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/internal/Utils")
	if err != nil {
		panic(err)
	}
	x := &AstInternalUtils{}
	x.Callable = &javabind.Callable{obj}
	return x
}


