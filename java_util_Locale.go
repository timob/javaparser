package javaparser

import "github.com/timob/javabind"

type JavaUtilLocaleInterface interface {
	JavaLangObjectInterface

	// public static java.util.Locale getDefault()
	GetDefault() *JavaUtilLocale

	// public static synchronized void setDefault(java.util.Locale)
	SetDefault(a JavaUtilLocaleInterface) 

	// public static java.util.Locale[] getAvailableLocales()
	GetAvailableLocales() []*JavaUtilLocale

	// public static java.lang.String[] getISOCountries()
	GetISOCountries() []string

	// public static java.lang.String[] getISOLanguages()
	GetISOLanguages() []string

	// public java.lang.String getLanguage()
	GetLanguage() string

	// public java.lang.String getCountry()
	GetCountry() string

	// public java.lang.String getVariant()
	GetVariant() string

	// public final java.lang.String getDisplayLanguage()
	GetDisplayLanguage() string

	// public java.lang.String getDisplayLanguage(java.util.Locale)
	GetDisplayLanguage2(a JavaUtilLocaleInterface) string

	// public final java.lang.String getDisplayCountry()
	GetDisplayCountry() string

	// public java.lang.String getDisplayCountry(java.util.Locale)
	GetDisplayCountry2(a JavaUtilLocaleInterface) string

	// public final java.lang.String getDisplayVariant()
	GetDisplayVariant() string

	// public java.lang.String getDisplayVariant(java.util.Locale)
	GetDisplayVariant2(a JavaUtilLocaleInterface) string

	// public final java.lang.String getDisplayName()
	GetDisplayName() string

	// public java.lang.String getDisplayName(java.util.Locale)
	GetDisplayName2(a JavaUtilLocaleInterface) string

	// public java.lang.Object clone()
	Clone() *JavaLangObject
}

type JavaUtilLocale struct {
	JavaLangObject
}

