package javaparser

import "github.com/timob/javabind"

type TokenInterface interface {
	JavaLangObjectInterface

	// public java.lang.Object getValue()
	GetValue() *JavaLangObject

	// public static com.github.javaparser.Token newToken(int, java.lang.String)
	NewToken(a int, b string) *Token

	// public static com.github.javaparser.Token newToken(int)
	NewToken2(a int) *Token
}

type Token struct {
	JavaLangObject
}

// public com.github.javaparser.Token()
func NewToken() (*Token) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/Token")
	if err != nil {
		panic(err)
	}
	x := &Token{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.Token(int)
func NewToken2(a int) (*Token) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/Token", a)
	if err != nil {
		panic(err)
	}
	x := &Token{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.Token(int, java.lang.String)
func NewToken3(a int, b string) (*Token) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/Token", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &Token{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.Object getValue()
func (jbobject *Token) GetValue() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "java/lang/Object")
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

// public java.lang.String toString()
func (jbobject *Token) ToString() string {
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

// public static com.github.javaparser.Token newToken(int, java.lang.String)
func (jbobject *Token) NewToken(a int, b string) *Token {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/Token", "newToken", "com/github/javaparser/Token", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &Token{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.Token newToken(int)
func (jbobject *Token) NewToken2(a int) *Token {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/Token", "newToken", "com/github/javaparser/Token", a)
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
	unique_x := &Token{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Token) Kind() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "kind", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Token) SetFieldKind(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "kind", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *Token) BeginLine() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "beginLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Token) SetFieldBeginLine(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "beginLine", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *Token) BeginColumn() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "beginColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Token) SetFieldBeginColumn(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "beginColumn", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *Token) EndLine() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "endLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Token) SetFieldEndLine(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "endLine", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *Token) EndColumn() int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "endColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Token) SetFieldEndColumn(val int) {
	err := jbobject.SetField(javabind.GetEnv(), "endColumn", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *Token) Image() string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "image", "java/lang/String")
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

func (jbobject *Token) SetFieldImage(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "image", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *Token) Next() *Token {
	jret, err := jbobject.GetField(javabind.GetEnv(), "next", "com/github/javaparser/Token")
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
	unique_x := &Token{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Token) SetFieldNext(val TokenInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "next", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *Token) SpecialToken() *Token {
	jret, err := jbobject.GetField(javabind.GetEnv(), "specialToken", "com/github/javaparser/Token")
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
	unique_x := &Token{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Token) SetFieldSpecialToken(val TokenInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "specialToken", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


