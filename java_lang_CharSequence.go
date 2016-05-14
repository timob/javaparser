package javaparser

import "github.com/timob/javabind"

type JavaLangCharSequenceInterface interface {

	// public abstract int length()
	Length() int

	// public abstract java.lang.CharSequence subSequence(int, int)
	SubSequence(a int, b int) *JavaLangCharSequence
}

type JavaLangCharSequence struct {
	JavaLangObject
}

// public abstract int length()
func (jbobject *JavaLangCharSequence) Length() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.CharSequence subSequence(int, int)
func (jbobject *JavaLangCharSequence) SubSequence(a int, b int) *JavaLangCharSequence {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subSequence", "java/lang/CharSequence", a, b)
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
	unique_x := &JavaLangCharSequence{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.lang.String toString()
func (jbobject *JavaLangCharSequence) ToString() string {
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


