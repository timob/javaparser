package javaparser

import "github.com/timob/javabind"

type AstNodeInterface interface {
	JavaLangObjectInterface

	// public final int getBeginColumn()
	GetBeginColumn() int

	// public final int getBeginLine()
	GetBeginLine() int

	// public final com.github.javaparser.ast.comments.Comment getComment()
	GetComment() *AstCommentsComment

	// public final java.lang.Object getData()
	GetData() *JavaLangObject

	// public final int getEndColumn()
	GetEndColumn() int

	// public final int getEndLine()
	GetEndLine() int

	// public final void setBeginColumn(int)
	SetBeginColumn(a int) 

	// public final void setBeginLine(int)
	SetBeginLine(a int) 

	// public final void setComment(com.github.javaparser.ast.comments.Comment)
	SetComment(a AstCommentsCommentInterface) 

	// public final void setData(java.lang.Object)
	SetData(a interface{}) 

	// public final void setEndColumn(int)
	SetEndColumn(a int) 

	// public final void setEndLine(int)
	SetEndLine(a int) 

	// public final java.lang.String toStringWithoutComments()
	ToStringWithoutComments() string

	// public com.github.javaparser.ast.Node clone()
	Clone() *AstNode

	// public com.github.javaparser.ast.Node getParentNode()
	GetParentNode() *AstNode

	// public java.util.List<com.github.javaparser.ast.Node> getChildrenNodes()
	GetChildrenNodes() []*AstNode

	// public boolean contains(com.github.javaparser.ast.Node)
	Contains(a AstNodeInterface) bool

	// public void addOrphanComment(com.github.javaparser.ast.comments.Comment)
	AddOrphanComment(a AstCommentsCommentInterface) 

	// public java.util.List<com.github.javaparser.ast.comments.Comment> getOrphanComments()
	GetOrphanComments() []*AstCommentsComment

	// public java.util.List<com.github.javaparser.ast.comments.Comment> getAllContainedComments()
	GetAllContainedComments() []*AstCommentsComment

	// public void setParentNode(com.github.javaparser.ast.Node)
	SetParentNode(a AstNodeInterface) 

	// public boolean isPositionedAfter(int, int)
	IsPositionedAfter(a int, b int) bool

	// public boolean isPositionedBefore(int, int)
	IsPositionedBefore(a int, b int) bool

	// public boolean hasComment()
	HasComment() bool
}

type AstNode struct {
	JavaLangObject
}

// public com.github.javaparser.ast.Node()
func NewAstNode() (*AstNode) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/Node")
	if err != nil {
		panic(err)
	}
	x := &AstNode{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public com.github.javaparser.ast.Node(int, int, int, int)
func NewAstNode2(a int, b int, c int, d int) (*AstNode) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/ast/Node", a, b, c, d)
	if err != nil {
		panic(err)
	}
	x := &AstNode{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public final int getBeginColumn()
func (jbobject *AstNode) GetBeginColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBeginColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final int getBeginLine()
func (jbobject *AstNode) GetBeginLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getBeginLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final com.github.javaparser.ast.comments.Comment getComment()
func (jbobject *AstNode) GetComment() *AstCommentsComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getComment", "com/github/javaparser/ast/comments/Comment")
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
	unique_x := &AstCommentsComment{}
	unique_x.Callable = dst
	return unique_x
}

// public final java.lang.Object getData()
func (jbobject *AstNode) GetData() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getData", "java/lang/Object")
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
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x
}

