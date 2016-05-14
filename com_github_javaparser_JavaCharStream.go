package javaparser

import "github.com/timob/javabind"

type JavaCharStreamInterface interface {
	JavaLangObjectInterface

	// public int getColumn()
	GetColumn() int

	// public int getLine()
	GetLine() int

	// public int getEndColumn()
	GetEndColumn() int

	// public int getEndLine()
	GetEndLine() int

	// public int getBeginColumn()
	GetBeginColumn() int

	// public int getBeginLine()
	GetBeginLine() int

	// public void backup(int)
	Backup(a int) 

	// public void ReInit(java.io.Reader, int, int, int)
	ReInit(a JavaIoReaderInterface, b int, c int, d int) 

	// public void ReInit(java.io.Reader, int, int)
	ReInit2(a JavaIoReaderInterface, b int, c int) 

	// public void ReInit(java.io.Reader)
	ReInit3(a JavaIoReaderInterface) 

	// public void ReInit(java.io.InputStream, int, int, int)
	ReInit5(a JavaIoInputStreamInterface, b int, c int, d int) 

	// public void ReInit(java.io.InputStream, int, int)
	ReInit7(a JavaIoInputStreamInterface, b int, c int) 

	// public void ReInit(java.io.InputStream)
	ReInit9(a JavaIoInputStreamInterface) 

	// public java.lang.String GetImage()
	GetImage() string

	// public void Done()
	Done() 

	// public void adjustBeginLineColumn(int, int)
	AdjustBeginLineColumn(a int, b int) 
}

type JavaCharStream struct {
	JavaLangObject
}

// public com.github.javaparser.JavaCharStream(java.io.Reader, int, int, int)
func NewJavaCharStream(a JavaIoReaderInterface, b int, c int, d int) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/Reader"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.JavaCharStream(java.io.Reader, int, int)
func NewJavaCharStream2(a JavaIoReaderInterface, b int, c int) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/Reader"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.JavaCharStream(java.io.Reader)
func NewJavaCharStream3(a JavaIoReaderInterface) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/Reader"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream, java.lang.String, int, int, int) throws java.io.UnsupportedEncodingException
func NewJavaCharStream4(a JavaIoInputStreamInterface, b string, c int, d int, e int) (*JavaCharStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"), c, d, e)
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream, int, int, int)
func NewJavaCharStream5(a JavaIoInputStreamInterface, b int, c int, d int) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream, java.lang.String, int, int) throws java.io.UnsupportedEncodingException
func NewJavaCharStream6(a JavaIoInputStreamInterface, b string, c int, d int) (*JavaCharStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"), c, d)
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream, int, int)
func NewJavaCharStream7(a JavaIoInputStreamInterface, b int, c int) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream, java.lang.String) throws java.io.UnsupportedEncodingException
func NewJavaCharStream8(a JavaIoInputStreamInterface, b string) (*JavaCharStream, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return nil, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x, nil
}

// public com.github.javaparser.JavaCharStream(java.io.InputStream)
func NewJavaCharStream9(a JavaIoInputStreamInterface) (*JavaCharStream) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/JavaCharStream", conv_a.Value().Cast("java/io/InputStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaCharStream{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getColumn()
func (jbobject *JavaCharStream) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getLine()
func (jbobject *JavaCharStream) GetLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getEndColumn()
func (jbobject *JavaCharStream) GetEndColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getEndLine()
func (jbobject *JavaCharStream) GetEndLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getBeginColumn()
func (jbobject *JavaCharStream) GetBeginColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBeginColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getBeginLine()
func (jbobject *JavaCharStream) GetBeginLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBeginLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public void backup(int)
func (jbobject *JavaCharStream) Backup(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "backup", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void ReInit(java.io.Reader, int, int, int)
func (jbobject *JavaCharStream) ReInit(a JavaIoReaderInterface, b int, c int, d int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/Reader"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(java.io.Reader, int, int)
func (jbobject *JavaCharStream) ReInit2(a JavaIoReaderInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/Reader"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(java.io.Reader)
func (jbobject *JavaCharStream) ReInit3(a JavaIoReaderInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/Reader"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(java.io.InputStream, java.lang.String, int, int, int) throws java.io.UnsupportedEncodingException
func (jbobject *JavaCharStream) ReInit4(a JavaIoInputStreamInterface, b string, c int, d int, e int) error {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"), c, d, e)
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return nil
}

// public void ReInit(java.io.InputStream, int, int, int)
func (jbobject *JavaCharStream) ReInit5(a JavaIoInputStreamInterface, b int, c int, d int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"), b, c, d)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(java.io.InputStream, java.lang.String, int, int) throws java.io.UnsupportedEncodingException
func (jbobject *JavaCharStream) ReInit6(a JavaIoInputStreamInterface, b string, c int, d int) error {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"), c, d)
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return nil
}

// public void ReInit(java.io.InputStream, int, int)
func (jbobject *JavaCharStream) ReInit7(a JavaIoInputStreamInterface, b int, c int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(java.io.InputStream, java.lang.String) throws java.io.UnsupportedEncodingException
func (jbobject *JavaCharStream) ReInit8(a JavaIoInputStreamInterface, b string) error {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		return err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return nil
}

// public void ReInit(java.io.InputStream)
func (jbobject *JavaCharStream) ReInit9(a JavaIoInputStreamInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("java/io/InputStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.lang.String GetImage()
func (jbobject *JavaCharStream) GetImage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "GetImage", "java/lang/String")
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

// public void Done()
func (jbobject *JavaCharStream) Done()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "Done", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public void adjustBeginLineColumn(int, int)
func (jbobject *JavaCharStream) AdjustBeginLineColumn(a int, b int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "adjustBeginLineColumn", javabind.Void, a, b)
	if err != nil {
		panic(err)
	}

}

func (jbobject *JavaCharStream) StaticFlag() bool {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/JavaCharStream", "staticFlag", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

func (jbobject *JavaCharStream) SetFieldStaticFlag(val bool) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/JavaCharStream", "staticFlag", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *JavaCharStream) Bufpos() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "bufpos", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *JavaCharStream) SetFieldBufpos(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "bufpos", val)
	if err != nil {
		panic(err)
	}

}


