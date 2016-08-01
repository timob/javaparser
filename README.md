# javaparser
Golang binding package for https://github.com/javaparser/javaparser

Make sure the class path includes Javaparser classes.

## Example

```` go
package main

// example from https://github.com/javaparser/javaparser/wiki/Manual

import (
    "github.com/timob/javabind"
    "github.com/timob/javaparser"
    "os"
    "log"
    "strings"
    "fmt"
)

func main() {
    err := javabind.SetupJVM(os.Getenv("CLASSPATH"))
    if err != nil {
        log.Fatal(err)
    }

    var parser *javaparser.JavaParser
    var helper *javaparser.ASTHelper

    cu, err := parser.Parse3(javaparser.NewJavaIoByteArrayInputStream([]byte(javaStr)))
    if err != nil {
        log.Fatal(err)
    }

    for _, itype := range cu.GetTypes() {
        for _, jtype := range itype.GetMembers() {
            if jtype.InstanceOf("com/github/javaparser/ast/body/MethodDeclaration") {
                md := &javaparser.AstBodyMethodDeclaration{*jtype}
                md.SetName(strings.ToUpper(md.GetName()))
                newArg := helper.CreateParameter(helper.INT_TYPE(), "value")
                helper.AddParameter(md, newArg)
            }
        }
    }

    fmt.Print(cu.ToString())
}


var javaStr = `
class ExampleCLass {

    public ExampleClass() {

    }

    public boolean exampleMethod(int a) {
        return (a == 42);
    }

    int exampleMethod2(String a) {
        return a.length();
    }

}

`
````
