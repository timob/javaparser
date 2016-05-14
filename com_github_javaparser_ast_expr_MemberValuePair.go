package javaparser

import "github.com/timob/javabind"

type AstExprMemberValuePairInterface interface {
	AstNodeInterface

	// public java.lang.String getName()
	GetName() string

	// public com.github.javaparser.ast.expr.Expression getValue()
	GetValue() *AstExprExpression

	// public void setName(java.lang.String)
	SetName(a string) 

	// public void setValue(com.github.javaparser.ast.expr.Expression)
	SetValue(a AstExprExpressionInterface) 
}

type AstExprMemberValuePair struct {
	AstNode
}

// public com.github.javaparser.ast.expr.MemberValuePair()
func NewAstExprMemberValuePair() (*AstExprMemberValuePair) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MemberValuePair")
	if err != nil {
		panic(err)
	}
	x := &AstExprMemberValuePair{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MemberValuePair(java.lang.String, com.github.javaparser.ast.expr.Expression)
func NewAstExprMemberValuePair2(a string, b AstExprExpressionInterface) (*AstExprMemberValuePair) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MemberValuePair", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstExprMemberValuePair{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.expr.MemberValuePair(int, int, int, int, java.lang.String, com.github.javaparser.ast.expr.Expression)
func NewAstExprMemberValuePair3(a int, b int, c int, d int, e string, f AstExprExpressionInterface) (*AstExprMemberValuePair) {
	conv_e := javabind.NewGoToJavaString()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/expr/MemberValuePair", a, b, c, d, conv_e.Value().Cast("java/lang/String"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstExprMemberValuePair{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getName()
func (jbobject *AstExprMemberValuePair) GetName() string {
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

// public com.github.javaparser.ast.expr.Expression getValue()
func (jbobject *AstExprMemberValuePair) GetValue() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getValue", "com/github/javaparser/ast/expr/Expression")
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
	unique_x := &AstExprExpression{}
	unique_x.Callable = dst
	return unique_x
}

// public void setName(java.lang.String)
func (jbobject *AstExprMemberValuePair) SetName(a string)  {
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

// public void setValue(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstExprMemberValuePair) SetValue(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setValue", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


