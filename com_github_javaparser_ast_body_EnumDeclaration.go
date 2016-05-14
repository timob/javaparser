package javaparser

import "github.com/timob/javabind"

type AstBodyEnumDeclarationInterface interface {
	AstBodyTypeDeclarationInterface

	// public java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration> getEntries()
	GetEntries() []*AstBodyEnumConstantDeclaration

	// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getImplements()
	GetImplements() []*AstTypeClassOrInterfaceType

	// public void setEntries(java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration>)
	SetEntries(a []*AstBodyEnumConstantDeclaration) 

	// public void setImplements(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
	SetImplements(a []*AstTypeClassOrInterfaceType) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyEnumDeclaration struct {
	AstBodyTypeDeclaration
}

// public com.github.javaparser.ast.body.EnumDeclaration()
func NewAstBodyEnumDeclaration() (*AstBodyEnumDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyEnumDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumDeclaration(int, java.lang.String)
func NewAstBodyEnumDeclaration2(a int, b string) (*AstBodyEnumDeclaration) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumDeclaration", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyEnumDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyEnumDeclaration3(a int, b []*AstExprAnnotationExpr, c string, d []*AstTypeClassOrInterfaceType, e []*AstBodyEnumConstantDeclaration, f []*AstBodyBodyDeclaration) (*AstBodyEnumDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/util/List"), conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstBodyEnumDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>, java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyEnumDeclaration4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g string, h []*AstTypeClassOrInterfaceType, i []*AstBodyEnumConstantDeclaration, j []*AstBodyBodyDeclaration) (*AstBodyEnumDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_i := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_j := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/lang/String"), conv_h.Value().Cast("java/util/List"), conv_i.Value().Cast("java/util/List"), conv_j.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	conv_j.CleanUp()
	x := &AstBodyEnumDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration> getEntries()
func (jbobject *AstBodyEnumDeclaration) GetEntries() []*AstBodyEnumConstantDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEntries", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyEnumConstantDeclaration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType> getImplements()
func (jbobject *AstBodyEnumDeclaration) GetImplements() []*AstTypeClassOrInterfaceType {
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

// public void setEntries(java.util.List<com.github.javaparser.ast.body.EnumConstantDeclaration>)
func (jbobject *AstBodyEnumDeclaration) SetEntries(a []*AstBodyEnumConstantDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEntries", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImplements(java.util.List<com.github.javaparser.ast.type.ClassOrInterfaceType>)
func (jbobject *AstBodyEnumDeclaration) SetImplements(a []*AstTypeClassOrInterfaceType)  {
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

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyEnumDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyEnumDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


