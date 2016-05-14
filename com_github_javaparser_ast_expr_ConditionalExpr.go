package javaparser

import "github.com/timob/javabind"

type AstExprConditionalExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getCondition()
	GetCondition() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getElseExpr()
	GetElseExpr() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getThenExpr()
	GetThenExpr() *AstExprExpression

	// public void setCondition(com.github.javaparser.ast.expr.Expression)
	SetCondition(a AstExprExpressionInterface) 

	// public void setElseExpr(com.github.javaparser.ast.expr.Expression)
	SetElseExpr(a AstExprExpressionInterface) 

	// public void setThenExpr(com.github.javaparser.ast.expr.Expression)
	SetThenExpr(a AstExprExpressionInterface) 
}

type AstExprConditionalExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ConditionalExpr()
func NewAstExprConditionalExpr() (*AstExprConditionalExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ConditionalExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprConditionalExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ConditionalExpr(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstExprConditionalExpr2(a AstExprExpressionInterface, b AstExprExpressionInterface, c AstExprExpressionInterface) (*AstExprConditionalExpr) {
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ConditionalExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_c.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstExprConditionalExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ConditionalExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstExprConditionalExpr3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstExprExpressionInterface, g AstExprExpressionInterface) (*AstExprConditionalExpr) {
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ConditionalExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_g.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstExprConditionalExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getCondition()
func (jbobject *AstExprConditionalExpr) GetCondition() *AstExprExpression {
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

// public com.github.javaparser.ast.expr.Expression getElseExpr()
func (jbobject *AstExprConditionalExpr) GetElseExpr() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getElseExpr", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.Expression getThenExpr()
func (jbobject *AstExprConditionalExpr) GetThenExpr() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThenExpr", "com/github/javaparser/ast/expr/Expression")
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

// public void setCondition(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprConditionalExpr) SetCondition(a AstExprExpressionInterface)  {
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

// public void setElseExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprConditionalExpr) SetElseExpr(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setElseExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setThenExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprConditionalExpr) SetThenExpr(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThenExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


