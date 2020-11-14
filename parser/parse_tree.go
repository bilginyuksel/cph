package parser

var idx int

var words = map[string]string{
	"function": "function",
	"const":    "variable",
	"let":      "variable",
	"var":      "variable",
}

var punc map[string]interface{}
var step4 func(string, *TreeNode) *TreeNode
var step5 func(string, *TreeNode) *TreeNode
var step6 func(string, *TreeNode) *TreeNode

func init() {
	step4 = func(token string, node *TreeNode) *TreeNode {
		node.DataType = token
		value, _ := punc[next()]
		return value.(func(string, *TreeNode) *TreeNode)(next(), node)
	}
	step5 = func(token string, node *TreeNode) *TreeNode {
		node.DataValue = token
		next()
		return step6(next(), node)
	}

	step6 = func(token string, node *TreeNode) *TreeNode {
		if !has() {
			return node
		}
		step1(token, root)
		return node
	}
	punc = make(map[string]interface{})
	punc[":"] = step4
	punc["="] = step5
	punc[";"] = step6
}

var root = &TreeNode{}

func step1(token string, node *TreeNode) {
	node.Children = append(node.Children, *step2(token))
}
func step2(token string) *TreeNode {
	node := &TreeNode{Value: token, Type: words[token]}
	return step3(next(), node)
}

func step3(name string, parent *TreeNode) *TreeNode {
	parent.Name = name
	token := next()
	value, _ := punc[token]
	return value.(func(string, *TreeNode) *TreeNode)(next(), parent)
}

func has() bool {
	return idx < len(tokens)
}

func Parse() {
	step1(next(), root)
}

func next() string {
	if !has() {
		panic("End of the stream.")
	}
	token := tokens[idx]
	idx++
	return token
}

type TreeNode struct {
	Children  []TreeNode
	Value     string
	Type      string
	Name      string
	DataType  string
	DataValue string
}
