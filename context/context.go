package context

import "io"

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
}

// Write writes the byte data to the context writer
func (c *Context) Write(data []byte) {
	if c.Error == nil {
		_, c.Error = c.Writer.Write(data)
	}
}
