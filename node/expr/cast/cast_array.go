package cast

import (
	"github.com/z7zmey/php-parser/node"
)

func (n CastArray) Name() string {
	return "CastArray"
}

type CastArray struct {
	Cast
}

func NewCastArray(expr node.Node) node.Node {
	return CastArray{
		Cast{
			"CastArray",
			expr,
		},
	}
}

func (n CastArray) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.expr != nil {
		vv := v.Children("expr")
		n.expr.Walk(vv)
	}
}
