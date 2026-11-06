package extract

import (
	"context"
)

// Point represents a position in source code (0-indexed line and column).
type Point struct {
	Row    uint32
	Column uint32
}

// Node provides an abstract view over an AST node.
type Node interface {
	Type() string
	Text() string
	ChildCount() int
	Child(i int) Node
	NamedChildCount() int
	NamedChild(i int) Node
	ChildByFieldName(name string) Node
	HasError() bool
	IsError() bool
	StartByte() uint32
	EndByte() uint32
	StartPoint() Point
	EndPoint() Point
}

// Tree represents a parsed syntax tree.
type Tree interface {
	RootNode() Node
	Source() []byte
	Release()
}

// Parser abstracts source code parsing across different language backends.
type Parser interface {
	Parse(ctx context.Context, source []byte) (Tree, error)
}

