package javaparser

import "github.com/timob/javabind"

type AstExprFieldAccessExprInterface interface {
	AstExprExpressionInterface

	// public java.lang.String getField()
	GetField() string

	// public com.github.javaparser.ast.expr.NameExpr getFieldExpr()
	GetFieldExpr() *AstExprNameExpr

	// public com.github.javaparser.ast.expr.Expression getScope()
	GetScope() *AstExprExpression

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
	GetTypeArgs() []*AstTypeType

	// public void setField(java.lang.String)
	SetField(a string) 

	// public void setFieldExpr(com.github.javaparser.ast.expr.NameExpr)
	SetFieldExpr(a AstExprNameExprInterface) 

	// public void setScope(com.github.javaparser.ast.expr.Expression)
	SetScope(a AstExprExpressionInterface) 

	// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
	SetTypeArgs(a []*AstTypeType) 
}

type AstExprFieldAccessExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.FieldAccessExpr()
func NewAstExprFieldAccessExpr() (*AstExprFieldAccessExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/FieldAccessExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprFieldAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.FieldAccessExpr(com.github.javaparser.ast.expr.Expression, java.lang.String)
func NewAstExprFieldAccessExpr2(a AstExprExpressionInterface, b string) (*AstExprFieldAccessExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/FieldAccessExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprFieldAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.FieldAccessExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.type.Type>, java.lang.String)
func NewAstExprFieldAccessExpr3(a int, b int, c int, d int, e AstExprExpressionInterface, f []*AstTypeType, g string) (*AstExprFieldAccessExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/FieldAccessExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstExprFieldAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getField()
func (jbobject *AstExprFieldAccessExpr) GetField() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getField", "java/lang/String")
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

// public com.github.javaparser.ast.expr.NameExpr getFieldExpr()
func (jbobject *AstExprFieldAccessExpr) GetFieldExpr() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFieldExpr", "com/github/javaparser/ast/expr/NameExpr")
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
func (jbobject *AstExprFieldAccessExpr) GetScope() *AstExprExpression {
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
func (jbobject *AstExprFieldAccessExpr) GetTypeArgs() []*AstTypeType {
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

// public void setField(java.lang.String)
func (jbobject *AstExprFieldAccessExpr) SetField(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setField", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFieldExpr(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstExprFieldAccessExpr) SetFieldExpr(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFieldExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setScope(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprFieldAccessExpr) SetScope(a AstExprExpressionInterface)  {
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
func (jbobject *AstExprFieldAccessExpr) SetTypeArgs(a []*AstTypeType)  {
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


