package javaparser

import "github.com/timob/javabind"

type AstStmtBlockStmtInterface interface {
	AstStmtStatementInterface

	// public java.util.List<com.github.javaparser.ast.stmt.Statement> getStmts()
	GetStmts() []*AstStmtStatement

	// public void setStmts(java.util.List<com.github.javaparser.ast.stmt.Statement>)
	SetStmts(a []*AstStmtStatement) 
}

type AstStmtBlockStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.BlockStmt()
func NewAstStmtBlockStmt() (*AstStmtBlockStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BlockStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtBlockStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt(java.util.List<com.github.javaparser.ast.stmt.Statement>)
func NewAstStmtBlockStmt2(a []*AstStmtStatement) (*AstStmtBlockStmt) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BlockStmt", conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtBlockStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt(int, int, int, int, java.util.List<com.github.javaparser.ast.stmt.Statement>)
func NewAstStmtBlockStmt3(a int, b int, c int, d int, e []*AstStmtStatement) (*AstStmtBlockStmt) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BlockStmt", a, b, c, d, conv_e.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtBlockStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.stmt.Statement> getStmts()
func (jbobject *AstStmtBlockStmt) GetStmts() []*AstStmtStatement {
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

// public void setStmts(java.util.List<com.github.javaparser.ast.stmt.Statement>)
func (jbobject *AstStmtBlockStmt) SetStmts(a []*AstStmtStatement)  {
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


