package javaparser

import "github.com/timob/javabind"

type AstStmtDoStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.stmt.Statement getBody()
	GetBody() *AstStmtStatement

	// public com.github.javaparser.ast.expr.Expression getCondition()
	GetCondition() *AstExprExpression

	// public void setBody(com.github.javaparser.ast.stmt.Statement)
	SetBody(a AstStmtStatementInterface) 

	// public void setCondition(com.github.javaparser.ast.expr.Expression)
	SetCondition(a AstExprExpressionInterface) 
}

type AstStmtDoStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.DoStmt()
func NewAstStmtDoStmt() (*AstStmtDoStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/DoStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtDoStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.DoStmt(com.github.javaparser.ast.stmt.Statement, com.github.javaparser.ast.expr.Expression)
func NewAstStmtDoStmt2(a AstStmtStatementInterface, b AstExprExpressionInterface) (*AstStmtDoStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/DoStmt", conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtDoStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.DoStmt(int, int, int, int, com.github.javaparser.ast.stmt.Statement, com.github.javaparser.ast.expr.Expression)
func NewAstStmtDoStmt3(a int, b int, c int, d int, e AstStmtStatementInterface, f AstExprExpressionInterface) (*AstStmtDoStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/DoStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/stmt/Statement"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtDoStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.Statement getBody()
func (jbobject *AstStmtDoStmt) GetBody() *AstStmtStatement {
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

// public com.github.javaparser.ast.expr.Expression getCondition()
func (jbobject *AstStmtDoStmt) GetCondition() *AstExprExpression {
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

// public void setBody(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtDoStmt) SetBody(a AstStmtStatementInterface)  {
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

// public void setCondition(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtDoStmt) SetCondition(a AstExprExpressionInterface)  {
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


