package javaparser

import "github.com/timob/javabind"

type AstBodyVariableDeclaratorInterface interface {
	AstNodeInterface

	// public com.github.javaparser.ast.body.VariableDeclaratorId getId()
	GetId() *AstBodyVariableDeclaratorId

	// public com.github.javaparser.ast.expr.Expression getInit()
	GetInit() *AstExprExpression

	// public void setId(com.github.javaparser.ast.body.VariableDeclaratorId)
	SetId(a AstBodyVariableDeclaratorIdInterface) 

	// public void setInit(com.github.javaparser.ast.expr.Expression)
	SetInit(a AstExprExpressionInterface) 
}

type AstBodyVariableDeclarator struct {
	AstNode
}

// public com.github.javaparser.ast.body.VariableDeclarator()
func NewAstBodyVariableDeclarator() (*AstBodyVariableDeclarator) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclarator")
	if err != nil {
		panic(err)
	}
	x := &AstBodyVariableDeclarator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclarator(com.github.javaparser.ast.body.VariableDeclaratorId)
func NewAstBodyVariableDeclarator2(a AstBodyVariableDeclaratorIdInterface) (*AstBodyVariableDeclarator) {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclarator", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &AstBodyVariableDeclarator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclarator(com.github.javaparser.ast.body.VariableDeclaratorId, com.github.javaparser.ast.expr.Expression)
func NewAstBodyVariableDeclarator3(a AstBodyVariableDeclaratorIdInterface, b AstExprExpressionInterface) (*AstBodyVariableDeclarator) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclarator", conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_b.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &AstBodyVariableDeclarator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclarator(int, int, int, int, com.github.javaparser.ast.body.VariableDeclaratorId, com.github.javaparser.ast.expr.Expression)
func NewAstBodyVariableDeclarator4(a int, b int, c int, d int, e AstBodyVariableDeclaratorIdInterface, f AstExprExpressionInterface) (*AstBodyVariableDeclarator) {
	conv_e := javabind.NewGoToJavaCallable()
	conv_f := javabind.NewGoToJavaCallable()
	if err := conv_e.Convert(e); err != nil {
		panic(err)
	}
	if err := conv_f.Convert(f); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/body/VariableDeclarator", a, b, c, d, conv_e.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"), conv_f.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_e.CleanUp()
	conv_f.CleanUp()
	x := &AstBodyVariableDeclarator{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.body.VariableDeclaratorId getId()
func (jbobject *AstBodyVariableDeclarator) GetId() *AstBodyVariableDeclaratorId {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getId", "com/github/javaparser/ast/body/VariableDeclaratorId")
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
	unique_x := &AstBodyVariableDeclaratorId{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.expr.Expression getInit()
func (jbobject *AstBodyVariableDeclarator) GetInit() *AstExprExpression {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getInit", "com/github/javaparser/ast/expr/Expression")
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

// public void setId(com.github.javaparser.ast.body.VariableDeclaratorId)
func (jbobject *AstBodyVariableDeclarator) SetId(a AstBodyVariableDeclaratorIdInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setId", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/body/VariableDeclaratorId"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public void setInit(com.github.javaparser.ast.expr.Expression)
func (jbobject *AstBodyVariableDeclarator) SetInit(a AstExprExpressionInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setInit", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/expr/Expression"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


