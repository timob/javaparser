package javaparser

import "github.com/timob/javabind"

type AstExprObjectCreationExprInterface interface {
	AstExprExpressionInterface

	// public java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getAnonymousClassBody()
	GetAnonymousClassBody() []*AstBodyBodyDeclaration

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
	GetArgs() []*AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getScope()
	GetScope() *AstExprExpression

	// public com.github.javaparser.ast.type.ClassOrInterfaceType getType()
	GetType() *AstTypeClassOrInterfaceType

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
	GetTypeArgs() []*AstTypeType

	// public void setAnonymousClassBody(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
	SetAnonymousClassBody(a []*AstBodyBodyDeclaration) 

	// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetArgs(a []*AstExprExpression) 

	// public void setScope(com.github.javaparser.ast.expr.Expression)
	SetScope(a AstExprExpressionInterface) 

	// public void setType(com.github.javaparser.ast.type.ClassOrInterfaceType)
	SetType(a AstTypeClassOrInterfaceTypeInterface) 

	// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
	SetTypeArgs(a []*AstTypeType) 
}

type AstExprObjectCreationExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ObjectCreationExpr()
func NewAstExprObjectCreationExpr() (*AstExprObjectCreationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ObjectCreationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprObjectCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ObjectCreationExpr(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.type.ClassOrInterfaceType, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstExprObjectCreationExpr2(a AstExprExpressionInterface, b AstTypeClassOrInterfaceTypeInterface, c []*AstExprExpression) (*AstExprObjectCreationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ObjectCreationExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstExprObjectCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ObjectCreationExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.type.ClassOrInterfaceType, java.util.List<com.github.javaparser.ast.type.Type>, java.util.List<com.github.javaparser.ast.expr.Expression>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstExprObjectCreationExpr3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstTypeClassOrInterfaceTypeInterface, g []*AstTypeType, h []*AstExprExpression, i []*AstBodyBodyDeclaration) (*AstExprObjectCreationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_i := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ObjectCreationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("java/util/List"), conv_i.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	x := &AstExprObjectCreationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getAnonymousClassBody()
func (jbobject *AstExprObjectCreationExpr) GetAnonymousClassBody() []*AstBodyBodyDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAnonymousClassBody", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyBodyDeclaration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
func (jbobject *AstExprObjectCreationExpr) GetArgs() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArgs", "java/util/List")
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

// public com.github.javaparser.ast.expr.Expression getScope()
func (jbobject *AstExprObjectCreationExpr) GetScope() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getScope", "com/github/javaparser/ast/expr/Expression")
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
	unique_x := &AstExprExpression{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.type.ClassOrInterfaceType getType()
func (jbobject *AstExprObjectCreationExpr) GetType() *AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/ClassOrInterfaceType")
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
func (jbobject *AstExprObjectCreationExpr) GetTypeArgs() []*AstTypeType {
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

// public void setAnonymousClassBody(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func (jbobject *AstExprObjectCreationExpr) SetAnonymousClassBody(a []*AstBodyBodyDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAnonymousClassBody", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstExprObjectCreationExpr) SetArgs(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArgs", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setScope(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprObjectCreationExpr) SetScope(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setScope", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setType(com.github.javaparser.ast.type.ClassOrInterfaceType)
func (jbobject *AstExprObjectCreationExpr) SetType(a AstTypeClassOrInterfaceTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
func (jbobject *AstExprObjectCreationExpr) SetTypeArgs(a []*AstTypeType)  {
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


