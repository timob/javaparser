package javaparser

import "github.com/timob/javabind"

type JavaLangThrowableInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getMessage()
	GetMessage() string

	// public java.lang.String getLocalizedMessage()
	GetLocalizedMessage() string

	// public java.lang.Throwable getCause()
	GetCause() *JavaLangThrowable

	// public synchronized java.lang.Throwable initCause(java.lang.Throwable)
	InitCause(a JavaLangThrowableInterface) *JavaLangThrowable

	// public void printStackTrace()
	PrintStackTrace() 

	// public void printStackTrace(java.io.PrintStream)
	PrintStackTrace2(a JavaIoPrintStreamInterface) 

	// public void printStackTrace(java.io.PrintWriter)
	PrintStackTrace3(a JavaIoPrintWriterInterface) 

	// public synchronized native java.lang.Throwable fillInStackTrace()
	FillInStackTrace() *JavaLangThrowable

	// public java.lang.StackTraceElement[] getStackTrace()
	GetStackTrace() []*JavaLangStackTraceElement

	// public void setStackTrace(java.lang.StackTraceElement[])
	SetStackTrace(a []*JavaLangStackTraceElement) 
}

type JavaLangThrowable struct {
	JavaLangObject
}

// public java.lang.Throwable()
func NewJavaLangThrowable() (*JavaLangThrowable) {

	obj, err := javabind.GetEnv().NewObject("java/lang/Throwable")
	if err != nil {
		panic(err)
	}
	x := &JavaLangThrowable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Throwable(java.lang.String)
func NewJavaLangThrowable2(a string) (*JavaLangThrowable) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Throwable", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangThrowable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Throwable(java.lang.String, java.lang.Throwable)
func NewJavaLangThrowable3(a string, b JavaLangThrowableInterface) (*JavaLangThrowable) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Throwable", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Throwable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaLangThrowable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Throwable(java.lang.Throwable)
func NewJavaLangThrowable4(a JavaLangThrowableInterface) (*JavaLangThrowable) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/lang/Throwable", conv_a.Value().Cast("java/lang/Throwable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaLangThrowable{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getMessage()
func (jbobject *JavaLangThrowable) GetMessage() string {
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

// public java.lang.String getLocalizedMessage()
func (jbobject *JavaLangThrowable) GetLocalizedMessage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLocalizedMessage", "java/lang/String")
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

// public java.lang.Throwable getCause()
func (jbobject *JavaLangThrowable) GetCause() *JavaLangThrowable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCause", "java/lang/Throwable")
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
	unique_x := &JavaLangThrowable{}
	unique_x.Callable = dst
	return unique_x
}

// public synchronized java.lang.Throwable initCause(java.lang.Throwable)
func (jbobject *JavaLangThrowable) InitCause(a JavaLangThrowableInterface) *JavaLangThrowable {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "initCause", "java/lang/Throwable", conv_a.Value().Cast("java/lang/Throwable"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangThrowable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String toString()
func (jbobject *JavaLangThrowable) ToString() string {
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

// public void printStackTrace()
func (jbobject *JavaLangThrowable) PrintStackTrace()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "printStackTrace", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void printStackTrace(java.io.PrintStream)
func (jbobject *JavaLangThrowable) PrintStackTrace2(a JavaIoPrintStreamInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "printStackTrace", javabind.Void, conv_a.Value().Cast("java/io/PrintStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void printStackTrace(java.io.PrintWriter)
func (jbobject *JavaLangThrowable) PrintStackTrace3(a JavaIoPrintWriterInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "printStackTrace", javabind.Void, conv_a.Value().Cast("java/io/PrintWriter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public synchronized native java.lang.Throwable fillInStackTrace()
func (jbobject *JavaLangThrowable) FillInStackTrace() *JavaLangThrowable {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "fillInStackTrace", "java/lang/Throwable")
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
	unique_x := &JavaLangThrowable{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.StackTraceElement[] getStackTrace()
func (jbobject *JavaLangThrowable) GetStackTrace() []*JavaLangStackTraceElement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStackTrace", javabind.ObjectArrayType("java/lang/StackTraceElement"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/lang/StackTraceElement")
	dst := new([]*JavaLangStackTraceElement)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setStackTrace(java.lang.StackTraceElement[])
func (jbobject *JavaLangThrowable) SetStackTrace(a []*JavaLangStackTraceElement)  {
	conv_a := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/StackTraceElement")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStackTrace", javabind.Void, conv_a.Value().Cast("java/lang/StackTraceElement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


