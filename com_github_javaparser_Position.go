package javaparser

import "github.com/timob/javabind"

type PositionInterface interface {
	JavaLangObjectInterface

	// public static com.github.javaparser.Position beginOf(com.github.javaparser.ast.Node)
	BeginOf(a AstNodeInterface) *Position

	// public static com.github.javaparser.Position endOf(com.github.javaparser.ast.Node)
	EndOf(a AstNodeInterface) *Position

	// public int getLine()
	GetLine() int

	// public int getColumn()
	GetColumn() int
}

type Position struct {
	JavaLangObject
}

// public com.github.javaparser.Position(int, int)
func NewPosition(a int, b int) (*Position) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/Position", a, b)
	if err != nil {
		panic(err)
	}
	x := &Position{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static com.github.javaparser.Position beginOf(com.github.javaparser.ast.Node)
func (jbobject *Position) BeginOf(a AstNodeInterface) *Position {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/Position", "beginOf", "com/github/javaparser/Position", conv_a.Value().Cast("com/github/javaparser/ast/Node"))
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
	unique_x := &Position{}
	unique_x.Callable = dst
	return unique_x
}

// public static com.github.javaparser.Position endOf(com.github.javaparser.ast.Node)
func (jbobject *Position) EndOf(a AstNodeInterface) *Position {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/Position", "endOf", "com/github/javaparser/Position", conv_a.Value().Cast("com/github/javaparser/ast/Node"))
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
	unique_x := &Position{}
	unique_x.Callable = dst
	return unique_x
}

// public int getLine()
func (jbobject *Position) GetLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public int getColumn()
func (jbobject *Position) GetColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *Position) ABSOLUTE_START() *Position {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/Position", "ABSOLUTE_START", "com/github/javaparser/Position")
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
	unique_x := &Position{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Position) SetFieldABSOLUTE_START(val PositionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/Position", "ABSOLUTE_START", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *Position) ABSOLUTE_END() *Position {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/Position", "ABSOLUTE_END", "com/github/javaparser/Position")
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
	unique_x := &Position{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *Position) SetFieldABSOLUTE_END(val PositionInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/Position", "ABSOLUTE_END", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


