package javaparser

import "github.com/timob/javabind"

type JavaIoInputStreamInterface interface {
	JavaLangObjectInterface

	// public synchronized void mark(int)
	Mark(a int) 

	// public boolean markSupported()
	MarkSupported() bool
}

type JavaIoInputStream struct {
	JavaLangObject
}

// public java.io.InputStream()
func NewJavaIoInputStream() (*JavaIoInputStream) {

	obj, err := javabind.GetEnv().NewObject("java/io/InputStream")
	if err != nil {
		panic(err)
	}
	x := &JavaIoInputStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract int read() throws java.io.IOException
func (jbobject *JavaIoInputStream) Read() (int, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int)
	if err != nil {
		var zero int
		return zero, err
	}
	return jret.(int), nil
}

// public int read(byte[]) throws java.io.IOException
func (jbobject *JavaIoInputStream) Read2(a []byte) (int, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int, a)
	if err != nil {
		var zero int
		return zero, err
	}
	return jret.(int), nil
}

// public int read(byte[], int, int) throws java.io.IOException
func (jbobject *JavaIoInputStream) Read3(a []byte, b int, c int) (int, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int, a, b, c)
	if err != nil {
		var zero int
		return zero, err
	}
	return jret.(int), nil
}

// public long skip(long) throws java.io.IOException
func (jbobject *JavaIoInputStream) Skip(a int64) (int64, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "skip", javabind.Long, a)
	if err != nil {
		var zero int64
		return zero, err
	}
	return jret.(int64), nil
}

// public int available() throws java.io.IOException
func (jbobject *JavaIoInputStream) Available() (int, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "available", javabind.Int)
	if err != nil {
		var zero int
		return zero, err
	}
	return jret.(int), nil
}

// public void close() throws java.io.IOException
func (jbobject *JavaIoInputStream) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public synchronized void mark(int)
func (jbobject *JavaIoInputStream) Mark(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mark", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public synchronized void reset() throws java.io.IOException
func (jbobject *JavaIoInputStream) Reset() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reset", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public boolean markSupported()
func (jbobject *JavaIoInputStream) MarkSupported() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "markSupported", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


