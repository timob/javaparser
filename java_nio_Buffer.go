package javaparser

import "github.com/timob/javabind"

type JavaNioBufferInterface interface {
	JavaLangObjectInterface

	// public final int capacity()
	Capacity() int

	// public final int position()
	Position() int

	// public final java.nio.Buffer position(int)
	Position2(a int) *JavaNioBuffer

	// public final int limit()
	Limit() int

	// public final java.nio.Buffer limit(int)
	Limit2(a int) *JavaNioBuffer

	// public final java.nio.Buffer mark()
	Mark() *JavaNioBuffer

	// public final java.nio.Buffer reset()
	Reset() *JavaNioBuffer

	// public final java.nio.Buffer clear()
	Clear() *JavaNioBuffer

	// public final java.nio.Buffer flip()
	Flip() *JavaNioBuffer

	// public final java.nio.Buffer rewind()
	Rewind() *JavaNioBuffer

	// public final int remaining()
	Remaining() int

	// public final boolean hasRemaining()
	HasRemaining() bool

	// public abstract boolean isReadOnly()
	IsReadOnly() bool

	// public abstract boolean hasArray()
	HasArray() bool

	// public abstract java.lang.Object array()
	Array() *JavaLangObject

	// public abstract int arrayOffset()
	ArrayOffset() int

	// public abstract boolean isDirect()
	IsDirect() bool
}

type JavaNioBuffer struct {
	JavaLangObject
}

// public final int capacity()
func (jbobject *JavaNioBuffer) Capacity() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "capacity", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final int position()
func (jbobject *JavaNioBuffer) Position() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "position", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final java.nio.Buffer position(int)
func (jbobject *JavaNioBuffer) Position2(a int) *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "position", "java/nio/Buffer", a)
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final int limit()
func (jbobject *JavaNioBuffer) Limit() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "limit", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final java.nio.Buffer limit(int)
func (jbobject *JavaNioBuffer) Limit2(a int) *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "limit", "java/nio/Buffer", a)
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.Buffer mark()
func (jbobject *JavaNioBuffer) Mark() *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mark", "java/nio/Buffer")
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.Buffer reset()
func (jbobject *JavaNioBuffer) Reset() *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "reset", "java/nio/Buffer")
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.Buffer clear()
func (jbobject *JavaNioBuffer) Clear() *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clear", "java/nio/Buffer")
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.Buffer flip()
func (jbobject *JavaNioBuffer) Flip() *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "flip", "java/nio/Buffer")
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.nio.Buffer rewind()
func (jbobject *JavaNioBuffer) Rewind() *JavaNioBuffer {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "rewind", "java/nio/Buffer")
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
	unique_x := &JavaNioBuffer{}
	unique_x.Callable = dst
	return unique_x
}

// public final int remaining()
func (jbobject *JavaNioBuffer) Remaining() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "remaining", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final boolean hasRemaining()
func (jbobject *JavaNioBuffer) HasRemaining() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasRemaining", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract boolean isReadOnly()
func (jbobject *JavaNioBuffer) IsReadOnly() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isReadOnly", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract boolean hasArray()
func (jbobject *JavaNioBuffer) HasArray() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasArray", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public abstract java.lang.Object array()
func (jbobject *JavaNioBuffer) Array() *JavaLangObject {
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

// public abstract int arrayOffset()
func (jbobject *JavaNioBuffer) ArrayOffset() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "arrayOffset", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public abstract boolean isDirect()
func (jbobject *JavaNioBuffer) IsDirect() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDirect", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}