// public java.util.Locale(java.lang.String, java.lang.String, java.lang.String)
func NewJavaUtilLocale(a string, b string, c string) (*JavaUtilLocale) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/util/Locale", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	x := &JavaUtilLocale{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Locale(java.lang.String, java.lang.String)
func NewJavaUtilLocale2(a string, b string) (*JavaUtilLocale) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/util/Locale", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaUtilLocale{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.util.Locale(java.lang.String)
func NewJavaUtilLocale3(a string) (*JavaUtilLocale) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/util/Locale", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaUtilLocale{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public static java.util.Locale getDefault()
func (jbobject *JavaUtilLocale) GetDefault() *JavaUtilLocale {
	jret, err := javabind.GetEnv().CallStaticMethod("java/util/Locale", "getDefault", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

// public static synchronized void setDefault(java.util.Locale)
func (jbobject *JavaUtilLocale) SetDefault(a JavaUtilLocaleInterface)  {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	_, err := javabind.GetEnv().CallStaticMethod("java/util/Locale", "setDefault", javabind.Void, conv_a.Value().Cast("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()

}

// public static java.util.Locale[] getAvailableLocales()
func (jbobject *JavaUtilLocale) GetAvailableLocales() []*JavaUtilLocale {
	jret, err := javabind.GetEnv().CallStaticMethod("java/util/Locale", "getAvailableLocales", javabind.ObjectArrayType("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/util/Locale")
	dst := new([]*JavaUtilLocale)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String[] getISOCountries()
func (jbobject *JavaUtilLocale) GetISOCountries() []string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/util/Locale", "getISOCountries", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public static java.lang.String[] getISOLanguages()
func (jbobject *JavaUtilLocale) GetISOLanguages() []string {
	jret, err := javabind.GetEnv().CallStaticMethod("java/util/Locale", "getISOLanguages", javabind.ObjectArrayType("java/lang/String"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.String getLanguage()
func (jbobject *JavaUtilLocale) GetLanguage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getLanguage", "java/lang/String")
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

// public java.lang.String getCountry()
func (jbobject *JavaUtilLocale) GetCountry() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCountry", "java/lang/String")
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

// public java.lang.String getVariant()
func (jbobject *JavaUtilLocale) GetVariant() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getVariant", "java/lang/String")
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

// public final java.lang.String toString()
func (jbobject *JavaUtilLocale) ToString() string {
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

// public java.lang.String getISO3Language() throws java.util.MissingResourceException
func (jbobject *JavaUtilLocale) GetISO3Language() (string, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getISO3Language", "java/lang/String")
	if err != nil {
		var zero string
		return zero, err
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public java.lang.String getISO3Country() throws java.util.MissingResourceException
func (jbobject *JavaUtilLocale) GetISO3Country() (string, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getISO3Country", "java/lang/String")
	if err != nil {
		var zero string
		return zero, err
	}
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst, nil
}

// public final java.lang.String getDisplayLanguage()
func (jbobject *JavaUtilLocale) GetDisplayLanguage() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayLanguage", "java/lang/String")
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

// public java.lang.String getDisplayLanguage(java.util.Locale)
func (jbobject *JavaUtilLocale) GetDisplayLanguage2(a JavaUtilLocaleInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayLanguage", "java/lang/String", conv_a.Value().Cast("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final java.lang.String getDisplayCountry()
func (jbobject *JavaUtilLocale) GetDisplayCountry() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayCountry", "java/lang/String")
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

// public java.lang.String getDisplayCountry(java.util.Locale)
func (jbobject *JavaUtilLocale) GetDisplayCountry2(a JavaUtilLocaleInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayCountry", "java/lang/String", conv_a.Value().Cast("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final java.lang.String getDisplayVariant()
func (jbobject *JavaUtilLocale) GetDisplayVariant() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayVariant", "java/lang/String")
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

// public java.lang.String getDisplayVariant(java.util.Locale)
func (jbobject *JavaUtilLocale) GetDisplayVariant2(a JavaUtilLocaleInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayVariant", "java/lang/String", conv_a.Value().Cast("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public final java.lang.String getDisplayName()
func (jbobject *JavaUtilLocale) GetDisplayName() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayName", "java/lang/String")
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

// public java.lang.String getDisplayName(java.util.Locale)
func (jbobject *JavaUtilLocale) GetDisplayName2(a JavaUtilLocaleInterface) string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getDisplayName", "java/lang/String", conv_a.Value().Cast("java/util/Locale"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoString()
	dst := new(string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.lang.Object clone()
func (jbobject *JavaUtilLocale) Clone() *JavaLangObject {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "clone", "java/lang/Object")
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

// public int hashCode()
func (jbobject *JavaUtilLocale) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaUtilLocale) Equals(a interface{}) bool {
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

func (jbobject *JavaUtilLocale) ENGLISH() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "ENGLISH", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldENGLISH(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "ENGLISH", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) FRENCH() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "FRENCH", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldFRENCH(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "FRENCH", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) GERMAN() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "GERMAN", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldGERMAN(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "GERMAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) ITALIAN() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "ITALIAN", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldITALIAN(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "ITALIAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) JAPANESE() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "JAPANESE", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldJAPANESE(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "JAPANESE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) KOREAN() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "KOREAN", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldKOREAN(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "KOREAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) CHINESE() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "CHINESE", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldCHINESE(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "CHINESE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) SIMPLIFIED_CHINESE() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "SIMPLIFIED_CHINESE", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldSIMPLIFIED_CHINESE(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "SIMPLIFIED_CHINESE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) TRADITIONAL_CHINESE() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "TRADITIONAL_CHINESE", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldTRADITIONAL_CHINESE(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "TRADITIONAL_CHINESE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) FRANCE() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "FRANCE", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldFRANCE(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "FRANCE", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) GERMANY() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "GERMANY", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldGERMANY(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "GERMANY", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) ITALY() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "ITALY", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldITALY(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "ITALY", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) JAPAN() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "JAPAN", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldJAPAN(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "JAPAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) KOREA() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "KOREA", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldKOREA(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "KOREA", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) CHINA() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "CHINA", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldCHINA(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "CHINA", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) PRC() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "PRC", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldPRC(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "PRC", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) TAIWAN() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "TAIWAN", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldTAIWAN(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "TAIWAN", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) UK() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "UK", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldUK(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "UK", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) US() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "US", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldUS(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "US", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) CANADA() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "CANADA", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldCANADA(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "CANADA", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) CANADA_FRENCH() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "CANADA_FRENCH", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldCANADA_FRENCH(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "CANADA_FRENCH", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaUtilLocale) ROOT() *JavaUtilLocale {
	jret, err := javabind.GetEnv().GetStaticField("java/util/Locale", "ROOT", "java/util/Locale")
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
	unique_x := &JavaUtilLocale{}
	unique_x.Callable = dst
	return unique_x
}

func (jbobject *JavaUtilLocale) SetFieldROOT(val JavaUtilLocaleInterface) {
	conv_val := javabind.NewGoToJavaCallable()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/util/Locale", "ROOT", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


