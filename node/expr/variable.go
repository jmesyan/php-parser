package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Variable) Name() string {
	return "Variable"
}

type Variable struct {
	name    string
	varName node.Node
}

func NewVariable(varName node.Node) node.Node {
	return Variable{
		"Variable",
		varName,
	}
}

func (n Variable) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.varName != nil {
		vv := v.Children("varName")
		n.varName.Walk(vv)
	}
}
