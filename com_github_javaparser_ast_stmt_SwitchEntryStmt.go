package javaparser

import "github.com/timob/javabind"

type AstStmtSwitchEntryStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.expr.Expression getLabel()
	GetLabel() *AstExprExpression

	// public java.util.List<com.github.javaparser.ast.stmt.Statement> getStmts()
	GetStmts() []*AstStmtStatement

	// public void setLabel(com.github.javaparser.ast.expr.Expression)
	SetLabel(a AstExprExpressionInterface) 

	// public void setStmts(java.util.List<com.github.javaparser.ast.stmt.Statement>)
	SetStmts(a []*AstStmtStatement) 
}

type AstStmtSwitchEntryStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.SwitchEntryStmt()
func NewAstStmtSwitchEntryStmt() (*AstStmtSwitchEntryStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchEntryStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtSwitchEntryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SwitchEntryStmt(com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.stmt.Statement>)
func NewAstStmtSwitchEntryStmt2(a AstExprExpressionInterface, b []*AstStmtStatement) (*AstStmtSwitchEntryStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchEntryStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtSwitchEntryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SwitchEntryStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.stmt.Statement>)
func NewAstStmtSwitchEntryStmt3(a int, b int, c int, d int, e AstExprExpressionInterface, f []*AstStmtStatement) (*AstStmtSwitchEntryStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchEntryStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtSwitchEntryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getLabel()
func (jbobject *AstStmtSwitchEntryStmt) GetLabel() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLabel", "com/github/javaparser/ast/expr/Expression")
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

// public java.util.List<com.github.javaparser.ast.stmt.Statement> getStmts()
func (jbobject *AstStmtSwitchEntryStmt) GetStmts() []*AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStmts", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstStmtStatement)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setLabel(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtSwitchEntryStmt) SetLabel(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLabel", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStmts(java.util.List<com.github.javaparser.ast.stmt.Statement>)
func (jbobject *AstStmtSwitchEntryStmt) SetStmts(a []*AstStmtStatement)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStmts", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


