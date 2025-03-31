package main

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"os"
	"strconv"
)

func idIncrement(ctx *context.Context) []byte {
	id := ctx.Data["id"].(int) + 1
	ctx.Data["id"] = id
	return []byte(strconv.Itoa(id))
}

var template = aitch.Collection(
	html.P("First", html.Id(idIncrement)),
	html.P("Second", html.Id(idIncrement)),
	html.P("Third", html.Id(idIncrement)))

func main() {
	data := map[string]any{
		"id": 0,
	}
	ctx := context.New(os.Stdout, data, nil)
	_ = template.Render(ctx)
}
