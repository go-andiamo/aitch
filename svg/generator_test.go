package svg

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestGenerateCode(t *testing.T) {
	svgData := `<svg width="200" height="200" viewBox="0 0 200 200" xmlns="http://www.w3.org/2000/svg">
	<!-- comment -->
	<unknown unknown="">text</unknown><unknown/>
	<circle cx="100" cy="100" r="50" stroke="black" stroke-width="3" fill="red"/>
	<path d="M 50 150 Q 100 50, 150 150" stroke="blue" stroke-width="3" fill="none"/>
</svg>`
	var writer bytes.Buffer
	err := GenerateCode(strings.NewReader(svgData), &writer)
	require.NoError(t, err)
	require.Equal(t, `var template = aitch.Collection(
	svg.Svg(svg.Width("200"), svg.Height("200"), aitch.Attribute("viewbox" ,"0 0 200 200"), aitch.Attribute("xmlns" ,"http://www.w3.org/2000/svg"),
		svg.Comment(" comment "),
		aitch.Element("unknown", aitch.Attribute("unknown" ,""),
			svg.Text("text"),
		),
		aitch.Element("unknown", ),
		svg.Circle(svg.Cx("100"), svg.Cy("100"), svg.R("50"), svg.Stroke("black"), svg.StrokeWidth("3"), svg.Fill("red")),
		svg.Path(svg.D("M 50 150 Q 100 50, 150 150"), svg.Stroke("blue"), svg.StrokeWidth("3"), svg.Fill("none")),
	),
)

`, writer.String())

	err = GenerateCode(strings.NewReader(svgData), nil)
	require.NoError(t, err)
}
