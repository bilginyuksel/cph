package parser

import "fmt"

var idx int

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
	export bool
	name   string
	rtype  string
	params []param
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
	for token := next(); token != "{"; {
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
	if token != ")" {
		collectParameterDataOfFunction(&fun, token)
	}
	token = next()
	if token == "{" {
		fun.rtype = "any"
	} else {
		fun.rtype = readFunctionReturnType()
	}

	return fun
}

var punc map[string]interface{}
var giveDataTypeToTreeNode func(string, *TreeNode) *TreeNode
var giveDataValueToTreeNode func(string, *TreeNode) *TreeNode
var endOfVariableParsing func(string, *TreeNode) *TreeNode
var endOrStartTheParametersOfAFunction func(string, *TreeNode) *TreeNode
var endOfFunctionParameters func(string, *TreeNode) *TreeNode

func init() {
	giveDataTypeToTreeNode = func(token string, node *TreeNode) *TreeNode {
		node.DataType = token
		value, _ := punc[next()]
		return value.(func(string, *TreeNode) *TreeNode)(next(), node)
	}
	giveDataValueToTreeNode = func(token string, node *TreeNode) *TreeNode {
		node.DataValue = token
		next()
		return endOfVariableParsing(next(), node)
	}

	endOfVariableParsing = func(token string, node *TreeNode) *TreeNode {
		if !has() {
			return node
		}
		treeInitializer(token, root)
		return node
	}

	endOrStartTheParametersOfAFunction = func(token string, node *TreeNode) *TreeNode {
		if token == ")" {
			// no parameters or end of parameters.
			// return type so we are waiting for the colon punctuation ':'
			// if no return type then we are waiting for the curly bracket '{'
			value, _ := punc[next()]
			return value.(func(string, *TreeNode) *TreeNode)(next(), node)
		}
		return startOfParameter(token, node)
	}

	punc = make(map[string]interface{})
	punc[":"] = giveDataTypeToTreeNode
	punc["="] = giveDataValueToTreeNode
	punc[";"] = endOfVariableParsing
	punc["("] = endOrStartTheParametersOfAFunction
	punc[")"] = endOfFunctionParameters
}

var root = &TreeNode{}

/*
[function, something, (, ), {, }]
function something(){}
function something(): string{}
function something(data: any): string{}
function something(data: any, data2: ()=>void): string {}
[function, something, (, data, :, any, data2, :, (, ), =, >, void, :, string, { })]
*/

/*

callback :  = > void  )
openParanthesis := 0
close openParanthesi > 0
open -= 1


*/

func startOfParameter(token string, node *TreeNode) *TreeNode {
	param := Parameter{Name: token}
	dataType := ""
	nextToken := next()
	if nextToken == ")" {
		return nil
	} else if nextToken == ":" {
		// fill the data type
		// comma, or closing paranthesis
		// any, string, ()=>void, callback:(data1, data2)=>void, (callback:(callback1:()=>void)=>void)=>void;
		// stk := []string{next()}
		// for len(stk) > 0 {
		// 	potentialToken := next()
		// 	if potentialToken == ")" {
		// 		stk = stk[:len(stk)-1]
		// 	} else {
		// 		stk = append(stk, potentialToken)
		// 	}
		// }

		// callback: (data1,data2) => void
		open := 0
		for true {
			potential := next()
			if potential == ")" && open > 0 {
				open--
				dataType += ")"
			} else if potential == ")" && open <= 0 {
				break
			} else if potential == "," && open <= 0 {
				param.DataType = dataType
				node.Parameters = append(node.Parameters, param)
				return startOfParameter(next(), node)
			} else {
				dataType += potential
			}
		}
	}
	param.DataType = dataType
	node.Parameters = append(node.Parameters, param)
	return node
}

func treeInitializer(token string, node *TreeNode) {
	node.Children = append(node.Children, *startTreeNode(token))
}
func startTreeNode(token string) *TreeNode {
	node := &TreeNode{Value: token, Type: words[token]}
	return nameTheTreeNode(next(), node)
}

func nameTheTreeNode(name string, parent *TreeNode) *TreeNode {
	parent.Name = name
	token := next()
	value, _ := punc[token]
	return value.(func(string, *TreeNode) *TreeNode)(next(), parent)
}

func has() bool {
	return idx < len(tokens)
}

// Parse ...
func Parse() {
	treeInitializer(next(), root)
}

func next() string {
	if !has() {
		panic("End of the stream.")
	}
	token := tokens[idx]
	idx++
	return token
}

// TreeNode ...
type TreeNode struct {
	Children   []TreeNode
	Value      string
	Type       string
	Name       string
	DataType   string
	DataValue  string
	Parameters []Parameter
}

// Parameter ...
type Parameter struct {
	Name     string
	DataType string
}
