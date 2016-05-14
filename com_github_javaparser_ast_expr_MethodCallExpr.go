package javaparser

import "github.com/timob/javabind"

type AstExprMethodCallExprInterface interface {
	AstExprExpressionInterface

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
	GetArgs() []*AstExprExpression

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.expr.NameExpr getNameExpr()
	GetNameExpr() *AstExprNameExpr

	// public com.github.javaparser.ast.expr.Expression getScope()
	GetScope() *AstExprExpression

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
	GetTypeArgs() []*AstTypeType

	// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetArgs(a []*AstExprExpression) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
	SetNameExpr(a AstExprNameExprInterface) 

	// public void setScope(com.github.javaparser.ast.expr.Expression)
	SetScope(a AstExprExpressionInterface) 

	// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
	SetTypeArgs(a []*AstTypeType) 
}

type AstExprMethodCallExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.MethodCallExpr()
func NewAstExprMethodCallExpr() (*AstExprMethodCallExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodCallExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprMethodCallExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MethodCallExpr(com.github.javaparser.ast.expr.Expression, java.lang.String)
func NewAstExprMethodCallExpr2(a AstExprExpressionInterface, b string) (*AstExprMethodCallExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodCallExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprMethodCallExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MethodCallExpr(com.github.javaparser.ast.expr.Expression, java.lang.String, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstExprMethodCallExpr3(a AstExprExpressionInterface, b string, c []*AstExprExpression) (*AstExprMethodCallExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodCallExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstExprMethodCallExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MethodCallExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.type.Type>, java.lang.String, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstExprMethodCallExpr4(a int, b int, c int, d int, e AstExprExpressionInterface, f []*AstTypeType, g string, h []*AstExprExpression) (*AstExprMethodCallExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MethodCallExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/lang/String"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstExprMethodCallExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
func (jbobject *AstExprMethodCallExpr) GetArgs() []*AstExprExpression {
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

// public java.lang.String getName()
func (jbobject *AstExprMethodCallExpr) GetName() string {
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

// public com.github.javaparser.ast.expr.NameExpr getNameExpr()
func (jbobject *AstExprMethodCallExpr) GetNameExpr() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNameExpr", "com/github/javaparser/ast/expr/NameExpr")
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
	unique_x := &AstExprNameExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.expr.Expression getScope()
func (jbobject *AstExprMethodCallExpr) GetScope() *AstExprExpression {
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

// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
func (jbobject *AstExprMethodCallExpr) GetTypeArgs() []*AstTypeType {
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

// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstExprMethodCallExpr) SetArgs(a []*AstExprExpression)  {
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

// public void setName(java.lang.String)
func (jbobject *AstExprMethodCallExpr) SetName(a string)  {
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

// public void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstExprMethodCallExpr) SetNameExpr(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNameExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setScope(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprMethodCallExpr) SetScope(a AstExprExpressionInterface)  {
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

// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
func (jbobject *AstExprMethodCallExpr) SetTypeArgs(a []*AstTypeType)  {
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


