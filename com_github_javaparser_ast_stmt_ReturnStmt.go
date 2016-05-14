package javaparser

import "github.com/timob/javabind"

type AstStmtReturnStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.expr.Expression getExpr()
	GetExpr() *AstExprExpression

	// public void setExpr(com.github.javaparser.ast.expr.Expression)
	SetExpr(a AstExprExpressionInterface) 
}

type AstStmtReturnStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ReturnStmt()
func NewAstStmtReturnStmt() (*AstStmtReturnStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ReturnStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtReturnStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ReturnStmt(com.github.javaparser.ast.expr.Expression)
func NewAstStmtReturnStmt2(a AstExprExpressionInterface) (*AstStmtReturnStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ReturnStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtReturnStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ReturnStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression)
func NewAstStmtReturnStmt3(a int, b int, c int, d int, e AstExprExpressionInterface) (*AstStmtReturnStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ReturnStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtReturnStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getExpr()
func (jbobject *AstStmtReturnStmt) GetExpr() *AstExprExpression {
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

// public void setExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtReturnStmt) SetExpr(a AstExprExpressionInterface)  {
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


