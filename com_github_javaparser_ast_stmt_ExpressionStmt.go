package javaparser

import "github.com/timob/javabind"

type AstStmtExpressionStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.expr.Expression getExpression()
	GetExpression() *AstExprExpression

	// public void setExpression(com.github.javaparser.ast.expr.Expression)
	SetExpression(a AstExprExpressionInterface) 
}

type AstStmtExpressionStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ExpressionStmt()
func NewAstStmtExpressionStmt() (*AstStmtExpressionStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExpressionStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtExpressionStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ExpressionStmt(com.github.javaparser.ast.expr.Expression)
func NewAstStmtExpressionStmt2(a AstExprExpressionInterface) (*AstStmtExpressionStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExpressionStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtExpressionStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ExpressionStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression)
func NewAstStmtExpressionStmt3(a int, b int, c int, d int, e AstExprExpressionInterface) (*AstStmtExpressionStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ExpressionStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtExpressionStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getExpression()
func (jbobject *AstStmtExpressionStmt) GetExpression() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExpression", "com/github/javaparser/ast/expr/Expression")
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

// public void setExpression(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtExpressionStmt) SetExpression(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExpression", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


