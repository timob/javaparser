package javaparser

import "github.com/timob/javabind"

type AstBodyClassOrInterfaceDeclarationInterface interface {
	AstBodyTypeDeclarationInterface

	// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getExtends()
	GetExtends() []*AstTypeClassOrInterfaceType

	// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getImplements()
	GetImplements() []*AstTypeClassOrInterfaceType

	// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
	GetTypeParameters() []*AstTypeParameter

	// public boolean isInterface()
	IsInterface() bool

	// public void setExtends(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
	SetExtends(a []*AstTypeClassOrInterfaceType) 

	// public void setImplements(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
	SetImplements(a []*AstTypeClassOrInterfaceType) 

	// public void setInterface(boolean)
	SetInterface(a bool) 

	// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
	SetTypeParameters(a []*AstTypeParameter) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyClassOrInterfaceDeclaration struct {
	AstBodyTypeDeclaration
}

// public com.github.javaparser.ast.body.ClassOrInterfaceDeclaration()
func NewAstBodyClassOrInterfaceDeclaration() (*AstBodyClassOrInterfaceDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyClassOrInterfaceDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ClassOrInterfaceDeclaration(int, boolean, java.lang.String)
func NewAstBodyClassOrInterfaceDeclaration2(a int, b bool, c string) (*AstBodyClassOrInterfaceDeclaration) {
	conv_c := javabind.NewGoToJavaString()
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration", a, b, conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_c.CleanUp()
	x := &AstBodyClassOrInterfaceDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ClassOrInterfaceDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, boolean, java.lang.String, java.util.List<com.github.javaparser.ast.TypeParameter>, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyClassOrInterfaceDeclaration3(a int, b []*AstExprAnnotationExpr, c bool, d string, e []*AstTypeParameter, f []*AstTypeClassOrInterfaceType, g []*AstTypeClassOrInterfaceType, h []*AstBodyBodyDeclaration) (*AstBodyClassOrInterfaceDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_d := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration", a, conv_b.Value().Cast("java/util/List"), c, conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyClassOrInterfaceDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ClassOrInterfaceDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, boolean, java.lang.String, java.util.List<com.github.javaparser.ast.TypeParameter>, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyClassOrInterfaceDeclaration4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g bool, h string, i []*AstTypeParameter, j []*AstTypeClassOrInterfaceType, k []*AstTypeClassOrInterfaceType, l []*AstBodyBodyDeclaration) (*AstBodyClassOrInterfaceDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaString()
	conv_i := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_j := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_k := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_l := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}
	if err := conv_j.Convert(j); err != nil {
		panic(err)
	}
	if err := conv_k.Convert(k); err != nil {
		panic(err)
	}
	if err := conv_l.Convert(l); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), g, conv_h.Value().Cast("java/lang/String"), conv_i.Value().Cast("java/util/List"), conv_j.Value().Cast("java/util/List"), conv_k.Value().Cast("java/util/List"), conv_l.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	conv_j.CleanUp()
	conv_k.CleanUp()
	conv_l.CleanUp()
	x := &AstBodyClassOrInterfaceDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getExtends()
func (jbobject *AstBodyClassOrInterfaceDeclaration) GetExtends() []*AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getExtends", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeClassOrInterfaceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getImplements()
func (jbobject *AstBodyClassOrInterfaceDeclaration) GetImplements() []*AstTypeClassOrInterfaceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImplements", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeClassOrInterfaceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
func (jbobject *AstBodyClassOrInterfaceDeclaration) GetTypeParameters() []*AstTypeParameter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypeParameters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeParameter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean isInterface()
func (jbobject *AstBodyClassOrInterfaceDeclaration) IsInterface() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isInterface", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setExtends(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func (jbobject *AstBodyClassOrInterfaceDeclaration) SetExtends(a []*AstTypeClassOrInterfaceType)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setExtends", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImplements(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func (jbobject *AstBodyClassOrInterfaceDeclaration) SetImplements(a []*AstTypeClassOrInterfaceType)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImplements", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInterface(boolean)
func (jbobject *AstBodyClassOrInterfaceDeclaration) SetInterface(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInterface", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
func (jbobject *AstBodyClassOrInterfaceDeclaration) SetTypeParameters(a []*AstTypeParameter)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypeParameters", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyClassOrInterfaceDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyClassOrInterfaceDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


