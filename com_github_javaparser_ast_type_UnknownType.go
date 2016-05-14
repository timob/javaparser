package javaparser

import "github.com/timob/javabind"

type AstTypeUnknownTypeInterface interface {
	AstTypeTypeInterface
}

type AstTypeUnknownType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.UnknownType()
func NewAstTypeUnknownType() (*AstTypeUnknownType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/UnknownType")
	if err != nil {
		panic(err)
	}
	x := &AstTypeUnknownType{}
	x.Callable = &javabind.Callable{obj}
	return x
}


