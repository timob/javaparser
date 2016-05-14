package javaparser

import "github.com/timob/javabind"

type AstTreeVisitorInterface interface {
	JavaLangObjectInterface

	// public void visitDepthFirst(com.github.javaparser.ast.Node)
	VisitDepthFirst(a AstNodeInterface) 

	// public abstract void process(com.github.javaparser.ast.Node)
	Process(a AstNodeInterface) 
}

type AstTreeVisitor struct {
	JavaLangObject
}

// public com.github.javaparser.ast.TreeVisitor()
func NewAstTreeVisitor() (*AstTreeVisitor) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/TreeVisitor")
	if err != nil {
		panic(err)
	}
	x := &AstTreeVisitor{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public void visitDepthFirst(com.github.javaparser.ast.Node)
func (jbobject *AstTreeVisitor) VisitDepthFirst(a AstNodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "visitDepthFirst", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public abstract void process(com.github.javaparser.ast.Node)
func (jbobject *AstTreeVisitor) Process(a AstNodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "process", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}


