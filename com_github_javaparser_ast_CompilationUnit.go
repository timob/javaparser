package javaparser

import "github.com/timob/javabind"

type AstCompilationUnitInterface interface {
	AstNodeInterface

	// public java.util.List<com.github.javaparser.ast.comments.Comment> getComments()
	GetComments() []*AstCommentsComment

	// public java.util.List<com.github.javaparser.ast.ImportDeclaration> getImports()
	GetImports() []*AstImportDeclaration

	// public com.github.javaparser.ast.PackageDeclaration getPackage()
	GetPackage() *AstPackageDeclaration

	// public java.util.List<com.github.javaparser.ast.body.TypeDeclaration> getTypes()
	GetTypes() []*AstBodyTypeDeclaration

	// public void setComments(java.util.List<com.github.javaparser.ast.comments.Comment>)
	SetComments(a []*AstCommentsComment) 

	// public void setImports(java.util.List<com.github.javaparser.ast.ImportDeclaration>)
	SetImports(a []*AstImportDeclaration) 

	// public void setPackage(com.github.javaparser.ast.PackageDeclaration)
	SetPackage(a AstPackageDeclarationInterface) 

	// public void setTypes(java.util.List<com.github.javaparser.ast.body.TypeDeclaration>)
	SetTypes(a []*AstBodyTypeDeclaration) 
}

type AstCompilationUnit struct {
	AstNode
}

// public com.github.javaparser.ast.CompilationUnit()
func NewAstCompilationUnit() (*AstCompilationUnit) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/CompilationUnit")
	if err != nil {
		panic(err)
	}
	x := &AstCompilationUnit{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.CompilationUnit(com.github.javaparser.ast.PackageDeclaration, java.util.List<com.github.javaparser.ast.ImportDeclaration>, java.util.List<com.github.javaparser.ast.body.TypeDeclaration>)
func NewAstCompilationUnit2(a AstPackageDeclarationInterface, b []*AstImportDeclaration, c []*AstBodyTypeDeclaration) (*AstCompilationUnit) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_c := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_b.Value().Cast("java/util/List"), conv_c.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &AstCompilationUnit{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.CompilationUnit(int, int, int, int, com.github.javaparser.ast.PackageDeclaration, java.util.List<com.github.javaparser.ast.ImportDeclaration>, java.util.List<com.github.javaparser.ast.body.TypeDeclaration>)
func NewAstCompilationUnit3(a int, b int, c int, d int, e AstPackageDeclarationInterface, f []*AstImportDeclaration, g []*AstBodyTypeDeclaration) (*AstCompilationUnit) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	conv_g := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}
	if err := conv_g.Convert(g); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/CompilationUnit", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_f.Value().Cast("java/util/List"), conv_g.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	conv_g.CleanUp()
	x := &AstCompilationUnit{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.List<com.github.javaparser.ast.comments.Comment> getComments()
func (jbobject *AstCompilationUnit) GetComments() []*AstCommentsComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.ImportDeclaration> getImports()
func (jbobject *AstCompilationUnit) GetImports() []*AstImportDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getImports", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstImportDeclaration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public com.github.javaparser.ast.PackageDeclaration getPackage()
func (jbobject *AstCompilationUnit) GetPackage() *AstPackageDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPackage", "com/github/javaparser/ast/PackageDeclaration")
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
	unique_x := &AstPackageDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.github.javaparser.ast.body.TypeDeclaration> getTypes()
func (jbobject *AstCompilationUnit) GetTypes() []*AstBodyTypeDeclaration {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTypes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstBodyTypeDeclaration)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setComments(java.util.List<com.github.javaparser.ast.comments.Comment>)
func (jbobject *AstCompilationUnit) SetComments(a []*AstCommentsComment)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setComments", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setImports(java.util.List<com.github.javaparser.ast.ImportDeclaration>)
func (jbobject *AstCompilationUnit) SetImports(a []*AstImportDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setImports", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setPackage(com.github.javaparser.ast.PackageDeclaration)
func (jbobject *AstCompilationUnit) SetPackage(a AstPackageDeclarationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setPackage", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setTypes(java.util.List<com.github.javaparser.ast.body.TypeDeclaration>)
func (jbobject *AstCompilationUnit) SetTypes(a []*AstBodyTypeDeclaration)  {
	conv_a := javabind.NewGoToJavaList(javabind.NewGoToJavaCallable())
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setTypes", javabind.Void, conv_a.Value().Cast("java/util/List"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


