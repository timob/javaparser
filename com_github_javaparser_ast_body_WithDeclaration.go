package javaparser

import "github.com/timob/javabind"

type AstBodyWithDeclarationInterface interface {

	// public abstract java.lang.String getDeclarationAsString()
	GetDeclarationAsString() string

	// public abstract java.lang.String getDeclarationAsString(boolean, boolean)
	GetDeclarationAsString2(a bool, b bool) string

	// public abstract java.lang.String getDeclarationAsString(boolean, boolean, boolean)
	GetDeclarationAsString3(a bool, b bool, c bool) string
}

type AstBodyWithDeclaration struct {
	JavaLangObject
}

// public abstract java.lang.String getDeclarationAsString()
func (jbobject *AstBodyWithDeclaration) GetDeclarationAsString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String")
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

// public abstract java.lang.String getDeclarationAsString(boolean, boolean)
func (jbobject *AstBodyWithDeclaration) GetDeclarationAsString2(a bool, b bool) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String", a, b)
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

// public abstract java.lang.String getDeclarationAsString(boolean, boolean, boolean)
func (jbobject *AstBodyWithDeclaration) GetDeclarationAsString3(a bool, b bool, c bool) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String", a, b, c)
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


