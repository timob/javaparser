package javaparser

import "github.com/timob/javabind"

type JavaIoFilenameFilterInterface interface {

	// public abstract boolean accept(java.io.File, java.lang.String)
	Accept(a JavaIoFileInterface, b string) bool
}

type JavaIoFilenameFilter struct {
	JavaLangObject
}

// public abstract boolean accept(java.io.File, java.lang.String)
func (jbobject *JavaIoFilenameFilter) Accept(a JavaIoFileInterface, b string) bool {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "accept", javabind.Boolean, conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	return jret.(bool)
}


