package javaparser

import "github.com/timob/javabind"

type AstVisitorEqualsVisitorInterface interface {
	JavaLangObjectInterface

	// public static boolean equals(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node)
	Equals2(a AstNodeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.CompilationUnit, com.github.javaparser.ast.Node)
	Visit(a AstCompilationUnitInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.PackageDeclaration, com.github.javaparser.ast.Node)
	Visit2(a AstPackageDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.ImportDeclaration, com.github.javaparser.ast.Node)
	Visit3(a AstImportDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.TypeParameter, com.github.javaparser.ast.Node)
	Visit4(a AstTypeParameterInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.comments.LineComment, com.github.javaparser.ast.Node)
	Visit5(a AstCommentsLineCommentInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.comments.BlockComment, com.github.javaparser.ast.Node)
	Visit6(a AstCommentsBlockCommentInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, com.github.javaparser.ast.Node)
	Visit7(a AstBodyClassOrInterfaceDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.EnumDeclaration, com.github.javaparser.ast.Node)
	Visit8(a AstBodyEnumDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, com.github.javaparser.ast.Node)
	Visit9(a AstBodyEmptyTypeDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.EnumConstantDeclaration, com.github.javaparser.ast.Node)
	Visit10(a AstBodyEnumConstantDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.AnnotationDeclaration, com.github.javaparser.ast.Node)
	Visit11(a AstBodyAnnotationDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, com.github.javaparser.ast.Node)
	Visit12(a AstBodyAnnotationMemberDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.FieldDeclaration, com.github.javaparser.ast.Node)
	Visit13(a AstBodyFieldDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.VariableDeclarator, com.github.javaparser.ast.Node)
	Visit14(a AstBodyVariableDeclaratorInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.VariableDeclaratorId, com.github.javaparser.ast.Node)
	Visit15(a AstBodyVariableDeclaratorIdInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.ConstructorDeclaration, com.github.javaparser.ast.Node)
	Visit16(a AstBodyConstructorDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.MethodDeclaration, com.github.javaparser.ast.Node)
	Visit17(a AstBodyMethodDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.Parameter, com.github.javaparser.ast.Node)
	Visit18(a AstBodyParameterInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.MultiTypeParameter, com.github.javaparser.ast.Node)
	Visit19(a AstBodyMultiTypeParameterInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, com.github.javaparser.ast.Node)
	Visit20(a AstBodyEmptyMemberDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.body.InitializerDeclaration, com.github.javaparser.ast.Node)
	Visit21(a AstBodyInitializerDeclarationInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.comments.JavadocComment, com.github.javaparser.ast.Node)
	Visit22(a AstCommentsJavadocCommentInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.ClassOrInterfaceType, com.github.javaparser.ast.Node)
	Visit23(a AstTypeClassOrInterfaceTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.PrimitiveType, com.github.javaparser.ast.Node)
	Visit24(a AstTypePrimitiveTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.ReferenceType, com.github.javaparser.ast.Node)
	Visit25(a AstTypeReferenceTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.IntersectionType, com.github.javaparser.ast.Node)
	Visit26(a AstTypeIntersectionTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.UnionType, com.github.javaparser.ast.Node)
	Visit27(a AstTypeUnionTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.VoidType, com.github.javaparser.ast.Node)
	Visit28(a AstTypeVoidTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.WildcardType, com.github.javaparser.ast.Node)
	Visit29(a AstTypeWildcardTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.type.UnknownType, com.github.javaparser.ast.Node)
	Visit30(a AstTypeUnknownTypeInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayAccessExpr, com.github.javaparser.ast.Node)
	Visit31(a AstExprArrayAccessExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayCreationExpr, com.github.javaparser.ast.Node)
	Visit32(a AstExprArrayCreationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, com.github.javaparser.ast.Node)
	Visit33(a AstExprArrayInitializerExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.AssignExpr, com.github.javaparser.ast.Node)
	Visit34(a AstExprAssignExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.BinaryExpr, com.github.javaparser.ast.Node)
	Visit35(a AstExprBinaryExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.CastExpr, com.github.javaparser.ast.Node)
	Visit36(a AstExprCastExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ClassExpr, com.github.javaparser.ast.Node)
	Visit37(a AstExprClassExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ConditionalExpr, com.github.javaparser.ast.Node)
	Visit38(a AstExprConditionalExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.EnclosedExpr, com.github.javaparser.ast.Node)
	Visit39(a AstExprEnclosedExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.FieldAccessExpr, com.github.javaparser.ast.Node)
	Visit40(a AstExprFieldAccessExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.InstanceOfExpr, com.github.javaparser.ast.Node)
	Visit41(a AstExprInstanceOfExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.StringLiteralExpr, com.github.javaparser.ast.Node)
	Visit42(a AstExprStringLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, com.github.javaparser.ast.Node)
	Visit43(a AstExprIntegerLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LongLiteralExpr, com.github.javaparser.ast.Node)
	Visit44(a AstExprLongLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, com.github.javaparser.ast.Node)
	Visit45(a AstExprIntegerLiteralMinValueExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, com.github.javaparser.ast.Node)
	Visit46(a AstExprLongLiteralMinValueExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.CharLiteralExpr, com.github.javaparser.ast.Node)
	Visit47(a AstExprCharLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, com.github.javaparser.ast.Node)
	Visit48(a AstExprDoubleLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, com.github.javaparser.ast.Node)
	Visit49(a AstExprBooleanLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NullLiteralExpr, com.github.javaparser.ast.Node)
	Visit50(a AstExprNullLiteralExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MethodCallExpr, com.github.javaparser.ast.Node)
	Visit51(a AstExprMethodCallExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NameExpr, com.github.javaparser.ast.Node)
	Visit52(a AstExprNameExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ObjectCreationExpr, com.github.javaparser.ast.Node)
	Visit53(a AstExprObjectCreationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.QualifiedNameExpr, com.github.javaparser.ast.Node)
	Visit54(a AstExprQualifiedNameExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ThisExpr, com.github.javaparser.ast.Node)
	Visit55(a AstExprThisExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.SuperExpr, com.github.javaparser.ast.Node)
	Visit56(a AstExprSuperExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.UnaryExpr, com.github.javaparser.ast.Node)
	Visit57(a AstExprUnaryExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, com.github.javaparser.ast.Node)
	Visit58(a AstExprVariableDeclarationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, com.github.javaparser.ast.Node)
	Visit59(a AstExprMarkerAnnotationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, com.github.javaparser.ast.Node)
	Visit60(a AstExprSingleMemberAnnotationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, com.github.javaparser.ast.Node)
	Visit61(a AstExprNormalAnnotationExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MemberValuePair, com.github.javaparser.ast.Node)
	Visit62(a AstExprMemberValuePairInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, com.github.javaparser.ast.Node)
	Visit63(a AstStmtExplicitConstructorInvocationStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, com.github.javaparser.ast.Node)
	Visit64(a AstStmtTypeDeclarationStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.AssertStmt, com.github.javaparser.ast.Node)
	Visit65(a AstStmtAssertStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.Node)
	Visit66(a AstStmtBlockStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.LabeledStmt, com.github.javaparser.ast.Node)
	Visit67(a AstStmtLabeledStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.EmptyStmt, com.github.javaparser.ast.Node)
	Visit68(a AstStmtEmptyStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ExpressionStmt, com.github.javaparser.ast.Node)
	Visit69(a AstStmtExpressionStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SwitchStmt, com.github.javaparser.ast.Node)
	Visit70(a AstStmtSwitchStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, com.github.javaparser.ast.Node)
	Visit71(a AstStmtSwitchEntryStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.BreakStmt, com.github.javaparser.ast.Node)
	Visit72(a AstStmtBreakStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ReturnStmt, com.github.javaparser.ast.Node)
	Visit73(a AstStmtReturnStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.IfStmt, com.github.javaparser.ast.Node)
	Visit74(a AstStmtIfStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.WhileStmt, com.github.javaparser.ast.Node)
	Visit75(a AstStmtWhileStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ContinueStmt, com.github.javaparser.ast.Node)
	Visit76(a AstStmtContinueStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.DoStmt, com.github.javaparser.ast.Node)
	Visit77(a AstStmtDoStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ForeachStmt, com.github.javaparser.ast.Node)
	Visit78(a AstStmtForeachStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ForStmt, com.github.javaparser.ast.Node)
	Visit79(a AstStmtForStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ThrowStmt, com.github.javaparser.ast.Node)
	Visit80(a AstStmtThrowStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SynchronizedStmt, com.github.javaparser.ast.Node)
	Visit81(a AstStmtSynchronizedStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.TryStmt, com.github.javaparser.ast.Node)
	Visit82(a AstStmtTryStmtInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.CatchClause, com.github.javaparser.ast.Node)
	Visit83(a AstStmtCatchClauseInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LambdaExpr, com.github.javaparser.ast.Node)
	Visit84(a AstExprLambdaExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MethodReferenceExpr, com.github.javaparser.ast.Node)
	Visit85(a AstExprMethodReferenceExprInterface, b AstNodeInterface) bool

	// public java.lang.Boolean visit(com.github.javaparser.ast.expr.TypeExpr, com.github.javaparser.ast.Node)
	Visit86(a AstExprTypeExprInterface, b AstNodeInterface) bool

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

type AstVisitorEqualsVisitor struct {
	JavaLangObject
}

// public static boolean equals(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Equals2(a AstNodeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/visitor/EqualsVisitor", "equals", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/Node"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}

// public java.lang.Boolean visit(com.github.javaparser.ast.CompilationUnit, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit(a AstCompilationUnitInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/CompilationUnit"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.PackageDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit2(a AstPackageDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/PackageDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.ImportDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit3(a AstImportDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/ImportDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.TypeParameter, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit4(a AstTypeParameterInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/TypeParameter"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.comments.LineComment, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit5(a AstCommentsLineCommentInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/comments/LineComment"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.comments.BlockComment, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit6(a AstCommentsBlockCommentInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/comments/BlockComment"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.ClassOrInterfaceDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit7(a AstBodyClassOrInterfaceDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/ClassOrInterfaceDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.EnumDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit8(a AstBodyEnumDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.EmptyTypeDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit9(a AstBodyEmptyTypeDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyTypeDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.EnumConstantDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit10(a AstBodyEnumConstantDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/EnumConstantDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.AnnotationDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit11(a AstBodyAnnotationDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.AnnotationMemberDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit12(a AstBodyAnnotationMemberDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/AnnotationMemberDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.FieldDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit13(a AstBodyFieldDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/FieldDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.VariableDeclarator, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit14(a AstBodyVariableDeclaratorInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclarator"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.VariableDeclaratorId, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit15(a AstBodyVariableDeclaratorIdInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.ConstructorDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit16(a AstBodyConstructorDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/ConstructorDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.MethodDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit17(a AstBodyMethodDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/MethodDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.Parameter, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit18(a AstBodyParameterInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/Parameter"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.MultiTypeParameter, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit19(a AstBodyMultiTypeParameterInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/MultiTypeParameter"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.EmptyMemberDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit20(a AstBodyEmptyMemberDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/EmptyMemberDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.body.InitializerDeclaration, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit21(a AstBodyInitializerDeclarationInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/body/InitializerDeclaration"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.comments.JavadocComment, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit22(a AstCommentsJavadocCommentInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/comments/JavadocComment"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.ClassOrInterfaceType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit23(a AstTypeClassOrInterfaceTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/ClassOrInterfaceType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.PrimitiveType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit24(a AstTypePrimitiveTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/PrimitiveType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.ReferenceType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit25(a AstTypeReferenceTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/ReferenceType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.IntersectionType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit26(a AstTypeIntersectionTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/IntersectionType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.UnionType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit27(a AstTypeUnionTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/UnionType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.VoidType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit28(a AstTypeVoidTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/VoidType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.WildcardType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit29(a AstTypeWildcardTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/WildcardType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.type.UnknownType, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit30(a AstTypeUnknownTypeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/type/UnknownType"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayAccessExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit31(a AstExprArrayAccessExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayAccessExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayCreationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit32(a AstExprArrayCreationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayCreationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ArrayInitializerExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit33(a AstExprArrayInitializerExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ArrayInitializerExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.AssignExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit34(a AstExprAssignExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/AssignExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.BinaryExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit35(a AstExprBinaryExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/BinaryExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.CastExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit36(a AstExprCastExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/CastExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ClassExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit37(a AstExprClassExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ClassExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ConditionalExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit38(a AstExprConditionalExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ConditionalExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.EnclosedExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit39(a AstExprEnclosedExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/EnclosedExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.FieldAccessExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit40(a AstExprFieldAccessExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/FieldAccessExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.InstanceOfExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit41(a AstExprInstanceOfExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/InstanceOfExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.StringLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit42(a AstExprStringLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/StringLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.IntegerLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit43(a AstExprIntegerLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LongLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit44(a AstExprLongLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.IntegerLiteralMinValueExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit45(a AstExprIntegerLiteralMinValueExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/IntegerLiteralMinValueExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LongLiteralMinValueExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit46(a AstExprLongLiteralMinValueExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/LongLiteralMinValueExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.CharLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit47(a AstExprCharLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/CharLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.DoubleLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit48(a AstExprDoubleLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/DoubleLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.BooleanLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit49(a AstExprBooleanLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/BooleanLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NullLiteralExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit50(a AstExprNullLiteralExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/NullLiteralExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MethodCallExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit51(a AstExprMethodCallExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodCallExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NameExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit52(a AstExprNameExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/NameExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ObjectCreationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit53(a AstExprObjectCreationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ObjectCreationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.QualifiedNameExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit54(a AstExprQualifiedNameExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/QualifiedNameExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.ThisExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit55(a AstExprThisExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/ThisExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.SuperExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit56(a AstExprSuperExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/SuperExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.UnaryExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit57(a AstExprUnaryExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/UnaryExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.VariableDeclarationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit58(a AstExprVariableDeclarationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/VariableDeclarationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MarkerAnnotationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit59(a AstExprMarkerAnnotationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/MarkerAnnotationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.SingleMemberAnnotationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit60(a AstExprSingleMemberAnnotationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/SingleMemberAnnotationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.NormalAnnotationExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit61(a AstExprNormalAnnotationExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/NormalAnnotationExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MemberValuePair, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit62(a AstExprMemberValuePairInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/MemberValuePair"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ExplicitConstructorInvocationStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit63(a AstStmtExplicitConstructorInvocationStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExplicitConstructorInvocationStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.TypeDeclarationStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit64(a AstStmtTypeDeclarationStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TypeDeclarationStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.AssertStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit65(a AstStmtAssertStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/AssertStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.BlockStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit66(a AstStmtBlockStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BlockStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.LabeledStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit67(a AstStmtLabeledStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/LabeledStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.EmptyStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit68(a AstStmtEmptyStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/EmptyStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ExpressionStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit69(a AstStmtExpressionStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ExpressionStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SwitchStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit70(a AstStmtSwitchStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SwitchEntryStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit71(a AstStmtSwitchEntryStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SwitchEntryStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.BreakStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit72(a AstStmtBreakStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/BreakStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ReturnStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit73(a AstStmtReturnStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ReturnStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.IfStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit74(a AstStmtIfStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/IfStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.WhileStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit75(a AstStmtWhileStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/WhileStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ContinueStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit76(a AstStmtContinueStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ContinueStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.DoStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit77(a AstStmtDoStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/DoStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ForeachStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit78(a AstStmtForeachStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForeachStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ForStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit79(a AstStmtForStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ForStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.ThrowStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit80(a AstStmtThrowStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/ThrowStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.SynchronizedStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit81(a AstStmtSynchronizedStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/SynchronizedStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.TryStmt, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit82(a AstStmtTryStmtInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/TryStmt"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.stmt.CatchClause, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit83(a AstStmtCatchClauseInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/stmt/CatchClause"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.LambdaExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit84(a AstExprLambdaExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/LambdaExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.MethodReferenceExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit85(a AstExprMethodReferenceExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/MethodReferenceExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Boolean visit(com.github.javaparser.ast.expr.TypeExpr, com.github.javaparser.ast.Node)
func (jbobject *AstVisitorEqualsVisitor) Visit86(a AstExprTypeExprInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "visit", "java/lang/Boolean", conv_a.Value().Cast("com/github/javaparser/ast/expr/TypeExpr"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoBoolean()
	dst := new(bool)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Object visit(com.github.javaparser.ast.expr.TypeExpr, java.lang.Object)
func (jbobject *AstVisitorEqualsVisitor) Visit87(a AstExprTypeExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit88(a AstExprMethodReferenceExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit89(a AstExprLambdaExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit90(a AstStmtCatchClauseInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit91(a AstStmtTryStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit92(a AstStmtSynchronizedStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit93(a AstStmtThrowStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit94(a AstStmtForStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit95(a AstStmtForeachStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit96(a AstStmtDoStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit97(a AstStmtContinueStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit98(a AstStmtWhileStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit99(a AstStmtIfStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit100(a AstStmtReturnStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit101(a AstStmtBreakStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit102(a AstStmtSwitchEntryStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit103(a AstStmtSwitchStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit104(a AstStmtExpressionStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit105(a AstStmtEmptyStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit106(a AstStmtLabeledStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit107(a AstStmtBlockStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit108(a AstStmtAssertStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit109(a AstStmtTypeDeclarationStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit110(a AstStmtExplicitConstructorInvocationStmtInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit111(a AstExprMemberValuePairInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit112(a AstExprNormalAnnotationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit113(a AstExprSingleMemberAnnotationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit114(a AstExprMarkerAnnotationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit115(a AstExprVariableDeclarationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit116(a AstExprUnaryExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit117(a AstExprSuperExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit118(a AstExprThisExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit119(a AstExprQualifiedNameExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit120(a AstExprObjectCreationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit121(a AstExprNameExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit122(a AstExprMethodCallExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit123(a AstExprNullLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit124(a AstExprBooleanLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit125(a AstExprDoubleLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit126(a AstExprCharLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit127(a AstExprLongLiteralMinValueExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit128(a AstExprIntegerLiteralMinValueExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit129(a AstExprLongLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit130(a AstExprIntegerLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit131(a AstExprStringLiteralExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit132(a AstExprInstanceOfExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit133(a AstExprFieldAccessExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit134(a AstExprEnclosedExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit135(a AstExprConditionalExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit136(a AstExprClassExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit137(a AstExprCastExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit138(a AstExprBinaryExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit139(a AstExprAssignExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit140(a AstExprArrayInitializerExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit141(a AstExprArrayCreationExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit142(a AstExprArrayAccessExprInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit143(a AstTypeUnknownTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit144(a AstTypeWildcardTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit145(a AstTypeVoidTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit146(a AstTypeUnionTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit147(a AstTypeIntersectionTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit148(a AstTypeReferenceTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit149(a AstTypePrimitiveTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit150(a AstTypeClassOrInterfaceTypeInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit151(a AstCommentsJavadocCommentInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit152(a AstBodyInitializerDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit153(a AstBodyEmptyMemberDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit154(a AstBodyMultiTypeParameterInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit155(a AstBodyParameterInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit156(a AstBodyMethodDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit157(a AstBodyConstructorDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit158(a AstBodyVariableDeclaratorIdInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit159(a AstBodyVariableDeclaratorInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit160(a AstBodyFieldDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit161(a AstBodyAnnotationMemberDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit162(a AstBodyAnnotationDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit163(a AstBodyEnumConstantDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit164(a AstBodyEmptyTypeDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit165(a AstBodyEnumDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit166(a AstBodyClassOrInterfaceDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit167(a AstCommentsBlockCommentInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit168(a AstCommentsLineCommentInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit169(a AstTypeParameterInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit170(a AstImportDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit171(a AstPackageDeclarationInterface, b interface{}) *JavaLangObject {
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
func (jbobject *AstVisitorEqualsVisitor) Visit172(a AstCompilationUnitInterface, b interface{}) *JavaLangObject {
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


