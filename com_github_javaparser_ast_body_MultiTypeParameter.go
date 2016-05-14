package javaparser

import "github.com/timob/javabind"

type AstBodyMultiTypeParameterInterface interface {
	AstBodyBaseParameterInterface

	// public com.github.javaparser.ast.type.UnionType getType()
	GetType() *AstTypeUnionType

	// public void setType(com.github.javaparser.ast.type.UnionType)
	SetType(a AstTypeUnionTypeInterface) 
}

type AstBodyMultiTypeParameter struct {
	AstBodyBaseParameter
}

// public com.github.javaparser.ast.body.MultiTypeParameter()
func NewAstBodyMultiTypeParameter() (*AstBodyMultiTypeParameter) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MultiTypeParameter")
	if err != nil {
		panic(err)
	}
	x := &AstBodyMultiTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MultiTypeParameter(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.UnionType, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyMultiTypeParameter2(a int, b []*AstExprAnnotationExpr, c AstTypeUnionTypeInterface, d AstBodyVariableDeclaratorIdInterface) (*AstBodyMultiTypeParameter) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MultiTypeParameter", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_d.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyMultiTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MultiTypeParameter(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.UnionType, com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyMultiTypeParameter3(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeUnionTypeInterface, h AstBodyVariableDeclaratorIdInterface) (*AstBodyMultiTypeParameter) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaCallable()
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MultiTypeParameter", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_h.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyMultiTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.UnionType getType()
func (jbobject *AstBodyMultiTypeParameter) GetType() *AstTypeUnionType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/UnionType")
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
	unique_x := &AstTypeUnionType{}
	unique_x.Callable = dst
	return unique_x
}

// public void setType(com.github.javaparser.ast.type.UnionType)
func (jbobject *AstBodyMultiTypeParameter) SetType(a AstTypeUnionTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/UnionType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


