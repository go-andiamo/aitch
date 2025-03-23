package html

import "github.com/go-andiamo/aitch"

// Text creates a new text Node
func Text(contents ...any) aitch.Node {
	return aitch.Text(contents...)
}

// Comment creates a new comment Node
func Comment(contents ...any) aitch.Node {
	return aitch.Comment(contents...)
}

// PI creates a processing instruction Node
func PI(name string, contents ...any) aitch.Node {
	return aitch.PI(name, contents...)
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
