package javaparser

import "github.com/timob/javabind"

type JavaParserInterface interface {
	JavaLangObjectInterface

	// public static boolean getDoNotConsiderAnnotationsAsNodeStartForCodeAttribution()
	GetDoNotConsiderAnnotationsAsNodeStartForCodeAttribution() bool

	// public static void setDoNotConsiderAnnotationsAsNodeStartForCodeAttribution(boolean)
	SetDoNotConsiderAnnotationsAsNodeStartForCodeAttribution(a bool) 

	// public static boolean getDoNotAssignCommentsPreceedingEmptyLines()
	GetDoNotAssignCommentsPreceedingEmptyLines() bool

	// public static void setDoNotAssignCommentsPreceedingEmptyLines(boolean)
	SetDoNotAssignCommentsPreceedingEmptyLines(a bool) 
}

type JavaParser struct {
	JavaLangObject
}

// public static boolean getDoNotConsiderAnnotationsAsNodeStartForCodeAttribution()
func (jbobject *JavaParser) GetDoNotConsiderAnnotationsAsNodeStartForCodeAttribution() bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "getDoNotConsiderAnnotationsAsNodeStartForCodeAttribution", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static void setDoNotConsiderAnnotationsAsNodeStartForCodeAttribution(boolean)
func (jbobject *JavaParser) SetDoNotConsiderAnnotationsAsNodeStartForCodeAttribution(a bool)  {
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "setDoNotConsiderAnnotationsAsNodeStartForCodeAttribution", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public static boolean getDoNotAssignCommentsPreceedingEmptyLines()
func (jbobject *JavaParser) GetDoNotAssignCommentsPreceedingEmptyLines() bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "getDoNotAssignCommentsPreceedingEmptyLines", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static void setDoNotAssignCommentsPreceedingEmptyLines(boolean)
func (jbobject *JavaParser) SetDoNotAssignCommentsPreceedingEmptyLines(a bool)  {
	_, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "setDoNotAssignCommentsPreceedingEmptyLines", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.InputStream, java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) Parse(a JavaIoInputStreamInterface, b string) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
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
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.InputStream, java.lang.String, boolean) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) Parse2(a JavaIoInputStreamInterface, b string, c bool) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/InputStream"), conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
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
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.InputStream) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) Parse3(a JavaIoInputStreamInterface) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/InputStream"))
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.File, java.lang.String) throws com.github.javaparser.ParseException, java.io.IOException
func (jbobject *JavaParser) Parse4(a JavaIoFileInterface, b string) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
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
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.File, java.lang.String, boolean) throws com.github.javaparser.ParseException, java.io.IOException
func (jbobject *JavaParser) Parse5(a JavaIoFileInterface, b string, c bool) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"), c)
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
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
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.File) throws com.github.javaparser.ParseException, java.io.IOException
func (jbobject *JavaParser) Parse6(a JavaIoFileInterface) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/File"))
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.CompilationUnit parse(java.io.Reader, boolean) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) Parse7(a JavaIoReaderInterface, b bool) (*AstCompilationUnit, error) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parse", "com/github/javaparser/ast/CompilationUnit", conv_a.Value().Cast("java/io/Reader"), b)
	if err != nil {
		var zero *AstCompilationUnit
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstCompilationUnit{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.stmt.BlockStmt parseBlock(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseBlock(a string) (*AstStmtBlockStmt, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseBlock", "com/github/javaparser/ast/stmt/BlockStmt", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstStmtBlockStmt
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstStmtBlockStmt{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.stmt.Statement parseStatement(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseStatement(a string) (*AstStmtStatement, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseStatement", "com/github/javaparser/ast/stmt/Statement", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstStmtStatement
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstStmtStatement{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.ImportDeclaration parseImport(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseImport(a string) (*AstImportDeclaration, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseImport", "com/github/javaparser/ast/ImportDeclaration", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstImportDeclaration
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstImportDeclaration{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.expr.Expression parseExpression(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseExpression(a string) (*AstExprExpression, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseExpression", "com/github/javaparser/ast/expr/Expression", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstExprExpression
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstExprExpression{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.expr.AnnotationExpr parseAnnotation(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseAnnotation(a string) (*AstExprAnnotationExpr, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseAnnotation", "com/github/javaparser/ast/expr/AnnotationExpr", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstExprAnnotationExpr
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstExprAnnotationExpr{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static com.github.javaparser.ast.body.BodyDeclaration parseBodyDeclaration(java.lang.String) throws com.github.javaparser.ParseException
func (jbobject *JavaParser) ParseBodyDeclaration(a string) (*AstBodyBodyDeclaration, error) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/JavaParser", "parseBodyDeclaration", "com/github/javaparser/ast/body/BodyDeclaration", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *AstBodyBodyDeclaration
		return zero, err
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &AstBodyBodyDeclaration{}
	unique_x.Callable = dst
	return unique_x, nil
}


