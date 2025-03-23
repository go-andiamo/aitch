package svg

import "github.com/go-andiamo/aitch"

// Text creates a new text Node
func Text(contents ...any) aitch.Node {
	return aitch.Text(contents...)
}

// DynamicValueFunc is a func that can be called to determine actual byte data to be written
//
// A DynamicValueFunc is typically used as a content arg passed to an Attribute, Text, Comment, Fragment etc.
type DynamicValueFunc = aitch.DynamicValueFunc

// DynamicValueKey is a key into the [context.Context.Data] and is used to provide a dynamic
// value that is read from the context
//
// A DynamicValueKey is typically used as a content arg passed to an Attribute, Text, Comment, Fragment etc.
type DynamicValueKey = aitch.DynamicValueKey
