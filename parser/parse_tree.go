package parser

// func parse() {
// 	token := next()
// 	isVariable := token == "const" || token == "var" || token == "let"
// 	if isVariable {
// 		name := next()
// 		token = next()
// 		if token == ":" {
// 			getType()
// 			token = next()
// 			if token == "=" {
// 				getValue()
// 			} else {
// 				return
// 			}
// 		} else if token == "=" {
// 			getValue()
// 		} else {
// 			return
// 		} // end of the variable
// 	}
// }

// func getValue() {
// 	// collect everything until finding an assignment operator,
// 	// then check if the next one is '>' greater than operator
// 	// if it is greater than operator continue getting valeu data
// 	// if it is not greater than operator then stop collection get value
// 	// and go for getValue:: this statement is for getType :))
// }
// func getType() {}

// func next() {
// 	if idx >= len(tokens) {
// 		panic("error")
// 	}
// 	token := tokens[idx]
// 	idx++
// 	return token
// }

// func buildTree() {
// 	root := TreeNode{}
// }

// func sampleParser(tokens []string) {
// 	for i := 0; i < len(tokens); i++ {
// 		tokens[i]
// 	}
// }

// func buildTSFunctionTrie() {
// 	trie := &TrieNode{}
// 	functionTrie := &TrieNode{End: false, Value: "function"}
// 	functionTrie.Children["<any-name>"] = TrieNode{}

// 	trie.Children["function"] = functionTrie
// }

// // design a nfa
// func functionState() {
// 	funcName := next()
// 	paranthesis := next()
// 	isParanthesis := next()
// 	back()
// 	if !isParanthesis {
// 		collectParam()
// 	}
// }

// func collectParam() {
// 	params := []string{}
// 	for next() != ")" {
// 		if params != "," {
// 		}
// 	}
// 	param1 := next()
// }

// func variableState() {
// 	return namingState()
// }

// func namingState() {

// }

// TreeNode ...
type TreeNode struct {
	Children []TreeNode
	Value    string
	Type     string
}

// TrieNode ...
type TrieNode struct {
	End      bool
	Value    string
	Children map[string]*TrieNode
}
