package javaparser

import "github.com/timob/javabind"

type AstBodyAnnotationMemberDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public com.github.javaparser.ast.expr.Expression getDefaultValue()
	GetDefaultValue() *AstExprExpression

	// public int getModifiers()
	GetModifiers() int

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public void setDefaultValue(com.github.javaparser.ast.expr.Expression)
	SetDefaultValue(a AstExprExpressionInterface) 

	// public void setModifiers(int)
	SetModifiers(a int) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyAnnotationMemberDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.AnnotationMemberDeclaration()
func NewAstBodyAnnotationMemberDeclaration() (*AstBodyAnnotationMemberDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationMemberDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyAnnotationMemberDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationMemberDeclaration(int, com.github.javaparser.ast.type.Type, java.lang.String, com.github.javaparser.ast.expr.Expression)
func NewAstBodyAnnotationMemberDeclaration2(a int, b AstTypeTypeInterface, c string, d AstExprExpressionInterface) (*AstBodyAnnotationMemberDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	conv_d := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	if err := conv_d.Convert(d); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationMemberDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyAnnotationMemberDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationMemberDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, java.lang.String, com.github.javaparser.ast.expr.Expression)
func NewAstBodyAnnotationMemberDeclaration3(a int, b []*AstExprAnnotationExpr, c AstTypeTypeInterface, d string, e AstExprExpressionInterface) (*AstBodyAnnotationMemberDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaCallable()
	conv_d := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationMemberDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("com/github/javaparser/ast/type/Type"), conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	x := &AstBodyAnnotationMemberDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.AnnotationMemberDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, com.github.javaparser.ast.type.Type, java.lang.String, com.github.javaparser.ast.expr.Expression)
func NewAstBodyAnnotationMemberDeclaration4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g AstTypeTypeInterface, h string, i AstExprExpressionInterface) (*AstBodyAnnotationMemberDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
	conv_h := javabind.NewGoToJavaString()
	conv_i := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/AnnotationMemberDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/type/Type"), conv_h.Value().Cast("java/lang/String"), conv_i.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	x := &AstBodyAnnotationMemberDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.Expression getDefaultValue()
func (jbobject *AstBodyAnnotationMemberDeclaration) GetDefaultValue() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDefaultValue", "com/github/javaparser/ast/expr/Expression")
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
	unique_x := &AstExprExpression{}
	unique_x.Callable = dst
	return unique_x
}

// public int getModifiers()
func (jbobject *AstBodyAnnotationMemberDeclaration) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getName()
func (jbobject *AstBodyAnnotationMemberDeclaration) GetName() string {
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

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstBodyAnnotationMemberDeclaration) GetType() *AstTypeType {
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

// public void setDefaultValue(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstBodyAnnotationMemberDeclaration) SetDefaultValue(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefaultValue", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setModifiers(int)
func (jbobject *AstBodyAnnotationMemberDeclaration) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setName(java.lang.String)
func (jbobject *AstBodyAnnotationMemberDeclaration) SetName(a string)  {
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

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstBodyAnnotationMemberDeclaration) SetType(a AstTypeTypeInterface)  {
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

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyAnnotationMemberDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyAnnotationMemberDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


