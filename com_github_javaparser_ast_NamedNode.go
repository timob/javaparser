package javaparser

import "github.com/timob/javabind"

type AstNamedNodeInterface interface {

	// public abstract java.lang.String getName()
	GetName() string
}

type AstNamedNode struct {
	JavaLangObject
}

// public abstract java.lang.String getName()
func (jbobject *AstNamedNode) GetName() string {
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


