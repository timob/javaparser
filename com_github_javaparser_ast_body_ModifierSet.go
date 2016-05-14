package javaparser

import "github.com/timob/javabind"

type AstBodyModifierSetInterface interface {
	JavaLangObjectInterface

	// public static com.github.javaparser.ast.AccessSpecifier getAccessSpecifier(int)
	GetAccessSpecifier(a int) *AstAccessSpecifier

	// public static int addModifier(int, int)
	AddModifier(a int, b int) int

	// public static boolean hasModifier(int, int)
	HasModifier(a int, b int) bool

	// public static boolean isAbstract(int)
	IsAbstract(a int) bool

	// public static boolean isFinal(int)
	IsFinal(a int) bool

	// public static boolean isNative(int)
	IsNative(a int) bool

	// public static boolean isPrivate(int)
	IsPrivate(a int) bool

	// public static boolean isProtected(int)
	IsProtected(a int) bool

	// public static boolean hasPackageLevelAccess(int)
	HasPackageLevelAccess(a int) bool

	// public static boolean isPublic(int)
	IsPublic(a int) bool

	// public static boolean isStatic(int)
	IsStatic(a int) bool

	// public static boolean isStrictfp(int)
	IsStrictfp(a int) bool

	// public static boolean isSynchronized(int)
	IsSynchronized(a int) bool

	// public static boolean isTransient(int)
	IsTransient(a int) bool

	// public static boolean isVolatile(int)
	IsVolatile(a int) bool

	// public static int removeModifier(int, int)
	RemoveModifier(a int, b int) int
}

type AstBodyModifierSet struct {
	JavaLangObject
}

// public static com.github.javaparser.ast.AccessSpecifier getAccessSpecifier(int)
func (jbobject *AstBodyModifierSet) GetAccessSpecifier(a int) *AstAccessSpecifier {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "getAccessSpecifier", "com/github/javaparser/ast/AccessSpecifier", a)
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
	unique_x := &AstAccessSpecifier{}
	unique_x.Callable = dst
	return unique_x
}

// public static int addModifier(int, int)
func (jbobject *AstBodyModifierSet) AddModifier(a int, b int) int {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "addModifier", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public static boolean hasModifier(int, int)
func (jbobject *AstBodyModifierSet) HasModifier(a int, b int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "hasModifier", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isAbstract(int)
func (jbobject *AstBodyModifierSet) IsAbstract(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isAbstract", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isFinal(int)
func (jbobject *AstBodyModifierSet) IsFinal(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isFinal", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isNative(int)
func (jbobject *AstBodyModifierSet) IsNative(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isNative", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isPrivate(int)
func (jbobject *AstBodyModifierSet) IsPrivate(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isPrivate", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isProtected(int)
func (jbobject *AstBodyModifierSet) IsProtected(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isProtected", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean hasPackageLevelAccess(int)
func (jbobject *AstBodyModifierSet) HasPackageLevelAccess(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "hasPackageLevelAccess", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isPublic(int)
func (jbobject *AstBodyModifierSet) IsPublic(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isPublic", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isStatic(int)
func (jbobject *AstBodyModifierSet) IsStatic(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isStatic", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isStrictfp(int)
func (jbobject *AstBodyModifierSet) IsStrictfp(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isStrictfp", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isSynchronized(int)
func (jbobject *AstBodyModifierSet) IsSynchronized(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isSynchronized", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isTransient(int)
func (jbobject *AstBodyModifierSet) IsTransient(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isTransient", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static boolean isVolatile(int)
func (jbobject *AstBodyModifierSet) IsVolatile(a int) bool {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "isVolatile", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static int removeModifier(int, int)
func (jbobject *AstBodyModifierSet) RemoveModifier(a int, b int) int {
	jret, err := javabind.GetEnv().CallStaticMethod("com/github/javaparser/ast/body/ModifierSet", "removeModifier", javabind.Int, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) PUBLIC() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "PUBLIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldPUBLIC(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "PUBLIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) PRIVATE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "PRIVATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldPRIVATE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "PRIVATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) PROTECTED() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "PROTECTED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldPROTECTED(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "PROTECTED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) STATIC() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "STATIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldSTATIC(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "STATIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) FINAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "FINAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldFINAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "FINAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) SYNCHRONIZED() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "SYNCHRONIZED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldSYNCHRONIZED(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "SYNCHRONIZED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) VOLATILE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "VOLATILE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldVOLATILE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "VOLATILE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) TRANSIENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "TRANSIENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldTRANSIENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "TRANSIENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) NATIVE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "NATIVE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldNATIVE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "NATIVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) ABSTRACT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "ABSTRACT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldABSTRACT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "ABSTRACT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *AstBodyModifierSet) STRICTFP() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ast/body/ModifierSet", "STRICTFP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *AstBodyModifierSet) SetFieldSTRICTFP(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ast/body/ModifierSet", "STRICTFP", val)
	if err != nil {
		panic(err)
	}

}


