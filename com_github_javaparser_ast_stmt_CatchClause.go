package javaparser

import "github.com/timob/javabind"

type AstStmtCatchClauseInterface interface {
	AstNodeInterface

	// public com.github.javaparser.ast.stmt.BlockStmt getCatchBlock()
	GetCatchBlock() *AstStmtBlockStmt

	// public com.github.javaparser.ast.body.Parameter getParam()
	GetParam() *AstBodyParameter

	// public void setCatchBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetCatchBlock(a AstStmtBlockStmtInterface) 

	// public void setParam(com.github.javaparser.ast.body.Parameter)
	SetParam(a AstBodyParameterInterface) 
}

type AstStmtCatchClause struct {
	AstNode
}

// public com.github.javaparser.ast.stmt.CatchClause()
func NewAstStmtCatchClause() (*AstStmtCatchClause) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/CatchClause")
	if err != nil {
		panic(err)
	}
	x := &AstStmtCatchClause{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.CatchClause(com.github.javaparser.ast.body.Parameter, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtCatchClause2(a AstBodyParameterInterface, b AstStmtBlockStmtInterface) (*AstStmtCatchClause) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/CatchClause", conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstStmtCatchClause{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.CatchClause(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclaratorId, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstStmtCatchClause3(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeTypeInterface, h AstBodyVariableDeclaratorIdInterface, i AstStmtBlockStmtInterface) (*AstStmtCatchClause) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaCallable()
	conv_i := javabind.NewGoToJavaCallable()
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/stmt/CatchClause", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/Type"), conv_h.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_i.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	x := &AstStmtCatchClause{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt getCatchBlock()
func (jbobject *AstStmtCatchClause) GetCatchBlock() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCatchBlock", "com/github/javaparser/ast/stmt/BlockStmt")
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

// public com.github.javaparser.ast.body.Parameter getParam()
func (jbobject *AstStmtCatchClause) GetParam() *AstBodyParameter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParam", "com/github/javaparser/ast/body/Parameter")
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
	unique_x := &AstBodyParameter{}
	unique_x.Callable = dst
	return unique_x
}

// public void setCatchBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstStmtCatchClause) SetCatchBlock(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setCatchBlock", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setParam(com.github.javaparser.ast.body.Parameter)
func (jbobject *AstStmtCatchClause) SetParam(a AstBodyParameterInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setParam", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


