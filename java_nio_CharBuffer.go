package javaparser

import "github.com/timob/javabind"

type JavaNioCharBufferInterface interface {
	JavaNioBufferInterface

	// public static java.nio.CharBuffer allocate(int)
	Allocate(a int) *JavaNioCharBuffer

	// public static java.nio.CharBuffer wrap(java.lang.CharSequence, int, int)
	Wrap(a JavaLangCharSequenceInterface, b int, c int) *JavaNioCharBuffer

	// public static java.nio.CharBuffer wrap(java.lang.CharSequence)
	Wrap2(a JavaLangCharSequenceInterface) *JavaNioCharBuffer

	// public abstract java.nio.CharBuffer slice()
	Slice() *JavaNioCharBuffer

	// public abstract java.nio.CharBuffer duplicate()
	Duplicate() *JavaNioCharBuffer

	// public abstract java.nio.CharBuffer asReadOnlyBuffer()
	AsReadOnlyBuffer() *JavaNioCharBuffer

	// public java.nio.CharBuffer put(java.nio.CharBuffer)
	Put(a JavaNioCharBufferInterface) *JavaNioCharBuffer

	// public java.nio.CharBuffer put(java.lang.String, int, int)
	Put2(a string, b int, c int) *JavaNioCharBuffer

	// public final java.nio.CharBuffer put(java.lang.String)
	Put3(a string) *JavaNioCharBuffer

	// public abstract java.nio.CharBuffer compact()
	Compact() *JavaNioCharBuffer

	// public int compareTo(java.nio.CharBuffer)
	CompareTo(a JavaNioCharBufferInterface) int

	// public final int length()
	Length() int

	// public abstract java.lang.CharSequence subSequence(int, int)
	SubSequence(a int, b int) *JavaLangCharSequence

	// public java.nio.CharBuffer append(java.lang.CharSequence)
	Append(a JavaLangCharSequenceInterface) *JavaNioCharBuffer

	// public java.nio.CharBuffer append(java.lang.CharSequence, int, int)
	Append2(a JavaLangCharSequenceInterface, b int, c int) *JavaNioCharBuffer

	// public abstract java.nio.ByteOrder order()
	Order() *JavaNioByteOrder

	// public int compareTo(java.lang.Object)
	CompareTo2(a interface{}) int
}

type JavaNioCharBuffer struct {
	JavaNioBuffer
}

// public static java.nio.CharBuffer allocate(int)
func (jbobject *JavaNioCharBuffer) Allocate(a int) *JavaNioCharBuffer {
	jret, err := javabind.GetEnv().CallStaticMethod("java/nio/CharBuffer", "allocate", "java/nio/CharBuffer", a)
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public int read(java.nio.CharBuffer) throws java.io.IOException
func (jbobject *JavaNioCharBuffer) Read(a JavaNioCharBufferInterface) (int, error) {
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

// public static java.nio.CharBuffer wrap(java.lang.CharSequence, int, int)
func (jbobject *JavaNioCharBuffer) Wrap(a JavaLangCharSequenceInterface, b int, c int) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/nio/CharBuffer", "wrap", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public static java.nio.CharBuffer wrap(java.lang.CharSequence)
func (jbobject *JavaNioCharBuffer) Wrap2(a JavaLangCharSequenceInterface) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/nio/CharBuffer", "wrap", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/CharSequence"))
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.nio.CharBuffer slice()
func (jbobject *JavaNioCharBuffer) Slice() *JavaNioCharBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "slice", "java/nio/CharBuffer")
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.nio.CharBuffer duplicate()
func (jbobject *JavaNioCharBuffer) Duplicate() *JavaNioCharBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "duplicate", "java/nio/CharBuffer")
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.nio.CharBuffer asReadOnlyBuffer()
func (jbobject *JavaNioCharBuffer) AsReadOnlyBuffer() *JavaNioCharBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "asReadOnlyBuffer", "java/nio/CharBuffer")
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public java.nio.CharBuffer put(java.nio.CharBuffer)
func (jbobject *JavaNioCharBuffer) Put(a JavaNioCharBufferInterface) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "put", "java/nio/CharBuffer", conv_a.Value().Cast("java/nio/CharBuffer"))
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public java.nio.CharBuffer put(java.lang.String, int, int)
func (jbobject *JavaNioCharBuffer) Put2(a string, b int, c int) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "put", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/String"), b, c)
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.CharBuffer put(java.lang.String)
func (jbobject *JavaNioCharBuffer) Put3(a string) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "put", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final boolean hasArray()
func (jbobject *JavaNioCharBuffer) HasArray() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasArray", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public final int arrayOffset()
func (jbobject *JavaNioCharBuffer) ArrayOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "arrayOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.nio.CharBuffer compact()
func (jbobject *JavaNioCharBuffer) Compact() *JavaNioCharBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compact", "java/nio/CharBuffer")
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract boolean isDirect()
func (jbobject *JavaNioCharBuffer) IsDirect() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDirect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public int hashCode()
func (jbobject *JavaNioCharBuffer) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaNioCharBuffer) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public int compareTo(java.nio.CharBuffer)
func (jbobject *JavaNioCharBuffer) CompareTo(a JavaNioCharBufferInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/nio/CharBuffer"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *JavaNioCharBuffer) ToString() string {
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

// public final int length()
func (jbobject *JavaNioCharBuffer) Length() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "length", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract java.lang.CharSequence subSequence(int, int)
func (jbobject *JavaNioCharBuffer) SubSequence(a int, b int) *JavaLangCharSequence {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "subSequence", "java/lang/CharSequence", a, b)
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
	unique_x := &JavaLangCharSequence{}
	unique_x.Callable = dst
	return unique_x
}

// public java.nio.CharBuffer append(java.lang.CharSequence)
func (jbobject *JavaNioCharBuffer) Append(a JavaLangCharSequenceInterface) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/CharSequence"))
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public java.nio.CharBuffer append(java.lang.CharSequence, int, int)
func (jbobject *JavaNioCharBuffer) Append2(a JavaLangCharSequenceInterface, b int, c int) *JavaNioCharBuffer {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "append", "java/nio/CharBuffer", conv_a.Value().Cast("java/lang/CharSequence"), b, c)
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
	unique_x := &JavaNioCharBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public abstract java.nio.ByteOrder order()
func (jbobject *JavaNioCharBuffer) Order() *JavaNioByteOrder {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "order", "java/nio/ByteOrder")
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
	unique_x := &JavaNioByteOrder{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object array()
func (jbobject *JavaNioCharBuffer) Array() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "array", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public int compareTo(java.lang.Object)
func (jbobject *JavaNioCharBuffer) CompareTo2(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public java.lang.Appendable append(java.lang.CharSequence, int, int) throws java.io.IOException
func (jbobject *JavaNioCharBuffer) Append3(a JavaLangCharSequenceInterface, b int, c int) (*JavaLangAppendable, error) {
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
func (jbobject *JavaNioCharBuffer) Append4(a JavaLangCharSequenceInterface) (*JavaLangAppendable, error) {
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


