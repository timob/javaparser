package javaparser

import "github.com/timob/javabind"

type JavaIoFilterOutputStreamInterface interface {
	JavaIoOutputStreamInterface
}

type JavaIoFilterOutputStream struct {
	JavaIoOutputStream
}

// public java.io.FilterOutputStream(java.io.OutputStream)
func NewJavaIoFilterOutputStream(a JavaIoOutputStreamInterface) (*JavaIoFilterOutputStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/FilterOutputStream", conv_a.Value().Cast("java/io/OutputStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoFilterOutputStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void write(int) throws java.io.IOException
func (jbobject *JavaIoFilterOutputStream) Write(a int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void write(byte[]) throws java.io.IOException
func (jbobject *JavaIoFilterOutputStream) Write2(a []byte) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a)
	if err != nil {
		return err
	}
	return nil
}

// public void write(byte[], int, int) throws java.io.IOException
func (jbobject *JavaIoFilterOutputStream) Write3(a []byte, b int, c int) error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "write", javabind.Void, a, b, c)
	if err != nil {
		return err
	}
	return nil
}

// public void flush() throws java.io.IOException
func (jbobject *JavaIoFilterOutputStream) Flush() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "flush", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}

// public void close() throws java.io.IOException
func (jbobject *JavaIoFilterOutputStream) Close() error {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "close", javabind.Void)
	if err != nil {
		return err
	}
	return nil
}


