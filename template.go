package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"html/template"
	"io"
)

// Template is a wrapper around a html/template.Template
type Template struct {
	template *template.Template
}

// Execute applies a parsed template to the specified data object,
// writing the output to wr.
// If an error occurs executing the template or writing its output,
// execution stops, but partial results may already have been written to
// the output writer.
// A template may be executed safely in parallel, although if parallel
// executions share a Writer the output may be interleaved.
func (t *Template) Execute(wr io.Writer, data map[string]any, cargo any) error {
	ctx := &context.Context{
		Data:   data,
		Cargo:  cargo,
		Writer: wr,
	}
	return t.template.Execute(wr, ctx)
}

// NodeMap is a map of Node's - where the key is the {{token}} marker in the template
type NodeMap map[string]Node

// NewTemplate creates a new Template with the given name and text to parse
//
// the NodeMap is used to map {{token .}} to Node's that will replace them
func NewTemplate(name string, text string, nodeMap NodeMap) (*Template, error) {
	var err error
	var tmp *template.Template
	var result *Template
	if tmp, err = template.New(name).Funcs(contextFuncMap(nodeMap)).Parse(text); err == nil {
		result = &Template{
			template: tmp,
		}
	}
	return result, err
}

func MustNewTemplate(name string, text string, nodeMap NodeMap) *Template {
	if t, err := NewTemplate(name, text, nodeMap); err == nil {
		return t
	} else {
		panic(err)
	}
}

func contextFuncMap(nodeMap NodeMap) template.FuncMap {
	result := make(template.FuncMap, len(nodeMap))
	for k, v := range nodeMap {
		if v != nil {
			result[k] = contextFunc(v)
		}
	}
	return result
}

func contextFunc(node Node) any {
	if node.Type() == AttributeNode {
		return func(ctx *context.Context) (template.HTMLAttr, error) {
			err := node.Render(ctx)
			return "", err
		}
	} else {
		return func(ctx *context.Context) (template.HTML, error) {
			err := node.Render(ctx)
			return "", err
		}
	}
}
