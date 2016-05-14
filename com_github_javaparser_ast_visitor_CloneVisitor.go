package javaparser

import "github.com/timob/javabind"

type AstVisitorCloneVisitorInterface interface {
	JavaLangObjectInterface

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
	Visit(a AstCompilationUnitInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
	Visit2(a AstPackageDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
	Visit3(a AstImportDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
	Visit4(a AstTypeParameterInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
	Visit5(a AstCommentsLineCommentInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
	Visit6(a AstCommentsBlockCommentInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
	Visit7(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
	Visit8(a AstBodyEnumDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
	Visit9(a AstBodyEmptyTypeDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
	Visit10(a AstBodyEnumConstantDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
	Visit11(a AstBodyAnnotationDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
	Visit12(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
	Visit13(a AstBodyFieldDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
	Visit14(a AstBodyVariableDeclaratorInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
	Visit15(a AstBodyVariableDeclaratorIdInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
	Visit16(a AstBodyConstructorDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
	Visit17(a AstBodyMethodDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
	Visit18(a AstBodyParameterInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
	Visit19(a AstBodyMultiTypeParameterInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
	Visit20(a AstBodyEmptyMemberDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
	Visit21(a AstBodyInitializerDeclarationInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
	Visit22(a AstCommentsJavadocCommentInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
	Visit23(a AstTypeClassOrInterfaceTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
	Visit24(a AstTypePrimitiveTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
	Visit25(a AstTypeReferenceTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
	Visit26(a AstTypeIntersectionTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
	Visit27(a AstTypeUnionTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
	Visit28(a AstTypeVoidTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
	Visit29(a AstTypeWildcardTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
	Visit30(a AstTypeUnknownTypeInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
	Visit31(a AstExprArrayAccessExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
	Visit32(a AstExprArrayCreationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
	Visit33(a AstExprArrayInitializerExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
	Visit34(a AstExprAssignExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
	Visit35(a AstExprBinaryExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
	Visit36(a AstExprCastExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
	Visit37(a AstExprClassExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
	Visit38(a AstExprConditionalExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
	Visit39(a AstExprEnclosedExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
	Visit40(a AstExprFieldAccessExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
	Visit41(a AstExprInstanceOfExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
	Visit42(a AstExprStringLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
	Visit43(a AstExprIntegerLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
	Visit44(a AstExprLongLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
	Visit45(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
	Visit46(a AstExprLongLiteralMinValueExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
	Visit47(a AstExprCharLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
	Visit48(a AstExprDoubleLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
	Visit49(a AstExprBooleanLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
	Visit50(a AstExprNullLiteralExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
	Visit51(a AstExprMethodCallExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
	Visit52(a AstExprNameExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
	Visit53(a AstExprObjectCreationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
	Visit54(a AstExprQualifiedNameExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
	Visit55(a AstExprThisExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
	Visit56(a AstExprSuperExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
	Visit57(a AstExprUnaryExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
	Visit58(a AstExprVariableDeclarationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
	Visit59(a AstExprMarkerAnnotationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
	Visit60(a AstExprSingleMemberAnnotationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
	Visit61(a AstExprNormalAnnotationExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
	Visit62(a AstExprMemberValuePairInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
	Visit63(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
	Visit64(a AstStmtTypeDeclarationStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
	Visit65(a AstStmtAssertStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
	Visit66(a AstStmtBlockStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
	Visit67(a AstStmtLabeledStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
	Visit68(a AstStmtEmptyStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
	Visit69(a AstStmtExpressionStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
	Visit70(a AstStmtSwitchStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
	Visit71(a AstStmtSwitchEntryStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
	Visit72(a AstStmtBreakStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
	Visit73(a AstStmtReturnStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
	Visit74(a AstStmtIfStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
	Visit75(a AstStmtWhileStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
	Visit76(a AstStmtContinueStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
	Visit77(a AstStmtDoStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
	Visit78(a AstStmtForeachStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
	Visit79(a AstStmtForStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
	Visit80(a AstStmtThrowStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
	Visit81(a AstStmtSynchronizedStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
	Visit82(a AstStmtTryStmtInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
	Visit83(a AstStmtCatchClauseInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
	Visit84(a AstExprLambdaExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
	Visit85(a AstExprMethodReferenceExprInterface, b interface{}) *AstNode

	// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
	Visit86(a AstExprTypeExprInterface, b interface{}) *AstNode

	// public java.lang.Object visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
	Visit87(a AstExprTypeExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
	Visit88(a AstExprMethodReferenceExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
	Visit89(a AstExprLambdaExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
	Visit90(a AstStmtCatchClauseInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
	Visit91(a AstStmtTryStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
	Visit92(a AstStmtSynchronizedStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
	Visit93(a AstStmtThrowStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
	Visit94(a AstStmtForStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
	Visit95(a AstStmtForeachStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
	Visit96(a AstStmtDoStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
	Visit97(a AstStmtContinueStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
	Visit98(a AstStmtWhileStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
	Visit99(a AstStmtIfStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
	Visit100(a AstStmtReturnStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
	Visit101(a AstStmtBreakStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
	Visit102(a AstStmtSwitchEntryStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
	Visit103(a AstStmtSwitchStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
	Visit104(a AstStmtExpressionStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
	Visit105(a AstStmtEmptyStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
	Visit106(a AstStmtLabeledStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
	Visit107(a AstStmtBlockStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
	Visit108(a AstStmtAssertStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
	Visit109(a AstStmtTypeDeclarationStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
	Visit110(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
	Visit111(a AstExprMemberValuePairInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
	Visit112(a AstExprNormalAnnotationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
	Visit113(a AstExprSingleMemberAnnotationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
	Visit114(a AstExprMarkerAnnotationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
	Visit115(a AstExprVariableDeclarationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
	Visit116(a AstExprUnaryExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
	Visit117(a AstExprSuperExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
	Visit118(a AstExprThisExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
	Visit119(a AstExprQualifiedNameExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
	Visit120(a AstExprObjectCreationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
	Visit121(a AstExprNameExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
	Visit122(a AstExprMethodCallExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
	Visit123(a AstExprNullLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
	Visit124(a AstExprBooleanLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
	Visit125(a AstExprDoubleLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
	Visit126(a AstExprCharLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
	Visit127(a AstExprLongLiteralMinValueExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
	Visit128(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
	Visit129(a AstExprLongLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
	Visit130(a AstExprIntegerLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
	Visit131(a AstExprStringLiteralExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
	Visit132(a AstExprInstanceOfExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
	Visit133(a AstExprFieldAccessExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
	Visit134(a AstExprEnclosedExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
	Visit135(a AstExprConditionalExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
	Visit136(a AstExprClassExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
	Visit137(a AstExprCastExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
	Visit138(a AstExprBinaryExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
	Visit139(a AstExprAssignExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
	Visit140(a AstExprArrayInitializerExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
	Visit141(a AstExprArrayCreationExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
	Visit142(a AstExprArrayAccessExprInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
	Visit143(a AstTypeUnknownTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
	Visit144(a AstTypeWildcardTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
	Visit145(a AstTypeVoidTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
	Visit146(a AstTypeUnionTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
	Visit147(a AstTypeIntersectionTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
	Visit148(a AstTypeReferenceTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
	Visit149(a AstTypePrimitiveTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
	Visit150(a AstTypeClassOrInterfaceTypeInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
	Visit151(a AstCommentsJavadocCommentInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
	Visit152(a AstBodyInitializerDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
	Visit153(a AstBodyEmptyMemberDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
	Visit154(a AstBodyMultiTypeParameterInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
	Visit155(a AstBodyParameterInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
	Visit156(a AstBodyMethodDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
	Visit157(a AstBodyConstructorDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
	Visit158(a AstBodyVariableDeclaratorIdInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
	Visit159(a AstBodyVariableDeclaratorInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
	Visit160(a AstBodyFieldDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
	Visit161(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
	Visit162(a AstBodyAnnotationDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
	Visit163(a AstBodyEnumConstantDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
	Visit164(a AstBodyEmptyTypeDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
	Visit165(a AstBodyEnumDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
	Visit166(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
	Visit167(a AstCommentsBlockCommentInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
	Visit168(a AstCommentsLineCommentInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
	Visit169(a AstTypeParameterInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
	Visit170(a AstImportDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
	Visit171(a AstPackageDeclarationInterface, b interface{}) *JavaLangObject

	// public java.lang.Object visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
	Visit172(a AstCompilationUnitInterface, b interface{}) *JavaLangObject
}

type AstVisitorCloneVisitor struct {
	JavaLangObject
}

// public com.github.javaparser.ast.visitor.CloneVisitor()
func NewAstVisitorCloneVisitor() (*AstVisitorCloneVisitor) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/visitor/CloneVisitor")
	if err != nil {
		panic(err)
	}
	x := &AstVisitorCloneVisitor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit(a AstCompilationUnitInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/CompilationUnit"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit2(a AstPackageDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit3(a AstImportDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/ImportDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit4(a AstTypeParameterInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/TypeParameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit5(a AstCommentsLineCommentInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/comments/LineComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit6(a AstCommentsBlockCommentInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/comments/BlockComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit7(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit8(a AstBodyEnumDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit9(a AstBodyEmptyTypeDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyTypeDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit10(a AstBodyEnumConstantDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumConstantDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit11(a AstBodyAnnotationDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit12(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit13(a AstBodyFieldDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/FieldDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit14(a AstBodyVariableDeclaratorInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit15(a AstBodyVariableDeclaratorIdInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit16(a AstBodyConstructorDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/ConstructorDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit17(a AstBodyMethodDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/MethodDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit18(a AstBodyParameterInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit19(a AstBodyMultiTypeParameterInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/MultiTypeParameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit20(a AstBodyEmptyMemberDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit21(a AstBodyInitializerDeclarationInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/body/InitializerDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit22(a AstCommentsJavadocCommentInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit23(a AstTypeClassOrInterfaceTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit24(a AstTypePrimitiveTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/PrimitiveType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit25(a AstTypeReferenceTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit26(a AstTypeIntersectionTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/IntersectionType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit27(a AstTypeUnionTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit28(a AstTypeVoidTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/VoidType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit29(a AstTypeWildcardTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/WildcardType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit30(a AstTypeUnknownTypeInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/type/UnknownType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit31(a AstExprArrayAccessExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit32(a AstExprArrayCreationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit33(a AstExprArrayInitializerExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit34(a AstExprAssignExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/AssignExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit35(a AstExprBinaryExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/BinaryExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit36(a AstExprCastExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/CastExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit37(a AstExprClassExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ClassExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit38(a AstExprConditionalExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ConditionalExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit39(a AstExprEnclosedExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/EnclosedExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit40(a AstExprFieldAccessExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/FieldAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit41(a AstExprInstanceOfExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/InstanceOfExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit42(a AstExprStringLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/StringLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit43(a AstExprIntegerLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit44(a AstExprLongLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit45(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit46(a AstExprLongLiteralMinValueExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit47(a AstExprCharLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/CharLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit48(a AstExprDoubleLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/DoubleLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit49(a AstExprBooleanLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/BooleanLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit50(a AstExprNullLiteralExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/NullLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit51(a AstExprMethodCallExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodCallExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit52(a AstExprNameExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit53(a AstExprObjectCreationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ObjectCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit54(a AstExprQualifiedNameExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/QualifiedNameExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit55(a AstExprThisExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/ThisExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit56(a AstExprSuperExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/SuperExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit57(a AstExprUnaryExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/UnaryExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit58(a AstExprVariableDeclarationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit59(a AstExprMarkerAnnotationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/MarkerAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit60(a AstExprSingleMemberAnnotationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit61(a AstExprNormalAnnotationExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/NormalAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit62(a AstExprMemberValuePairInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/MemberValuePair"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit63(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit64(a AstStmtTypeDeclarationStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TypeDeclarationStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit65(a AstStmtAssertStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/AssertStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit66(a AstStmtBlockStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit67(a AstStmtLabeledStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/LabeledStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit68(a AstStmtEmptyStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/EmptyStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit69(a AstStmtExpressionStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExpressionStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit70(a AstStmtSwitchStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit71(a AstStmtSwitchEntryStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchEntryStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit72(a AstStmtBreakStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BreakStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit73(a AstStmtReturnStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ReturnStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit74(a AstStmtIfStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/IfStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit75(a AstStmtWhileStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/WhileStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit76(a AstStmtContinueStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ContinueStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit77(a AstStmtDoStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/DoStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit78(a AstStmtForeachStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForeachStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit79(a AstStmtForStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit80(a AstStmtThrowStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ThrowStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit81(a AstStmtSynchronizedStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SynchronizedStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit82(a AstStmtTryStmtInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TryStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit83(a AstStmtCatchClauseInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/stmt/CatchClause"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit84(a AstExprLambdaExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/LambdaExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit85(a AstExprMethodReferenceExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodReferenceExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit86(a AstExprTypeExprInterface, b interface{}) *AstNode {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "com/github/javaparser/ast/Node", conv_a.Value().Cast("com/github/javaparser/ast/expr/TypeExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit87(a AstExprTypeExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/TypeExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.MethodReferenceExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit88(a AstExprMethodReferenceExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodReferenceExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.LambdaExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit89(a AstExprLambdaExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/LambdaExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.CatchClause, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit90(a AstStmtCatchClauseInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/CatchClause"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.TryStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit91(a AstStmtTryStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TryStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.SynchronizedStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit92(a AstStmtSynchronizedStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SynchronizedStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ThrowStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit93(a AstStmtThrowStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ThrowStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ForStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit94(a AstStmtForStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ForeachStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit95(a AstStmtForeachStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForeachStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.DoStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit96(a AstStmtDoStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/DoStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ContinueStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit97(a AstStmtContinueStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ContinueStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.WhileStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit98(a AstStmtWhileStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/WhileStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.IfStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit99(a AstStmtIfStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/IfStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ReturnStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit100(a AstStmtReturnStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ReturnStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.BreakStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit101(a AstStmtBreakStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BreakStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit102(a AstStmtSwitchEntryStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchEntryStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.SwitchStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit103(a AstStmtSwitchStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ExpressionStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit104(a AstStmtExpressionStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExpressionStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.EmptyStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit105(a AstStmtEmptyStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/EmptyStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.LabeledStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit106(a AstStmtLabeledStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/LabeledStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.BlockStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit107(a AstStmtBlockStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.AssertStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit108(a AstStmtAssertStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/AssertStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit109(a AstStmtTypeDeclarationStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TypeDeclarationStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit110(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.MemberValuePair, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit111(a AstExprMemberValuePairInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/MemberValuePair"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit112(a AstExprNormalAnnotationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/NormalAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit113(a AstExprSingleMemberAnnotationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit114(a AstExprMarkerAnnotationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/MarkerAnnotationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit115(a AstExprVariableDeclarationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.UnaryExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit116(a AstExprUnaryExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/UnaryExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.SuperExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit117(a AstExprSuperExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/SuperExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ThisExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit118(a AstExprThisExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ThisExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.QualifiedNameExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit119(a AstExprQualifiedNameExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/QualifiedNameExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ObjectCreationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit120(a AstExprObjectCreationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ObjectCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.NameExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit121(a AstExprNameExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.MethodCallExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit122(a AstExprMethodCallExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodCallExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.NullLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit123(a AstExprNullLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/NullLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit124(a AstExprBooleanLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/BooleanLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit125(a AstExprDoubleLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/DoubleLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.CharLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit126(a AstExprCharLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/CharLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit127(a AstExprLongLiteralMinValueExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit128(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.LongLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit129(a AstExprLongLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit130(a AstExprIntegerLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.StringLiteralExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit131(a AstExprStringLiteralExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/StringLiteralExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.InstanceOfExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit132(a AstExprInstanceOfExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/InstanceOfExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.FieldAccessExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit133(a AstExprFieldAccessExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/FieldAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.EnclosedExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit134(a AstExprEnclosedExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/EnclosedExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ConditionalExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit135(a AstExprConditionalExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ConditionalExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ClassExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit136(a AstExprClassExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ClassExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.CastExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit137(a AstExprCastExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/CastExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.BinaryExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit138(a AstExprBinaryExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/BinaryExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.AssignExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit139(a AstExprAssignExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/AssignExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit140(a AstExprArrayInitializerExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayCreationExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit141(a AstExprArrayCreationExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayCreationExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.ArrayAccessExpr, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit142(a AstExprArrayAccessExprInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayAccessExpr"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.UnknownType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit143(a AstTypeUnknownTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/UnknownType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.WildcardType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit144(a AstTypeWildcardTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/WildcardType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.VoidType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit145(a AstTypeVoidTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/VoidType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.UnionType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit146(a AstTypeUnionTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.IntersectionType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit147(a AstTypeIntersectionTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/IntersectionType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.ReferenceType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit148(a AstTypeReferenceTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.PrimitiveType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit149(a AstTypePrimitiveTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/PrimitiveType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.type.ClassOrInterfaceType, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit150(a AstTypeClassOrInterfaceTypeInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.comments.JavadocComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit151(a AstCommentsJavadocCommentInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.InitializerDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit152(a AstBodyInitializerDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/InitializerDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit153(a AstBodyEmptyMemberDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.MultiTypeParameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit154(a AstBodyMultiTypeParameterInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/MultiTypeParameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.Parameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit155(a AstBodyParameterInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.MethodDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit156(a AstBodyMethodDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/MethodDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.ConstructorDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit157(a AstBodyConstructorDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/ConstructorDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.VariableDeclaratorId, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit158(a AstBodyVariableDeclaratorIdInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.VariableDeclarator, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit159(a AstBodyVariableDeclaratorInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.FieldDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit160(a AstBodyFieldDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/FieldDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit161(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationMemberDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.AnnotationDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit162(a AstBodyAnnotationDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.EnumConstantDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit163(a AstBodyEnumConstantDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumConstantDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit164(a AstBodyEmptyTypeDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyTypeDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.EnumDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit165(a AstBodyEnumDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit166(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.comments.BlockComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit167(a AstCommentsBlockCommentInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/comments/BlockComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.comments.LineComment, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit168(a AstCommentsLineCommentInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/comments/LineComment"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.TypeParameter, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit169(a AstTypeParameterInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/TypeParameter"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.ImportDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit170(a AstImportDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/ImportDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.PackageDeclaration, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit171(a AstPackageDeclarationInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.Object visit(com.github.javaparser.ast.CompilationUnit, java.lang.Object)
func (jbobject *AstVisitorCloneVisitor) Visit172(a AstCompilationUnitInterface, b interface{}) *JavaLangObject {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Object", conv_a.Value().Cast("com/github/javaparser/ast/CompilationUnit"), conv_b.Value().Cast("java/lang/Object"))
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}


