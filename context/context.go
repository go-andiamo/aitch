package context

import (
	"fmt"
	"io"
	"reflect"
)

// Context is the rendering context
type Context struct {
	// Error is the first encountered error during rendering
	Error error
	// Data is the data that may be referenced during rendering to provide dynamic content
	Data map[string]any
	// Cargo is any additional data (e.g. to carry data normally passed to html/template.Execute)
	Cargo any
	// Writer is the context writer
	Writer io.Writer
	// Parent is the parent context (used when iterating)
	Parent *Context
}

func New(writer io.Writer, data map[string]any, cargo any) *Context {
	return &Context{
		Data:   data,
		Cargo:  cargo,
		Writer: writer,
	}
}

// Write writes the byte data to the context writer
func (c *Context) Write(data []byte) {
	if c.Error == nil {
		_, c.Error = c.Writer.Write(data)
	}
}

// Get is a generic function for extracting typed data from a Context
//
// # If the key exists in Context.Data (and is the correct type) it's returned
//
// It then checks Context.Cargo (as a map or struct) to see if that key/property exists (and is the correct type)
func Get[T any](ctx *Context, key string) (T, bool) {
	var zero T
	if ctx == nil {
		return zero, false
	}
	if ctx.Data != nil {
		if val, ok := ctx.Data[key]; ok {
			if typed, ok := val.(T); ok {
				return typed, true
			}
		}
	}
	if ctx.Cargo != nil {
		v := reflect.ValueOf(ctx.Cargo)
		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		switch v.Kind() {
		case reflect.Map:
			if v.Type().Key().Kind() == reflect.String {
				if val := v.MapIndex(reflect.ValueOf(key)); val.IsValid() {
					if typed, ok := val.Interface().(T); ok {
						return typed, true
					}
				}
			}
		case reflect.Struct:
			if f, ok := v.Type().FieldByName(key); ok {
				if val := v.FieldByIndex(f.Index); val.IsValid() {
					if typed, ok := val.Interface().(T); ok {
						return typed, true
					}
				}
			}
		}
	}
	return zero, false
}

// MustGet is the same as GET - except that it panics if the key isn't there or incorrect type
func MustGet[T any](ctx *Context, key string) T {
	if val, ok := Get[T](ctx, key); !ok {
		panic(fmt.Sprintf("missing or invalid context data for key %q", key))
	} else {
		return val
	}
}
