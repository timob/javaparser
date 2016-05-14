package javaparser

import "github.com/timob/javabind"

type AstTypedNodeInterface interface {

	// public abstract com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public abstract void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 
}

type AstTypedNode struct {
	JavaLangObject
}

// public abstract com.github.javaparser.ast.type.Type getType()
func (jbobject *AstTypedNode) GetType() *AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/Type")
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
	unique_x := &AstTypeType{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstTypedNode) SetType(a AstTypeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


