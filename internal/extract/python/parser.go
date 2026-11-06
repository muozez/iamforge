package python

import (
	"context"

	"github.com/muozez/iamforge/internal/extract"
	"github.com/odvcencio/gotreesitter"
	pyGrammar "github.com/odvcencio/gotreesitter/grammars/python"
)

// Parser implements extract.Parser using a pure-Go tree-sitter Python grammar.
type Parser struct {
	lang *gotreesitter.Language
}

// NewParser creates a new pure-Go Python parser.
func NewParser() *Parser {
	return &Parser{
		lang: pyGrammar.Language(),
	}
}

// Parse parses Python source code into an abstract syntax tree.
func (p *Parser) Parse(ctx context.Context, source []byte) (extract.Tree, error) {
	tsParser := gotreesitter.NewParser(p.lang)
	tree, err := tsParser.Parse(source)
	if err != nil {
		return nil, err
	}
	return &treeWrapper{
		tree:   tree,
		lang:   p.lang,
		source: source,
	}, nil
}

type treeWrapper struct {
	tree   *gotreesitter.Tree
	lang   *gotreesitter.Language
	source []byte
}

func (t *treeWrapper) RootNode() extract.Node {
	if t.tree == nil {
		return nil
	}
	return &nodeWrapper{
		node:   t.tree.RootNode(),
		lang:   t.lang,
		source: t.source,
	}
}

func (t *treeWrapper) Source() []byte {
	return t.source
}

func (t *treeWrapper) Release() {
	if t.tree != nil {
		t.tree.Release()
	}
}

type nodeWrapper struct {
	node   *gotreesitter.Node
	lang   *gotreesitter.Language
	source []byte
}

func (n *nodeWrapper) Type() string {
	if n.node == nil {
		return ""
	}
	return n.node.Type(n.lang)
}

func (n *nodeWrapper) Text() string {
	if n.node == nil {
		return ""
	}
	return n.node.Text(n.source)
}

func (n *nodeWrapper) ChildCount() int {
	if n.node == nil {
		return 0
	}
	return n.node.ChildCount()
}

func (n *nodeWrapper) Child(i int) extract.Node {
	if n.node == nil {
		return nil
	}
	ch := n.node.Child(i)
	if ch == nil {
		return nil
	}
	return &nodeWrapper{node: ch, lang: n.lang, source: n.source}
}

func (n *nodeWrapper) NamedChildCount() int {
	if n.node == nil {
		return 0
	}
	return n.node.NamedChildCount()
}

func (n *nodeWrapper) NamedChild(i int) extract.Node {
	if n.node == nil {
		return nil
	}
	ch := n.node.NamedChild(i)
	if ch == nil {
		return nil
	}
	return &nodeWrapper{node: ch, lang: n.lang, source: n.source}
}

func (n *nodeWrapper) ChildByFieldName(name string) extract.Node {
	if n.node == nil {
		return nil
	}
	ch := n.node.ChildByFieldName(name, n.lang)
	if ch == nil {
		return nil
	}
	return &nodeWrapper{node: ch, lang: n.lang, source: n.source}
}

func (n *nodeWrapper) HasError() bool {
	if n.node == nil {
		return false
	}
	return n.node.HasError()
}

func (n *nodeWrapper) IsError() bool {
	if n.node == nil {
		return false
	}
	return n.node.IsError()
}

func (n *nodeWrapper) StartByte() uint32 {
	if n.node == nil {
		return 0
	}
	return n.node.StartByte()
}

func (n *nodeWrapper) EndByte() uint32 {
	if n.node == nil {
		return 0
	}
	return n.node.EndByte()
}

func (n *nodeWrapper) StartPoint() extract.Point {
	if n.node == nil {
		return extract.Point{}
	}
	pt := n.node.StartPoint()
	return extract.Point{Row: pt.Row, Column: pt.Column}
}

func (n *nodeWrapper) EndPoint() extract.Point {
	if n.node == nil {
		return extract.Point{}
	}
	pt := n.node.EndPoint()
	return extract.Point{Row: pt.Row, Column: pt.Column}
}

