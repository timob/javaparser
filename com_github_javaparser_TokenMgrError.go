package javaparser

import "github.com/timob/javabind"

type TokenMgrErrorInterface interface {
	JavaLangErrorInterface
}

type TokenMgrError struct {
	JavaLangError
}

// public com.github.javaparser.TokenMgrError()
func NewTokenMgrError() (*TokenMgrError) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/TokenMgrError")
	if err != nil {
		panic(err)
	}
	x := &TokenMgrError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.TokenMgrError(java.lang.String, int)
func NewTokenMgrError2(a string, b int) (*TokenMgrError) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/TokenMgrError", conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &TokenMgrError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getMessage()
func (jbobject *TokenMgrError) GetMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "java/lang/String")
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


