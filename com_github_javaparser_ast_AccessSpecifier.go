package javaparser

import "github.com/timob/javabind"

type AstAccessSpecifierInterface interface {

	// public static com.github.javaparser.ast.AccessSpecifier[] values()
	Values() []*AstAccessSpecifier

	// public static com.github.javaparser.ast.AccessSpecifier valueOf(java.lang.String)
	ValueOf(a string) *AstAccessSpecifier

	// public java.lang.String getCodeRepresenation()
	GetCodeRepresenation() string
}

type AstAccessSpecifier struct {
	*javabind.Callable
}

// public static com.github.javaparser.ast.AccessSpecifier[] values()
func (jbobject *AstAccessSpecifier) Values() []*AstAccessSpecifier {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/AccessSpecifier", "values", javabind.ObjectArrayType("com/github/javaparser/ast/AccessSpecifier"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "com/github/javaparser/ast/AccessSpecifier")
	dst := new([]*AstAccessSpecifier)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static com.github.javaparser.ast.AccessSpecifier valueOf(java.lang.String)
func (jbobject *AstAccessSpecifier) ValueOf(a string) *AstAccessSpecifier {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/AccessSpecifier", "valueOf", "com/github/javaparser/ast/AccessSpecifier", conv_a.Value().Cast("java/lang/String"))
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getCodeRepresenation()
func (jbobject *AstAccessSpecifier) GetCodeRepresenation() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCodeRepresenation", "java/lang/String")
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

func (jbobject *AstAccessSpecifier) PUBLIC() *AstAccessSpecifier {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/AccessSpecifier", "PUBLIC", "com/github/javaparser/ast/AccessSpecifier")
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstAccessSpecifier) SetFieldPUBLIC(val AstAccessSpecifierInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/AccessSpecifier", "PUBLIC", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *AstAccessSpecifier) PRIVATE() *AstAccessSpecifier {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/AccessSpecifier", "PRIVATE", "com/github/javaparser/ast/AccessSpecifier")
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstAccessSpecifier) SetFieldPRIVATE(val AstAccessSpecifierInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/AccessSpecifier", "PRIVATE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *AstAccessSpecifier) PROTECTED() *AstAccessSpecifier {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/AccessSpecifier", "PROTECTED", "com/github/javaparser/ast/AccessSpecifier")
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstAccessSpecifier) SetFieldPROTECTED(val AstAccessSpecifierInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/AccessSpecifier", "PROTECTED", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *AstAccessSpecifier) DEFAULT() *AstAccessSpecifier {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/AccessSpecifier", "DEFAULT", "com/github/javaparser/ast/AccessSpecifier")
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstAccessSpecifier) SetFieldDEFAULT(val AstAccessSpecifierInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/AccessSpecifier", "DEFAULT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


