package javaparser

import "github.com/timob/javabind"

type AstBodyVariableDeclaratorIdInterface interface {
	AstNodeInterface

	// public int getArrayCount()
	GetArrayCount() int

	// public java.lang.String getName()
	GetName() string

	// public void setArrayCount(int)
	SetArrayCount(a int) 

	// public void setName(java.lang.String)
	SetName(a string) 
}

type AstBodyVariableDeclaratorId struct {
	AstNode
}

// public com.github.javaparser.ast.body.VariableDeclaratorId()
func NewAstBodyVariableDeclaratorId() (*AstBodyVariableDeclaratorId) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclaratorId")
	if err != nil {
		panic(err)
	}
	x := &AstBodyVariableDeclaratorId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclaratorId(java.lang.String)
func NewAstBodyVariableDeclaratorId2(a string) (*AstBodyVariableDeclaratorId) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclaratorId", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstBodyVariableDeclaratorId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclaratorId(int, int, int, int, java.lang.String, int)
func NewAstBodyVariableDeclaratorId3(a int, b int, c int, d int, e string, f int) (*AstBodyVariableDeclaratorId) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclaratorId", a, b, c, d, conv_e.Value().Cast("java/lang/String"), f)
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstBodyVariableDeclaratorId{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getArrayCount()
func (jbobject *AstBodyVariableDeclaratorId) GetArrayCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArrayCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getName()
func (jbobject *AstBodyVariableDeclaratorId) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public void setArrayCount(int)
func (jbobject *AstBodyVariableDeclaratorId) SetArrayCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArrayCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setName(java.lang.String)
func (jbobject *AstBodyVariableDeclaratorId) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


