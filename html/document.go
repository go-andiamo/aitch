package html

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"io"
)

type document struct {
	prologue aitch.Node
	contents []aitch.Node
}

func (d *document) Render(w io.Writer, ctx *context.Context) error {
	if ctx == nil {
		ctx = context.DefaultContext(w)
	} else {
		ctx.Writer = w
	}
	if d.prologue != nil {
		_ = d.prologue.Render(w, ctx)
	}
	for _, n := range d.contents {
		_ = n.Render(w, ctx)
	}
	return ctx.Error
}

func (d *document) Type() aitch.NodeType {
	return aitch.DocumentNode
}

func (d *document) Name() string {
	return "#document"
}

// Document creates a new HTML document Node
func Document(lang any, head []aitch.Node, body []aitch.Node) aitch.Node {
	var langAtt aitch.Node
	if lang != nil {
		langAtt = Lang(lang)
	}
	var headNode aitch.Node
	if len(head) > 0 {
		headNode = Head(head...)
	}
	bodyNode := Body(body...)
	return &document{
		prologue: aitch.DocType(),
		contents: []aitch.Node{Html(langAtt, headNode, bodyNode)},
	}
}
