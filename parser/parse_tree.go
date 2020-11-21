package parser

import "fmt"

var words = map[string]string{
	"function": "function",
	"const":    "variable",
	"let":      "variable",
	"var":      "variable",
	"class":    "class",
}

type param struct {
	name  string
	dtype string
}

type function struct {
	export         bool
	accessModifier string
	declare        bool
	async          bool
	annotations    []annotation
	docs           docstring
	name           string
	rtype          string
	body           *fbody
	sbody          string
	params         []param
	static         bool
}

type docstring struct {
	paramDocs map[string]string
	desc      string
	rtype     string
}

type annotation struct {
}

type fbody struct {
	statements []string
	variables  []string
	returns    []string
}
type variable struct {
	accessModifier string
	readonly       bool
	name           string
	dType          string
	dValue         string
	static         bool
	declare        bool
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
}

func ParseLoop() {
	for has() {
		token := next()
		if value, ok := words[token]; ok {
			// read function or read variable or class
			if value == "class" {
				class := readClass()
				fmt.Println(class)
			}
		}
	}
}

func test(function *function) {
	function.export = true
}

var accessModifiers = map[string]string{
	"public":    "public",
	"private":   "private",
	"protected": "protected",
}

func readClass() class {
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
		token = readInterfaces(&class)
	}
	for token := next(); token != "}"; token = next() {
		accessModifier, token := getAccessModifier(token)
		if accessModifier != "" {
			token = next()
		}
		identifiers, token := getIdentifiersAndReturnToken(token)
		name := token
		token = next()
		if token == "(" { // class element is a function
			fun := readFunction(name)
			fun.attachAccessModifierAndIdentifiers(accessModifier, identifiers)
			class.functions = append(class.functions, fun)
		} else { // class element is a variable
			variable := readVariable(name)
			variable.attachAccessModifierAndIdentifiers(accessModifier, identifiers)
			class.variables = append(class.variables, variable)
		}

	}

	return class
}

func (fun *function) attachAccessModifierAndIdentifiers(accessModifier string, identifiers []string) {
	fun.accessModifier = accessModifier
	for _, val := range identifiers {
		if val == "async" {
			fun.async = true
		} else if val == "static" {
			fun.static = true
		}
	}
}
func (variable *variable) attachAccessModifierAndIdentifiers(accessModifier string, identifiers []string) {
	variable.accessModifier = accessModifier
	for _, val := range identifiers {
		if val == "static" {
			variable.static = true
		} else if val == "readonly" {
			variable.readonly = true
		} else if val == "declare" {
			variable.declare = true
		}
	}
}

var identifiersMap = map[string]bool{
	"readonly": true,
	"static":   true,
	"declare":  true,
	"async":    true,
}

func getIdentifiersAndReturnToken(token string) ([]string, string) {
	identifiers := []string{}
	for ; identifiersMap[token]; token = next() {
		identifiers = append(identifiers, token)
	}
	return identifiers, token
}

func getAccessModifier(token string) (string, string) {
	return accessModifiers[token], token
}

func readVariable(name string) variable {
	vr := variable{}
	vr.name = name
	vr.dType = "any"
	token := next()
	if token == ":" {
		vr.dType, token = readVariableDataType()
	}
	if token == "=" {
		vr.dValue = readVariableDataValue()
	}
	return vr
}

//think about how to implement anonymous function case
func readVariableDataValue() string {
	dataValue := ""
	for token := next(); token != ";"; token = next() {
		dataValue += token
	}
	return dataValue
}

//think about how to implement anonymous function case
func readVariableDataType() (string, string) {
	dataType := ""
	var token string
	for token = next(); token != "=" && token != ";"; token = next() {
		dataType += token
	}
	return dataType, token
}
func readInterfaces(class *class) string {
	token := next()
	for ; token != "{"; token = next() {
		if token == "," {
			continue
		}
		class.implementedClasses = append(class.implementedClasses, token)
	}
	return token
}

func readFunctionReturnType() string {
	returnType := ""
	for token := next(); token != "{"; token = next() {
		returnType += token
	}
	return returnType
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
		for true {
			token = next()
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
		fun.params = append(fun.params, prm)
	}
}

func readFunction(name string) function {
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
		fun.rtype = readFunctionReturnType()
	}
	fun.sbody = getFunctionBodyToString()
	return fun
}

func getFunctionBodyToString() string {
	open := 1
	body := ""
	for token := next(); token != "}" && open > 0; token = next() {
		if token == "{" {
			open++
		} else if token == "}" {
			open--
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
