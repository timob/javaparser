package javaparser

import "github.com/timob/javabind"

type AstTypePrimitiveTypeInterface interface {
	AstTypeTypeInterface

	// public com.github.javaparser.ast.type.ClassOrInterfaceType toBoxedType()
	ToBoxedType() *AstTypeClassOrInterfaceType
}

type AstTypePrimitiveType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.PrimitiveType()
func NewAstTypePrimitiveType() (*AstTypePrimitiveType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/PrimitiveType")
	if err != nil {
		panic(err)
	}
	x := &AstTypePrimitiveType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType toBoxedType()
func (jbobject *AstTypePrimitiveType) ToBoxedType() *AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toBoxedType", "com/github/javaparser/ast/type/ClassOrInterfaceType")
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
	unique_x := &AstTypeClassOrInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}


