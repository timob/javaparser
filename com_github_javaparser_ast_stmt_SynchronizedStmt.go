package javaparser

import "github.com/timob/javabind"

type AstStmtSynchronizedStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
	GetBlock() *AstStmtBlockStmt

	// public com.github.javaparser.ast.expr.Expression getExpr()
	GetExpr() *AstExprExpression

	// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetBlock(a AstStmtBlockStmtInterface) 

	// public void setExpr(com.github.javaparser.ast.expr.Expression)
	SetExpr(a AstExprExpressionInterface) 
}

type AstStmtSynchronizedStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.SynchronizedStmt()
func NewAstStmtSynchronizedStmt() (*AstStmtSynchronizedStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SynchronizedStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtSynchronizedStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SynchronizedStmt(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtSynchronizedStmt2(a AstExprExpressionInterface, b AstStmtBlockStmtInterface) (*AstStmtSynchronizedStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SynchronizedStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtSynchronizedStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.SynchronizedStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtSynchronizedStmt3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstStmtBlockStmtInterface) (*AstStmtSynchronizedStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/SynchronizedStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtSynchronizedStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
func (jbobject *AstStmtSynchronizedStmt) GetBlock() *AstStmtBlockStmt {
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

// public com.github.javaparser.ast.expr.Expression getExpr()
func (jbobject *AstStmtSynchronizedStmt) GetExpr() *AstExprExpression {
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

// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstStmtSynchronizedStmt) SetBlock(a AstStmtBlockStmtInterface)  {
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

// public void setExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtSynchronizedStmt) SetExpr(a AstExprExpressionInterface)  {
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


