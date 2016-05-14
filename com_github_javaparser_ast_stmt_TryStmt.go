package javaparser

import "github.com/timob/javabind"

type AstStmtTryStmtInterface interface {
	AstStmtStatementInterface

	// public java.util.List<com.github.javaparser.ast.stmt.CatchClause> getCatchs()
	GetCatchs() []*AstStmtCatchClause

	// public com.github.javaparser.ast.stmt.BlockStmt getFinallyBlock()
	GetFinallyBlock() *AstStmtBlockStmt

	// public com.github.javaparser.ast.stmt.BlockStmt getTryBlock()
	GetTryBlock() *AstStmtBlockStmt

	// public java.util.List<com.github.javaparser.ast.expr.VariableDeclarationExpr> getResources()
	GetResources() []*AstExprVariableDeclarationExpr

	// public void setCatchs(java.util.List<com.github.javaparser.ast.stmt.CatchClause>)
	SetCatchs(a []*AstStmtCatchClause) 

	// public void setFinallyBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetFinallyBlock(a AstStmtBlockStmtInterface) 

	// public void setTryBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetTryBlock(a AstStmtBlockStmtInterface) 

	// public void setResources(java.util.List<com.github.javaparser.ast.expr.VariableDeclarationExpr>)
	SetResources(a []*AstExprVariableDeclarationExpr) 
}

type AstStmtTryStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.TryStmt()
func NewAstStmtTryStmt() (*AstStmtTryStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TryStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtTryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.TryStmt(com.github.javaparser.ast.stmt.BlockStmt, java.util.List<com.github.javaparser.ast.stmt.CatchClause>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtTryStmt2(a AstStmtBlockStmtInterface, b []*AstStmtCatchClause, c AstStmtBlockStmtInterface) (*AstStmtTryStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TryStmt", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstStmtTryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.TryStmt(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.VariableDeclarationExpr>, com.github.javaparser.ast.stmt.BlockStmt, java.util.List<com.github.javaparser.ast.stmt.CatchClause>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtTryStmt3(a int, b int, c int, d int, e []*AstExprVariableDeclarationExpr, f AstStmtBlockStmtInterface, g []*AstStmtCatchClause, h AstStmtBlockStmtInterface) (*AstStmtTryStmt) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/TryStmt", a, b, c, d, conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstStmtTryStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.stmt.CatchClause> getCatchs()
func (jbobject *AstStmtTryStmt) GetCatchs() []*AstStmtCatchClause {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCatchs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstStmtCatchClause)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.stmt.BlockStmt getFinallyBlock()
func (jbobject *AstStmtTryStmt) GetFinallyBlock() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFinallyBlock", "com/github/javaparser/ast/stmt/BlockStmt")
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

// public com.github.javaparser.ast.stmt.BlockStmt getTryBlock()
func (jbobject *AstStmtTryStmt) GetTryBlock() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTryBlock", "com/github/javaparser/ast/stmt/BlockStmt")
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

// public java.util.List<com.github.javaparser.ast.expr.VariableDeclarationExpr> getResources()
func (jbobject *AstStmtTryStmt) GetResources() []*AstExprVariableDeclarationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getResources", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprVariableDeclarationExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setCatchs(java.util.List<com.github.javaparser.ast.stmt.CatchClause>)
func (jbobject *AstStmtTryStmt) SetCatchs(a []*AstStmtCatchClause)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCatchs", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setFinallyBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstStmtTryStmt) SetFinallyBlock(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setFinallyBlock", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTryBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstStmtTryStmt) SetTryBlock(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTryBlock", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setResources(java.util.List<com.github.javaparser.ast.expr.VariableDeclarationExpr>)
func (jbobject *AstStmtTryStmt) SetResources(a []*AstExprVariableDeclarationExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setResources", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


