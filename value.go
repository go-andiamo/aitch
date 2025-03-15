package aitch

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"strconv"
	"strings"
)

type value struct {
	value       []byte
	dynamicFunc DynamicFunc
}

func (v value) render(w io.Writer, ctx *Context) (bool, error) {
	var l int
	if v.dynamicFunc != nil {
		data := v.dynamicFunc(ctx)
		l, ctx.Error = w.Write(data)
	} else {
		l, ctx.Error = w.Write(v.value)
	}
	return l > 0, ctx.Error
}

type DynamicFunc func(ctx *Context) []byte

type DynamicKey string

func newValue(v any) []value {
	switch vt := v.(type) {
	case DynamicKey:
		return []value{{nil, func(ctx *Context) []byte {
			if ctx != nil {
				if dv, ok := ctx.Data[string(vt)]; ok {
					return valueToBytes(dv)
				}
			}
			return nil
		}}}
	case Node:
		switch vt.Type() {
		case TextNode:
			if tn, ok := vt.(*text); ok {
				return tn.values
			}
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
		return HTMLEscapeString(vt)
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

func HTMLEscapeString(s string) []byte {
	if !strings.ContainsAny(s, "'\"&<>\000") {
		return []byte(s)
	}
	var b bytes.Buffer
	template.HTMLEscape(&b, []byte(s))
	return b.Bytes()
}
