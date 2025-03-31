package html

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
)

type document struct {
	prologue aitch.Node
	contents []aitch.Node
}

func (d *document) Render(ctx *context.Context) error {
	if d.prologue != nil {
		_ = d.prologue.Render(ctx)
	}
	for _, n := range d.contents {
		_ = n.Render(ctx)
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
func Document(lang any, head []any, body []any) aitch.Node {
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
