package javaparser

import "github.com/timob/javabind"

type AstStmtIfStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.expr.Expression getCondition()
	GetCondition() *AstExprExpression

	// public com.github.javaparser.ast.stmt.Statement getElseStmt()
	GetElseStmt() *AstStmtStatement

	// public com.github.javaparser.ast.stmt.Statement getThenStmt()
	GetThenStmt() *AstStmtStatement

	// public void setCondition(com.github.javaparser.ast.expr.Expression)
	SetCondition(a AstExprExpressionInterface) 

	// public void setElseStmt(com.github.javaparser.ast.stmt.Statement)
	SetElseStmt(a AstStmtStatementInterface) 

	// public void setThenStmt(com.github.javaparser.ast.stmt.Statement)
	SetThenStmt(a AstStmtStatementInterface) 
}

type AstStmtIfStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.IfStmt()
func NewAstStmtIfStmt() (*AstStmtIfStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/IfStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtIfStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.IfStmt(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtIfStmt2(a AstExprExpressionInterface, b AstStmtStatementInterface, c AstStmtStatementInterface) (*AstStmtIfStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/IfStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/Statement"), conv_c.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstStmtIfStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.IfStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtIfStmt3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstStmtStatementInterface, g AstStmtStatementInterface) (*AstStmtIfStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/IfStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/Statement"), conv_g.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstStmtIfStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getCondition()
func (jbobject *AstStmtIfStmt) GetCondition() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCondition", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.stmt.Statement getElseStmt()
func (jbobject *AstStmtIfStmt) GetElseStmt() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getElseStmt", "com/github/javaparser/ast/stmt/Statement")
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

// public com.github.javaparser.ast.stmt.Statement getThenStmt()
func (jbobject *AstStmtIfStmt) GetThenStmt() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThenStmt", "com/github/javaparser/ast/stmt/Statement")
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

// public void setCondition(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtIfStmt) SetCondition(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCondition", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setElseStmt(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtIfStmt) SetElseStmt(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setElseStmt", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setThenStmt(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtIfStmt) SetThenStmt(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThenStmt", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


