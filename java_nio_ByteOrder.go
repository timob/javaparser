package javaparser

import "github.com/timob/javabind"

type JavaNioByteOrderInterface interface {
	JavaLangObjectInterface

	// public static java.nio.ByteOrder nativeOrder()
	NativeOrder() *JavaNioByteOrder
}

type JavaNioByteOrder struct {
	JavaLangObject
}

// public static java.nio.ByteOrder nativeOrder()
func (jbobject *JavaNioByteOrder) NativeOrder() *JavaNioByteOrder {
	jret, err := javabind.GetEnv().CallStaticMethod("java/nio/ByteOrder", "nativeOrder", "java/nio/ByteOrder")
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

// public java.lang.String toString()
func (jbobject *JavaNioByteOrder) ToString() string {
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

func (jbobject *JavaNioByteOrder) BIG_ENDIAN() *JavaNioByteOrder {
	jret, err := javabind.GetEnv().GetStaticField("java/nio/ByteOrder", "BIG_ENDIAN", "java/nio/ByteOrder")
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

func (jbobject *JavaNioByteOrder) SetFieldBIG_ENDIAN(val JavaNioByteOrderInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/nio/ByteOrder", "BIG_ENDIAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaNioByteOrder) LITTLE_ENDIAN() *JavaNioByteOrder {
	jret, err := javabind.GetEnv().GetStaticField("java/nio/ByteOrder", "LITTLE_ENDIAN", "java/nio/ByteOrder")
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

func (jbobject *JavaNioByteOrder) SetFieldLITTLE_ENDIAN(val JavaNioByteOrderInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/nio/ByteOrder", "LITTLE_ENDIAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


