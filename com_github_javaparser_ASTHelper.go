package javaparser

import "github.com/timob/javabind"

type ASTHelperInterface interface {
	JavaLangObjectInterface

	// public static com.github.javaparser.ast.expr.NameExpr createNameExpr(java.lang.String)
	CreateNameExpr(a string) *AstExprNameExpr

	// public static com.github.javaparser.ast.body.Parameter createParameter(com.github.javaparser.ast.type.Type, java.lang.String)
	CreateParameter(a AstTypeTypeInterface, b string) *AstBodyParameter

	// public static com.github.javaparser.ast.body.FieldDeclaration createFieldDeclaration(int, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclarator)
	CreateFieldDeclaration(a int, b AstTypeTypeInterface, c AstBodyVariableDeclaratorInterface) *AstBodyFieldDeclaration

	// public static com.github.javaparser.ast.body.FieldDeclaration createFieldDeclaration(int, com.github.javaparser.ast.type.Type, java.lang.String)
	CreateFieldDeclaration2(a int, b AstTypeTypeInterface, c string) *AstBodyFieldDeclaration

	// public static com.github.javaparser.ast.expr.VariableDeclarationExpr createVariableDeclarationExpr(com.github.javaparser.ast.type.Type, java.lang.String)
	CreateVariableDeclarationExpr(a AstTypeTypeInterface, b string) *AstExprVariableDeclarationExpr

	// public static void addParameter(com.github.javaparser.ast.body.MethodDeclaration, com.github.javaparser.ast.body.Parameter)
	AddParameter(a AstBodyMethodDeclarationInterface, b AstBodyParameterInterface) 

	// public static void addArgument(com.github.javaparser.ast.expr.MethodCallExpr, com.github.javaparser.ast.expr.Expression)
	AddArgument(a AstExprMethodCallExprInterface, b AstExprExpressionInterface) 

	// public static void addTypeDeclaration(com.github.javaparser.ast.CompilationUnit, com.github.javaparser.ast.body.TypeDeclaration)
	AddTypeDeclaration(a AstCompilationUnitInterface, b AstBodyTypeDeclarationInterface) 

	// public static com.github.javaparser.ast.type.ReferenceType createReferenceType(java.lang.String, int)
	CreateReferenceType(a string, b int) *AstTypeReferenceType

	// public static com.github.javaparser.ast.type.ReferenceType createReferenceType(com.github.javaparser.ast.type.PrimitiveType, int)
	CreateReferenceType2(a AstTypePrimitiveTypeInterface, b int) *AstTypeReferenceType

	// public static void addStmt(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.stmt.Statement)
	AddStmt(a AstStmtBlockStmtInterface, b AstStmtStatementInterface) 

	// public static void addStmt(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.expr.Expression)
	AddStmt2(a AstStmtBlockStmtInterface, b AstExprExpressionInterface) 

	// public static void addMember(com.github.javaparser.ast.body.TypeDeclaration, com.github.javaparser.ast.body.BodyDeclaration)
	AddMember(a AstBodyTypeDeclarationInterface, b AstBodyBodyDeclarationInterface) 
}

type ASTHelper struct {
	JavaLangObject
}

