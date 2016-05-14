package javaparser

import "github.com/timob/javabind"

type JavaIoReaderInterface interface {
	JavaLangObjectInterface

	// public boolean markSupported()
	MarkSupported() bool
}

type JavaIoReader struct {
	JavaLangObject
}

// public int read(java.nio.CharBuffer) throws java.io.IOException
func (jbobject *JavaIoReader) Read(a JavaNioCharBufferInterface) (int, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int, conv_a.Value().Cast("java/nio/CharBuffer"))
	if err != nil {
		var zero int
		return zero, err
	}
	conv_a.CleanUp()
	return jret.(int), nil
}

// public int read() throws java.io.IOException
func (jbobject *JavaIoReader) Read2() (int, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int)
	if err != nil {
		var zero int
		return zero, err
	}
	return jret.(int), nil
}

// public long skip(long) throws java.io.IOException
func (jbobject *JavaIoReader) Skip(a int64) (int64, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "skip", javabind.Long, a)
	if err != nil {
		var zero int64
		return zero, err
	}
	return jret.(int64), nil
}

// public boolean ready() throws java.io.IOException
func (jbobject *JavaIoReader) Ready() (bool, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "ready", javabind.Boolean)
	if err != nil {
		var zero bool
		return zero, err
	}
	return jret.(bool), nil
}

// public boolean markSupported()
func (jbobject *JavaIoReader) MarkSupported() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "markSupported", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void mark(int) throws java.io.IOException
func (jbobject *JavaIoReader) Mark(a int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mark", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void reset() throws java.io.IOException
func (jbobject *JavaIoReader) Reset() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reset", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public abstract void close() throws java.io.IOException
func (jbobject *JavaIoReader) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}


