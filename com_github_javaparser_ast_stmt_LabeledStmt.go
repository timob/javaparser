package javaparser

import "github.com/timob/javabind"

type AstStmtLabeledStmtInterface interface {
	AstStmtStatementInterface

	// public java.lang.String getLabel()
	GetLabel() string

	// public com.github.javaparser.ast.stmt.Statement getStmt()
	GetStmt() *AstStmtStatement

	// public void setLabel(java.lang.String)
	SetLabel(a string) 

	// public void setStmt(com.github.javaparser.ast.stmt.Statement)
	SetStmt(a AstStmtStatementInterface) 
}

type AstStmtLabeledStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.LabeledStmt()
func NewAstStmtLabeledStmt() (*AstStmtLabeledStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/LabeledStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtLabeledStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.LabeledStmt(java.lang.String, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtLabeledStmt2(a string, b AstStmtStatementInterface) (*AstStmtLabeledStmt) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/LabeledStmt", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtLabeledStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.LabeledStmt(int, int, int, int, java.lang.String, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtLabeledStmt3(a int, b int, c int, d int, e string, f AstStmtStatementInterface) (*AstStmtLabeledStmt) {
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/LabeledStmt", a, b, c, d, conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtLabeledStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getLabel()
func (jbobject *AstStmtLabeledStmt) GetLabel() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLabel", "java/lang/String")
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

// public com.github.javaparser.ast.stmt.Statement getStmt()
func (jbobject *AstStmtLabeledStmt) GetStmt() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getStmt", "com/github/javaparser/ast/stmt/Statement")
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

// public void setLabel(java.lang.String)
func (jbobject *AstStmtLabeledStmt) SetLabel(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setLabel", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStmt(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtLabeledStmt) SetStmt(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStmt", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


