package javaparser

import "github.com/timob/javabind"

type ASTParserTokenManagerInterface interface {
	JavaLangObjectInterface

	// public void setDebugStream(java.io.PrintStream)
	SetDebugStream(a JavaIoPrintStreamInterface) 

	// public void ReInit(com.github.javaparser.JavaCharStream)
	ReInit(a JavaCharStreamInterface) 

	// public void ReInit(com.github.javaparser.JavaCharStream, int)
	ReInit2(a JavaCharStreamInterface, b int) 

	// public void SwitchTo(int)
	SwitchTo(a int) 

	// public com.github.javaparser.Token getNextToken()
	GetNextToken() *Token
}

type ASTParserTokenManager struct {
	JavaLangObject
}

// public com.github.javaparser.ASTParserTokenManager(com.github.javaparser.JavaCharStream)
func NewASTParserTokenManager(a JavaCharStreamInterface) (*ASTParserTokenManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ASTParserTokenManager", conv_a.Value().Cast("com/github/javaparser/JavaCharStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ASTParserTokenManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ASTParserTokenManager(com.github.javaparser.JavaCharStream, int)
func NewASTParserTokenManager2(a JavaCharStreamInterface, b int) (*ASTParserTokenManager) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ASTParserTokenManager", conv_a.Value().Cast("com/github/javaparser/JavaCharStream"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &ASTParserTokenManager{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setDebugStream(java.io.PrintStream)
func (jbobject *ASTParserTokenManager) SetDebugStream(a JavaIoPrintStreamInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDebugStream", javabind.Void, conv_a.Value().Cast("java/io/PrintStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(com.github.javaparser.JavaCharStream)
func (jbobject *ASTParserTokenManager) ReInit(a JavaCharStreamInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/JavaCharStream"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void ReInit(com.github.javaparser.JavaCharStream, int)
func (jbobject *ASTParserTokenManager) ReInit2(a JavaCharStreamInterface, b int)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "ReInit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/JavaCharStream"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void SwitchTo(int)
func (jbobject *ASTParserTokenManager) SwitchTo(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "SwitchTo", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public com.github.javaparser.Token getNextToken()
func (jbobject *ASTParserTokenManager) GetNextToken() *Token {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNextToken", "com/github/javaparser/Token")
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

func (jbobject *ASTParserTokenManager) DebugStream() *JavaIoPrintStream {
	jret, err := jbobject.GetField(javabind.GetEnv(), "debugStream", "java/io/PrintStream")
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
	unique_x := &JavaIoPrintStream{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTParserTokenManager) SetFieldDebugStream(val JavaIoPrintStreamInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := jbobject.SetField(javabind.GetEnv(), "debugStream", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTParserTokenManager) JjstrLiteralImages() []string {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserTokenManager", "jjstrLiteralImages", javabind.ObjectArrayType("java/lang/String"))
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

func (jbobject *ASTParserTokenManager) SetFieldJjstrLiteralImages(val []string) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserTokenManager", "jjstrLiteralImages", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTParserTokenManager) LexStateNames() []string {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserTokenManager", "lexStateNames", javabind.ObjectArrayType("java/lang/String"))
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

func (jbobject *ASTParserTokenManager) SetFieldLexStateNames(val []string) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserTokenManager", "lexStateNames", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTParserTokenManager) JjnewLexState() []int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserTokenManager", "jjnewLexState", javabind.Int | javabind.Array)
	if err != nil {
		panic(err)
	}
	return jret.([]int)
}

func (jbobject *ASTParserTokenManager) SetFieldJjnewLexState(val []int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserTokenManager", "jjnewLexState", val)
	if err != nil {
		panic(err)
	}

}


