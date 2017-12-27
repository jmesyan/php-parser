package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n AltElseIf) Name() string {
	return "AltElseIf"
}

type AltElseIf struct {
	name  string
	token token.Token
	cond  node.Node
	stmt  node.Node
}

func NewAltElseIf(token token.Token, cond node.Node, stmt node.Node) node.Node {
	return AltElseIf{
		"AltElseIf",
		token,
		cond,
		stmt,
	}
}

func (n AltElseIf) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.cond != nil {
		vv := v.Children("cond")
		n.cond.Walk(vv)
	}

	if n.stmt != nil {
		vv := v.Children("stmt")
		n.stmt.Walk(vv)
	}
}
