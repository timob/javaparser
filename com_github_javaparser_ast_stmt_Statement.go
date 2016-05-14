package javaparser

import "github.com/timob/javabind"

type AstStmtStatementInterface interface {
	AstNodeInterface
}

type AstStmtStatement struct {
	AstNode
}

// public com.github.javaparser.ast.stmt.Statement()
func NewAstStmtStatement() (*AstStmtStatement) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/Statement")
	if err != nil {
		panic(err)
	}
	x := &AstStmtStatement{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.Statement(int, int, int, int)
func NewAstStmtStatement2(a int, b int, c int, d int) (*AstStmtStatement) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/Statement", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstStmtStatement{}
	x.Callable = &javabind.Callable{obj}
	return x
}


