package javaparser

import "github.com/timob/javabind"

type AstTypeParameterInterface interface {
	AstNodeInterface

	// public java.lang.String getName()
	GetName() string

	// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getTypeBound()
	GetTypeBound() []*AstTypeClassOrInterfaceType

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setTypeBound(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
	SetTypeBound(a []*AstTypeClassOrInterfaceType) 

	// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 
}

type AstTypeParameter struct {
	AstNode
}

// public com.github.javaparser.ast.TypeParameter()
func NewAstTypeParameter() (*AstTypeParameter) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/TypeParameter")
	if err != nil {
		panic(err)
	}
	x := &AstTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.TypeParameter(java.lang.String, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func NewAstTypeParameter2(a string, b []*AstTypeClassOrInterfaceType) (*AstTypeParameter) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/TypeParameter", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.TypeParameter(int, int, int, int, java.lang.String, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func NewAstTypeParameter3(a int, b int, c int, d int, e string, f []*AstTypeClassOrInterfaceType) (*AstTypeParameter) {
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/TypeParameter", a, b, c, d, conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.TypeParameter(int, int, int, int, java.lang.String, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func NewAstTypeParameter4(a int, b int, c int, d int, e string, f []*AstTypeClassOrInterfaceType, g []*AstExprAnnotationExpr) (*AstTypeParameter) {
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/TypeParameter", a, b, c, d, conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstTypeParameter{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getName()
func (jbobject *AstTypeParameter) GetName() string {
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

// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getTypeBound()
func (jbobject *AstTypeParameter) GetTypeBound() []*AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeBound", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeClassOrInterfaceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setName(java.lang.String)
func (jbobject *AstTypeParameter) SetName(a string)  {
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

// public void setTypeBound(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func (jbobject *AstTypeParameter) SetTypeBound(a []*AstTypeClassOrInterfaceType)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeBound", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstTypeParameter) GetAnnotations() []*AstExprAnnotationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAnnotations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprAnnotationExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstTypeParameter) SetAnnotations(a []*AstExprAnnotationExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAnnotations", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


