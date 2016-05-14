package javaparser

import "github.com/timob/javabind"

type AstBodyTypeDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public final java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getMembers()
	GetMembers() []*AstBodyBodyDeclaration

	// public final int getModifiers()
	GetModifiers() int

	// public final java.lang.String getName()
	GetName() string

	// public void setMembers(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
	SetMembers(a []*AstBodyBodyDeclaration) 

	// public final void setModifiers(int)
	SetModifiers(a int) 

	// public final void setName(java.lang.String)
	SetName(a string) 

	// public final void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
	SetNameExpr(a AstExprNameExprInterface) 

	// public final com.github.javaparser.ast.expr.NameExpr getNameExpr()
	GetNameExpr() *AstExprNameExpr
}

type AstBodyTypeDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.TypeDeclaration()
func NewAstBodyTypeDeclaration() (*AstBodyTypeDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/TypeDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.TypeDeclaration(int, java.lang.String)
func NewAstBodyTypeDeclaration2(a int, b string) (*AstBodyTypeDeclaration) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/TypeDeclaration", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.TypeDeclaration(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, int, java.lang.String, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyTypeDeclaration3(a []*AstExprAnnotationExpr, b int, c string, d []*AstBodyBodyDeclaration) (*AstBodyTypeDeclaration) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/TypeDeclaration", conv_a.Value().Cast("java/util/List"), b, conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.TypeDeclaration(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, int, java.lang.String, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyTypeDeclaration4(a int, b int, c int, d int, e []*AstExprAnnotationExpr, f int, g string, h []*AstBodyBodyDeclaration) (*AstBodyTypeDeclaration) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/TypeDeclaration", a, b, c, d, conv_e.Value().Cast("java/util/List"), f, conv_g.Value().Cast("java/lang/String"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyTypeDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getMembers()
func (jbobject *AstBodyTypeDeclaration) GetMembers() []*AstBodyBodyDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getMembers", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyBodyDeclaration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final int getModifiers()
func (jbobject *AstBodyTypeDeclaration) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final java.lang.String getName()
func (jbobject *AstBodyTypeDeclaration) GetName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "java/lang/String")
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

// public void setMembers(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func (jbobject *AstBodyTypeDeclaration) SetMembers(a []*AstBodyBodyDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setMembers", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final void setModifiers(int)
func (jbobject *AstBodyTypeDeclaration) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final void setName(java.lang.String)
func (jbobject *AstBodyTypeDeclaration) SetName(a string)  {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstBodyTypeDeclaration) SetNameExpr(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setNameExpr", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final com.github.javaparser.ast.expr.NameExpr getNameExpr()
func (jbobject *AstBodyTypeDeclaration) GetNameExpr() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getNameExpr", "com/github/javaparser/ast/expr/NameExpr")
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


