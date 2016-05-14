package javaparser

import "github.com/timob/javabind"

type JavaIoFileInterface interface {
	JavaLangObjectInterface

	// public java.lang.String getName()
	GetName() string

	// public java.lang.String getParent()
	GetParent() string

	// public java.io.File getParentFile()
	GetParentFile() *JavaIoFile

	// public java.lang.String getPath()
	GetPath() string

	// public boolean isAbsolute()
	IsAbsolute() bool

	// public java.lang.String getAbsolutePath()
	GetAbsolutePath() string

	// public java.io.File getAbsoluteFile()
	GetAbsoluteFile() *JavaIoFile

	// public boolean canRead()
	CanRead() bool

	// public boolean canWrite()
	CanWrite() bool

	// public boolean exists()
	Exists() bool

	// public boolean isDirectory()
	IsDirectory() bool

	// public boolean isFile()
	IsFile() bool

	// public boolean isHidden()
	IsHidden() bool

	// public long lastModified()
	LastModified() int64

	// public long length()
	Length() int64

	// public boolean delete()
	Delete() bool

	// public void deleteOnExit()
	DeleteOnExit() 

	// public java.lang.String[] list()
	List() []string

	// public java.lang.String[] list(java.io.FilenameFilter)
	List2(a JavaIoFilenameFilterInterface) []string

	// public java.io.File[] listFiles()
	ListFiles() []*JavaIoFile

	// public java.io.File[] listFiles(java.io.FilenameFilter)
	ListFiles2(a JavaIoFilenameFilterInterface) []*JavaIoFile

	// public java.io.File[] listFiles(java.io.FileFilter)
	ListFiles3(a JavaIoFileFilterInterface) []*JavaIoFile

	// public boolean mkdir()
	Mkdir() bool

	// public boolean mkdirs()
	Mkdirs() bool

	// public boolean renameTo(java.io.File)
	RenameTo(a JavaIoFileInterface) bool

	// public boolean setLastModified(long)
	SetLastModified(a int64) bool

	// public boolean setReadOnly()
	SetReadOnly() bool

	// public boolean setWritable(boolean, boolean)
	SetWritable(a bool, b bool) bool

	// public boolean setWritable(boolean)
	SetWritable2(a bool) bool

	// public boolean setReadable(boolean, boolean)
	SetReadable(a bool, b bool) bool

	// public boolean setReadable(boolean)
	SetReadable2(a bool) bool

	// public boolean setExecutable(boolean, boolean)
	SetExecutable(a bool, b bool) bool

	// public boolean setExecutable(boolean)
	SetExecutable2(a bool) bool

	// public boolean canExecute()
	CanExecute() bool

	// public static java.io.File[] listRoots()
	ListRoots() []*JavaIoFile

	// public long getTotalSpace()
	GetTotalSpace() int64

	// public long getFreeSpace()
	GetFreeSpace() int64

	// public long getUsableSpace()
	GetUsableSpace() int64

	// public int compareTo(java.io.File)
	CompareTo(a JavaIoFileInterface) int

	// public int compareTo(java.lang.Object)
	CompareTo2(a interface{}) int
}

type JavaIoFile struct {
	JavaLangObject
}

