package javaparser

import "github.com/timob/javabind"

type AstStmtForStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.stmt.Statement getBody()
	GetBody() *AstStmtStatement

	// public com.github.javaparser.ast.expr.Expression getCompare()
	GetCompare() *AstExprExpression

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getInit()
	GetInit() []*AstExprExpression

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getUpdate()
	GetUpdate() []*AstExprExpression

	// public void setBody(com.github.javaparser.ast.stmt.Statement)
	SetBody(a AstStmtStatementInterface) 

	// public void setCompare(com.github.javaparser.ast.expr.Expression)
	SetCompare(a AstExprExpressionInterface) 

	// public void setInit(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetInit(a []*AstExprExpression) 

	// public void setUpdate(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetUpdate(a []*AstExprExpression) 
}

type AstStmtForStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ForStmt()
func NewAstStmtForStmt() (*AstStmtForStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtForStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ForStmt(java.util.List<com.github.javaparser.ast.expr.Expression>, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.expr.Expression>, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtForStmt2(a []*AstExprExpression, b AstExprExpressionInterface, c []*AstExprExpression, d AstStmtStatementInterface) (*AstStmtForStmt) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForStmt", conv_a.Value().Cast("java/util/List"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_c.Value().Cast("java/util/List"), conv_d.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstStmtForStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ForStmt(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.Expression>, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.expr.Expression>, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtForStmt3(a int, b int, c int, d int, e []*AstExprExpression, f AstExprExpressionInterface, g []*AstExprExpression, h AstStmtStatementInterface) (*AstStmtForStmt) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForStmt", a, b, c, d, conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstStmtForStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.Statement getBody()
func (jbobject *AstStmtForStmt) GetBody() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBody", "com/github/javaparser/ast/stmt/Statement")
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
	unique_x := &AstStmtStatement{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.expr.Expression getCompare()
func (jbobject *AstStmtForStmt) GetCompare() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCompare", "com/github/javaparser/ast/expr/Expression")
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

// public java.util.List<com.github.javaparser.ast.expr.Expression> getInit()
func (jbobject *AstStmtForStmt) GetInit() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInit", "java/util/List")
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

// public java.util.List<com.github.javaparser.ast.expr.Expression> getUpdate()
func (jbobject *AstStmtForStmt) GetUpdate() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUpdate", "java/util/List")
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

// public void setBody(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtForStmt) SetBody(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBody", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setCompare(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtForStmt) SetCompare(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCompare", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInit(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstStmtForStmt) SetInit(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInit", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setUpdate(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstStmtForStmt) SetUpdate(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setUpdate", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


