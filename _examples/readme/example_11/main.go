package main

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/go-andiamo/aitch/html"
	"github.com/go-andiamo/aitch/htmx"
	"os"
)

var template = html.Button(
	html.Type("button"),
	aitch.When("enabled",
		htmx.Get("/action"),
		htmx.Trigger(htmx.TriggerClick),
		htmx.Swap(htmx.SwapOuterHTML),
	),
	aitch.When("!enabled",
		html.Disabled(),
	),
	"Click Me!",
)

func main() {
	data := map[string]any{
		"enabled": true,
	}
	ctx := context.New(os.Stdout, data, nil)
	_ = template.Render(ctx)
}