// public final int getEndColumn()
func (jbobject *AstNode) GetEndColumn() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndColumn", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final int getEndLine()
func (jbobject *AstNode) GetEndLine() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getEndLine", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public final void setBeginColumn(int)
func (jbobject *AstNode) SetBeginColumn(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBeginColumn", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final void setBeginLine(int)
func (jbobject *AstNode) SetBeginLine(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setBeginLine", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final void setComment(com.github.javaparser.ast.comments.Comment)
func (jbobject *AstNode) SetComment(a AstCommentsCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setComment", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/Comment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final void setData(java.lang.Object)
func (jbobject *AstNode) SetData(a interface{})  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setData", javabind.Void, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public final void setEndColumn(int)
func (jbobject *AstNode) SetEndColumn(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndColumn", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final void setEndLine(int)
func (jbobject *AstNode) SetEndLine(a int)  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setEndLine", javabind.Void, a)
	if err != nil {
		panic(err)
	}

}

// public final java.lang.String toString()
func (jbobject *AstNode) ToString() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toString", "java/lang/String")
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

// public final java.lang.String toStringWithoutComments()
func (jbobject *AstNode) ToStringWithoutComments() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "toStringWithoutComments", "java/lang/String")
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

// public final int hashCode()
func (jbobject *AstNode) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *AstNode) Equals(a interface{}) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "equals", javabind.Boolean, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public com.github.javaparser.ast.Node clone()
func (jbobject *AstNode) Clone() *AstNode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "com/github/javaparser/ast/Node")
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public com.github.javaparser.ast.Node getParentNode()
func (jbobject *AstNode) GetParentNode() *AstNode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentNode", "com/github/javaparser/ast/Node")
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
	unique_x := &AstNode{}
	unique_x.Callable = dst
	return unique_x
}

// public java.util.List<com.github.javaparser.ast.Node> getChildrenNodes()
func (jbobject *AstNode) GetChildrenNodes() []*AstNode {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getChildrenNodes", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstNode)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean contains(com.github.javaparser.ast.Node)
func (jbobject *AstNode) Contains(a AstNodeInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "contains", javabind.Boolean, conv_a.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public void addOrphanComment(com.github.javaparser.ast.comments.Comment)
func (jbobject *AstNode) AddOrphanComment(a AstCommentsCommentInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "addOrphanComment", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/comments/Comment"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public java.util.List<com.github.javaparser.ast.comments.Comment> getOrphanComments()
func (jbobject *AstNode) GetOrphanComments() []*AstCommentsComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getOrphanComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.util.List<com.github.javaparser.ast.comments.Comment> getAllContainedComments()
func (jbobject *AstNode) GetAllContainedComments() []*AstCommentsComment {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAllContainedComments", "java/util/List")
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoList(javabind.NewJavaToGoCallable())
	dst := new([]*AstCommentsComment)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public void setParentNode(com.github.javaparser.ast.Node)
func (jbobject *AstNode) SetParentNode(a AstNodeInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := jbobject.CallMethod(javabind.GetEnv(), "setParentNode", javabind.Void, conv_a.Value().Cast("com/github/javaparser/ast/Node"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public boolean isPositionedAfter(int, int)
func (jbobject *AstNode) IsPositionedAfter(a int, b int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPositionedAfter", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isPositionedBefore(int, int)
func (jbobject *AstNode) IsPositionedBefore(a int, b int) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isPositionedBefore", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean hasComment()
func (jbobject *AstNode) HasComment() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hasComment", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.Object clone() throws java.lang.CloneNotSupportedException
func (jbobject *AstNode) Clone2() (*JavaLangObject, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
	if err != nil {
		var zero *JavaLangObject
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaLangObject{}
	unique_x.Callable = dst
	return unique_x, nil
}

func (jbobject *AstNode) ABSOLUTE_BEGIN_LINE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/Node", "ABSOLUTE_BEGIN_LINE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstNode) SetFieldABSOLUTE_BEGIN_LINE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/Node", "ABSOLUTE_BEGIN_LINE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstNode) ABSOLUTE_END_LINE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/Node", "ABSOLUTE_END_LINE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstNode) SetFieldABSOLUTE_END_LINE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/Node", "ABSOLUTE_END_LINE", val)
	if err != nil {
		panic(err)
	}

}


