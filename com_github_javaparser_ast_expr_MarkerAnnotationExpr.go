package javaparser

import "github.com/timob/javabind"

type AstExprMarkerAnnotationExprInterface interface {
	AstExprAnnotationExprInterface
}

type AstExprMarkerAnnotationExpr struct {
	AstExprAnnotationExpr
}

// public com.github.javaparser.ast.expr.MarkerAnnotationExpr()
func NewAstExprMarkerAnnotationExpr() (*AstExprMarkerAnnotationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MarkerAnnotationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprMarkerAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MarkerAnnotationExpr(com.github.javaparser.ast.expr.NameExpr)
func NewAstExprMarkerAnnotationExpr2(a AstExprNameExprInterface) (*AstExprMarkerAnnotationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MarkerAnnotationExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstExprMarkerAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MarkerAnnotationExpr(int, int, int, int, com.github.javaparser.ast.expr.NameExpr)
func NewAstExprMarkerAnnotationExpr3(a int, b int, c int, d int, e AstExprNameExprInterface) (*AstExprMarkerAnnotationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MarkerAnnotationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstExprMarkerAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}


