package javaparser

import "github.com/timob/javabind"

type AstStmtContinueStmtInterface interface {
	AstStmtStatementInterface

	// public java.lang.String getId()
	GetId() string

	// public void setId(java.lang.String)
	SetId(a string) 
}

type AstStmtContinueStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ContinueStmt()
func NewAstStmtContinueStmt() (*AstStmtContinueStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ContinueStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtContinueStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ContinueStmt(java.lang.String)
func NewAstStmtContinueStmt2(a string) (*AstStmtContinueStmt) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ContinueStmt", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtContinueStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ContinueStmt(int, int, int, int, java.lang.String)
func NewAstStmtContinueStmt3(a int, b int, c int, d int, e string) (*AstStmtContinueStmt) {
	conv_e := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ContinueStmt", a, b, c, d, conv_e.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtContinueStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getId()
func (jbobject *AstStmtContinueStmt) GetId() string {
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
func (jbobject *AstStmtContinueStmt) SetId(a string)  {
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


