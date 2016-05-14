package javaparser

import "github.com/timob/javabind"

type AstVisitorDumpVisitorInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getSource()
	GetSource() string

	// public void visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
	Visit(a AstCompilationUnitInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
	Visit2(a AstPackageDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
	Visit3(a AstExprNameExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
	Visit4(a AstExprQualifiedNameExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
	Visit5(a AstImportDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
	Visit6(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
	Visit7(a AstBodyEmptyTypeDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
	Visit8(a AstCommentsJavadocCommentInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
	Visit9(a AstTypeClassOrInterfaceTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
	Visit10(a AstTypeParameterInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
	Visit11(a AstTypePrimitiveTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
	Visit12(a AstTypeReferenceTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
	Visit13(a AstTypeIntersectionTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
	Visit14(a AstTypeUnionTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
	Visit15(a AstTypeWildcardTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
	Visit16(a AstTypeUnknownTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
	Visit17(a AstBodyFieldDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
	Visit18(a AstBodyVariableDeclaratorInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
	Visit19(a AstBodyVariableDeclaratorIdInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
	Visit20(a AstExprArrayInitializerExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
	Visit21(a AstTypeVoidTypeInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
	Visit22(a AstExprArrayAccessExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
	Visit23(a AstExprArrayCreationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
	Visit24(a AstExprAssignExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
	Visit25(a AstExprBinaryExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
	Visit26(a AstExprCastExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
	Visit27(a AstExprClassExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
	Visit28(a AstExprConditionalExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
	Visit29(a AstExprEnclosedExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
	Visit30(a AstExprFieldAccessExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
	Visit31(a AstExprInstanceOfExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
	Visit32(a AstExprCharLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
	Visit33(a AstExprDoubleLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
	Visit34(a AstExprIntegerLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
	Visit35(a AstExprLongLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
	Visit36(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
	Visit37(a AstExprLongLiteralMinValueExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
	Visit38(a AstExprStringLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
	Visit39(a AstExprBooleanLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
	Visit40(a AstExprNullLiteralExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
	Visit41(a AstExprThisExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
	Visit42(a AstExprSuperExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
	Visit43(a AstExprMethodCallExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
	Visit44(a AstExprObjectCreationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
	Visit45(a AstExprUnaryExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
	Visit46(a AstBodyConstructorDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
	Visit47(a AstBodyMethodDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
	Visit48(a AstBodyParameterInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
	Visit49(a AstBodyMultiTypeParameterInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
	Visit50(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
	Visit51(a AstExprVariableDeclarationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
	Visit52(a AstStmtTypeDeclarationStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
	Visit53(a AstStmtAssertStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
	Visit54(a AstStmtBlockStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
	Visit55(a AstStmtLabeledStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
	Visit56(a AstStmtEmptyStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
	Visit57(a AstStmtExpressionStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
	Visit58(a AstStmtSwitchStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
	Visit59(a AstStmtSwitchEntryStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
	Visit60(a AstStmtBreakStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
	Visit61(a AstStmtReturnStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
	Visit62(a AstBodyEnumDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
	Visit63(a AstBodyEnumConstantDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
	Visit64(a AstBodyEmptyMemberDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
	Visit65(a AstBodyInitializerDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
	Visit66(a AstStmtIfStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
	Visit67(a AstStmtWhileStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
	Visit68(a AstStmtContinueStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
	Visit69(a AstStmtDoStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
	Visit70(a AstStmtForeachStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
	Visit71(a AstStmtForStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
	Visit72(a AstStmtThrowStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
	Visit73(a AstStmtSynchronizedStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
	Visit74(a AstStmtTryStmtInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
	Visit75(a AstStmtCatchClauseInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
	Visit76(a AstBodyAnnotationDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
	Visit77(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
	Visit78(a AstExprMarkerAnnotationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
	Visit79(a AstExprSingleMemberAnnotationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
	Visit80(a AstExprNormalAnnotationExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
	Visit81(a AstExprMemberValuePairInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
	Visit82(a AstCommentsLineCommentInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
	Visit83(a AstCommentsBlockCommentInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
	Visit84(a AstExprLambdaExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
	Visit85(a AstExprMethodReferenceExprInterface, b interface{}) 

	// public void visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
	Visit86(a AstExprTypeExprInterface, b interface{}) 
}

type AstVisitorDumpVisitor struct {
	JavaLangObject
}

// public com.github.javaparser.ast.visitor.DumpVisitor()
func NewAstVisitorDumpVisitor() (*AstVisitorDumpVisitor) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/visitor/DumpVisitor")
	if err != nil {
		panic(err)
	}
	x := &AstVisitorDumpVisitor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.visitor.DumpVisitor(boolean)
func NewAstVisitorDumpVisitor2(a bool) (*AstVisitorDumpVisitor) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/visitor/DumpVisitor", a)
	if err != nil {
		panic(err)
	}
	x := &AstVisitorDumpVisitor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getSource()
func (jbobject *AstVisitorDumpVisitor) GetSource() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getSource", "java/lang/String")
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

// public void visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit(a AstCompilationUnitInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/CompilationUnit"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit2(a AstPackageDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit3(a AstExprNameExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit4(a AstExprQualifiedNameExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/QualifiedNameExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit5(a AstImportDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/ImportDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit6(a AstBodyClassOrInterfaceDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit7(a AstBodyEmptyTypeDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyTypeDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit8(a AstCommentsJavadocCommentInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit9(a AstTypeClassOrInterfaceTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit10(a AstTypeParameterInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/TypeParameter"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit11(a AstTypePrimitiveTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/PrimitiveType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit12(a AstTypeReferenceTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit13(a AstTypeIntersectionTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/IntersectionType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit14(a AstTypeUnionTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit15(a AstTypeWildcardTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/WildcardType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit16(a AstTypeUnknownTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/UnknownType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit17(a AstBodyFieldDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/FieldDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit18(a AstBodyVariableDeclaratorInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit19(a AstBodyVariableDeclaratorIdInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit20(a AstExprArrayInitializerExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit21(a AstTypeVoidTypeInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/type/VoidType"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit22(a AstExprArrayAccessExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit23(a AstExprArrayCreationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit24(a AstExprAssignExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/AssignExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit25(a AstExprBinaryExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/BinaryExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit26(a AstExprCastExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/CastExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit27(a AstExprClassExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ClassExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit28(a AstExprConditionalExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ConditionalExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit29(a AstExprEnclosedExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/EnclosedExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit30(a AstExprFieldAccessExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/FieldAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit31(a AstExprInstanceOfExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/InstanceOfExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit32(a AstExprCharLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/CharLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit33(a AstExprDoubleLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/DoubleLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit34(a AstExprIntegerLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit35(a AstExprLongLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit36(a AstExprIntegerLiteralMinValueExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit37(a AstExprLongLiteralMinValueExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit38(a AstExprStringLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/StringLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit39(a AstExprBooleanLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/BooleanLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit40(a AstExprNullLiteralExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NullLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit41(a AstExprThisExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ThisExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit42(a AstExprSuperExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/SuperExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit43(a AstExprMethodCallExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodCallExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit44(a AstExprObjectCreationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/ObjectCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit45(a AstExprUnaryExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/UnaryExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit46(a AstBodyConstructorDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/ConstructorDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit47(a AstBodyMethodDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/MethodDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit48(a AstBodyParameterInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit49(a AstBodyMultiTypeParameterInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/MultiTypeParameter"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit50(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit51(a AstExprVariableDeclarationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit52(a AstStmtTypeDeclarationStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/TypeDeclarationStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit53(a AstStmtAssertStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/AssertStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit54(a AstStmtBlockStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit55(a AstStmtLabeledStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/LabeledStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit56(a AstStmtEmptyStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/EmptyStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit57(a AstStmtExpressionStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExpressionStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit58(a AstStmtSwitchStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit59(a AstStmtSwitchEntryStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchEntryStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit60(a AstStmtBreakStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/BreakStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit61(a AstStmtReturnStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ReturnStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit62(a AstBodyEnumDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/EnumDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit63(a AstBodyEnumConstantDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/EnumConstantDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit64(a AstBodyEmptyMemberDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit65(a AstBodyInitializerDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/InitializerDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit66(a AstStmtIfStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/IfStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit67(a AstStmtWhileStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/WhileStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit68(a AstStmtContinueStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ContinueStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit69(a AstStmtDoStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/DoStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit70(a AstStmtForeachStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForeachStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit71(a AstStmtForStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit72(a AstStmtThrowStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/ThrowStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit73(a AstStmtSynchronizedStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/SynchronizedStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit74(a AstStmtTryStmtInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/TryStmt"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit75(a AstStmtCatchClauseInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/stmt/CatchClause"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit76(a AstBodyAnnotationDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit77(a AstBodyAnnotationMemberDeclarationInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit78(a AstExprMarkerAnnotationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/MarkerAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit79(a AstExprSingleMemberAnnotationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit80(a AstExprNormalAnnotationExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/NormalAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit81(a AstExprMemberValuePairInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/MemberValuePair"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit82(a AstCommentsLineCommentInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/LineComment"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit83(a AstCommentsBlockCommentInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/BlockComment"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit84(a AstExprLambdaExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/LambdaExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit85(a AstExprMethodReferenceExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodReferenceExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}

// public void visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
func (jbobject *AstVisitorDumpVisitor) Visit86(a AstExprTypeExprInterface, b interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/TypeExpr"), conv_b.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()

}


