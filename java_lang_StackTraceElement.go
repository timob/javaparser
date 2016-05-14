package javaparser

import "github.com/timob/javabind"

type JavaLangStackTraceElementInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getFileName()
	GetFileName() string

	// public int getLineNumber()
	GetLineNumber() int

	// public java.lang.String getClassName()
	GetClassName() string

	// public java.lang.String getMethodName()
	GetMethodName() string

	// public boolean isNativeMethod()
	IsNativeMethod() bool
}

type JavaLangStackTraceElement struct {
	JavaLangObject
}

// public java.lang.StackTraceElement(java.lang.String, java.lang.String, java.lang.String, int)
func NewJavaLangStackTraceElement(a string, b string, c string, d int) (*JavaLangStackTraceElement) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/StackTraceElement", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"), d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &JavaLangStackTraceElement{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getFileName()
func (jbobject *JavaLangStackTraceElement) GetFileName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFileName", "java/lang/String")
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

// public int getLineNumber()
func (jbobject *JavaLangStackTraceElement) GetLineNumber() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLineNumber", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getClassName()
func (jbobject *JavaLangStackTraceElement) GetClassName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClassName", "java/lang/String")
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

// public java.lang.String getMethodName()
func (jbobject *JavaLangStackTraceElement) GetMethodName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMethodName", "java/lang/String")
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

// public boolean isNativeMethod()
func (jbobject *JavaLangStackTraceElement) IsNativeMethod() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isNativeMethod", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String toString()
func (jbobject *JavaLangStackTraceElement) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

// public boolean equals(java.lang.Object)
func (jbobject *JavaLangStackTraceElement) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int hashCode()
func (jbobject *JavaLangStackTraceElement) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}


