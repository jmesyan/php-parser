package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Require) Name() string {
	return "Require"
}

type Require struct {
	name string
	expr node.Node
}

func NewRequire(expression node.Node) node.Node {
	return Require{
		"Require",
		expression,
	}
}

func (n Require) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
