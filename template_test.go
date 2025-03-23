package aitch

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

const templateBasic = `<!DOCTYPE html>
<html {{lang}}>
<head>
    {{title}}
</head>
<body>
{{body}}
</body>
</html>`

func TitleElement(contents ...Node) Node {
	return NewElement([]byte("title"), contents...)
}

func Lang(v ...any) Node {
	return NewAttribute([]byte("lang"), v...)
}

func TestNewTemplate(t *testing.T) {
	atmp, err := NewTemplate("test", templateBasic, NodeMap{
		"lang":  Lang(DynamicValueKey("lang")),
		"title": TitleElement(Text("Test")),
		"body":  When("yes", P(Id(1), Text("Here!"))),
	})
	data1 := map[string]any{
		"lang": "fr",
	}
	data2 := map[string]any{
		"lang": "en",
		"yes":  true,
	}
	require.NoError(t, err)
	var w bytes.Buffer
	err = atmp.Execute(&w, data1, data2)
	require.NoError(t, err)
	require.Equal(t, `<!DOCTYPE html>
<html  lang="en">
<head>
    <title>Test</title>
</head>
<body>
<p id="1">Here!</p>
</body>
</html>`, w.String())

	w = bytes.Buffer{}
	err = atmp.Template.Execute(&w, data1)
	require.NoError(t, err)
	require.Equal(t, `<!DOCTYPE html>
<html  lang="">
<head>
    <title>Test</title>
</head>
<body>

</body>
</html>`, w.String())
}
