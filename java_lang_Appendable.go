package javaparser

import "github.com/timob/javabind"

type JavaLangAppendableInterface interface {
}

type JavaLangAppendable struct {
	JavaLangObject
}

// public abstract java.lang.Appendable append(java.lang.CharSequence) throws java.io.IOException
func (jbobject *JavaLangAppendable) Append(a JavaLangCharSequenceInterface) (*JavaLangAppendable, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/lang/Appendable", conv_a.Value().Cast("java/lang/CharSequence"))
	if err != nil {
		var zero *JavaLangAppendable
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangAppendable{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public abstract java.lang.Appendable append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaLangAppendable) Append2(a JavaLangCharSequenceInterface, b int, c int) (*JavaLangAppendable, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/lang/Appendable", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
	if err != nil {
		var zero *JavaLangAppendable
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangAppendable{}
	unique_x.Callable = dst
	return unique_x, nil
}


