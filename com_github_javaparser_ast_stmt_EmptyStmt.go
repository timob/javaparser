package javaparser

import "github.com/timob/javabind"

type AstStmtEmptyStmtInterface interface {
	AstStmtStatementInterface
}

type AstStmtEmptyStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.EmptyStmt()
func NewAstStmtEmptyStmt() (*AstStmtEmptyStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/EmptyStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtEmptyStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.EmptyStmt(int, int, int, int)
func NewAstStmtEmptyStmt2(a int, b int, c int, d int) (*AstStmtEmptyStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/EmptyStmt", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstStmtEmptyStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}


