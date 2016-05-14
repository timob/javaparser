package javaparser

import "github.com/timob/javabind"

type JavaIoPrintStreamInterface interface {
	JavaIoFilterOutputStreamInterface

	// public boolean checkError()
	CheckError() bool

	// public void print(boolean)
	Print(a bool) 

	// public void print(int)
	Print2(a int) 

	// public void print(long)
	Print3(a int64) 

	// public void print(float)
	Print4(a float32) 

	// public void print(double)
	Print5(a float64) 

	// public void print(java.lang.String)
	Print6(a string) 

	// public void print(java.lang.Object)
	Print7(a interface{}) 

	// public void println()
	Println() 

	// public void println(boolean)
	Println2(a bool) 

	// public void println(int)
	Println3(a int) 

	// public void println(long)
	Println4(a int64) 

	// public void println(float)
	Println5(a float32) 

	// public void println(double)
	Println6(a float64) 

	// public void println(java.lang.String)
	Println7(a string) 

	// public void println(java.lang.Object)
	Println8(a interface{}) 

	// public java.io.PrintStream printf(java.lang.String, java.lang.Object...)
	Printf(a string, b ...*JavaLangObject) *JavaIoPrintStream

	// public java.io.PrintStream printf(java.util.Locale, java.lang.String, java.lang.Object...)
	Printf2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintStream

	// public java.io.PrintStream format(java.lang.String, java.lang.Object...)
	Format(a string, b ...*JavaLangObject) *JavaIoPrintStream

	// public java.io.PrintStream format(java.util.Locale, java.lang.String, java.lang.Object...)
	Format2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintStream

	// public java.io.PrintStream append(java.lang.CharSequence)
	Append(a JavaLangCharSequenceInterface) *JavaIoPrintStream

	// public java.io.PrintStream append(java.lang.CharSequence, int, int)
	Append2(a JavaLangCharSequenceInterface, b int, c int) *JavaIoPrintStream
}

type JavaIoPrintStream struct {
	JavaIoFilterOutputStream
}

// public java.io.PrintStream(java.io.OutputStream)
func NewJavaIoPrintStream(a JavaIoOutputStreamInterface) (*JavaIoPrintStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/io/OutputStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintStream(java.io.OutputStream, boolean)
func NewJavaIoPrintStream2(a JavaIoOutputStreamInterface, b bool) (*JavaIoPrintStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/io/OutputStream"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintStream(java.io.OutputStream, boolean, java.lang.String) throws java.io.UnsupportedEncodingException
func NewJavaIoPrintStream3(a JavaIoOutputStreamInterface, b bool, c string) (*JavaIoPrintStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/io/OutputStream"), b, conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintStream(java.lang.String) throws java.io.FileNotFoundException
func NewJavaIoPrintStream4(a string) (*JavaIoPrintStream, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintStream(java.lang.String, java.lang.String) throws java.io.FileNotFoundException, java.io.UnsupportedEncodingException
func NewJavaIoPrintStream5(a string, b string) (*JavaIoPrintStream, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintStream(java.io.File) throws java.io.FileNotFoundException
func NewJavaIoPrintStream6(a JavaIoFileInterface) (*JavaIoPrintStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/io/File"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintStream(java.io.File, java.lang.String) throws java.io.FileNotFoundException, java.io.UnsupportedEncodingException
func NewJavaIoPrintStream7(a JavaIoFileInterface, b string) (*JavaIoPrintStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintStream", conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoPrintStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public void flush()
func (jbobject *JavaIoPrintStream) Flush()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "flush", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void close()
func (jbobject *JavaIoPrintStream) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean checkError()
func (jbobject *JavaIoPrintStream) CheckError() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkError", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void write(int)
func (jbobject *JavaIoPrintStream) Write(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void write(byte[], int, int)
func (jbobject *JavaIoPrintStream) Write3(a []byte, b int, c int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a, b, c)
	if err != nil {
		panic(err)
	}

}

// public void print(boolean)
func (jbobject *JavaIoPrintStream) Print(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(int)
func (jbobject *JavaIoPrintStream) Print2(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(long)
func (jbobject *JavaIoPrintStream) Print3(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(float)
func (jbobject *JavaIoPrintStream) Print4(a float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(double)
func (jbobject *JavaIoPrintStream) Print5(a float64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(java.lang.String)
func (jbobject *JavaIoPrintStream) Print6(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void print(java.lang.Object)
func (jbobject *JavaIoPrintStream) Print7(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void println()
func (jbobject *JavaIoPrintStream) Println()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void println(boolean)
func (jbobject *JavaIoPrintStream) Println2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(int)
func (jbobject *JavaIoPrintStream) Println3(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(long)
func (jbobject *JavaIoPrintStream) Println4(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(float)
func (jbobject *JavaIoPrintStream) Println5(a float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(double)
func (jbobject *JavaIoPrintStream) Println6(a float64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(java.lang.String)
func (jbobject *JavaIoPrintStream) Println7(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void println(java.lang.Object)
func (jbobject *JavaIoPrintStream) Println8(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.io.PrintStream printf(java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintStream) Printf(a string, b ...*JavaLangObject) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "printf", "java/io/PrintStream", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintStream printf(java.util.Locale, java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintStream) Printf2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "printf", "java/io/PrintStream", conv_a.Value().Cast("java/util/Locale"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintStream format(java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintStream) Format(a string, b ...*JavaLangObject) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "format", "java/io/PrintStream", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintStream format(java.util.Locale, java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintStream) Format2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "format", "java/io/PrintStream", conv_a.Value().Cast("java/util/Locale"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintStream append(java.lang.CharSequence)
func (jbobject *JavaIoPrintStream) Append(a JavaLangCharSequenceInterface) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/PrintStream", conv_a.Value().Cast("java/lang/CharSequence"))
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
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintStream append(java.lang.CharSequence, int, int)
func (jbobject *JavaIoPrintStream) Append2(a JavaLangCharSequenceInterface, b int, c int) *JavaIoPrintStream {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/PrintStream", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
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
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Appendable append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaIoPrintStream) Append3(a JavaLangCharSequenceInterface, b int, c int) (*JavaLangAppendable, error) {
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

// public java.lang.Appendable append(java.lang.CharSequence) throws java.io.IOException
func (jbobject *JavaIoPrintStream) Append4(a JavaLangCharSequenceInterface) (*JavaLangAppendable, error) {
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


