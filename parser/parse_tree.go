package parser

import (
	"fmt"
	"strings"
)

var words = map[string]string{
	"function":  "function",
	"const":     "variable",
	"let":       "variable",
	"var":       "variable",
	"class":     "class",
	"interface": "interface",
	"enum":      "enum",
}

var accessModifiers = map[string]bool{
	"public":    true,
	"private":   true,
	"protected": true,
}

// TSFile ...
type TSFile struct {
	Classes    []Class
	Functions  []Function
	Interfaces []Tinterface
	Variables  []Variable
	Enums      []Enum
}

// Class ...
type Class struct {
	Export             bool
	Name               string
	Extends            bool
	ExtendedClass      string
	Implements         bool
	ImplementedClasses []string
	Variables          []Variable
	Functions          []Function
	Abstract           bool
	Declare            bool
	Cdefault           bool
}

// Enum ...
type Enum struct {
	Export  bool
	Declare bool
	Name    string
	Items   []EnumItem
}

// EnumItem ...
type EnumItem struct {
	Name  string
	Value string
}

// Tinterface ...
type Tinterface struct {
	Export          bool
	Declare         bool
	Name            string
	Extends         bool
	ExtendedClasses []string
	Variables       []Variable
	Functions       []Function
}

// Docstring ...
type Docstring struct {
	ParamDocs map[string]string
	Desc      string
	Rtype     Docrtype
}

// Docrtype ...
type Docrtype struct {
	rtype string
	desc  string
}

// Variable ...
type Variable struct {
	Export         bool
	AccessModifier string
	Readonly       bool
	Name           string
	DType          string
	DValue         string
	Static         bool
	Declare        bool
}

// Annotation ...
type Annotation struct {
}

// Function ...
type Function struct {
	Export         bool
	AccessModifier string
	Declare        bool
	Async          bool
	Annotations    []Annotation
	Docs           *Docstring
	Name           string
	Rtype          string
	Body           *Fbody
	Sbody          string
	Params         []FParam
	Static         bool
}

// FParam ...
type FParam struct {
	Name  string
	Dtype string
	Value string
}

// Fbody ...
type Fbody struct {
	Statements []string
	Variables  []string
	Returns    []string
}

// Parse ...
func Parse() *TSFile {
	classes := []Class{}
	functions := []Function{}
	variables := []Variable{}
	tinterfaces := []Tinterface{}
	enums := []Enum{}
	for has() {
		token := next()
		if token == "//" || token == "/*" {
			docS := next()
			parseDoc(docS)
			continue
		}
		identifiers, token := getIdentifiersAndReturnToken(token)
		if value, ok := words[token]; ok {
			if value == "class" {
				classes = append(classes, readClass(identifiers))
			} else if value == "function" {
				name := next()
				next() // opening paranthesis
				functions = append(functions, readFunction(name, identifiers))
			} else if value == "variable" {
				name := next()
				token = next() // ':' or '=' or ';'
				variables = append(variables, readVariable(token, name, identifiers))
			} else if value == "interface" {
				tinterfaces = append(tinterfaces, readInterface(identifiers))
			} else if value == "enum" {
				enums = append(enums, readEnum(identifiers))
			}
		}
	}

	return &TSFile{classes, functions, tinterfaces, variables, enums}
}

func parseDoc(docS string) Docstring {
	docstring := Docstring{}
	docS = strings.Replace(docS, "*", "", -1)
	docstring.Desc = readDocDesc(docS)
	docstring.ParamDocs = readDocParams(docS)
	docstring.Rtype = readReturnType(docS)
	fmt.Println(docS)
	return docstring
}

func readReturnType(doc string) Docrtype {
	docrtype := Docrtype{}
	if !strings.Contains(doc, "@return") {
		return docrtype
	}
	startIdx := strings.Index(doc, "@return") + 7
	endIdx := findEndIdxOfDocReturnType(doc, startIdx)
	pieces := strings.SplitN(doc[startIdx+1:endIdx-1], " ", 2)
	docrtype.rtype = pieces[0]
	docrtype.desc = ""
	if len(pieces) > 1 {
		docrtype.desc = pieces[1]
	}
	return docrtype
}

func findEndIdxOfDocReturnType(doc string, startIdx int) int {
	endIdx := startIdx
	for i := startIdx; i < len(doc) && doc[i] != '@'; i++ {
		endIdx++
	}
	return endIdx
}

