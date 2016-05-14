package javaparser

import "github.com/timob/javabind"

type AstBodyInitializerDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
	GetBlock() *AstStmtBlockStmt

	// public boolean isStatic()
	IsStatic() bool

	// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetBlock(a AstStmtBlockStmtInterface) 

	// public void setStatic(boolean)
	SetStatic(a bool) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyInitializerDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.InitializerDeclaration()
func NewAstBodyInitializerDeclaration() (*AstBodyInitializerDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/InitializerDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyInitializerDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.InitializerDeclaration(boolean, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyInitializerDeclaration2(a bool, b AstStmtBlockStmtInterface) (*AstBodyInitializerDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/InitializerDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyInitializerDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.InitializerDeclaration(int, int, int, int, boolean, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyInitializerDeclaration3(a int, b int, c int, d int, e bool, f AstStmtBlockStmtInterface) (*AstBodyInitializerDeclaration) {
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/InitializerDeclaration", a, b, c, d, e, conv_f.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	x := &AstBodyInitializerDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
func (jbobject *AstBodyInitializerDeclaration) GetBlock() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlock", "com/github/javaparser/ast/stmt/BlockStmt")
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
	unique_x := &AstStmtBlockStmt{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isStatic()
func (jbobject *AstBodyInitializerDeclaration) IsStatic() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStatic", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstBodyInitializerDeclaration) SetBlock(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlock", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStatic(boolean)
func (jbobject *AstBodyInitializerDeclaration) SetStatic(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatic", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyInitializerDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJavaDoc", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
func (jbobject *AstBodyInitializerDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJavaDoc", "com/github/javaparser/ast/comments/JavadocComment")
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
	unique_x := &AstCommentsJavadocComment{}
	unique_x.Callable = dst
	return unique_x
}


