package javaparser

import "github.com/timob/javabind"

type JavaLangErrorInterface interface {
	JavaLangThrowableInterface
}

type JavaLangError struct {
	JavaLangThrowable
}

// public java.lang.Error()
func NewJavaLangError() (*JavaLangError) {

	obj, err := javabind.GetEnv().NewObject("java/lang/Error")
	if err != nil {
		panic(err)
	}
	x := &JavaLangError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Error(java.lang.String)
func NewJavaLangError2(a string) (*JavaLangError) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Error", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Error(java.lang.String, java.lang.Throwable)
func NewJavaLangError3(a string, b JavaLangThrowableInterface) (*JavaLangError) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Error", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Throwable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaLangError{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Error(java.lang.Throwable)
func NewJavaLangError4(a JavaLangThrowableInterface) (*JavaLangError) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Error", conv_a.Value().Cast("java/lang/Throwable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangError{}
	x.Callable = &javabind.Callable{obj}
	return x
}


