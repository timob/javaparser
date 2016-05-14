package javaparser

import "github.com/timob/javabind"

type AstExprQualifiedNameExprInterface interface {
	AstExprNameExprInterface

	// public com.github.javaparser.ast.expr.NameExpr getQualifier()
	GetQualifier() *AstExprNameExpr

	// public void setQualifier(com.github.javaparser.ast.expr.NameExpr)
	SetQualifier(a AstExprNameExprInterface) 
}

type AstExprQualifiedNameExpr struct {
	AstExprNameExpr
}

// public com.github.javaparser.ast.expr.QualifiedNameExpr()
func NewAstExprQualifiedNameExpr() (*AstExprQualifiedNameExpr) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/QualifiedNameExpr")
	if err != nil {
		panic(err)
	}
	x := &AstExprQualifiedNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.QualifiedNameExpr(com.github.javaparser.ast.expr.NameExpr, java.lang.String)
func NewAstExprQualifiedNameExpr2(a AstExprNameExprInterface, b string) (*AstExprQualifiedNameExpr) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/QualifiedNameExpr", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprQualifiedNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.QualifiedNameExpr(int, int, int, int, com.github.javaparser.ast.expr.NameExpr, java.lang.String)
func NewAstExprQualifiedNameExpr3(a int, b int, c int, d int, e AstExprNameExprInterface, f string) (*AstExprQualifiedNameExpr) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaString()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/QualifiedNameExpr", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_f.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprQualifiedNameExpr{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.NameExpr getQualifier()
func (jbobject *AstExprQualifiedNameExpr) GetQualifier() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getQualifier", "com/github/javaparser/ast/expr/NameExpr")
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
	unique_x := &AstExprNameExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public void setQualifier(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstExprQualifiedNameExpr) SetQualifier(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setQualifier", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


