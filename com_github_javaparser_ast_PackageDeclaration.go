package javaparser

import "github.com/timob/javabind"

type AstPackageDeclarationInterface interface {
	AstNodeInterface

	// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
	GetAnnotations() []*AstExprAnnotationExpr

	// public com.github.javaparser.ast.expr.NameExpr getName()
	GetName() *AstExprNameExpr

	// public java.lang.String getPackageName()
	GetPackageName() string

	// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
	SetAnnotations(a []*AstExprAnnotationExpr) 

	// public void setName(com.github.javaparser.ast.expr.NameExpr)
	SetName(a AstExprNameExprInterface) 
}

type AstPackageDeclaration struct {
	AstNode
}

// public com.github.javaparser.ast.PackageDeclaration()
func NewAstPackageDeclaration() (*AstPackageDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/PackageDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstPackageDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.PackageDeclaration(com.github.javaparser.ast.expr.NameExpr)
func NewAstPackageDeclaration2(a AstExprNameExprInterface) (*AstPackageDeclaration) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/PackageDeclaration", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstPackageDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.PackageDeclaration(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.expr.NameExpr)
func NewAstPackageDeclaration3(a []*AstExprAnnotationExpr, b AstExprNameExprInterface) (*AstPackageDeclaration) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/PackageDeclaration", conv_a.Value().Cast("java/util/List"), conv_b.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstPackageDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.PackageDeclaration(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.expr.NameExpr)
func NewAstPackageDeclaration4(a int, b int, c int, d int, e []*AstExprAnnotationExpr, f AstExprNameExprInterface) (*AstPackageDeclaration) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/PackageDeclaration", a, b, c, d, conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstPackageDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.AnnotationExpr> getAnnotations()
func (jbobject *AstPackageDeclaration) GetAnnotations() []*AstExprAnnotationExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAnnotations", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprAnnotationExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.expr.NameExpr getName()
func (jbobject *AstPackageDeclaration) GetName() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "com/github/javaparser/ast/expr/NameExpr")
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

// public java.lang.String getPackageName()
func (jbobject *AstPackageDeclaration) GetPackageName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPackageName", "java/lang/String")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setAnnotations(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>)
func (jbobject *AstPackageDeclaration) SetAnnotations(a []*AstExprAnnotationExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAnnotations", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setName(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstPackageDeclaration) SetName(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


