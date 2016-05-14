package javaparser

import "github.com/timob/javabind"

type AstBodyFieldDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public int getModifiers()
	GetModifiers() int

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public java.util.List<com.github.javaparser.ast.body.VariableDeclarator> getVariables()
	GetVariables() []*AstBodyVariableDeclarator

	// public void setModifiers(int)
	SetModifiers(a int) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public void setVariables(java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
	SetVariables(a []*AstBodyVariableDeclarator) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyFieldDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.FieldDeclaration()
func NewAstBodyFieldDeclaration() (*AstBodyFieldDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/FieldDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyFieldDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.FieldDeclaration(int, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclarator)
func NewAstBodyFieldDeclaration2(a int, b AstTypeTypeInterface, c AstBodyVariableDeclaratorInterface) (*AstBodyFieldDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/FieldDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstBodyFieldDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.FieldDeclaration(int, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstBodyFieldDeclaration3(a int, b AstTypeTypeInterface, c []*AstBodyVariableDeclarator) (*AstBodyFieldDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/FieldDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstBodyFieldDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.FieldDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstBodyFieldDeclaration4(a int, b []*AstExprAnnotationExpr, c AstTypeTypeInterface, d []*AstBodyVariableDeclarator) (*AstBodyFieldDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/FieldDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("com/github/javaparser/ast/type/Type"), conv_d.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyFieldDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.FieldDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func NewAstBodyFieldDeclaration5(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeTypeInterface, h []*AstBodyVariableDeclarator) (*AstBodyFieldDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/FieldDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/Type"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyFieldDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getModifiers()
func (jbobject *AstBodyFieldDeclaration) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstBodyFieldDeclaration) GetType() *AstTypeType {
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

// public java.util.List<com.github.javaparser.ast.body.VariableDeclarator> getVariables()
func (jbobject *AstBodyFieldDeclaration) GetVariables() []*AstBodyVariableDeclarator {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVariables", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyVariableDeclarator)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setModifiers(int)
func (jbobject *AstBodyFieldDeclaration) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstBodyFieldDeclaration) SetType(a AstTypeTypeInterface)  {
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

// public void setVariables(java.util.List<com.github.javaparser.ast.body.VariableDeclarator>)
func (jbobject *AstBodyFieldDeclaration) SetVariables(a []*AstBodyVariableDeclarator)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setVariables", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyFieldDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setJavaDoc", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
func (jbobject *AstBodyFieldDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getJavaDoc", "com/github/javaparser/ast/comments/JavadocComment")
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
	unique_x := &AstCommentsJavadocComment{}
	unique_x.Callable = dst
	return unique_x
}


