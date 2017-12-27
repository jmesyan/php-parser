package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n Include) Name() string {
	return "Include"
}

type Include struct {
	name string
	expr node.Node
}

func NewInclude(expression node.Node) node.Node {
	return Include{
		"Include",
		expression,
	}
}

func (n Include) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
