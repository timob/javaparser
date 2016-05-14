package javaparser

import "github.com/timob/javabind"

type AstStmtTypeDeclarationStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.body.TypeDeclaration getTypeDeclaration()
	GetTypeDeclaration() *AstBodyTypeDeclaration

	// public void setTypeDeclaration(com.github.javaparser.ast.body.TypeDeclaration)
	SetTypeDeclaration(a AstBodyTypeDeclarationInterface) 
}

type AstStmtTypeDeclarationStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.TypeDeclarationStmt()
func NewAstStmtTypeDeclarationStmt() (*AstStmtTypeDeclarationStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TypeDeclarationStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtTypeDeclarationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.TypeDeclarationStmt(com.github.javaparser.ast.body.TypeDeclaration)
func NewAstStmtTypeDeclarationStmt2(a AstBodyTypeDeclarationInterface) (*AstStmtTypeDeclarationStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TypeDeclarationStmt", conv_a.Value().Cast("com/github/javaparser/ast/body/TypeDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtTypeDeclarationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.TypeDeclarationStmt(int, int, int, int, com.github.javaparser.ast.body.TypeDeclaration)
func NewAstStmtTypeDeclarationStmt3(a int, b int, c int, d int, e AstBodyTypeDeclarationInterface) (*AstStmtTypeDeclarationStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TypeDeclarationStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/body/TypeDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstStmtTypeDeclarationStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.TypeDeclaration getTypeDeclaration()
func (jbobject *AstStmtTypeDeclarationStmt) GetTypeDeclaration() *AstBodyTypeDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeDeclaration", "com/github/javaparser/ast/body/TypeDeclaration")
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
	unique_x := &AstBodyTypeDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public void setTypeDeclaration(com.github.javaparser.ast.body.TypeDeclaration)
func (jbobject *AstStmtTypeDeclarationStmt) SetTypeDeclaration(a AstBodyTypeDeclarationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeDeclaration", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/TypeDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


