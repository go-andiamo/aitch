package aitch

import "io"

type document struct {
	prologue Node
	contents []Node
}

func (d *document) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext(w)
	} else {
		ctx.w = w
	}
	if d.prologue != nil {
		_ = d.prologue.Render(w, ctx)
	}
	for _, n := range d.contents {
		_ = n.Render(w, ctx)
	}
	return ctx.Error
}

func (d *document) Type() NodeType {
	return DocumentNode
}

func (d *document) Name() string {
	return "#document"
}

// Document creates a new HTML document Node
func Document(lang any, head []Node, body []Node) Node {
	var langAtt Node
	if lang != nil {
		langAtt = Lang(lang)
	}
	var headNode Node
	if len(head) > 0 {
		headNode = Head(head...)
	}
	bodyNode := Body(body...)
	return &document{
		prologue: DocType(),
		contents: []Node{Html(langAtt, headNode, bodyNode)},
	}
}
