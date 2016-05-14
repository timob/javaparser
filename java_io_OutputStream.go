package javaparser

import "github.com/timob/javabind"

type JavaIoOutputStreamInterface interface {
	JavaLangObjectInterface
}

type JavaIoOutputStream struct {
	JavaLangObject
}

// public java.io.OutputStream()
func NewJavaIoOutputStream() (*JavaIoOutputStream) {

	obj, err := javabind.GetEnv().NewObject("java/io/OutputStream")
	if err != nil {
		panic(err)
	}
	x := &JavaIoOutputStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public abstract void write(int) throws java.io.IOException
func (jbobject *JavaIoOutputStream) Write(a int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void write(byte[]) throws java.io.IOException
func (jbobject *JavaIoOutputStream) Write2(a []byte) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void write(byte[], int, int) throws java.io.IOException
func (jbobject *JavaIoOutputStream) Write3(a []byte, b int, c int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a, b, c)
	if err != nil {
		return err
	}
	return nil
}

// public void flush() throws java.io.IOException
func (jbobject *JavaIoOutputStream) Flush() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "flush", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public void close() throws java.io.IOException
func (jbobject *JavaIoOutputStream) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}


