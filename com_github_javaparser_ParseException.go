package javaparser

import "github.com/timob/javabind"

type ParseExceptionInterface interface {
}

type ParseException struct {
	*javabind.Callable
}

// public com.github.javaparser.ParseException(com.github.javaparser.Token, int[][], java.lang.String[])
func NewParseException(a TokenInterface, b [][]int, c []string) (*ParseException) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "int"), "int[]")
	conv_c := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ParseException", conv_a.Value().Cast("com/github/javaparser/Token"), conv_b.Value().Cast("[]"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &ParseException{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ParseException()
func NewParseException2() (*ParseException) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ParseException")
	if err != nil {
		panic(err)
	}
	x := &ParseException{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ParseException(java.lang.String)
func NewParseException3(a string) (*ParseException) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ParseException", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ParseException{}
	x.Callable = &javabind.Callable{obj}
	return x
}

func (jbobject *ParseException) CurrentToken() *Token {
	jret, err := jbobject.GetField(javabind.GetEnv(), "currentToken", "com/github/javaparser/Token")
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

func (jbobject *ParseException) SetFieldCurrentToken(val TokenInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "currentToken", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ParseException) ExpectedTokenSequences() [][]int {
	jret, err := jbobject.GetField(javabind.GetEnv(), "expectedTokenSequences", javabind.ObjectArrayType("int[]"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "int"), "int[]")
	dst := new([][]int)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *ParseException) SetFieldExpectedTokenSequences(val [][]int) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaCallable(), "int"), "int[]")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "expectedTokenSequences", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ParseException) TokenImage() []string {
	jret, err := jbobject.GetField(javabind.GetEnv(), "tokenImage", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

func (jbobject *ParseException) SetFieldTokenImage(val []string) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "tokenImage", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


