package aitch

import "io"

type document struct {
	prologue Node
	contents []Node
	body     Node
}

func (d *document) Render(w io.Writer, ctx *Context) error {
	if ctx == nil {
		ctx = defaultContext()
	}
	if ctx.Error == nil {
		if d.prologue != nil {
			if ctx.Error = d.prologue.Render(w, ctx); ctx.Error != nil {
				return ctx.Error
			}
		}
		for _, n := range d.contents {
			if ctx.Error = n.Render(w, ctx); ctx.Error != nil {
				return ctx.Error
			}
		}
	}
	return ctx.Error
}

func (d *document) Type() NodeType {
	return DocumentNode
}

func (d *document) Name() string {
	return "#document"
}

func (d *document) Append(nodes ...Node) Node {
	d.body.Append(nodes...)
	return d
}

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
