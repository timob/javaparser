package javaparser

import "github.com/timob/javabind"

type AstTypeVoidTypeInterface interface {
	AstTypeTypeInterface
}

type AstTypeVoidType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.VoidType()
func NewAstTypeVoidType() (*AstTypeVoidType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/VoidType")
	if err != nil {
		panic(err)
	}
	x := &AstTypeVoidType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.VoidType(int, int, int, int)
func NewAstTypeVoidType2(a int, b int, c int, d int) (*AstTypeVoidType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/VoidType", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstTypeVoidType{}
	x.Callable = &javabind.Callable{obj}
	return x
}


