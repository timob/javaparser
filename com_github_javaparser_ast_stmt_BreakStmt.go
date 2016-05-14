package javaparser

import "github.com/timob/javabind"

type AstStmtBreakStmtInterface interface {
	AstStmtStatementInterface

	// public java.lang.String getId()
	GetId() string

	// public void setId(java.lang.String)
	SetId(a string) 
}

type AstStmtBreakStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.BreakStmt()
func NewAstStmtBreakStmt() (*AstStmtBreakStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BreakStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtBreakStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BreakStmt(java.lang.String)
func NewAstStmtBreakStmt2(a string) (*AstStmtBreakStmt) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BreakStmt", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtBreakStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BreakStmt(int, int, int, int, java.lang.String)
func NewAstStmtBreakStmt3(a int, b int, c int, d int, e string) (*AstStmtBreakStmt) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/BreakStmt", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtBreakStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getId()
func (jbobject *AstStmtBreakStmt) GetId() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getId", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setId(java.lang.String)
func (jbobject *AstStmtBreakStmt) SetId(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setId", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


