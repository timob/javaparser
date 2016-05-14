package javaparser

import "github.com/timob/javabind"

type AstStmtExplicitConstructorInvocationStmtInterface interface {
	AstStmtStatementInterface

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
	GetArgs() []*AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getExpr()
	GetExpr() *AstExprExpression

	// public java.util.List<com.github.javaparser.ast.type.Type> getTypeArgs()
	GetTypeArgs() []*AstTypeType

	// public boolean isThis()
	IsThis() bool

	// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetArgs(a []*AstExprExpression) 

	// public void setExpr(com.github.javaparser.ast.expr.Expression)
	SetExpr(a AstExprExpressionInterface) 

	// public void setThis(boolean)
	SetThis(a bool) 

	// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
	SetTypeArgs(a []*AstTypeType) 
}

type AstStmtExplicitConstructorInvocationStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt()
func NewAstStmtExplicitConstructorInvocationStmt() (*AstStmtExplicitConstructorInvocationStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtExplicitConstructorInvocationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt(boolean, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstStmtExplicitConstructorInvocationStmt2(a bool, b AstExprExpressionInterface, c []*AstExprExpression) (*AstStmtExplicitConstructorInvocationStmt) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt", a, conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstStmtExplicitConstructorInvocationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt(int, int, int, int, java.util.List<com.github.javaparser.ast.type.Type>, boolean, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.expr.Expression>)
func NewAstStmtExplicitConstructorInvocationStmt3(a int, b int, c int, d int, e []*AstTypeType, f bool, g AstExprExpressionInterface, h []*AstExprExpression) (*AstStmtExplicitConstructorInvocationStmt) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt", a, b, c, d, conv_e.Value().Cast("java/util/List"), f, conv_g.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstStmtExplicitConstructorInvocationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
func (jbobject *AstStmtExplicitConstructorInvocationStmt) GetArgs() []*AstExprExpression {
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

// public com.github.javaparser.ast.expr.Expression getExpr()
func (jbobject *AstStmtExplicitConstructorInvocationStmt) GetExpr() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpr", "com/github/javaparser/ast/expr/Expression")
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
func (jbobject *AstStmtExplicitConstructorInvocationStmt) GetTypeArgs() []*AstTypeType {
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

// public boolean isThis()
func (jbobject *AstStmtExplicitConstructorInvocationStmt) IsThis() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isThis", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstStmtExplicitConstructorInvocationStmt) SetArgs(a []*AstExprExpression)  {
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

// public void setExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtExplicitConstructorInvocationStmt) SetExpr(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setThis(boolean)
func (jbobject *AstStmtExplicitConstructorInvocationStmt) SetThis(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThis", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTypeArgs(java.util.List<com.github.javaparser.ast.type.Type>)
func (jbobject *AstStmtExplicitConstructorInvocationStmt) SetTypeArgs(a []*AstTypeType)  {
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


