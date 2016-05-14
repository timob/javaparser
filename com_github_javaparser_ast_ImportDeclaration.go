package javaparser

import "github.com/timob/javabind"

type AstImportDeclarationInterface interface {
	AstNodeInterface

	// public static com.github.javaparser.ast.ImportDeclaration createEmptyDeclaration()
	CreateEmptyDeclaration() *AstImportDeclaration

	// public static com.github.javaparser.ast.ImportDeclaration createEmptyDeclaration(int, int, int, int)
	CreateEmptyDeclaration2(a int, b int, c int, d int) *AstImportDeclaration

	// public boolean isEmptyImportDeclaration()
	IsEmptyImportDeclaration() bool

	// public com.github.javaparser.ast.expr.NameExpr getName()
	GetName() *AstExprNameExpr

	// public boolean isAsterisk()
	IsAsterisk() bool

	// public boolean isStatic()
	IsStatic() bool

	// public void setAsterisk(boolean)
	SetAsterisk(a bool) 

	// public void setName(com.github.javaparser.ast.expr.NameExpr)
	SetName(a AstExprNameExprInterface) 

	// public void setStatic(boolean)
	SetStatic(a bool) 
}

type AstImportDeclaration struct {
	AstNode
}

// public com.github.javaparser.ast.ImportDeclaration(com.github.javaparser.ast.expr.NameExpr, boolean, boolean)
func NewAstImportDeclaration(a AstExprNameExprInterface, b bool, c bool) (*AstImportDeclaration) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/ImportDeclaration", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), b, c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstImportDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.ImportDeclaration(int, int, int, int, com.github.javaparser.ast.expr.NameExpr, boolean, boolean)
func NewAstImportDeclaration2(a int, b int, c int, d int, e AstExprNameExprInterface, f bool, g bool) (*AstImportDeclaration) {
	conv_e := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/ImportDeclaration", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), f, g)
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	x := &AstImportDeclaration{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static com.github.javaparser.ast.ImportDeclaration createEmptyDeclaration()
func (jbobject *AstImportDeclaration) CreateEmptyDeclaration() *AstImportDeclaration {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/ImportDeclaration", "createEmptyDeclaration", "com/github/javaparser/ast/ImportDeclaration")
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
	unique_x := &AstImportDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.ImportDeclaration createEmptyDeclaration(int, int, int, int)
func (jbobject *AstImportDeclaration) CreateEmptyDeclaration2(a int, b int, c int, d int) *AstImportDeclaration {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/ImportDeclaration", "createEmptyDeclaration", "com/github/javaparser/ast/ImportDeclaration", a, b, c, d)
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
	unique_x := &AstImportDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public boolean isEmptyImportDeclaration()
func (jbobject *AstImportDeclaration) IsEmptyImportDeclaration() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isEmptyImportDeclaration", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public com.github.javaparser.ast.expr.NameExpr getName()
func (jbobject *AstImportDeclaration) GetName() *AstExprNameExpr {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getName", "com/github/javaparser/ast/expr/NameExpr")
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

// public boolean isAsterisk()
func (jbobject *AstImportDeclaration) IsAsterisk() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAsterisk", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isStatic()
func (jbobject *AstImportDeclaration) IsStatic() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isStatic", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void setAsterisk(boolean)
func (jbobject *AstImportDeclaration) SetAsterisk(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setAsterisk", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public void setName(com.github.javaparser.ast.expr.NameExpr)
func (jbobject *AstImportDeclaration) SetName(a AstExprNameExprInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setName", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setStatic(boolean)
func (jbobject *AstImportDeclaration) SetStatic(a bool)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setStatic", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}


