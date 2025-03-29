package aitch

import (
	"bytes"
	"github.com/go-andiamo/aitch/context"
	"html/template"
	"io"
)

// Template is a wrapper around a html/template Template
type Template struct {
	Template *template.Template
	nodeMap  NodeMap
}

// Execute applies a parsed template to the specified data object,
// writing the output to wr.
// If an error occurs executing the template or writing its output,
// execution stops, but partial results may already have been written to
// the output writer.
// A template may be executed safely in parallel, although if parallel
// executions share a Writer the output may be interleaved.
func (t *Template) Execute(wr io.Writer, data any, ctxData ...map[string]any) error {
	ctx := templateContext(wr, data, ctxData)
	var tmp *template.Template
	var err error
	if tmp, err = t.Template.Clone(); err == nil {
		err = tmp.Funcs(writerFuncMap(t.nodeMap, ctx)).Execute(wr, data)
	}
	return err
}

func templateContext(wr io.Writer, data any, ctxData []map[string]any) *context.Context {
	useData := map[string]any{}
	if mData, ok := data.(map[string]any); ok {
		for k, v := range mData {
			useData[k] = v
		}
	}
	for _, m := range ctxData {
		for k, v := range m {
			useData[k] = v
		}
	}
	return &context.Context{
		Data:   useData,
		Writer: wr,
	}
}

func writerFuncMap(nodeMap NodeMap, ctx *context.Context) template.FuncMap {
	result := make(template.FuncMap, len(nodeMap))
	for k, v := range nodeMap {
		if v != nil {
			result[k] = writerFunc(v, ctx)
		}
	}
	return result
}

func writerFunc(node Node, ctx *context.Context) any {
	if node.Type() == AttributeNode {
		return func() (template.HTMLAttr, error) {
			err := node.Render(ctx)
			return "", err
		}
	} else {
		return func() (template.HTML, error) {
			err := node.Render(ctx)
			return "", err
		}
	}
}

// NodeMap is a map of Node's - where the key is the {{token}} marker in the template
type NodeMap map[string]Node

// NewTemplate creates a new Template with the given name and text to parse
//
// the NodeMap is used to map {{token}} to Node's that will replace them
func NewTemplate(name string, text string, nodeMap NodeMap) (*Template, error) {
	var err error
	var tmp *template.Template
	var result *Template
	if tmp, err = template.New(name).Funcs(rawFuncMap(nodeMap)).Parse(text); err == nil {
		result = &Template{
			Template: tmp,
			nodeMap:  nodeMap,
		}
	}
	return result, err
}

func rawFuncMap(nodeMap NodeMap) template.FuncMap {
	result := make(template.FuncMap, len(nodeMap))
	for k, v := range nodeMap {
		if v != nil {
			result[k] = rawFunc(v)
		}
	}
	return result
}

func rawFunc(node Node) any {
	if node.Type() == AttributeNode {
		return func() (template.HTMLAttr, error) {
			var w bytes.Buffer
			err := node.Render(&context.Context{Writer: &w})
			return template.HTMLAttr(w.String()), err
		}
	} else {
		return func() (template.HTML, error) {
			var w bytes.Buffer
			err := node.Render(&context.Context{Writer: &w})
			return template.HTML(w.String()), err
		}
	}
}
