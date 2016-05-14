package javaparser

import "github.com/timob/javabind"

type AstExprVariableDeclarationExprInterface interface {
	AstExprExpressionInterface

	// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public int getModifiers()
	GetModifiers() int

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public java.util.List<com.github.javaparser.ast.body.VariableDeclarator> getVars()
	GetVars() []*AstBodyVariableDeclarator

	// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 

	// public void setModifiers(int)
	SetModifiers(a int) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public void setVars(java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
	SetVars(a []*AstBodyVariableDeclarator) 
}

type AstExprVariableDeclarationExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.VariableDeclarationExpr()
func NewAstExprVariableDeclarationExpr() (*AstExprVariableDeclarationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/VariableDeclarationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprVariableDeclarationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.VariableDeclarationExpr(com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstExprVariableDeclarationExpr2(a AstTypeTypeInterface, b []*AstBodyVariableDeclarator) (*AstExprVariableDeclarationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/VariableDeclarationExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprVariableDeclarationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.VariableDeclarationExpr(int, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstExprVariableDeclarationExpr3(a int, b AstTypeTypeInterface, c []*AstBodyVariableDeclarator) (*AstExprVariableDeclarationExpr) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/VariableDeclarationExpr", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstExprVariableDeclarationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.VariableDeclarationExpr(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstExprVariableDeclarationExpr4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeTypeInterface, h []*AstBodyVariableDeclarator) (*AstExprVariableDeclarationExpr) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/VariableDeclarationExpr", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/Type"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstExprVariableDeclarationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstExprVariableDeclarationExpr) GetAnnotations() []*AstExprAnnotationExpr {
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

// public int getModifiers()
func (jbobject *AstExprVariableDeclarationExpr) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstExprVariableDeclarationExpr) GetType() *AstTypeType {
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

// public java.util.List<com.github.javaparser.ast.body.VariableDeclarator> getVars()
func (jbobject *AstExprVariableDeclarationExpr) GetVars() []*AstBodyVariableDeclarator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVars", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyVariableDeclarator)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstExprVariableDeclarationExpr) SetAnnotations(a []*AstExprAnnotationExpr)  {
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

// public void setModifiers(int)
func (jbobject *AstExprVariableDeclarationExpr) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstExprVariableDeclarationExpr) SetType(a AstTypeTypeInterface)  {
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

// public void setVars(java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func (jbobject *AstExprVariableDeclarationExpr) SetVars(a []*AstBodyVariableDeclarator)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVars", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


