package javaparser

import "github.com/timob/javabind"

type AstExprNormalAnnotationExprInterface interface {
	AstExprAnnotationExprInterface

	// public java.util.List<com.github.javaparser.ast.expr.MemberValuePair> getPairs()
	GetPairs() []*AstExprMemberValuePair

	// public void setPairs(java.util.List<com.github.javaparser.ast.expr.MemberValuePair>)
	SetPairs(a []*AstExprMemberValuePair) 
}

type AstExprNormalAnnotationExpr struct {
	AstExprAnnotationExpr
}

// public com.github.javaparser.ast.expr.NormalAnnotationExpr()
func NewAstExprNormalAnnotationExpr() (*AstExprNormalAnnotationExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NormalAnnotationExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprNormalAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NormalAnnotationExpr(com.github.javaparser.ast.expr.NameExpr, java.util.List<com.github.javaparser.ast.expr.MemberValuePair>)
func NewAstExprNormalAnnotationExpr2(a AstExprNameExprInterface, b []*AstExprMemberValuePair) (*AstExprNormalAnnotationExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NormalAnnotationExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprNormalAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NormalAnnotationExpr(int, int, int, int, com.github.javaparser.ast.expr.NameExpr, java.util.List<com.github.javaparser.ast.expr.MemberValuePair>)
func NewAstExprNormalAnnotationExpr3(a int, b int, c int, d int, e AstExprNameExprInterface, f []*AstExprMemberValuePair) (*AstExprNormalAnnotationExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/NormalAnnotationExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_f.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprNormalAnnotationExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.MemberValuePair> getPairs()
func (jbobject *AstExprNormalAnnotationExpr) GetPairs() []*AstExprMemberValuePair {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPairs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprMemberValuePair)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setPairs(java.util.List<com.github.javaparser.ast.expr.MemberValuePair>)
func (jbobject *AstExprNormalAnnotationExpr) SetPairs(a []*AstExprMemberValuePair)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPairs", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


