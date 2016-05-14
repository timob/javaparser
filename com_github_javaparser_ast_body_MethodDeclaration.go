package javaparser

import "github.com/timob/javabind"

type AstBodyMethodDeclarationInterface interface {
	AstBodyBodyDeclarationInterface

	// public int getArrayCount()
	GetArrayCount() int

	// public com.github.javaparser.ast.stmt.BlockStmt getBody()
	GetBody() *AstStmtBlockStmt

	// public int getModifiers()
	GetModifiers() int

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.expr.NameExpr getNameExpr()
	GetNameExpr() *AstExprNameExpr

	// public java.util.List<com.github.javaparser.ast.body.Parameter> getParameters()
	GetParameters() []*AstBodyParameter

	// public java.util.List<com.github.javaparser.ast.type.ReferenceType> getThrows()
	GetThrows() []*AstTypeReferenceType

	// public com.github.javaparser.ast.type.Type getType()
	GetType() *AstTypeType

	// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
	GetTypeParameters() []*AstTypeParameter

	// public void setArrayCount(int)
	SetArrayCount(a int) 

	// public void setBody(com.github.javaparser.ast.stmt.BlockStmt)
	SetBody(a AstStmtBlockStmtInterface) 

	// public void setModifiers(int)
	SetModifiers(a int) 

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setNameExpr(com.github.javaparser.ast.expr.NameExpr)
	SetNameExpr(a AstExprNameExprInterface) 

	// public void setParameters(java.util.List<com.github.javaparser.ast.body.Parameter>)
	SetParameters(a []*AstBodyParameter) 

	// public void setThrows(java.util.List<com.github.javaparser.ast.type.ReferenceType>)
	SetThrows(a []*AstTypeReferenceType) 

	// public void setType(com.github.javaparser.ast.type.Type)
	SetType(a AstTypeTypeInterface) 

	// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
	SetTypeParameters(a []*AstTypeParameter) 

	// public boolean isDefault()
	IsDefault() bool

	// public void setDefault(boolean)
	SetDefault(a bool) 

	// public java.lang.String getDeclarationAsString()
	GetDeclarationAsString() string

	// public java.lang.String getDeclarationAsString(boolean, boolean)
	GetDeclarationAsString2(a bool, b bool) string

	// public java.lang.String getDeclarationAsString(boolean, boolean, boolean)
	GetDeclarationAsString3(a bool, b bool, c bool) string

	// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
	SetJavaDoc(a AstCommentsJavadocCommentInterface) 

	// public com.github.javaparser.ast.comments.JavadocComment getJavaDoc()
	GetJavaDoc() *AstCommentsJavadocComment
}

type AstBodyMethodDeclaration struct {
	AstBodyBodyDeclaration
}

// public com.github.javaparser.ast.body.MethodDeclaration()
func NewAstBodyMethodDeclaration() (*AstBodyMethodDeclaration) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MethodDeclaration")
	if err != nil {
		panic(err)
	}
	x := &AstBodyMethodDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MethodDeclaration(int, com.github.javaparser.ast.type.Type, java.lang.String)
