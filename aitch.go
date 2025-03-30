package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"regexp"
)

// PanicOnInvalidName determines whether aitch panics when creating invalid element/attribute names
//
// used by Element() and Attribute()
var PanicOnInvalidName = false

var htmlTagRegex = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9._\-:]*$`)

// Node is the interface for all nodes (elements, attributes, etc.)
type Node interface {
	// Render renders the node into the writer
	Render(ctx *context.Context) error
	// Type returns the NodeType of the Node
	Type() NodeType
	// Name returns the name of the node (e.g. the element or attribute name)
	Name() string
}

// NodeType is the node type (e.g. ElementNode, AttributeNode etc.) for a Node
type NodeType int

const (
	ElementNode   NodeType = iota // ElementNode is the NodeType for elements
	AttributeNode                 // AttributeNode is the NodeType for attributes
	TextNode                      // TextNode is the NodeType for text nodes
	CommentNode                   // CommentNode is the NodeType for comment nodes
	PINode                        // PINode is the NodeType for processing instruction nodes (e.g. DocType prologue)
	DocumentNode                  // DocumentNode is the NodeType for a document node
	FragmentNode                  // FragmentNode is the NodeType for fragments (see Fragment())
	collectionNode
	conditionalNode
	dynamicNode
	imperativeNode
)

var (
	space             = []byte{' '}
	openAngleBracket  = []byte{'<'}
	closeAngleBracket = []byte{'>'}
	closeTagStart     = []byte{'<', '/'}
	attStart          = []byte{'=', '"'}
	attEnd            = []byte{'"'}
	openComment       = []byte{'<', '!', '-', '-'}
	closeComment      = []byte{'-', '-', '>'}
	openPi            = []byte{'<', '!'}
	closePi           = []byte{'>', '\n'}
	fixedTrue         = []byte("true")
	fixedFalse        = []byte("false")
)