// public java.io.File(java.lang.String)
func NewJavaIoFile(a string) (*JavaIoFile) {
	conv_a := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/File", conv_a.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	x := &JavaIoFile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.File(java.lang.String, java.lang.String)
func NewJavaIoFile2(a string, b string) (*JavaIoFile) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/File", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoFile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.io.File(java.io.File, java.lang.String)
func NewJavaIoFile3(a JavaIoFileInterface, b string) (*JavaIoFile) {
	conv_a := javabind.NewGoToJavaCallable()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}

	obj, err := javabind.GetEnv().NewObject("java/io/File", conv_a.Value().Cast("java/io/File"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	x := &JavaIoFile{}
	x.Callable = &javabind.Callable{obj}
	return x
}

// public java.lang.String getName()
func (jbobject *JavaIoFile) GetName() string {
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

// public java.lang.String getParent()
func (jbobject *JavaIoFile) GetParent() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParent", "java/lang/String")
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

// public java.io.File getParentFile()
func (jbobject *JavaIoFile) GetParentFile() *JavaIoFile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getParentFile", "java/io/File")
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
	unique_x := &JavaIoFile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getPath()
func (jbobject *JavaIoFile) GetPath() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getPath", "java/lang/String")
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

// public boolean isAbsolute()
func (jbobject *JavaIoFile) IsAbsolute() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isAbsolute", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public java.lang.String getAbsolutePath()
func (jbobject *JavaIoFile) GetAbsolutePath() string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAbsolutePath", "java/lang/String")
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

// public java.io.File getAbsoluteFile()
func (jbobject *JavaIoFile) GetAbsoluteFile() *JavaIoFile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getAbsoluteFile", "java/io/File")
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
	unique_x := &JavaIoFile{}
	unique_x.Callable = dst
	return unique_x
}

// public java.lang.String getCanonicalPath() throws java.io.IOException
func (jbobject *JavaIoFile) GetCanonicalPath() (string, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCanonicalPath", "java/lang/String")
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

// public java.io.File getCanonicalFile() throws java.io.IOException
func (jbobject *JavaIoFile) GetCanonicalFile() (*JavaIoFile, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getCanonicalFile", "java/io/File")
	if err != nil {
		var zero *JavaIoFile
		return zero, err
	}
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoFile{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public boolean canRead()
func (jbobject *JavaIoFile) CanRead() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canRead", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean canWrite()
func (jbobject *JavaIoFile) CanWrite() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canWrite", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean exists()
func (jbobject *JavaIoFile) Exists() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "exists", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isDirectory()
func (jbobject *JavaIoFile) IsDirectory() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isDirectory", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isFile()
func (jbobject *JavaIoFile) IsFile() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isFile", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean isHidden()
func (jbobject *JavaIoFile) IsHidden() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "isHidden", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public long lastModified()
func (jbobject *JavaIoFile) LastModified() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "lastModified", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long length()
func (jbobject *JavaIoFile) Length() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "length", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public boolean createNewFile() throws java.io.IOException
func (jbobject *JavaIoFile) CreateNewFile() (bool, error) {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "createNewFile", javabind.Boolean)
	if err != nil {
		var zero bool
		return zero, err
	}
	return jret.(bool), nil
}

// public boolean delete()
func (jbobject *JavaIoFile) Delete() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "delete", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public void deleteOnExit()
func (jbobject *JavaIoFile) DeleteOnExit()  {
	_, err := jbobject.CallMethod(javabind.GetEnv(), "deleteOnExit", javabind.Void)
	if err != nil {
		panic(err)
	}

}

// public java.lang.String[] list()
func (jbobject *JavaIoFile) List() []string {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "list", javabind.ObjectArrayType("java/lang/String"))
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

// public java.lang.String[] list(java.io.FilenameFilter)
func (jbobject *JavaIoFile) List2(a JavaIoFilenameFilterInterface) []string {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "list", javabind.ObjectArrayType("java/lang/String"), conv_a.Value().Cast("java/io/FilenameFilter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoString(), "java/lang/String")
	dst := new([]string)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.io.File[] listFiles()
func (jbobject *JavaIoFile) ListFiles() []*JavaIoFile {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listFiles", javabind.ObjectArrayType("java/io/File"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/io/File")
	dst := new([]*JavaIoFile)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.io.File[] listFiles(java.io.FilenameFilter)
func (jbobject *JavaIoFile) ListFiles2(a JavaIoFilenameFilterInterface) []*JavaIoFile {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listFiles", javabind.ObjectArrayType("java/io/File"), conv_a.Value().Cast("java/io/FilenameFilter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/io/File")
	dst := new([]*JavaIoFile)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public java.io.File[] listFiles(java.io.FileFilter)
func (jbobject *JavaIoFile) ListFiles3(a JavaIoFileFilterInterface) []*JavaIoFile {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "listFiles", javabind.ObjectArrayType("java/io/File"), conv_a.Value().Cast("java/io/FileFilter"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/io/File")
	dst := new([]*JavaIoFile)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public boolean mkdir()
func (jbobject *JavaIoFile) Mkdir() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mkdir", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean mkdirs()
func (jbobject *JavaIoFile) Mkdirs() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "mkdirs", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean renameTo(java.io.File)
func (jbobject *JavaIoFile) RenameTo(a JavaIoFileInterface) bool {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "renameTo", javabind.Boolean, conv_a.Value().Cast("java/io/File"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(bool)
}

// public boolean setLastModified(long)
func (jbobject *JavaIoFile) SetLastModified(a int64) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setLastModified", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setReadOnly()
func (jbobject *JavaIoFile) SetReadOnly() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setReadOnly", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setWritable(boolean, boolean)
func (jbobject *JavaIoFile) SetWritable(a bool, b bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setWritable", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setWritable(boolean)
func (jbobject *JavaIoFile) SetWritable2(a bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setWritable", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setReadable(boolean, boolean)
func (jbobject *JavaIoFile) SetReadable(a bool, b bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setReadable", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setReadable(boolean)
func (jbobject *JavaIoFile) SetReadable2(a bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setReadable", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setExecutable(boolean, boolean)
func (jbobject *JavaIoFile) SetExecutable(a bool, b bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setExecutable", javabind.Boolean, a, b)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean setExecutable(boolean)
func (jbobject *JavaIoFile) SetExecutable2(a bool) bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "setExecutable", javabind.Boolean, a)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public boolean canExecute()
func (jbobject *JavaIoFile) CanExecute() bool {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "canExecute", javabind.Boolean)
	if err != nil {
		panic(err)
	}
	return jret.(bool)
}

// public static java.io.File[] listRoots()
func (jbobject *JavaIoFile) ListRoots() []*JavaIoFile {
	jret, err := javabind.GetEnv().CallStaticMethod("java/io/File", "listRoots", javabind.ObjectArrayType("java/io/File"))
	if err != nil {
		panic(err)
	}
	retconv := javabind.NewJavaToGoObjectArray(javabind.NewJavaToGoCallable(), "java/io/File")
	dst := new([]*JavaIoFile)
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	return *dst
}

// public long getTotalSpace()
func (jbobject *JavaIoFile) GetTotalSpace() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getTotalSpace", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getFreeSpace()
func (jbobject *JavaIoFile) GetFreeSpace() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getFreeSpace", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public long getUsableSpace()
func (jbobject *JavaIoFile) GetUsableSpace() int64 {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "getUsableSpace", javabind.Long)
	if err != nil {
		panic(err)
	}
	return jret.(int64)
}

// public static java.io.File createTempFile(java.lang.String, java.lang.String, java.io.File) throws java.io.IOException
func (jbobject *JavaIoFile) CreateTempFile(a string, b string, c JavaIoFileInterface) (*JavaIoFile, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	conv_c := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	if err := conv_c.Convert(c); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/io/File", "createTempFile", "java/io/File", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"), conv_c.Value().Cast("java/io/File"))
	if err != nil {
		var zero *JavaIoFile
		return zero, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	conv_c.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoFile{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public static java.io.File createTempFile(java.lang.String, java.lang.String) throws java.io.IOException
func (jbobject *JavaIoFile) CreateTempFile2(a string, b string) (*JavaIoFile, error) {
	conv_a := javabind.NewGoToJavaString()
	conv_b := javabind.NewGoToJavaString()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	if err := conv_b.Convert(b); err != nil {
		panic(err)
	}
	jret, err := javabind.GetEnv().CallStaticMethod("java/io/File", "createTempFile", "java/io/File", conv_a.Value().Cast("java/lang/String"), conv_b.Value().Cast("java/lang/String"))
	if err != nil {
		var zero *JavaIoFile
		return zero, err
	}
	conv_a.CleanUp()
	conv_b.CleanUp()
	retconv := javabind.NewJavaToGoCallable()
	dst := &javabind.Callable{}
	retconv.Dest(dst)
	if err := retconv.Convert(javabind.ObjectRef(jret)); err != nil {
		panic(err)
	}
	retconv.CleanUp()
	unique_x := &JavaIoFile{}
	unique_x.Callable = dst
	return unique_x, nil
}

// public int compareTo(java.io.File)
func (jbobject *JavaIoFile) CompareTo(a JavaIoFileInterface) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/io/File"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

// public boolean equals(java.lang.Object)
func (jbobject *JavaIoFile) Equals(a interface{}) bool {
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

// public int hashCode()
func (jbobject *JavaIoFile) HashCode() int {
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "hashCode", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

// public java.lang.String toString()
func (jbobject *JavaIoFile) ToString() string {
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

// public int compareTo(java.lang.Object)
func (jbobject *JavaIoFile) CompareTo2(a interface{}) int {
	conv_a := javabind.NewGoToJavaCallable()
	if err := conv_a.Convert(a); err != nil {
		panic(err)
	}
	jret, err := jbobject.CallMethod(javabind.GetEnv(), "compareTo", javabind.Int, conv_a.Value().Cast("java/lang/Object"))
	if err != nil {
		panic(err)
	}
	conv_a.CleanUp()
	return jret.(int)
}

func (jbobject *JavaIoFile) Separator() string {
	jret, err := javabind.GetEnv().GetStaticField("java/io/File", "separator", "java/lang/String")
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

func (jbobject *JavaIoFile) SetFieldSeparator(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/io/File", "separator", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}

func (jbobject *JavaIoFile) PathSeparator() string {
	jret, err := javabind.GetEnv().GetStaticField("java/io/File", "pathSeparator", "java/lang/String")
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

func (jbobject *JavaIoFile) SetFieldPathSeparator(val string) {
	conv_val := javabind.NewGoToJavaString()
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("java/io/File", "pathSeparator", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


