package javaparser

import "github.com/timob/javabind"

type AstBodyParameterInterface interface {
	AstBodyBaseParameterInterface

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public boolean isVarArgs()
	IsVarArgs() bool

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public void setVarArgs(boolean)
	SetVarArgs(a bool) 
}

type AstBodyParameter struct {
	AstBodyBaseParameter
}

// public com.github.javaparser.ast.body.Parameter()
func NewAstBodyParameter() (*AstBodyParameter) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/Parameter")
	if err != nil {
		panic(err)
	}
	x := &AstBodyParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.Parameter(com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyParameter2(a AstTypeTypeInterface, b AstBodyVariableDeclaratorIdInterface) (*AstBodyParameter) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/Parameter", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstBodyParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.Parameter(int, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyParameter3(a int, b AstTypeTypeInterface, c AstBodyVariableDeclaratorIdInterface) (*AstBodyParameter) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/Parameter", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstBodyParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.Parameter(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, boolean, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyParameter4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeTypeInterface, h bool, i AstBodyVariableDeclaratorIdInterface) (*AstBodyParameter) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_i := javabind.NewGoToJavaCallable()
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/Parameter", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/Type"), h, conv_i.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_i.CleanUp()
	x := &AstBodyParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstBodyParameter) GetType() *AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/Type")
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
	unique_x := &AstTypeType{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isVarArgs()
func (jbobject *AstBodyParameter) IsVarArgs() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isVarArgs", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstBodyParameter) SetType(a AstTypeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setVarArgs(boolean)
func (jbobject *AstBodyParameter) SetVarArgs(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVarArgs", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


