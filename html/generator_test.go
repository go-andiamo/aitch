package html

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	htmlData := `<!DOCTYPE html><!-- this is 
		a "comment" --><html lang="en&quot;&lt;" foo="bar" disabled><body><unknown/><unknown unknown="foo"/><br/><h1 class="foo">Hello, World!</h1></body></html>`
	var writer bytes.Buffer
	err := GenerateCode(strings.NewReader(htmlData), &writer)
	require.NoError(t, err)
	require.Equal(t, `var template = aitch.Collection(
	html.PI("DOCTYPE", "html"),
	html.Comment(" this is \n\t\ta \"comment\" "),
	html.Html(html.Lang("en\"<"), aitch.Attribute("foo" ,"bar"), html.Disabled(),
		html.Body(
			aitch.Element("unknown", ),
			aitch.Element("unknown", aitch.Attribute("unknown" ,"foo")),
			html.Br(),
			html.H1(html.Class("foo"),
				html.Text("Hello, World!"),
			),
		),
	),
)

`, writer.String())
	err = GenerateCode(strings.NewReader(htmlData), nil)
	require.NoError(t, err)
}
