package javaparser

import "github.com/timob/javabind"

type SourcesHelperInterface interface {
	JavaLangObjectInterface
}

type SourcesHelper struct {
	JavaLangObject
}

// public com.github.javaparser.SourcesHelper()
func NewSourcesHelper() (*SourcesHelper) {

	obj, err := javabind.GetEnv().NewObject("com/github/javaparser/SourcesHelper")
	if err != nil {
		panic(err)
	}
	x := &SourcesHelper{}
	x.Callable = &javabind.Callable{obj}
	return x
}


