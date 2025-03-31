package main

import (
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"os"
)

var template = html.P(
	"This is a para",
	html.Id(1), html.Class("foo"),
	html.Br(), " >>> and this is more",
	html.Id(2), html.Class("bar"))

func main() {
	ctx := context.New(os.Stdout, nil, nil)
	_ = template.Render(ctx)
}
