package javaparser

import "github.com/timob/javabind"

type AstTypeClassOrInterfaceTypeInterface interface {
	AstTypeTypeInterface

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.type.ClassOrInterfaceType getScope()
	GetScope() *AstTypeClassOrInterfaceType

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
	GetTypeArgs() []*AstTypeType

	// public com.github.javaparser.ast.TypeArguments getTypeArguments()
	GetTypeArguments() *AstTypeArguments

	// public boolean isUsingDiamondOperator()
	IsUsingDiamondOperator() bool

	// public boolean isBoxedType()
	IsBoxedType() bool

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setScope(com.github.javaparser.ast.type.ClassOrInterfaceType)
	SetScope(a AstTypeClassOrInterfaceTypeInterface) 

	// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
	SetTypeArgs(a []*AstTypeType) 

	// public void setTypeArguments(com.github.javaparser.ast.TypeArguments)
	SetTypeArguments(a AstTypeArgumentsInterface) 
}

type AstTypeClassOrInterfaceType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType()
func NewAstTypeClassOrInterfaceType() (*AstTypeClassOrInterfaceType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ClassOrInterfaceType")
	if err != nil {
		panic(err)
	}
	x := &AstTypeClassOrInterfaceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType(java.lang.String)
func NewAstTypeClassOrInterfaceType2(a string) (*AstTypeClassOrInterfaceType) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ClassOrInterfaceType", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeClassOrInterfaceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.String)
func NewAstTypeClassOrInterfaceType3(a AstTypeClassOrInterfaceTypeInterface, b string) (*AstTypeClassOrInterfaceType) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ClassOrInterfaceType", conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstTypeClassOrInterfaceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType(int, int, int, int, com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.String, java.util.List<com.github.javaparser.ast.type.Type>)
func NewAstTypeClassOrInterfaceType4(a int, b int, c int, d int, e AstTypeClassOrInterfaceTypeInterface, f string, g []*AstTypeType) (*AstTypeClassOrInterfaceType) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaString()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ClassOrInterfaceType", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_f.Value().Cast("java/lang/String"), conv_g.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstTypeClassOrInterfaceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType(int, int, int, int, com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.String, com.github.javaparser.ast.TypeArguments)
func NewAstTypeClassOrInterfaceType5(a int, b int, c int, d int, e AstTypeClassOrInterfaceTypeInterface, f string, g AstTypeArgumentsInterface) (*AstTypeClassOrInterfaceType) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaString()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ClassOrInterfaceType", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_f.Value().Cast("java/lang/String"), conv_g.Value().Cast("com/github/javaparser/ast/TypeArguments"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstTypeClassOrInterfaceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getName()
func (jbobject *AstTypeClassOrInterfaceType) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public com.github.javaparser.ast.type.ClassOrInterfaceType getScope()
func (jbobject *AstTypeClassOrInterfaceType) GetScope() *AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScope", "com/github/javaparser/ast/type/ClassOrInterfaceType")
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
	unique_x := &AstTypeClassOrInterfaceType{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
func (jbobject *AstTypeClassOrInterfaceType) GetTypeArgs() []*AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeArgs", "java/util/List")
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

// public com.github.javaparser.ast.TypeArguments getTypeArguments()
func (jbobject *AstTypeClassOrInterfaceType) GetTypeArguments() *AstTypeArguments {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeArguments", "com/github/javaparser/ast/TypeArguments")
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

// public boolean isUsingDiamondOperator()
func (jbobject *AstTypeClassOrInterfaceType) IsUsingDiamondOperator() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isUsingDiamondOperator", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isBoxedType()
func (jbobject *AstTypeClassOrInterfaceType) IsBoxedType() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isBoxedType", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public com.github.javaparser.ast.type.PrimitiveType toUnboxedType() throws java.lang.UnsupportedOperationException
func (jbobject *AstTypeClassOrInterfaceType) ToUnboxedType() (*AstTypePrimitiveType, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toUnboxedType", "com/github/javaparser/ast/type/PrimitiveType")
	if err != nil {
		var zero *AstTypePrimitiveType
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public void setName(java.lang.String)
func (jbobject *AstTypeClassOrInterfaceType) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setScope(com.github.javaparser.ast.type.ClassOrInterfaceType)
func (jbobject *AstTypeClassOrInterfaceType) SetScope(a AstTypeClassOrInterfaceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScope", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
func (jbobject *AstTypeClassOrInterfaceType) SetTypeArgs(a []*AstTypeType)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeArgs", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTypeArguments(com.github.javaparser.ast.TypeArguments)
func (jbobject *AstTypeClassOrInterfaceType) SetTypeArguments(a AstTypeArgumentsInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeArguments", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/TypeArguments"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


