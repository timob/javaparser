package javaparser

import "github.com/timob/javabind"

type AstStmtSwitchStmtInterface interface {
	AstStmtStatementInterface

	// public java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt> getEntries()
	GetEntries() []*AstStmtSwitchEntryStmt

	// public com.github.javaparser.ast.expr.Expression getSelector()
	GetSelector() *AstExprExpression

	// public void setEntries(java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt>)
	SetEntries(a []*AstStmtSwitchEntryStmt) 

	// public void setSelector(com.github.javaparser.ast.expr.Expression)
	SetSelector(a AstExprExpressionInterface) 
}

type AstStmtSwitchStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.SwitchStmt()
func NewAstStmtSwitchStmt() (*AstStmtSwitchStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtSwitchStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SwitchStmt(com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt>)
func NewAstStmtSwitchStmt2(a AstExprExpressionInterface, b []*AstStmtSwitchEntryStmt) (*AstStmtSwitchStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtSwitchStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SwitchStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt>)
func NewAstStmtSwitchStmt3(a int, b int, c int, d int, e AstExprExpressionInterface, f []*AstStmtSwitchEntryStmt) (*AstStmtSwitchStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SwitchStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtSwitchStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt> getEntries()
func (jbobject *AstStmtSwitchStmt) GetEntries() []*AstStmtSwitchEntryStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEntries", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstStmtSwitchEntryStmt)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.expr.Expression getSelector()
func (jbobject *AstStmtSwitchStmt) GetSelector() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSelector", "com/github/javaparser/ast/expr/Expression")
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

// public void setEntries(java.util.List<com.github.javaparser.ast.stmt.SwitchEntryStmt>)
func (jbobject *AstStmtSwitchStmt) SetEntries(a []*AstStmtSwitchEntryStmt)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEntries", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setSelector(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtSwitchStmt) SetSelector(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setSelector", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