func readDocParams(doc string) map[string]string {
	paramMap := map[string]string{}
	if !strings.Contains(doc, "@param") {
		return paramMap
	}
	for i := 1; i < len(doc); i++ {
		if doc[i-1] == '@' && doc[i] == 'p' {
			key, value := readNextDocParam(doc, i+6)
			paramMap[key] = value
		}
	}
	return paramMap
}

func readNextDocParam(doc string, startIdx int) (string, string) {
	endIdx := startIdx
	for i := endIdx; i < len(doc) && doc[i] != '@'; i++ {
		endIdx++
	}
	pieces := strings.SplitN(doc[startIdx:endIdx], " ", 2)
	return pieces[0], pieces[1]
}

func readDocDesc(doc string) string {
	desc := ""
	for i := 0; i < len(doc); i++ {
		if doc[i] == '@' {
			return desc
		}
		desc += string(doc[i])
	}
	return desc
}

func readEnum(identifiers []string) Enum {
	enum := Enum{}
	enum.Name = next()
	next()
	for token := next(); token != "}"; token = next() {
		enum.Items = append(enum.Items, readEnumItem(token))
	}
	enum.attachIdentifiers(identifiers)
	return enum
}

func readEnumItem(name string) EnumItem {
	next() // This is =
	itemValue := ""
	for token := next(); token != ","; token = next() {
		if token == "}" {
			prev()
			break
		}
		itemValue += token
	}
	return EnumItem{name, itemValue}
}

func readInterface(identifiers []string) Tinterface {
	tinterface := Tinterface{}
	tinterface.Name = next()
	token := next()
	if token == "extends" {
		tinterface.Extends = true
		tinterface.ExtendedClasses = readExtendedInterfaces()
	}
	for token = next(); token != "}"; token = next() {
		name := token
		token = next() // This is => : or (
		if token == "(" {
			tinterface.Functions = append(tinterface.Functions, readFunction(name, []string{}))
		} else {
			tinterface.Variables = append(tinterface.Variables, readVariable(token, name, []string{}))
		}
	}
	tinterface.attachIdentifiers(identifiers)
	return tinterface
}

func readExtendedInterfaces() []string {
	tinterfaces := []string{}
	for token := next(); token != "{"; token = next() {
		if token == "," {
			continue
		}
		tinterfaces = append(tinterfaces, token)
	}
	return tinterfaces
}

func readClass(identifiers []string) Class {
	class := Class{}
	class.Name = next()
	token := next()
	if token == "extends" {
		class.Extends = true
		class.ExtendedClass = next()
		token = next()
	}
	if token == "implements" {
		class.Implements = true
		token = readImplementedInterfaces(&class)
	}
	for token := next(); token != "}"; token = next() {
		if token == "//" || token == "/*" {
			docS := next()
			fmt.Println(docS)
			continue
		}
		identifiers, token := getIdentifiersAndReturnToken(token)
		name := token
		token = next()
		if token == "(" { // class element is a function
			class.Functions = append(class.Functions, readFunction(name, identifiers))
		} else { // class element is a variable

			class.Variables = append(class.Variables, readVariable(token, name, identifiers))
		}

	}
	class.attachIdentifiers(identifiers)
	return class
}

func (class *Class) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			class.Export = true
		} else if val == "abstract" {
			class.Abstract = true
		} else if val == "default" {
			class.Cdefault = true
		} else if val == "declare" {
			class.Declare = true
		}
	}
}

func (tinterface *Tinterface) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			tinterface.Export = true
		} else if val == "declare" {
			tinterface.Declare = true
		}
	}
}

func (enum *Enum) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			enum.Export = true
		} else if val == "declare" {
			enum.Declare = true
		}
	}
}

func (fun *Function) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "async" {
			fun.Async = true
		} else if val == "static" {
			fun.Static = true
		} else if val == "export" {
			fun.Export = true
		} else if val == "declare" {
			fun.Declare = true
		} else if accessModifiers[val] {
			fun.AccessModifier = val
		}
	}
}
func (variable *Variable) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "static" {
			variable.Static = true
		} else if val == "readonly" {
			variable.Readonly = true
		} else if val == "declare" {
			variable.Declare = true
		} else if val == "export" {
			variable.Export = true
		} else if accessModifiers[val] {
			variable.AccessModifier = val
		}
	}
}

var identifiersMap = map[string]bool{
	"readonly":  true,
	"static":    true,
	"declare":   true,
	"async":     true,
	"export":    true,
	"public":    true,
	"private":   true,
	"protected": true,
	"abstract":  true,
	"default":   true,
}

