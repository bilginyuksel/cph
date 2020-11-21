package parser

import "fmt"

var words = map[string]string{
	"function": "function",
	"const":    "variable",
	"let":      "variable",
	"var":      "variable",
}

type param struct {
	name  string
	dtype string
}

type function struct {
	export      bool
	declare     bool
	async       bool
	annotations []annotation
	docs        docstring
	name        string
	rtype       string
	body        *fbody
	params      []param
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

func ParseLoop() {
	for has() {
		token := next()
		if _, ok := words[token]; ok {
			// read function or read variable
			fun := readFunction()
			fmt.Println(fun)
		}
	}
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

func readFunction() function {
	fun := function{}
	fun.name = next()
	next() // start paranthesis
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

	return fun
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
