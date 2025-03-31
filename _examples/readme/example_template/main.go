package main

import (
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/html"
	"os"
)

const baseHtml = `<!DOCTYPE html>
<html {{lang .}}>
<head>
{{headers .}}
</head>
<body>
{{body .}}
</body>
</html>`

var langAtt = html.Lang(aitch.DynamicValueKey("lang"))
var headers = aitch.Collection(
	html.Meta(html.Charset("utf-8")),
	html.TitleElement(aitch.DynamicValueKey("title")))
var body = aitch.Collection(
	html.P("Hello, ", aitch.DynamicValueKey("username")))
var template = aitch.MustNewTemplate("index", baseHtml, aitch.NodeMap{
	"lang":    langAtt,
	"headers": headers,
	"body":    body,
})

func main() {
	data := map[string]any{
		"lang":     "en",
		"title":    "Test rendered",
		"username": "Aitch",
	}
	_ = template.Execute(os.Stdout, data, nil)
}
