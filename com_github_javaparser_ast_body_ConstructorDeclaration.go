package javaparser

import "github.com/timob/javabind"

type AstBodyConstructorDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
	GetBlock() *AstStmtBlockStmt

	// public int getModifiers()
	GetModifiers() int

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.expr.NameExpr getNameExpr()
	GetNameExpr() *AstExprNameExpr

	// public java.util.List<com.github.javaparser.ast.body.Parameter> getParameters()
	GetParameters() []*AstBodyParameter

	// public java.util.List<com.github.javaparser.ast.expr.NameExpr> getThrows()
	GetThrows() []*AstExprNameExpr

	// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
	GetTypeParameters() []*AstTypeParameter

	// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
	SetBlock(a AstStmtBlockStmtInterface) 

	// public void setModifiers(int)
	SetModifiers(a int) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
	SetNameExpr(a AstExprNameExprInterface) 

	// public void setParameters(java.util.List<com.github.javaparser.ast.body.Parameter>)
	SetParameters(a []*AstBodyParameter) 

	// public void setThrows(java.util.List<com.github.javaparser.ast.expr.NameExpr>)
	SetThrows(a []*AstExprNameExpr) 

	// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
	SetTypeParameters(a []*AstTypeParameter) 

	// public java.lang.String getDeclarationAsString(boolean, boolean, boolean)
	GetDeclarationAsString(a bool, b bool, c bool) string

	// public java.lang.String getDeclarationAsString(boolean, boolean)
	GetDeclarationAsString2(a bool, b bool) string

	// public java.lang.String getDeclarationAsString()
	GetDeclarationAsString3() string

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyConstructorDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.ConstructorDeclaration()
func NewAstBodyConstructorDeclaration() (*AstBodyConstructorDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ConstructorDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyConstructorDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ConstructorDeclaration(int, java.lang.String)
func NewAstBodyConstructorDeclaration2(a int, b string) (*AstBodyConstructorDeclaration) {
	conv_b := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ConstructorDeclaration", a, conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	x := &AstBodyConstructorDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ConstructorDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.util.List<com.github.javaparser.ast.TypeParameter>, java.lang.String, java.util.List<com.github.javaparser.ast.body.Parameter>, java.util.List<com.github.javaparser.ast.expr.NameExpr>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyConstructorDeclaration3(a int, b []*AstExprAnnotationExpr, c []*AstTypeParameter, d string, e []*AstBodyParameter, f []*AstExprNameExpr, g AstStmtBlockStmtInterface) (*AstBodyConstructorDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_d := javabind.NewGoToJavaString()
	conv_e := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaCallable()
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
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ConstructorDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("java/util/List"), conv_d.Value().Cast("java/lang/String"), conv_e.Value().Cast("java/util/List"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstBodyConstructorDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.ConstructorDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.util.List<com.github.javaparser.ast.TypeParameter>, java.lang.String, java.util.List<com.github.javaparser.ast.body.Parameter>, java.util.List<com.github.javaparser.ast.expr.NameExpr>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyConstructorDeclaration4(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g []*AstTypeParameter, h string, i []*AstBodyParameter, j []*AstExprNameExpr, k AstStmtBlockStmtInterface) (*AstBodyConstructorDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaString()
	conv_i := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_j := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_k := javabind.NewGoToJavaCallable()
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
	if err := conv_k.Convert(k); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/ConstructorDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("java/lang/String"), conv_i.Value().Cast("java/util/List"), conv_j.Value().Cast("java/util/List"), conv_k.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	conv_j.CleanUp()
	conv_k.CleanUp()
	x := &AstBodyConstructorDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.stmt.BlockStmt getBlock()
func (jbobject *AstBodyConstructorDeclaration) GetBlock() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBlock", "com/github/javaparser/ast/stmt/BlockStmt")
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
	unique_x := &AstStmtBlockStmt{}
	unique_x.Callable = dst
	return unique_x
}

// public int getModifiers()
func (jbobject *AstBodyConstructorDeclaration) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getName()
func (jbobject *AstBodyConstructorDeclaration) GetName() string {
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

// public com.github.javaparser.ast.expr.NameExpr getNameExpr()
func (jbobject *AstBodyConstructorDeclaration) GetNameExpr() *AstExprNameExpr {
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

// public java.util.List<com.github.javaparser.ast.body.Parameter> getParameters()
func (jbobject *AstBodyConstructorDeclaration) GetParameters() []*AstBodyParameter {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParameters", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyParameter)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.expr.NameExpr> getThrows()
func (jbobject *AstBodyConstructorDeclaration) GetThrows() []*AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThrows", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstExprNameExpr)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
func (jbobject *AstBodyConstructorDeclaration) GetTypeParameters() []*AstTypeParameter {
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

// public void setBlock(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstBodyConstructorDeclaration) SetBlock(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBlock", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setModifiers(int)
func (jbobject *AstBodyConstructorDeclaration) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setName(java.lang.String)
func (jbobject *AstBodyConstructorDeclaration) SetName(a string)  {
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

// public void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstBodyConstructorDeclaration) SetNameExpr(a AstExprNameExprInterface)  {
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

// public void setParameters(java.util.List<com.github.javaparser.ast.body.Parameter>)
func (jbobject *AstBodyConstructorDeclaration) SetParameters(a []*AstBodyParameter)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setParameters", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setThrows(java.util.List<com.github.javaparser.ast.expr.NameExpr>)
func (jbobject *AstBodyConstructorDeclaration) SetThrows(a []*AstExprNameExpr)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setThrows", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
func (jbobject *AstBodyConstructorDeclaration) SetTypeParameters(a []*AstTypeParameter)  {
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

// public java.lang.String getDeclarationAsString(boolean, boolean, boolean)
func (jbobject *AstBodyConstructorDeclaration) GetDeclarationAsString(a bool, b bool, c bool) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String", a, b, c)
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

// public java.lang.String getDeclarationAsString(boolean, boolean)
func (jbobject *AstBodyConstructorDeclaration) GetDeclarationAsString2(a bool, b bool) string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String", a, b)
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

// public java.lang.String getDeclarationAsString()
func (jbobject *AstBodyConstructorDeclaration) GetDeclarationAsString3() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDeclarationAsString", "java/lang/String")
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

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyConstructorDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyConstructorDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


