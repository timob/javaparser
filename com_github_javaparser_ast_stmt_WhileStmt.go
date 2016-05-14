package javaparser

import "github.com/timob/javabind"

type AstStmtWhileStmtInterface interface {
	AstStmtStatementInterface

	// public com.github.javaparser.ast.stmt.Statement getBody()
	GetBody() *AstStmtStatement

	// public com.github.javaparser.ast.expr.Expression getCondition()
	GetCondition() *AstExprExpression

	// public void setBody(com.github.javaparser.ast.stmt.Statement)
	SetBody(a AstStmtStatementInterface) 

	// public void setCondition(com.github.javaparser.ast.expr.Expression)
	SetCondition(a AstExprExpressionInterface) 
}

type AstStmtWhileStmt struct {
	AstStmtStatement
}

// public com.github.javaparser.ast.stmt.WhileStmt()
func NewAstStmtWhileStmt() (*AstStmtWhileStmt) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/WhileStmt")
	if err != nil {
		panic(err)
	}
	x := &AstStmtWhileStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.WhileStmt(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtWhileStmt2(a AstExprExpressionInterface, b AstStmtStatementInterface) (*AstStmtWhileStmt) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/WhileStmt", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtWhileStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.WhileStmt(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.stmt.Statement)
func NewAstStmtWhileStmt3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstStmtStatementInterface) (*AstStmtWhileStmt) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/WhileStmt", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstStmtWhileStmt{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.Statement getBody()
func (jbobject *AstStmtWhileStmt) GetBody() *AstStmtStatement {
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

// public com.github.javaparser.ast.expr.Expression getCondition()
func (jbobject *AstStmtWhileStmt) GetCondition() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCondition", "com/github/javaparser/ast/expr/Expression")
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

// public void setBody(com.github.javaparser.ast.stmt.Statement)
func (jbobject *AstStmtWhileStmt) SetBody(a AstStmtStatementInterface)  {
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

// public void setCondition(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstStmtWhileStmt) SetCondition(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCondition", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