func NewAstBodyMethodDeclaration2(a int, b AstTypeTypeInterface, c string) (*AstBodyMethodDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MethodDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstBodyMethodDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MethodDeclaration(int, com.github.javaparser.ast.type.Type, java.lang.String, java.util.List<com.github.javaparser.ast.body.Parameter>)
func NewAstBodyMethodDeclaration3(a int, b AstTypeTypeInterface, c string, d []*AstBodyParameter) (*AstBodyMethodDeclaration) {
	conv_b := javabind.NewGoToJavaCallable()
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

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MethodDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/lang/String"), conv_d.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	x := &AstBodyMethodDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MethodDeclaration(int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.util.List<com.github.javaparser.ast.TypeParameter>, com.github.javaparser.ast.type.Type, java.lang.String, java.util.List<com.github.javaparser.ast.body.Parameter>, int, java.util.List<com.github.javaparser.ast.type.ReferenceType>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyMethodDeclaration4(a int, b []*AstExprAnnotationExpr, c []*AstTypeParameter, d AstTypeTypeInterface, e string, f []*AstBodyParameter, g int, h []*AstTypeReferenceType, i AstStmtBlockStmtInterface) (*AstBodyMethodDeclaration) {
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_d := javabind.NewGoToJavaCallable()
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_i := javabind.NewGoToJavaCallable()
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
	if err := conv_h.Convert(h); err != nil {
		panic(err)
	}
	if err := conv_i.Convert(i); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MethodDeclaration", a, conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("java/util/List"), conv_d.Value().Cast("com/github/javaparser/ast/type/Type"), conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("java/util/List"), g, conv_h.Value().Cast("java/util/List"), conv_i.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	conv_d.CleanUp()
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	x := &AstBodyMethodDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.MethodDeclaration(int, int, int, int, int, java.util.List<com.github.javaparser.ast.expr.AnnotationExpr>, java.util.List<com.github.javaparser.ast.TypeParameter>, com.github.javaparser.ast.type.Type, java.lang.String, java.util.List<com.github.javaparser.ast.body.Parameter>, int, java.util.List<com.github.javaparser.ast.type.ReferenceType>, com.github.javaparser.ast.stmt.BlockStmt)
func NewAstBodyMethodDeclaration5(a int, b int, c int, d int, e int, f []*AstExprAnnotationExpr, g []*AstTypeParameter, h AstTypeTypeInterface, i string, j []*AstBodyParameter, k int, l []*AstTypeReferenceType, m AstStmtBlockStmtInterface) (*AstBodyMethodDeclaration) {
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_h := javabind.NewGoToJavaCallable()
	conv_i := javabind.NewGoToJavaString()
	conv_j := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_l := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_m := javabind.NewGoToJavaCallable()
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
	if err := conv_l.Convert(l); err != nil {
		panic(err)
	}
	if err := conv_m.Convert(m); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/MethodDeclaration", a, b, c, d, e, conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/util/List"), conv_h.Value().Cast("com/github/javaparser/ast/type/Type"), conv_i.Value().Cast("java/lang/String"), conv_j.Value().Cast("java/util/List"), k, conv_l.Value().Cast("java/util/List"), conv_m.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_f.CleanUp()
	conv_g.CleanUp()
	conv_h.CleanUp()
	conv_i.CleanUp()
	conv_j.CleanUp()
	conv_l.CleanUp()
	conv_m.CleanUp()
	x := &AstBodyMethodDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public int getArrayCount()
func (jbobject *AstBodyMethodDeclaration) GetArrayCount() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getArrayCount", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public com.github.javaparser.ast.stmt.BlockStmt getBody()
func (jbobject *AstBodyMethodDeclaration) GetBody() *AstStmtBlockStmt {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBody", "com/github/javaparser/ast/stmt/BlockStmt")
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
func (jbobject *AstBodyMethodDeclaration) GetModifiers() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getModifiers", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String getName()
func (jbobject *AstBodyMethodDeclaration) GetName() string {
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
func (jbobject *AstBodyMethodDeclaration) GetNameExpr() *AstExprNameExpr {
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
func (jbobject *AstBodyMethodDeclaration) GetParameters() []*AstBodyParameter {
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

// public java.util.List<com.github.javaparser.ast.type.ReferenceType> getThrows()
func (jbobject *AstBodyMethodDeclaration) GetThrows() []*AstTypeReferenceType {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getThrows", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstTypeReferenceType)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.type.Type getType()
func (jbobject *AstBodyMethodDeclaration) GetType() *AstTypeType {
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

// public java.util.List<com.github.javaparser.ast.TypeParameter> getTypeParameters()
func (jbobject *AstBodyMethodDeclaration) GetTypeParameters() []*AstTypeParameter {
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

// public void setArrayCount(int)
func (jbobject *AstBodyMethodDeclaration) SetArrayCount(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setArrayCount", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setBody(com.github.javaparser.ast.stmt.BlockStmt)
func (jbobject *AstBodyMethodDeclaration) SetBody(a AstStmtBlockStmtInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBody", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setModifiers(int)
func (jbobject *AstBodyMethodDeclaration) SetModifiers(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setModifiers", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setName(java.lang.String)
func (jbobject *AstBodyMethodDeclaration) SetName(a string)  {
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
func (jbobject *AstBodyMethodDeclaration) SetNameExpr(a AstExprNameExprInterface)  {
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
func (jbobject *AstBodyMethodDeclaration) SetParameters(a []*AstBodyParameter)  {
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

// public void setThrows(java.util.List<com.github.javaparser.ast.type.ReferenceType>)
func (jbobject *AstBodyMethodDeclaration) SetThrows(a []*AstTypeReferenceType)  {
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

// public void setType(com.github.javaparser.ast.type.Type)
func (jbobject *AstBodyMethodDeclaration) SetType(a AstTypeTypeInterface)  {
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

// public void setTypeParameters(java.util.List<com.github.javaparser.ast.TypeParameter>)
func (jbobject *AstBodyMethodDeclaration) SetTypeParameters(a []*AstTypeParameter)  {
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

// public boolean isDefault()
func (jbobject *AstBodyMethodDeclaration) IsDefault() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDefault", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setDefault(boolean)
func (jbobject *AstBodyMethodDeclaration) SetDefault(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setDefault", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String getDeclarationAsString()
func (jbobject *AstBodyMethodDeclaration) GetDeclarationAsString() string {
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

// public java.lang.String getDeclarationAsString(boolean, boolean)
func (jbobject *AstBodyMethodDeclaration) GetDeclarationAsString2(a bool, b bool) string {
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

// public java.lang.String getDeclarationAsString(boolean, boolean, boolean)
func (jbobject *AstBodyMethodDeclaration) GetDeclarationAsString3(a bool, b bool, c bool) string {
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

// public void setJavaDoc(com.github.javaparser.ast.comments.JavadocComment)
func (jbobject *AstBodyMethodDeclaration) SetJavaDoc(a AstCommentsJavadocCommentInterface)  {
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
func (jbobject *AstBodyMethodDeclaration) GetJavaDoc() *AstCommentsJavadocComment {
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


