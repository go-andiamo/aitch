package main

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"os"
)

var template = html.P(
	"Hello, ", aitch.DynamicValueKey("name"),
	html.Class("greeting"), html.Class(aitch.DynamicValueKey("greeting-class")))

func main() {
	data := map[string]any{
		"name":           "Aitch!",
		"greeting-class": "admin",
	}
	ctx := context.New(os.Stdout, data, nil)
	_ = template.Render(ctx)
}