func getIdentifiersAndReturnToken(token string) ([]string, string) {
	identifiers := []string{}
	for ; identifiersMap[token]; token = next() {
		identifiers = append(identifiers, token)
	}
	return identifiers, token
}

func readVariable(token string, name string, identifiers []string) Variable {
	vr := Variable{}
	vr.Name = name
	vr.DType = "any"
	if token == ":" {
		vr.DType, token = readVariableDataType()
	}
	if token == "=" {
		vr.DValue = readVariableDataValue()
	}
	vr.attachIdentifiers(identifiers)
	return vr
}

//think about how to implement anonymous function case
func readVariableDataValue() string {
	dataValue := ""
	for token := next(); token != ";"; token = next() {
		if keywords[token] {
			// if there is no semi colon end of the statement.
			// it will check the next one so we one it is a keyword we need to
			// keep track of the token's index. To not lose it we are using the function
			// prev.
			token = prev()
			break
		}
		dataValue += token
		if !has() {
			break
		}
	}
	return dataValue
}

//think about how to implement anonymous function case
func readVariableDataType() (string, string) {
	dataType := ""
	var token string
	for token = next(); token != "=" && token != ";"; token = next() {
		if keywords[token] {
			// if there is no semi colon end of the statement.
			// it will check the next one so we one it is a keyword we need to
			// keep track of the token's index. To not lose it we are using the function
			// prev.
			token = prev()
			break
		}
		dataType += token
		if !has() {
			break
		}
	}
	return dataType, token
}
func readImplementedInterfaces(class *Class) string {
	token := next()
	for ; token != "{"; token = next() {
		if token == "," {
			continue
		}
		class.ImplementedClasses = append(class.ImplementedClasses, token)
	}
	return token
}

func readFunctionReturnType() (string, string) {
	returnType := ""
	var token string
	for token = next(); !keywords[token] && token != "{" && token != ";"; token = next() {
		returnType += token
		if !has() {
			break
		}
	}
	return returnType, token
}

func collectParameterDataOfFunction(fun *Function, nameOfTheFirstParam string) {
	prm := FParam{}
	prm.Name = nameOfTheFirstParam
	token := next()
	if token == "," {
		prm.Dtype = "any"
		fun.Params = append(fun.Params, prm)
		collectParameterDataOfFunction(fun, next())
	} else if token == ":" {
		// get data type
		open := 0
		dtype := ""
		previousToken := ""
		for true {
			previousToken = token
			token = next()
			if previousToken == "=" && token != ">" && open <= 0 {
				prev()
				dtype = dtype[:len(dtype)-1]
				break
			}
			if open <= 0 && token == ")" {
				break // end of the parameter collection
			} else if open <= 0 && token == "," {
				collectParameterDataOfFunction(fun, next()) // new parameter
				break
			} else if open > 0 && token == ")" {
				open--
			} else if token == "(" {
				open++
			}
			dtype += token
		}
		prm.Dtype = dtype
		if !(token == "," || token == ")") {
			prm.Value = readDefaultValueOfParameter()
		}
		fun.Params = append(fun.Params, prm)
	}
}

func readDefaultValueOfParameter() string {
	value := ""
	open := 1
	for true {
		token := next()
		if token == "(" {
			open++
		} else if token == ")" {
			open--
		}
		if open <= 0 {
			break
		}
		value += token
	}
	return value
}

func readFunction(name string, identifiers []string) Function {
	fun := Function{}
	fun.Name = name
	token := next()
	if token != ")" { // if there are parameters read them
		collectParameterDataOfFunction(&fun, token)
	}
	token = next() // if given is '{' start of the function body otherwise ':' expected get the rtype
	if token == "{" {
		fun.Rtype = "any" // if no return type given
	} else {
		fun.Rtype, token = readFunctionReturnType()
	}
	fun.attachIdentifiers(identifiers)
	if !fun.Declare && token != ";" {
		fun.Sbody = getFunctionBodyToString()
	}
	return fun
}

func getFunctionBodyToString() string {
	open := 1
	body := ""
	for token := next(); open > 0; token = next() {
		if token == "{" {
			open++
		} else if token == "}" {
			open--
		}
		if open <= 0 {
			break
		}
		body += token
	}
	return body
}

var idx int

func has() bool {
	return idx < len(tokens)
}
func next() string {
	if !has() {
		panic("End of the stream.")
	}
	token := tokens[idx]
	idx++
	return token
}

func prev() string {
	if idx <= 0 {
		panic("Index is smaller than 0")
	}
	token := tokens[idx-1]
	idx--
	return token
}
