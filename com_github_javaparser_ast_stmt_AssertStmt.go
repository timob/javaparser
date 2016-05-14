package javaparser

import "github.com/timob/javabind"

type AstStmtAssertStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.expr.Expression getCheck()
	GetCheck() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getMessage()
	GetMessage() *AstExprExpression

	// public void setCheck(com.github.javaparser.ast.expr.Expression)
	SetCheck(a AstExprExpressionInterface) 

	// public void setMessage(com.github.javaparser.ast.expr.Expression)
	SetMessage(a AstExprExpressionInterface) 
}

type AstStmtAssertStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.AssertStmt()
func NewAstStmtAssertStmt() (*AstStmtAssertStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/AssertStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtAssertStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.AssertStmt(com.github.javaparser.ast.expr.Expression)
func NewAstStmtAssertStmt2(a AstExprExpressionInterface) (*AstStmtAssertStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/AssertStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstStmtAssertStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.AssertStmt(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstStmtAssertStmt3(a AstExprExpressionInterface, b AstExprExpressionInterface) (*AstStmtAssertStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/AssertStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtAssertStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.AssertStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstStmtAssertStmt4(a int, b int, c int, d int, e AstExprExpressionInterface, f AstExprExpressionInterface) (*AstStmtAssertStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/AssertStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtAssertStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getCheck()
func (jbobject *AstStmtAssertStmt) GetCheck() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCheck", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.Expression getMessage()
func (jbobject *AstStmtAssertStmt) GetMessage() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMessage", "com/github/javaparser/ast/expr/Expression")
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

// public void setCheck(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtAssertStmt) SetCheck(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCheck", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setMessage(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtAssertStmt) SetMessage(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMessage", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


