package javaparser

import "github.com/timob/javabind"

type AstExprCastExprInterface interface {
	AstExprExpressionInterface

	// public com.github.javaparser.ast.expr.Expression getExpr()
	GetExpr() *AstExprExpression

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setExpr(com.github.javaparser.ast.expr.Expression)
	SetExpr(a AstExprExpressionInterface) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 
}

type AstExprCastExpr struct {
	AstExprExpression
}

// public com.github.javaparser.ast.expr.CastExpr()
func NewAstExprCastExpr() (*AstExprCastExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CastExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprCastExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.CastExpr(com.github.javaparser.ast.type.Type, com.github.javaparser.ast.expr.Expression)
func NewAstExprCastExpr2(a AstTypeTypeInterface, b AstExprExpressionInterface) (*AstExprCastExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CastExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprCastExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.CastExpr(int, int, int, int, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.expr.Expression)
func NewAstExprCastExpr3(a int, b int, c int, d int, e AstTypeTypeInterface, f AstExprExpressionInterface) (*AstExprCastExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/CastExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/type/Type"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprCastExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getExpr()
func (jbobject *AstExprCastExpr) GetExpr() *AstExprExpression {
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

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstExprCastExpr) GetType() *AstTypeType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getType", "com/github/javaparser/ast/type/Type")
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
	unique_x := &AstTypeType{}
	unique_x.Callable = dst
	return unique_x
}

// public void setExpr(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprCastExpr) SetExpr(a AstExprExpressionInterface)  {
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

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstExprCastExpr) SetType(a AstTypeTypeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setType", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/Type"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


