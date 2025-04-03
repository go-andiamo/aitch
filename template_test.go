package aitch

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

const templateBasic = `<!DOCTYPE html>
<html {{lang .}}>
<head>
    {{title .}}
</head>
<body>
{{body .}}
</body>
</html>`

func TitleElement(contents ...any) Node {
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
	data := map[string]any{
		"lang": "en",
		"yes":  true,
	}
	require.NoError(t, err)
	var w bytes.Buffer
	err = atmp.Execute(&w, data, nil)
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

	data["yes"] = false
	w = bytes.Buffer{}
	err = atmp.Execute(&w, data, nil)
	require.NoError(t, err)
	require.Equal(t, `<!DOCTYPE html>
<html  lang="en">
<head>
    <title>Test</title>
</head>
<body>

</body>
</html>`, w.String())

	w = bytes.Buffer{}
	err = atmp.template.Execute(&w, nil)
	require.Error(t, err)
}

func TestMustNewTemplate(t *testing.T) {
	require.NotPanics(t, func() {
		_ = MustNewTemplate("test", templateBasic, NodeMap{
			"lang":  Lang(DynamicValueKey("lang")),
			"title": TitleElement(Text("Test")),
			"body":  When("yes", P(Id(1), Text("Here!"))),
		})
	})
	require.Panics(t, func() {
		_ = MustNewTemplate("test", templateBasic, NodeMap{})
	})
}
