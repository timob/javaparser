package javaparser

import "github.com/timob/javabind"

type AstBodyEnumConstantDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
	GetArgs() []*AstExprExpression

	// public java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getClassBody()
	GetClassBody() []*AstBodyBodyDeclaration

	// public java.lang.String getName()
	GetName() string

	// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
	SetArgs(a []*AstExprExpression) 

	// public void setClassBody(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
	SetClassBody(a []*AstBodyBodyDeclaration) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyEnumConstantDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.EnumConstantDeclaration()
func NewAstBodyEnumConstantDeclaration() (*AstBodyEnumConstantDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumConstantDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyEnumConstantDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumConstantDeclaration(java.lang.String)
func NewAstBodyEnumConstantDeclaration2(a string) (*AstBodyEnumConstantDeclaration) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumConstantDeclaration", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstBodyEnumConstantDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumConstantDeclaration(java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.expr.Expression>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyEnumConstantDeclaration3(a []*AstExprAnnotationExpr, b string, c []*AstExprExpression, d []*AstBodyBodyDeclaration) (*AstBodyEnumConstantDeclaration) {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_d := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumConstantDeclaration", conv_a.Value().Cast("java/util/List"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/util/List"), conv_d.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyEnumConstantDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.EnumConstantDeclaration(int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.lang.String, java.util.List<com.github.javaparser.ast.expr.Expression>, java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func NewAstBodyEnumConstantDeclaration4(a int, b int, c int, d int, e []*AstExprAnnotationExpr, f string, g []*AstExprExpression, h []*AstBodyBodyDeclaration) (*AstBodyEnumConstantDeclaration) {
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaString()
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/EnumConstantDeclaration", a, b, c, d, conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("java/lang/String"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	x := &AstBodyEnumConstantDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.expr.Expression> getArgs()
func (jbobject *AstBodyEnumConstantDeclaration) GetArgs() []*AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArgs", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprExpression)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.body.BodyDeclaration> getClassBody()
func (jbobject *AstBodyEnumConstantDeclaration) GetClassBody() []*AstBodyBodyDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getClassBody", "java/util/List")
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

// public java.lang.String getName()
func (jbobject *AstBodyEnumConstantDeclaration) GetName() string {
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

// public void setArgs(java.util.List<com.github.javaparser.ast.expr.Expression>)
func (jbobject *AstBodyEnumConstantDeclaration) SetArgs(a []*AstExprExpression)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArgs", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setClassBody(java.util.List<com.github.javaparser.ast.body.BodyDeclaration>)
func (jbobject *AstBodyEnumConstantDeclaration) SetClassBody(a []*AstBodyBodyDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setClassBody", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setName(java.lang.String)
func (jbobject *AstBodyEnumConstantDeclaration) SetName(a string)  {
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

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyEnumConstantDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyEnumConstantDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


