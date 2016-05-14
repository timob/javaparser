package javaparser

import "github.com/timob/javabind"

type AstTypeUnionTypeInterface interface {
	AstTypeTypeInterface

	// public java.util.List<com.github.javaparser.ast.type.ReferenceType> getElements()
	GetElements() []*AstTypeReferenceType

	// public void setElements(java.util.List<com.github.javaparser.ast.type.ReferenceType>)
	SetElements(a []*AstTypeReferenceType) 
}

type AstTypeUnionType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.UnionType(int, int, int, int, java.util.List<com.github.javaparser.ast.type.ReferenceType>)
func NewAstTypeUnionType(a int, b int, c int, d int, e []*AstTypeReferenceType) (*AstTypeUnionType) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/UnionType", a, b, c, d, conv_e.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstTypeUnionType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.UnionType(java.util.List<com.github.javaparser.ast.type.ReferenceType>)
func NewAstTypeUnionType2(a []*AstTypeReferenceType) (*AstTypeUnionType) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/UnionType", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeUnionType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.type.ReferenceType> getElements()
func (jbobject *AstTypeUnionType) GetElements() []*AstTypeReferenceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getElements", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeReferenceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setElements(java.util.List<com.github.javaparser.ast.type.ReferenceType>)
func (jbobject *AstTypeUnionType) SetElements(a []*AstTypeReferenceType)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setElements", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


