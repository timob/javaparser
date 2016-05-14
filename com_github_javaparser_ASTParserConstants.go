package javaparser

import "github.com/timob/javabind"

type ASTParserConstantsInterface interface {
}

type ASTParserConstants struct {
	JavaLangObject
}

func (jbobject *ASTParserConstants) EOF() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "EOF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldEOF(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "EOF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SINGLE_LINE_COMMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SINGLE_LINE_COMMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSINGLE_LINE_COMMENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SINGLE_LINE_COMMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) JAVA_DOC_COMMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "JAVA_DOC_COMMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldJAVA_DOC_COMMENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "JAVA_DOC_COMMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) MULTI_LINE_COMMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "MULTI_LINE_COMMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldMULTI_LINE_COMMENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "MULTI_LINE_COMMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ABSTRACT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ABSTRACT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldABSTRACT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ABSTRACT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ASSERT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ASSERT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldASSERT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ASSERT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BOOLEAN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BOOLEAN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBOOLEAN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BOOLEAN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BREAK() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BREAK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBREAK(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BREAK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BYTE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BYTE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBYTE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BYTE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CASE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CASE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCASE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CASE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CATCH() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CATCH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCATCH(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CATCH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CHAR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CHAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCHAR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CHAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CLASS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CLASS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCLASS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CLASS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CONST() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CONST", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCONST(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CONST", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CONTINUE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CONTINUE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCONTINUE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CONTINUE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) _DEFAULT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "_DEFAULT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetField_DEFAULT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "_DEFAULT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DO() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDO(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DOUBLE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DOUBLE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDOUBLE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DOUBLE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ELSE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ELSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldELSE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ELSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ENUM() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ENUM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldENUM(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ENUM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) EXTENDS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "EXTENDS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldEXTENDS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "EXTENDS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FALSE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FALSE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFALSE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FALSE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FINAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FINAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFINAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FINAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FINALLY() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FINALLY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFINALLY(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FINALLY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FLOAT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FLOAT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFLOAT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FLOAT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FOR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FOR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFOR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FOR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) GOTO() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "GOTO", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldGOTO(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "GOTO", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IF() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIF(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IMPLEMENTS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IMPLEMENTS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIMPLEMENTS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IMPLEMENTS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IMPORT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IMPORT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIMPORT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IMPORT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) INSTANCEOF() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "INSTANCEOF", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldINSTANCEOF(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "INSTANCEOF", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) INT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "INT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldINT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "INT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) INTERFACE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "INTERFACE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldINTERFACE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "INTERFACE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LONG() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LONG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLONG(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LONG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) NATIVE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "NATIVE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldNATIVE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "NATIVE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) NEW() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "NEW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldNEW(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "NEW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) NULL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "NULL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldNULL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "NULL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PACKAGE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PACKAGE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPACKAGE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PACKAGE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PRIVATE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PRIVATE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPRIVATE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PRIVATE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PROTECTED() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PROTECTED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPROTECTED(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PROTECTED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PUBLIC() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PUBLIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPUBLIC(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PUBLIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RETURN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RETURN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRETURN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RETURN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SHORT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SHORT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSHORT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SHORT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) STATIC() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "STATIC", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSTATIC(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "STATIC", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) STRICTFP() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "STRICTFP", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSTRICTFP(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "STRICTFP", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SUPER() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SUPER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSUPER(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SUPER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SWITCH() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SWITCH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSWITCH(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SWITCH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SYNCHRONIZED() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SYNCHRONIZED", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSYNCHRONIZED(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SYNCHRONIZED", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) THIS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "THIS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTHIS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "THIS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) THROW() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "THROW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTHROW(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "THROW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) THROWS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "THROWS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTHROWS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "THROWS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) TRANSIENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "TRANSIENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTRANSIENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "TRANSIENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) TRUE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "TRUE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTRUE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "TRUE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) TRY() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "TRY", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTRY(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "TRY", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) VOID() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "VOID", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldVOID(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "VOID", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) VOLATILE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "VOLATILE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldVOLATILE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "VOLATILE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) WHILE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "WHILE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldWHILE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "WHILE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LONG_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LONG_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLONG_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LONG_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) INTEGER_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "INTEGER_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldINTEGER_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "INTEGER_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DECIMAL_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDECIMAL_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) HEX_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "HEX_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldHEX_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "HEX_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) OCTAL_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "OCTAL_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldOCTAL_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "OCTAL_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BINARY_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BINARY_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBINARY_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BINARY_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) FLOATING_POINT_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "FLOATING_POINT_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldFLOATING_POINT_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "FLOATING_POINT_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DECIMAL_FLOATING_POINT_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_FLOATING_POINT_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDECIMAL_FLOATING_POINT_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_FLOATING_POINT_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DECIMAL_EXPONENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_EXPONENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDECIMAL_EXPONENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DECIMAL_EXPONENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) HEXADECIMAL_FLOATING_POINT_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "HEXADECIMAL_FLOATING_POINT_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldHEXADECIMAL_FLOATING_POINT_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "HEXADECIMAL_FLOATING_POINT_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) HEXADECIMAL_EXPONENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "HEXADECIMAL_EXPONENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldHEXADECIMAL_EXPONENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "HEXADECIMAL_EXPONENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) CHARACTER_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "CHARACTER_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCHARACTER_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "CHARACTER_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) STRING_LITERAL() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "STRING_LITERAL", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSTRING_LITERAL(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "STRING_LITERAL", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IDENTIFIER() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IDENTIFIER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIDENTIFIER(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IDENTIFIER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LETTER() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LETTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLETTER(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LETTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PART_LETTER() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PART_LETTER", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPART_LETTER(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PART_LETTER", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LPAREN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LPAREN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLPAREN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LPAREN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RPAREN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RPAREN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRPAREN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RPAREN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LBRACE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LBRACE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLBRACE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LBRACE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RBRACE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RBRACE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRBRACE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RBRACE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LBRACKET() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LBRACKET", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLBRACKET(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LBRACKET", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RBRACKET() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RBRACKET", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRBRACKET(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RBRACKET", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SEMICOLON() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SEMICOLON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSEMICOLON(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SEMICOLON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) COMMA() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "COMMA", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCOMMA(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "COMMA", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DOT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DOT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDOT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DOT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) AT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "AT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldAT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "AT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BANG() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BANG", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBANG(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BANG", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) TILDE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "TILDE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldTILDE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "TILDE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) HOOK() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "HOOK", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldHOOK(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "HOOK", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) COLON() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "COLON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldCOLON(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "COLON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) EQ() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "EQ", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldEQ(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "EQ", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) GE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "GE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldGE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "GE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) NE() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "NE", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldNE(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "NE", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SC_OR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SC_OR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSC_OR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SC_OR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SC_AND() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SC_AND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSC_AND(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SC_AND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) INCR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "INCR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldINCR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "INCR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DECR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DECR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDECR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DECR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PLUS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PLUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPLUS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PLUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) MINUS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "MINUS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldMINUS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "MINUS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) STAR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "STAR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSTAR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "STAR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SLASH() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SLASH", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSLASH(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SLASH", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BIT_AND() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BIT_AND", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBIT_AND(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BIT_AND", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) BIT_OR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "BIT_OR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldBIT_OR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "BIT_OR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) XOR() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "XOR", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldXOR(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "XOR", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) REM() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "REM", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldREM(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "REM", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LSHIFT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LSHIFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLSHIFT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LSHIFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) PLUSASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "PLUSASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldPLUSASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "PLUSASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) MINUSASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "MINUSASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldMINUSASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "MINUSASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) STARASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "STARASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSTARASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "STARASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) SLASHASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "SLASHASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldSLASHASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "SLASHASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ANDASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ANDASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldANDASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ANDASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ORASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ORASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldORASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ORASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) XORASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "XORASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldXORASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "XORASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) REMASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "REMASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldREMASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "REMASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) LSHIFTASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "LSHIFTASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldLSHIFTASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "LSHIFTASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RSIGNEDSHIFTASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RSIGNEDSHIFTASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRSIGNEDSHIFTASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RSIGNEDSHIFTASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RUNSIGNEDSHIFTASSIGN() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RUNSIGNEDSHIFTASSIGN", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRUNSIGNEDSHIFTASSIGN(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RUNSIGNEDSHIFTASSIGN", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ELLIPSIS() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ELLIPSIS", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldELLIPSIS(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ELLIPSIS", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) ARROW() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "ARROW", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldARROW(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "ARROW", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DOUBLECOLON() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DOUBLECOLON", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDOUBLECOLON(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DOUBLECOLON", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RUNSIGNEDSHIFT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RUNSIGNEDSHIFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRUNSIGNEDSHIFT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RUNSIGNEDSHIFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) RSIGNEDSHIFT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "RSIGNEDSHIFT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldRSIGNEDSHIFT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "RSIGNEDSHIFT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) GT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "GT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldGT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "GT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) DEFAULT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "DEFAULT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldDEFAULT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "DEFAULT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IN_JAVA_DOC_COMMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IN_JAVA_DOC_COMMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIN_JAVA_DOC_COMMENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IN_JAVA_DOC_COMMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) IN_MULTI_LINE_COMMENT() int {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "IN_MULTI_LINE_COMMENT", javabind.Int)
	if err != nil {
		panic(err)
	}
	return jret.(int)
}

func (jbobject *ASTParserConstants) SetFieldIN_MULTI_LINE_COMMENT(val int) {
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "IN_MULTI_LINE_COMMENT", val)
	if err != nil {
		panic(err)
	}

}

func (jbobject *ASTParserConstants) TokenImage() []string {
	jret, err := javabind.GetEnv().GetStaticField("com/github/javaparser/ASTParserConstants", "tokenImage", javabind.ObjectArrayType("java/lang/String"))
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

func (jbobject *ASTParserConstants) SetFieldTokenImage(val []string) {
	conv_val := javabind.NewGoToJavaObjectArray(javabind.NewGoToJavaString(), "java/lang/String")
	if err := conv_val.Convert(val); err != nil {
		panic(err)
	}
	err := javabind.GetEnv().SetStaticField("com/github/javaparser/ASTParserConstants", "tokenImage", conv_val.Value())
	if err != nil {
		panic(err)
	}
	conv_val.CleanUp()

}


