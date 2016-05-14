package javaparser

import "github.com/timob/javabind"

type AstTypeWildcardTypeInterface interface {
	AstTypeTypeInterface

	// public com.github.javaparser.ast.type.ReferenceType getExtends()
	GetExtends() *AstTypeReferenceType

	// public com.github.javaparser.ast.type.ReferenceType getSuper()
	GetSuper() *AstTypeReferenceType

	// public void setExtends(com.github.javaparser.ast.type.ReferenceType)
	SetExtends(a AstTypeReferenceTypeInterface) 

	// public void setSuper(com.github.javaparser.ast.type.ReferenceType)
	SetSuper(a AstTypeReferenceTypeInterface) 
}

type AstTypeWildcardType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.WildcardType()
func NewAstTypeWildcardType() (*AstTypeWildcardType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/WildcardType")
	if err != nil {
		panic(err)
	}
	x := &AstTypeWildcardType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.WildcardType(com.github.javaparser.ast.type.ReferenceType)
func NewAstTypeWildcardType2(a AstTypeReferenceTypeInterface) (*AstTypeWildcardType) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/WildcardType", conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeWildcardType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.WildcardType(com.github.javaparser.ast.type.ReferenceType, com.github.javaparser.ast.type.ReferenceType)
func NewAstTypeWildcardType3(a AstTypeReferenceTypeInterface, b AstTypeReferenceTypeInterface) (*AstTypeWildcardType) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/WildcardType", conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_b.Value().Cast("com/github/javaparser/ast/type/ReferenceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstTypeWildcardType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.WildcardType(int, int, int, int, com.github.javaparser.ast.type.ReferenceType, com.github.javaparser.ast.type.ReferenceType)
func NewAstTypeWildcardType4(a int, b int, c int, d int, e AstTypeReferenceTypeInterface, f AstTypeReferenceTypeInterface) (*AstTypeWildcardType) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/WildcardType", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_f.Value().Cast("com/github/javaparser/ast/type/ReferenceType"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstTypeWildcardType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ReferenceType getExtends()
func (jbobject *AstTypeWildcardType) GetExtends() *AstTypeReferenceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExtends", "com/github/javaparser/ast/type/ReferenceType")
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
	unique_x := &AstTypeReferenceType{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.type.ReferenceType getSuper()
func (jbobject *AstTypeWildcardType) GetSuper() *AstTypeReferenceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSuper", "com/github/javaparser/ast/type/ReferenceType")
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
	unique_x := &AstTypeReferenceType{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExtends(com.github.javaparser.ast.type.ReferenceType)
func (jbobject *AstTypeWildcardType) SetExtends(a AstTypeReferenceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExtends", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSuper(com.github.javaparser.ast.type.ReferenceType)
func (jbobject *AstTypeWildcardType) SetSuper(a AstTypeReferenceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSuper", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


