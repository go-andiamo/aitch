package aitch

import (
	"bytes"
	"fmt"
	"github.com/go-andiamo/aitch/context"
	"strconv"
	"strings"
)

type valuesNode interface {
	getValues() []Value
}

type Value struct {
	value       []byte
	dynamicFunc DynamicValueFunc
	concats     []Value
}

// NewValue creates a new Value for rendering
func NewValue(value any) Value {
	if values := newValue(value); len(values) == 1 {
		return values[0]
	} else {
		return Value{nil, nil, values}
	}
}

// NewConcatValue creates a new concatenated Value for rendering
func NewConcatValue(values ...any) Value {
	return Value{nil, nil, newValues(values...)}
}

func (v Value) render(ctx *context.Context) (bool, error) {
	var l int
	if v.dynamicFunc != nil {
		data := v.dynamicFunc(ctx)
		l, ctx.Error = ctx.Writer.Write(data)
	} else if len(v.concats) > 0 {
		var written bool
		for _, concat := range v.concats {
			if written, ctx.Error = concat.render(ctx); written {
				l++
			}
		}
	} else {
		l, ctx.Error = ctx.Writer.Write(v.value)
	}
	return l > 0, ctx.Error
}

// DynamicValueFunc is a func that can be called to determine actual byte data to be written
//
// A DynamicValueFunc is typically used as a content arg passed to an Attribute, Text, Comment, Fragment etc.
type DynamicValueFunc func(ctx *context.Context) []byte

// DynamicValueKey is a key into the [context.Context.Data] and is used to provide a dynamic
// Value that is read from the context
//
// A DynamicValueKey is typically used as a content arg passed to an Attribute, Text, Comment, Fragment etc.
//
// Note: If the key is "." - the value is the context.Cargo
type DynamicValueKey string

func newValues(contents ...any) (result []Value) {
	result = make([]Value, 0, len(contents))
	for _, v := range contents {
		if v != nil {
			result = append(result, newValue(v)...)
		}
	}
	return result
}

func newValue(v any) []Value {
	if v == nil {
		return nil
	}
	switch vt := v.(type) {
	case Value:
		return []Value{vt}
	case DynamicValueFunc:
		return []Value{{nil, vt, nil}}
	case func(*context.Context) []byte:
		return []Value{{nil, vt, nil}}
	case DynamicValueKey:
		if vt == "." {
			return []Value{{nil, func(ctx *context.Context) (result []byte) {
				if ctx != nil {
					result = valueToBytes(ctx.Cargo)
				}
				return
			}, nil}}
		}
		return []Value{{nil, func(ctx *context.Context) []byte {
			if ctx != nil {
				if dv, ok := ctx.Data[string(vt)]; ok {
					return valueToBytes(dv)
				}
			}
			return nil
		}, nil}}
	case valuesNode:
		return vt.getValues()
	case Node:
		switch vt.Type() {
		case collectionNode:
			if cn, ok := vt.(*collection); ok {
				result := make([]Value, 0, len(cn.nodes))
				for _, node := range cn.nodes {
					result = append(result, newValue(node)...)
				}
				return result
			}
		}
		return nil
	default:
		return []Value{{valueToBytes(v), nil, nil}}
	}
}

func valueToBytes(v any) []byte {
	switch vt := v.(type) {
	case bool:
		if vt {
			return fixedTrue
		} else {
			return fixedFalse
		}
	case []byte:
		return vt
	case string:
		return htmlEscapeString(vt)
	case int:
		return []byte(strconv.FormatInt(int64(vt), 10))
	case int8:
		return []byte(strconv.FormatInt(int64(vt), 10))
	case int16:
		return []byte(strconv.FormatInt(int64(vt), 10))
	case int32:
		return []byte(strconv.FormatInt(int64(vt), 10))
	case int64:
		return []byte(strconv.FormatInt(vt, 10))
	case uint:
		return []byte(strconv.FormatUint(uint64(vt), 10))
	case uint8:
		return []byte(strconv.FormatUint(uint64(vt), 10))
	case uint16:
		return []byte(strconv.FormatUint(uint64(vt), 10))
	case uint32:
		return []byte(strconv.FormatUint(uint64(vt), 10))
	case uint64:
		return []byte(strconv.FormatUint(vt, 10))
	case float32:
		return []byte(strconv.FormatFloat(float64(vt), 'g', -1, 32))
	case float64:
		return []byte(strconv.FormatFloat(vt, 'g', -1, 64))
	case fmt.Stringer:
		return []byte(vt.String())
	case func() any:
		return valueToBytes(vt())
	}
	return nil
}

var (
	htmlQuot = []byte("&quot;")
	htmlApos = []byte("&apos;")
	htmlAmp  = []byte("&amp;")
	htmlLt   = []byte("&lt;")
	htmlGt   = []byte("&gt;")
	htmlNull = []byte{}
)

func htmlEscapeString(s string) []byte {
	if !strings.ContainsAny(s, "'\"&<>\000") {
		return []byte(s)
	}

	sb := []byte(s)
	var b bytes.Buffer
	last := 0
	for i, c := range sb {
		var html []byte
		switch c {
		case '\000':
			html = htmlNull
		case '"':
			html = htmlQuot
		case '\'':
			html = htmlApos
		case '&':
			html = htmlAmp
		case '<':
			html = htmlLt
		case '>':
			html = htmlGt
		default:
			continue
		}
		_, _ = b.Write(sb[last:i])
		_, _ = b.Write(html)
		last = i + 1
	}
	_, _ = b.Write(sb[last:])
	return b.Bytes()
}
