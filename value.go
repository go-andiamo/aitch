package aitch

import (
	"bytes"
	"fmt"
	"github.com/go-andiamo/aitch/context"
	"strconv"
	"strings"
)

type valuesNode interface {
	getValues() []value
}

type value struct {
	value       []byte
	dynamicFunc DynamicFunc
}

func (v value) render(ctx *context.Context) (bool, error) {
	var l int
	if v.dynamicFunc != nil {
		data := v.dynamicFunc(ctx)
		l, ctx.Error = ctx.Writer.Write(data)
	} else {
		l, ctx.Error = ctx.Writer.Write(v.value)
	}
	return l > 0, ctx.Error
}

type DynamicFunc func(ctx *context.Context) []byte

type DynamicKey string

func newValues(contents ...any) (result []value) {
	result = make([]value, 0, len(contents))
	for _, v := range contents {
		if v != nil {
			result = append(result, newValue(v)...)
		}
	}
	return result
}

func newValue(v any) []value {
	if v == nil {
		return nil
	}
	switch vt := v.(type) {
	case DynamicFunc:
		return []value{{nil, vt}}
	case func(*context.Context) []byte:
		return []value{{nil, vt}}
	case DynamicKey:
		return []value{{nil, func(ctx *context.Context) []byte {
			if ctx != nil {
				if dv, ok := ctx.Data[string(vt)]; ok {
					return valueToBytes(dv)
				}
			}
			return nil
		}}}
	case valuesNode:
		return vt.getValues()
	case Node:
		switch vt.Type() {
		case collectionNode:
			if cn, ok := vt.(*collection); ok {
				result := make([]value, 0, len(cn.nodes))
				for _, node := range cn.nodes {
					result = append(result, newValue(node)...)
				}
				return result
			}
		}
		return nil
	default:
		return []value{{valueToBytes(v), nil}}
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
