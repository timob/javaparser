package javaparser

import "github.com/timob/javabind"

type JavaIoFileFilterInterface interface {

	// public abstract boolean accept(java.io.File)
	Accept(a JavaIoFileInterface) bool
}

type JavaIoFileFilter struct {
	JavaLangObject
}

// public abstract boolean accept(java.io.File)
func (jbobject *JavaIoFileFilter) Accept(a JavaIoFileInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accept", javabind.Boolean, conv_a.Value().Cast("java/io/File"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}


