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
	classes    []class
	functions  []function
	interfaces []tinterface
	variables  []variable
	enums      []enum
}

type class struct {
	export             bool
	name               string
	extends            bool
	extendedClass      string
	implements         bool
	implementedClasses []string
	variables          []variable
	functions          []function
	abstract           bool
	declare            bool
	cdefault           bool
}

type enum struct {
	export  bool
	declare bool
	name    string
	items   []enumItem
}
type enumItem struct {
	name  string
	value string
}

type tinterface struct {
	export          bool
	declare         bool
	name            string
	extends         bool
	extendedClasses []string
	variables       []variable
	functions       []function
}

type docstring struct {
	paramDocs map[string]string
	desc      string
	rtype     docrtype
}
type docrtype struct {
	rtype string
	desc  string
}

type variable struct {
	export         bool
	accessModifier string
	readonly       bool
	name           string
	dType          string
	dValue         string
	static         bool
	declare        bool
}

type annotation struct {
}

type function struct {
	export         bool
	accessModifier string
	declare        bool
	async          bool
	annotations    []annotation
	docs           *docstring
	name           string
	rtype          string
	body           *fbody
	sbody          string
	params         []param
	static         bool
}
type param struct {
	name  string
	dtype string
	value string
}
type fbody struct {
	statements []string
	variables  []string
	returns    []string
}

// ParseLoop ...
func ParseLoop() *TSFile {
	classes := []class{}
	functions := []function{}
	variables := []variable{}
	tinterfaces := []tinterface{}
	enums := []enum{}
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

	return &TSFile{classes: classes, functions: functions, variables: variables, interfaces: tinterfaces, enums: enums}
}

func parseDoc(docS string) docstring {
	docstring := docstring{}
	docS = strings.Replace(docS, "*", "", -1)
	docstring.desc = readDocDesc(docS)
	docstring.paramDocs = readDocParams(docS)
	docstring.rtype = readReturnType(docS)
	fmt.Println(docS)
	return docstring
}

func readReturnType(doc string) docrtype {
	docrtype := docrtype{}
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

func readEnum(identifiers []string) enum {
	enum := enum{}
	enum.name = next()
	next()
	for token := next(); token != "}"; token = next() {
		enum.items = append(enum.items, readEnumItem(token))
	}
	enum.attachIdentifiers(identifiers)
	return enum
}

func readEnumItem(name string) enumItem {
	next() // This is =
	itemValue := ""
	for token := next(); token != ","; token = next() {
		if token == "}" {
			prev()
			break
		}
		itemValue += token
	}
	return enumItem{name: name, value: itemValue}
}

func readInterface(identifiers []string) tinterface {
	tinterface := tinterface{}
	tinterface.name = next()
	token := next()
	if token == "extends" {
		tinterface.extends = true
		tinterface.extendedClasses = readExtendedInterfaces()
	}
	for token = next(); token != "}"; token = next() {
		name := token
		token = next() // This is => : or (
		if token == "(" {
			tinterface.functions = append(tinterface.functions, readFunction(name, []string{}))
		} else {
			tinterface.variables = append(tinterface.variables, readVariable(token, name, []string{}))
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

func readClass(identifiers []string) class {
	class := class{}
	class.name = next()
	token := next()
	if token == "extends" {
		class.extends = true
		class.extendedClass = next()
		token = next()
	}
	if token == "implements" {
		class.implements = true
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
			class.functions = append(class.functions, readFunction(name, identifiers))
		} else { // class element is a variable

			class.variables = append(class.variables, readVariable(token, name, identifiers))
		}

	}
	class.attachIdentifiers(identifiers)
	return class
}

func (class *class) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			class.export = true
		} else if val == "abstract" {
			class.abstract = true
		} else if val == "default" {
			class.cdefault = true
		} else if val == "declare" {
			class.declare = true
		}
	}
}

func (tinterface *tinterface) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			tinterface.export = true
		} else if val == "declare" {
			tinterface.declare = true
		}
	}
}

func (enum *enum) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "export" {
			enum.export = true
		} else if val == "declare" {
			enum.declare = true
		}
	}
}

func (fun *function) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "async" {
			fun.async = true
		} else if val == "static" {
			fun.static = true
		} else if val == "export" {
			fun.export = true
		} else if val == "declare" {
			fun.declare = true
		} else if accessModifiers[val] {
			fun.accessModifier = val
		}
	}
}
func (variable *variable) attachIdentifiers(identifiers []string) {
	for _, val := range identifiers {
		if val == "static" {
			variable.static = true
		} else if val == "readonly" {
			variable.readonly = true
		} else if val == "declare" {
			variable.declare = true
		} else if val == "export" {
			variable.export = true
		} else if accessModifiers[val] {
			variable.accessModifier = val
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

func readVariable(token string, name string, identifiers []string) variable {
	vr := variable{}
	vr.name = name
	vr.dType = "any"
	if token == ":" {
		vr.dType, token = readVariableDataType()
	}
	if token == "=" {
		vr.dValue = readVariableDataValue()
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
func readImplementedInterfaces(class *class) string {
	token := next()
	for ; token != "{"; token = next() {
		if token == "," {
			continue
		}
		class.implementedClasses = append(class.implementedClasses, token)
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

func collectParameterDataOfFunction(fun *function, nameOfTheFirstParam string) {
	prm := param{}
	prm.name = nameOfTheFirstParam
	token := next()
	if token == "," {
		prm.dtype = "any"
		fun.params = append(fun.params, prm)
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
		prm.dtype = dtype
		if !(token == "," || token == ")") {
			prm.value = readDefaultValueOfParameter()
		}
		fun.params = append(fun.params, prm)
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

func readFunction(name string, identifiers []string) function {
	fun := function{}
	fun.name = name
	token := next()
	if token != ")" { // if there are parameters read them
		collectParameterDataOfFunction(&fun, token)
	}
	token = next() // if given is '{' start of the function body otherwise ':' expected get the rtype
	if token == "{" {
		fun.rtype = "any" // if no return type given
	} else {
		fun.rtype, token = readFunctionReturnType()
	}
	fun.attachIdentifiers(identifiers)
	if !fun.declare && token != ";" {
		fun.sbody = getFunctionBodyToString()
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
