package aitch

import (
	"github.com/go-andiamo/aitch/context"
	"reflect"
)

// Iterate iterates over the result of the provided function and renders the given nodes for each item.
//
// Any attribute nodes in `nodes` are ignored.
//
// For each iteration:
//   - Sets the current item as context.Cargo.
//   - If the current item is a map[string]any, it also sets context.Data.
//
// Panics if the provided function is nil.
func Iterate(fn func(ctx *context.Context) []any, nodes ...Node) Node {
	if fn == nil {
		panic("iterate fn cannot be nil")
	}
	result := &iterator{
		nodes: make([]Node, 0, len(nodes)),
		fn:    fn,
	}
	for _, n := range nodes {
		if n.Type() != AttributeNode {
			result.nodes = append(result.nodes, n)
		}
	}
	return result
}

// IterateKey extracts a slice from context.Data[key] and renders the given nodes for each item.
//
// Any attribute nodes in `nodes` are ignored.
//
// For each iteration:
//   - Sets the current item as context.Cargo.
//   - If the current item is a map[string]any, it also sets context.Data.
//
// The key lookup supports []any values, or any typed slice (e.g., []string, []int) via reflection.
func IterateKey(key string, nodes ...Node) Node {
	result := &iterator{
		nodes: make([]Node, 0, len(nodes)),
		fn: func(ctx *context.Context) []any {
			if rv, ok := ctx.Data[key]; ok {
				if sv, ok := rv.([]any); ok {
					return sv
				} else if rrv := reflect.ValueOf(rv); rrv.Kind() == reflect.Slice && rrv.Len() > 0 {
					l := rrv.Len()
					r := make([]any, l)
					for i := 0; i < l; i++ {
						r[i] = rrv.Index(i).Interface()
					}
					return r
				}
			}
			return nil
		},
	}
	for _, n := range nodes {
		if n.Type() != AttributeNode {
			result.nodes = append(result.nodes, n)
		}
	}
	return result
}

type iterator struct {
	nodes []Node
	fn    func(ctx *context.Context) []any
}

var _ Node = (*iterator)(nil)

func (i *iterator) Render(ctx *context.Context) error {
	if ctx.Error != nil {
		return ctx.Error
	}
	if items := i.fn(ctx); len(items) > 0 {
		ictx := &context.Context{
			Error:  ctx.Error,
			Writer: ctx.Writer,
			Parent: ctx,
		}
		for _, item := range items {
			ictx.Cargo = item
			if mi, ok := item.(map[string]any); ok {
				ictx.Data = mi
			} else {
				ictx.Data = ctx.Data
			}
			for _, n := range i.nodes {
				if err := n.Render(ictx); err != nil {
					ctx.Error = err
					return err
				}
			}
		}
	}
	return nil
}

func (i *iterator) Type() NodeType {
	return iteratorNode
}

func (i *iterator) Name() string {
	return "!iterator"
}

// IterateYield creates a node that yields items one at a time using a yield-style function.
// The context.Cargo is set for each item, and if the item is a map[string]any, it is also set as context.Data.
//
// Example:
//
//	IterateYield(func(ctx *Context, yield func(any)) {
//		for _, v := range []string{"a", "b", "c"} {
//			yield(v)
//		}
//	}, P(Id("foo"), Text(DynamicValueKey("."))))
func IterateYield(fn func(ctx *context.Context, yield func(any)), nodes ...Node) Node {
	if fn == nil {
		panic("iterateYield fn cannot be nil")
	}
	result := &yieldIterator{
		fn:    fn,
		nodes: make([]Node, 0, len(nodes)),
	}
	for _, n := range nodes {
		if n.Type() != AttributeNode {
			result.nodes = append(result.nodes, n)
		}
	}
	return result
}

type yieldIterator struct {
	fn    func(ctx *context.Context, yield func(any))
	nodes []Node
}

var _ Node = (*yieldIterator)(nil)

func (i *yieldIterator) Render(ctx *context.Context) error {
	if ctx.Error != nil {
		return ctx.Error
	}
	ictx := &context.Context{
		Error:  ctx.Error,
		Writer: ctx.Writer,
		Parent: ctx,
	}
	i.fn(ctx, func(item any) {
		ictx.Cargo = item
		if mi, ok := item.(map[string]any); ok {
			ictx.Data = mi
		}
		for _, n := range i.nodes {
			if err := n.Render(ictx); err != nil {
				ctx.Error = err
				return
			}
		}
	})
	return ctx.Error
}

func (i *yieldIterator) Type() NodeType {
	return iteratorNode
}

func (i *yieldIterator) Name() string {
	return "!yield_iterator"
}