// public static com.github.javaparser.ast.expr.NameExpr createNameExpr(java.lang.String)
func (jbobject *ASTHelper) CreateNameExpr(a string) *AstExprNameExpr {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createNameExpr", "com/github/javaparser/ast/expr/NameExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
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

// public static com.github.javaparser.ast.body.Parameter createParameter(com.github.javaparser.ast.type.Type, java.lang.String)
func (jbobject *ASTHelper) CreateParameter(a AstTypeTypeInterface, b string) *AstBodyParameter {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createParameter", "com/github/javaparser/ast/body/Parameter", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstBodyParameter{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.body.FieldDeclaration createFieldDeclaration(int, com.github.javaparser.ast.type.Type, com.github.javaparser.ast.body.VariableDeclarator)
func (jbobject *ASTHelper) CreateFieldDeclaration(a int, b AstTypeTypeInterface, c AstBodyVariableDeclaratorInterface) *AstBodyFieldDeclaration {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createFieldDeclaration", "com/github/javaparser/ast/body/FieldDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstBodyFieldDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.body.FieldDeclaration createFieldDeclaration(int, com.github.javaparser.ast.type.Type, java.lang.String)
func (jbobject *ASTHelper) CreateFieldDeclaration2(a int, b AstTypeTypeInterface, c string) *AstBodyFieldDeclaration {
	conv_b := javabind.NewGoToJavaCallable()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createFieldDeclaration", "com/github/javaparser/ast/body/FieldDeclaration", a, conv_b.Value().Cast("com/github/javaparser/ast/type/Type"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstBodyFieldDeclaration{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.expr.VariableDeclarationExpr createVariableDeclarationExpr(com.github.javaparser.ast.type.Type, java.lang.String)
func (jbobject *ASTHelper) CreateVariableDeclarationExpr(a AstTypeTypeInterface, b string) *AstExprVariableDeclarationExpr {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createVariableDeclarationExpr", "com/github/javaparser/ast/expr/VariableDeclarationExpr", conv_a.Value().Cast("com/github/javaparser/ast/type/Type"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstExprVariableDeclarationExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public static void addParameter(com.github.javaparser.ast.body.MethodDeclaration, com.github.javaparser.ast.body.Parameter)
func (jbobject *ASTHelper) AddParameter(a AstBodyMethodDeclarationInterface, b AstBodyParameterInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addParameter", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/MethodDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/body/Parameter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static void addArgument(com.github.javaparser.ast.expr.MethodCallExpr, com.github.javaparser.ast.expr.Expression)
func (jbobject *ASTHelper) AddArgument(a AstExprMethodCallExprInterface, b AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addArgument", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodCallExpr"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static void addTypeDeclaration(com.github.javaparser.ast.CompilationUnit, com.github.javaparser.ast.body.TypeDeclaration)
func (jbobject *ASTHelper) AddTypeDeclaration(a AstCompilationUnitInterface, b AstBodyTypeDeclarationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addTypeDeclaration", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/CompilationUnit"), conv_b.Value().Cast("com/github/javaparser/ast/body/TypeDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static com.github.javaparser.ast.type.ReferenceType createReferenceType(java.lang.String, int)
func (jbobject *ASTHelper) CreateReferenceType(a string, b int) *AstTypeReferenceType {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createReferenceType", "com/github/javaparser/ast/type/ReferenceType", conv_a.Value().Cast("java/lang/String"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstTypeReferenceType{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.ast.type.ReferenceType createReferenceType(com.github.javaparser.ast.type.PrimitiveType, int)
func (jbobject *ASTHelper) CreateReferenceType2(a AstTypePrimitiveTypeInterface, b int) *AstTypeReferenceType {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "createReferenceType", "com/github/javaparser/ast/type/ReferenceType", conv_a.Value().Cast("com/github/javaparser/ast/type/PrimitiveType"), b)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstTypeReferenceType{}
	unique_x.Callable = dst
	return unique_x
}

// public static void addStmt(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.stmt.Statement)
func (jbobject *ASTHelper) AddStmt(a AstStmtBlockStmtInterface, b AstStmtStatementInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addStmt", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("com/github/javaparser/ast/stmt/Statement"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static void addStmt(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.expr.Expression)
func (jbobject *ASTHelper) AddStmt2(a AstStmtBlockStmtInterface, b AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addStmt", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public static void addMember(com.github.javaparser.ast.body.TypeDeclaration, com.github.javaparser.ast.body.BodyDeclaration)
func (jbobject *ASTHelper) AddMember(a AstBodyTypeDeclarationInterface, b AstBodyBodyDeclarationInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ASTHelper", "addMember", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/TypeDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/body/BodyDeclaration"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

func (jbobject *ASTHelper) BYTE_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "BYTE_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldBYTE_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "BYTE_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) SHORT_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "SHORT_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldSHORT_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "SHORT_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) INT_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "INT_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldINT_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "INT_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) LONG_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "LONG_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldLONG_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "LONG_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) FLOAT_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "FLOAT_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldFLOAT_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "FLOAT_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) DOUBLE_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "DOUBLE_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldDOUBLE_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "DOUBLE_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) BOOLEAN_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "BOOLEAN_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldBOOLEAN_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "BOOLEAN_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) CHAR_TYPE() *AstTypePrimitiveType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "CHAR_TYPE", "com/github/javaparser/ast/type/PrimitiveType")
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
	unique_x := &AstTypePrimitiveType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldCHAR_TYPE(val AstTypePrimitiveTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "CHAR_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *ASTHelper) VOID_TYPE() *AstTypeVoidType {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTHelper", "VOID_TYPE", "com/github/javaparser/ast/type/VoidType")
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
	unique_x := &AstTypeVoidType{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *ASTHelper) SetFieldVOID_TYPE(val AstTypeVoidTypeInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTHelper", "VOID_TYPE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


