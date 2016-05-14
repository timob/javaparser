package javaparser

import "github.com/timob/javabind"

type JavaIoPrintWriterInterface interface {
	JavaIoWriterInterface

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

	// public java.io.PrintWriter printf(java.lang.String, java.lang.Object...)
	Printf(a string, b ...*JavaLangObject) *JavaIoPrintWriter

	// public java.io.PrintWriter printf(java.util.Locale, java.lang.String, java.lang.Object...)
	Printf2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintWriter

	// public java.io.PrintWriter format(java.lang.String, java.lang.Object...)
	Format(a string, b ...*JavaLangObject) *JavaIoPrintWriter

	// public java.io.PrintWriter format(java.util.Locale, java.lang.String, java.lang.Object...)
	Format2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintWriter

	// public java.io.PrintWriter append(java.lang.CharSequence)
	Append5(a JavaLangCharSequenceInterface) *JavaIoPrintWriter

	// public java.io.PrintWriter append(java.lang.CharSequence, int, int)
	Append6(a JavaLangCharSequenceInterface, b int, c int) *JavaIoPrintWriter
}

type JavaIoPrintWriter struct {
	JavaIoWriter
}

// public java.io.PrintWriter(java.io.Writer)
func NewJavaIoPrintWriter(a JavaIoWriterInterface) (*JavaIoPrintWriter) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/Writer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintWriter(java.io.Writer, boolean)
func NewJavaIoPrintWriter2(a JavaIoWriterInterface, b bool) (*JavaIoPrintWriter) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/Writer"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintWriter(java.io.OutputStream)
func NewJavaIoPrintWriter3(a JavaIoOutputStreamInterface) (*JavaIoPrintWriter) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/OutputStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintWriter(java.io.OutputStream, boolean)
func NewJavaIoPrintWriter4(a JavaIoOutputStreamInterface, b bool) (*JavaIoPrintWriter) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/OutputStream"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.PrintWriter(java.lang.String) throws java.io.FileNotFoundException
func NewJavaIoPrintWriter5(a string) (*JavaIoPrintWriter, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintWriter(java.lang.String, java.lang.String) throws java.io.FileNotFoundException, java.io.UnsupportedEncodingException
func NewJavaIoPrintWriter6(a string, b string) (*JavaIoPrintWriter, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintWriter(java.io.File) throws java.io.FileNotFoundException
func NewJavaIoPrintWriter7(a JavaIoFileInterface) (*JavaIoPrintWriter, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/File"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public java.io.PrintWriter(java.io.File, java.lang.String) throws java.io.FileNotFoundException, java.io.UnsupportedEncodingException
func NewJavaIoPrintWriter8(a JavaIoFileInterface, b string) (*JavaIoPrintWriter, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/PrintWriter", conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoPrintWriter{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public void flush()
func (jbobject *JavaIoPrintWriter) Flush()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "flush", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void close()
func (jbobject *JavaIoPrintWriter) Close()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public boolean checkError()
func (jbobject *JavaIoPrintWriter) CheckError() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "checkError", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void write(int)
func (jbobject *JavaIoPrintWriter) Write(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void write(java.lang.String, int, int)
func (jbobject *JavaIoPrintWriter) Write3(a string, b int, c int)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void write(java.lang.String)
func (jbobject *JavaIoPrintWriter) Write2(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void print(boolean)
func (jbobject *JavaIoPrintWriter) Print(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(int)
func (jbobject *JavaIoPrintWriter) Print2(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(long)
func (jbobject *JavaIoPrintWriter) Print3(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(float)
func (jbobject *JavaIoPrintWriter) Print4(a float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(double)
func (jbobject *JavaIoPrintWriter) Print5(a float64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "print", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void print(java.lang.String)
func (jbobject *JavaIoPrintWriter) Print6(a string)  {
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
func (jbobject *JavaIoPrintWriter) Print7(a interface{})  {
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
func (jbobject *JavaIoPrintWriter) Println()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void println(boolean)
func (jbobject *JavaIoPrintWriter) Println2(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(int)
func (jbobject *JavaIoPrintWriter) Println3(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(long)
func (jbobject *JavaIoPrintWriter) Println4(a int64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(float)
func (jbobject *JavaIoPrintWriter) Println5(a float32)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(double)
func (jbobject *JavaIoPrintWriter) Println6(a float64)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "println", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void println(java.lang.String)
func (jbobject *JavaIoPrintWriter) Println7(a string)  {
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
func (jbobject *JavaIoPrintWriter) Println8(a interface{})  {
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

// public java.io.PrintWriter printf(java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintWriter) Printf(a string, b ...*JavaLangObject) *JavaIoPrintWriter {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "printf", "java/io/PrintWriter", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintWriter printf(java.util.Locale, java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintWriter) Printf2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintWriter {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "printf", "java/io/PrintWriter", conv_a.Value().Cast("java/util/Locale"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintWriter format(java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintWriter) Format(a string, b ...*JavaLangObject) *JavaIoPrintWriter {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "java/lang/Object")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "format", "java/io/PrintWriter", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintWriter format(java.util.Locale, java.lang.String, java.lang.Object...)
func (jbobject *JavaIoPrintWriter) Format2(a JavaUtilLocaleInterface, b string, c ...*JavaLangObject) *JavaIoPrintWriter {
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
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "format", "java/io/PrintWriter", conv_a.Value().Cast("java/util/Locale"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintWriter append(java.lang.CharSequence)
func (jbobject *JavaIoPrintWriter) Append5(a JavaLangCharSequenceInterface) *JavaIoPrintWriter {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/PrintWriter", conv_a.Value().Cast("java/lang/CharSequence"))
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.PrintWriter append(java.lang.CharSequence, int, int)
func (jbobject *JavaIoPrintWriter) Append6(a JavaLangCharSequenceInterface, b int, c int) *JavaIoPrintWriter {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/PrintWriter", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
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
	unique_x := &JavaIoPrintWriter{}
	unique_x.Callable = dst
	return unique_x
}

// public java.io.Writer append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaIoPrintWriter) Append2(a JavaLangCharSequenceInterface, b int, c int) (*JavaIoWriter, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/Writer", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
	if err != nil {
		var zero *JavaIoWriter
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
	unique_x := &JavaIoWriter{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public java.io.Writer append(java.lang.CharSequence) throws java.io.IOException
func (jbobject *JavaIoPrintWriter) Append(a JavaLangCharSequenceInterface) (*JavaIoWriter, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/io/Writer", conv_a.Value().Cast("java/lang/CharSequence"))
	if err != nil {
		var zero *JavaIoWriter
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
	unique_x := &JavaIoWriter{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public java.lang.Appendable append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaIoPrintWriter) Append3(a JavaLangCharSequenceInterface, b int, c int) (*JavaLangAppendable, error) {
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
func (jbobject *JavaIoPrintWriter) Append4(a JavaLangCharSequenceInterface) (*JavaLangAppendable, error) {
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


