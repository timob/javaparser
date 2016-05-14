package javaparser

import "github.com/timob/javabind"

type AstBodyAnnotationDeclarationInterface interface {
	AstBodyTypeDeclarationInterface

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyAnnotationDeclaration struct {
	AstBodyTypeDeclaration
}

// public com.github.javaparser.ast.body.AnnotationDeclaration()
func NewAstBodyAnnotationDeclaration() (*AstBodyAnnotationDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyAnnotationDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationDeclaration(int, java.lang.String)
func NewAstBodyAnnotationDeclaration2(a int, b string) (*AstBodyAnnotationDeclaration) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationDeclaration", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyAnnotationDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyAnnotationDeclaration3(a int, b []*AstExprAnnotationExpr, c string, d []*AstBodyBodyDeclaration) (*AstBodyAnnotationDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaString()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyAnnotationDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyAnnotationDeclaration4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g string, h []*AstBodyBodyDeclaration) (*AstBodyAnnotationDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaString()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/lang/String"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyAnnotationDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyAnnotationDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyAnnotationDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


