package stmt

import (
	"github.com/z7zmey/php-parser/node"
	"github.com/z7zmey/php-parser/token"
)

func (n Try) Name() string {
	return "Try"
}

type Try struct {
	name    string
	token   token.Token
	stmts   []node.Node
	catches []node.Node
	finally node.Node
}

func NewTry(token token.Token, stmts []node.Node, catches []node.Node, finally node.Node) node.Node {
	return Try{
		"Try",
		token,
		stmts,
		catches,
		finally,
	}
}

func (n Try) Walk(v node.Visitor) {
	if v.Visit(n) == false {
		return
	}

	if n.stmts != nil {
		vv := v.Children("stmts")
		for _, nn := range n.stmts {
			nn.Walk(vv)
		}
	}

	if n.catches != nil {
		vv := v.Children("catches")
		for _, nn := range n.catches {
			nn.Walk(vv)
		}
	}

	if n.finally != nil {
		vv := v.Children("finally")
		n.finally.Walk(vv)
	}
}
