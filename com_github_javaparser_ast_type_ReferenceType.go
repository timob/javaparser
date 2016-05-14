package javaparser

import "github.com/timob/javabind"

type AstTypeReferenceTypeInterface interface {
	AstTypeTypeInterface

	// public int getArrayCount()
	GetArrayCount() int

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setArrayCount(int)
	SetArrayCount(a int) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>> getArraysAnnotations()
	GetArraysAnnotations() []*[]*AstExprAnnotationExpr

	// public void setArraysAnnotations(java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>>)
	SetArraysAnnotations(a []*[]*AstExprAnnotationExpr) 
}

type AstTypeReferenceType struct {
	AstTypeType
}

// public com.github.javaparser.ast.type.ReferenceType()
func NewAstTypeReferenceType() (*AstTypeReferenceType) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ReferenceType")
	if err != nil {
		panic(err)
	}
	x := &AstTypeReferenceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ReferenceType(com.github.javaparser.ast.type.Type)
func NewAstTypeReferenceType2(a AstTypeTypeInterface) (*AstTypeReferenceType) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ReferenceType", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeReferenceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ReferenceType(com.github.javaparser.ast.type.Type, int)
func NewAstTypeReferenceType3(a AstTypeTypeInterface, b int) (*AstTypeReferenceType) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ReferenceType", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstTypeReferenceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ReferenceType(int, int, int, int, com.github.javaparser.ast.type.Type, int)
func NewAstTypeReferenceType4(a int, b int, c int, d int, e AstTypeTypeInterface, f int) (*AstTypeReferenceType) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ReferenceType", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"), f)
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstTypeReferenceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.type.ReferenceType(int, int, int, int, com.github.javaparser.ast.type.Type, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>>)
func NewAstTypeReferenceType5(a int, b int, c int, d int, e AstTypeTypeInterface, f int, g []*AstExprAnnotationExpr, h []*[]*AstExprAnnotationExpr) (*AstTypeReferenceType) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaList(javabind.NewGoToJavaCallable()))
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/type/ReferenceType", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"), f, conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstTypeReferenceType{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getArrayCount()
func (jbobject *AstTypeReferenceType) GetArrayCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArrayCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstTypeReferenceType) GetType() *AstTypeType {
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

// public void setArrayCount(int)
func (jbobject *AstTypeReferenceType) SetArrayCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArrayCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstTypeReferenceType) SetType(a AstTypeTypeInterface)  {
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

// public java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>> getArraysAnnotations()
func (jbobject *AstTypeReferenceType) GetArraysAnnotations() []*[]*AstExprAnnotationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArraysAnnotations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoList(javabind.NewJavaToGoCallable()))
	dst := new([]*[]*AstExprAnnotationExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setArraysAnnotations(java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>>)
func (jbobject *AstTypeReferenceType) SetArraysAnnotations(a []*[]*AstExprAnnotationExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaList(javabind.NewGoToJavaCallable()))
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArraysAnnotations", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


