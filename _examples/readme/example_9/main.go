package main

import (
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"os"
)

var template = html.P(
	"Hello ", []byte("there "), true, " ", 1, " ", 1.2, " ", uint16(3), func() any { return " something" },
	html.Class(true), html.Class(1), html.Class(2.2))

func main() {
	ctx := context.New(os.Stdout, nil, nil)
	_ = template.Render(ctx)
}
