package javaparser

import "github.com/timob/javabind"

type JavaIoWriterInterface interface {
	JavaLangObjectInterface
}

type JavaIoWriter struct {
	JavaLangObject
}

// public void write(int) throws java.io.IOException
func (jbobject *JavaIoWriter) Write(a int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void write(java.lang.String) throws java.io.IOException
func (jbobject *JavaIoWriter) Write2(a string) error {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	return nil
}

// public void write(java.lang.String, int, int) throws java.io.IOException
func (jbobject *JavaIoWriter) Write3(a string, b int, c int) error {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, conv_a.Value().Cast("java/lang/String"), b, c)
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	return nil
}

// public java.io.Writer append(java.lang.CharSequence) throws java.io.IOException
func (jbobject *JavaIoWriter) Append(a JavaLangCharSequenceInterface) (*JavaIoWriter, error) {
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

// public java.io.Writer append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaIoWriter) Append2(a JavaLangCharSequenceInterface, b int, c int) (*JavaIoWriter, error) {
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

// public abstract void flush() throws java.io.IOException
func (jbobject *JavaIoWriter) Flush() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "flush", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public abstract void close() throws java.io.IOException
func (jbobject *JavaIoWriter) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public java.lang.Appendable append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaIoWriter) Append3(a JavaLangCharSequenceInterface, b int, c int) (*JavaLangAppendable, error) {
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
func (jbobject *JavaIoWriter) Append4(a JavaLangCharSequenceInterface) (*JavaLangAppendable, error) {
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


