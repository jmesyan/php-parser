package expr

import (
	"github.com/z7zmey/php-parser/node"
)

func (n BooleanNot) Name() string {
	return "BooleanNot"
}

type BooleanNot struct {
	name string
	expr node.Node
}

func NewBooleanNot(expression node.Node) node.Node {
	return BooleanNot{
		"BooleanNot",
		expression,
	}
}

func (n BooleanNot) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
