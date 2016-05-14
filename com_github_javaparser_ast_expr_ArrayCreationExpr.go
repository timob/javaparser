package javaparser

import "github.com/timob/javabind"

type AstExprArrayCreationExprInterface interface {
	AstExprExpressionInterface

	// public int getArrayCount()
	GetArrayCount() int

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getDimensions()
	GetDimensions() []*AstExprExpression

	// public com.github.javaparser.ast.expr.ArrayInitializerExpr getInitializer()
	GetInitializer() *AstExprArrayInitializerExpr

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setArrayCount(int)
	SetArrayCount(a int) 

	// public void setDimensions(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetDimensions(a []*AstExprExpression) 

	// public void setInitializer(com.github.javaparser.ast.expr.ArrayInitializerExpr)
	SetInitializer(a AstExprArrayInitializerExprInterface) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>> getArraysAnnotations()
	GetArraysAnnotations() []*[]*AstExprAnnotationExpr

	// public void setArraysAnnotations(java.util.List<java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>>)
	SetArraysAnnotations(a []*[]*AstExprAnnotationExpr) 
}

type AstExprArrayCreationExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ArrayCreationExpr()
func NewAstExprArrayCreationExpr() (*AstExprArrayCreationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayCreationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprArrayCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayCreationExpr(com.github.javaparser.ast.type.Type, int, com.github.javaparser.ast.expr.ArrayInitializerExpr)
func NewAstExprArrayCreationExpr2(a AstTypeTypeInterface, b int, c AstExprArrayInitializerExprInterface) (*AstExprArrayCreationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayCreationExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), b, conv_c.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	x := &AstExprArrayCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayCreationExpr(int, int, int, int, com.github.javaparser.ast.type.Type, int, com.github.javaparser.ast.expr.ArrayInitializerExpr)
func NewAstExprArrayCreationExpr3(a int, b int, c int, d int, e AstTypeTypeInterface, f int, g AstExprArrayInitializerExprInterface) (*AstExprArrayCreationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayCreationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"), f, conv_g.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_g.CleanUp()
	x := &AstExprArrayCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayCreationExpr(com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.expr.Expression>, int)
func NewAstExprArrayCreationExpr4(a AstTypeTypeInterface, b []*AstExprExpression, c int) (*AstExprArrayCreationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayCreationExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("java/util/List"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprArrayCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayCreationExpr(int, int, int, int, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.expr.Expression>, int)
func NewAstExprArrayCreationExpr5(a int, b int, c int, d int, e AstTypeTypeInterface, f []*AstExprExpression, g int) (*AstExprArrayCreationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayCreationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"), conv_f.Value().Cast("java/util/List"), g)
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprArrayCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getArrayCount()
func (jbobject *AstExprArrayCreationExpr) GetArrayCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArrayCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getDimensions()
func (jbobject *AstExprArrayCreationExpr) GetDimensions() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDimensions", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprExpression)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.expr.ArrayInitializerExpr getInitializer()
func (jbobject *AstExprArrayCreationExpr) GetInitializer() *AstExprArrayInitializerExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInitializer", "com/github/javaparser/ast/expr/ArrayInitializerExpr")
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
	unique_x := &AstExprArrayInitializerExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstExprArrayCreationExpr) GetType() *AstTypeType {
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
func (jbobject *AstExprArrayCreationExpr) SetArrayCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArrayCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setDimensions(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstExprArrayCreationExpr) SetDimensions(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDimensions", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInitializer(com.github.javaparser.ast.expr.ArrayInitializerExpr)
func (jbobject *AstExprArrayCreationExpr) SetInitializer(a AstExprArrayInitializerExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInitializer", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstExprArrayCreationExpr) SetType(a AstTypeTypeInterface)  {
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
func (jbobject *AstExprArrayCreationExpr) GetArraysAnnotations() []*[]*AstExprAnnotationExpr {
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
func (jbobject *AstExprArrayCreationExpr) SetArraysAnnotations(a []*[]*AstExprAnnotationExpr)  {
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


