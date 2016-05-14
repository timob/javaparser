package javaparser

import "github.com/timob/javabind"

type AstExprArrayAccessExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getIndex()
	GetIndex() *AstExprExpression

	// public com.github.javaparser.ast.expr.Expression getName()
	GetName() *AstExprExpression

	// public void setIndex(com.github.javaparser.ast.expr.Expression)
	SetIndex(a AstExprExpressionInterface) 

	// public void setName(com.github.javaparser.ast.expr.Expression)
	SetName(a AstExprExpressionInterface) 
}

type AstExprArrayAccessExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.ArrayAccessExpr()
func NewAstExprArrayAccessExpr() (*AstExprArrayAccessExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayAccessExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprArrayAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayAccessExpr(com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstExprArrayAccessExpr2(a AstExprExpressionInterface, b AstExprExpressionInterface) (*AstExprArrayAccessExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayAccessExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprArrayAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.ArrayAccessExpr(int, int, int, int, com.github.javaparser.ast.expr.Expression, com.github.javaparser.ast.expr.Expression)
func NewAstExprArrayAccessExpr3(a int, b int, c int, d int, e AstExprExpressionInterface, f AstExprExpressionInterface) (*AstExprArrayAccessExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/ArrayAccessExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprArrayAccessExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getIndex()
func (jbobject *AstExprArrayAccessExpr) GetIndex() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getIndex", "com/github/javaparser/ast/expr/Expression")
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

// public com.github.javaparser.ast.expr.Expression getName()
func (jbobject *AstExprArrayAccessExpr) GetName() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "com/github/javaparser/ast/expr/Expression")
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

// public void setIndex(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprArrayAccessExpr) SetIndex(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setIndex", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setName(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprArrayAccessExpr) SetName(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


