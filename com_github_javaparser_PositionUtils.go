package javaparser

import "github.com/timob/javabind"

type PositionUtilsInterface interface {
	JavaLangObjectInterface

	// public static boolean areInOrder(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node)
	AreInOrder(a AstNodeInterface, b AstNodeInterface) bool

	// public static boolean areInOrder(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node, boolean)
	AreInOrder2(a AstNodeInterface, b AstNodeInterface, c bool) bool

	// public static com.github.javaparser.ast.expr.AnnotationExpr getLastAnnotation(com.github.javaparser.ast.Node)
	GetLastAnnotation(a AstNodeInterface) *AstExprAnnotationExpr

	// public static boolean nodeContains(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node, boolean)
	NodeContains(a AstNodeInterface, b AstNodeInterface, c bool) bool
}

type PositionUtils struct {
	JavaLangObject
}

// public static boolean areInOrder(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node)
func (jbobject *PositionUtils) AreInOrder(a AstNodeInterface, b AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/PositionUtils", "areInOrder", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/Node"), conv_b.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}

// public static boolean areInOrder(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node, boolean)
func (jbobject *PositionUtils) AreInOrder2(a AstNodeInterface, b AstNodeInterface, c bool) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/PositionUtils", "areInOrder", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/Node"), conv_b.Value().Cast("com/github/javaparser/ast/Node"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}

// public static com.github.javaparser.ast.expr.AnnotationExpr getLastAnnotation(com.github.javaparser.ast.Node)
func (jbobject *PositionUtils) GetLastAnnotation(a AstNodeInterface) *AstExprAnnotationExpr {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/PositionUtils", "getLastAnnotation", "com/github/javaparser/ast/expr/AnnotationExpr", conv_a.Value().Cast("com/github/javaparser/ast/Node"))
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
	unique_x := &AstExprAnnotationExpr{}
	unique_x.Callable = dst
	return unique_x
}

// public static boolean nodeContains(com.github.javaparser.ast.Node, com.github.javaparser.ast.Node, boolean)
func (jbobject *PositionUtils) NodeContains(a AstNodeInterface, b AstNodeInterface, c bool) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/PositionUtils", "nodeContains", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/Node"), conv_b.Value().Cast("com/github/javaparser/ast/Node"), c)
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}


