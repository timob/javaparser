package javaparser

import "github.com/timob/javabind"

type AstTypeArgumentsInterface interface {
	JavaLangObjectInterface

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArguments()
	GetTypeArguments() []*AstTypeType

	// public boolean isUsingDiamondOperator()
	IsUsingDiamondOperator() bool

	// public static com.github.javaparser.ast.TypeArguments withDiamondOperator()
	WithDiamondOperator() *AstTypeArguments

	// public static com.github.javaparser.ast.TypeArguments withArguments(java.util.List<com.github.javaparser.ast.type.Type>)
	WithArguments(a []*AstTypeType) *AstTypeArguments
}

type AstTypeArguments struct {
	JavaLangObject
}

// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArguments()
func (jbobject *AstTypeArguments) GetTypeArguments() []*AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeArguments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean isUsingDiamondOperator()
func (jbobject *AstTypeArguments) IsUsingDiamondOperator() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUsingDiamondOperator", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static com.github.javaparser.ast.TypeArguments withDiamondOperator()
func (jbobject *AstTypeArguments) WithDiamondOperator() *AstTypeArguments {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/TypeArguments", "withDiamondOperator", "com/github/javaparser/ast/TypeArguments")
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
	unique_x := &AstTypeArguments{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.TypeArguments withArguments(java.util.List<com.github.javaparser.ast.type.Type>)
func (jbobject *AstTypeArguments) WithArguments(a []*AstTypeType) *AstTypeArguments {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/TypeArguments", "withArguments", "com/github/javaparser/ast/TypeArguments", conv_a.Value().Cast("java/util/List"))
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
	unique_x := &AstTypeArguments{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstTypeArguments) EMPTY() *AstTypeArguments {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/TypeArguments", "EMPTY", "com/github/javaparser/ast/TypeArguments")
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
	unique_x := &AstTypeArguments{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *AstTypeArguments) SetFieldEMPTY(val AstTypeArgumentsInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/TypeArguments", "EMPTY", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


