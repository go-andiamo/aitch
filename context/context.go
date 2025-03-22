package context

import "io"

// Context is the rendering context
type Context struct {
	// Error is the first encountered error during rendering
	Error error
	// Data is the data that may be referenced during rendering to provide dynamic content
	Data map[string]any
	// Writer is the context writer
	Writer io.Writer
}

func (c *Context) Write(data []byte) {
	if c.Error == nil {
		_, c.Error = c.Writer.Write(data)
	}
}

func DefaultContext(w io.Writer) *Context {
	return &Context{
		Data:   make(map[string]any),
		Writer: w,
	}
}
