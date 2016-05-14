package javaparser

import "github.com/timob/javabind"

type JavaIoByteArrayInputStreamInterface interface {
	JavaIoInputStreamInterface
}

type JavaIoByteArrayInputStream struct {
	JavaIoInputStream
}

// public java.io.ByteArrayInputStream(byte[])
func NewJavaIoByteArrayInputStream(a []byte) (*JavaIoByteArrayInputStream) {

	obj, err := javabind.GetEnv().NewObject("java/io/ByteArrayInputStream", a)
	if err != nil {
		panic(err)
	}
	x := &JavaIoByteArrayInputStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.ByteArrayInputStream(byte[], int, int)
func NewJavaIoByteArrayInputStream2(a []byte, b int, c int) (*JavaIoByteArrayInputStream) {

	obj, err := javabind.GetEnv().NewObject("java/io/ByteArrayInputStream", a, b, c)
	if err != nil {
		panic(err)
	}
	x := &JavaIoByteArrayInputStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public synchronized int read()
func (jbobject *JavaIoByteArrayInputStream) Read() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public synchronized int read(byte[], int, int)
func (jbobject *JavaIoByteArrayInputStream) Read3(a []byte, b int, c int) int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "read", javabind.Int, a, b, c)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public synchronized long skip(long)
func (jbobject *JavaIoByteArrayInputStream) Skip(a int64) int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "skip", javabind.Long, a)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public synchronized int available()
func (jbobject *JavaIoByteArrayInputStream) Available() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "available", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean markSupported()
func (jbobject *JavaIoByteArrayInputStream) MarkSupported() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "markSupported", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void mark(int)
func (jbobject *JavaIoByteArrayInputStream) Mark(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "mark", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public synchronized void reset()
func (jbobject *JavaIoByteArrayInputStream) Reset()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "reset", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void close() throws java.io.IOException
func (jbobject *JavaIoByteArrayInputStream) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}


