package javaparser

import "github.com/timob/javabind"

type AstStmtForeachStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.stmt.Statement getBody()
	GetBody() *AstStmtStatement

	// public com.github.javaparser.ast.expr.Expression getIterable()
	GetIterable() *AstExprExpression

	// public com.github.javaparser.ast.expr.VariableDeclarationExpr getVariable()
	GetVariable() *AstExprVariableDeclarationExpr

	// public void setBody(com.github.javaparser.ast.stmt.Statement)
	SetBody(a AstStmtStatementInterface) 

	// public void setIterable(com.github.javaparser.ast.expr.Expression)
	SetIterable(a AstExprExpressionInterface) 

	// public void setVariable(com.github.javaparser.ast.expr.VariableDeclarationExpr)
	SetVariable(a AstExprVariableDeclarationExprInterface) 
}

type AstStmtForeachStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.ForeachStmt()
func NewAstStmtForeachStmt() (*AstStmtForeachStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForeachStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtForeachStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ForeachStmt(com.github.javaparser.ast.expr.VariableDeclarationExpr, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtForeachStmt2(a AstExprVariableDeclarationExprInterface, b AstExprExpressionInterface, c AstStmtStatementInterface) (*AstStmtForeachStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForeachStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_c.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstStmtForeachStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.ForeachStmt(int, int, int, int, com.github.javaparser.ast.expr.VariableDeclarationExpr, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtForeachStmt3(a int, b int, c int, d int, e AstExprVariableDeclarationExprInterface, f AstExprExpressionInterface, g AstStmtStatementInterface) (*AstStmtForeachStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	conv_g := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/ForeachStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_g.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstStmtForeachStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.Statement getBody()
func (jbobject *AstStmtForeachStmt) GetBody() *AstStmtStatement {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBody", "com/github/javaparser/ast/stmt/Statement")
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

// public com.github.javaparser.ast.expr.Expression getIterable()
func (jbobject *AstStmtForeachStmt) GetIterable() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIterable", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.VariableDeclarationExpr getVariable()
func (jbobject *AstStmtForeachStmt) GetVariable() *AstExprVariableDeclarationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVariable", "com/github/javaparser/ast/expr/VariableDeclarationExpr")
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
	unique_x := &AstExprVariableDeclarationExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public void setBody(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtForeachStmt) SetBody(a AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBody", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setIterable(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtForeachStmt) SetIterable(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIterable", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setVariable(com.github.javaparser.ast.expr.VariableDeclarationExpr)
func (jbobject *AstStmtForeachStmt) SetVariable(a AstExprVariableDeclarationExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVariable", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


