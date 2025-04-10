package css

import (
	"bytes"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func testRender(node aitch.Node, t *testing.T) string {
	var w bytes.Buffer
	err := node.Render(&context.Context{Writer: &w})
	require.NoError(t, err)
	return w.String()
}

func TestPx(t *testing.T) {
	a := Width(Px(10))
	assert.Equal(t, ` style="width:10px"`, testRender(a, t))
}

func TestPt(t *testing.T) {
	a := Width(Pt(10))
	assert.Equal(t, ` style="width:10pt"`, testRender(a, t))
}

func TestPc(t *testing.T) {
	a := Width(Pc(10))
	assert.Equal(t, ` style="width:10pc"`, testRender(a, t))
}

func TestIn(t *testing.T) {
	a := Width(In(10))
	assert.Equal(t, ` style="width:10in"`, testRender(a, t))
}

func TestCm(t *testing.T) {
	a := Width(Cm(10))
	assert.Equal(t, ` style="width:10cm"`, testRender(a, t))
}

func TestMm(t *testing.T) {
	a := Width(Mm(10))
	assert.Equal(t, ` style="width:10mm"`, testRender(a, t))
}

func TestEm(t *testing.T) {
	a := Width(Em(10))
	assert.Equal(t, ` style="width:10em"`, testRender(a, t))
}

func TestRem(t *testing.T) {
	a := Width(Rem(10))
	assert.Equal(t, ` style="width:10rem"`, testRender(a, t))
}

func TestPercent(t *testing.T) {
	a := Width(Percent(10))
	assert.Equal(t, ` style="width:10%"`, testRender(a, t))
}

func TestVw(t *testing.T) {
	a := Width(Vw(10))
	assert.Equal(t, ` style="width:10vw"`, testRender(a, t))
}

func TestVh(t *testing.T) {
	a := Width(Vh(10))
	assert.Equal(t, ` style="width:10vh"`, testRender(a, t))
}

func TestVMin(t *testing.T) {
	a := Width(VMin(10))
	assert.Equal(t, ` style="width:10vmin"`, testRender(a, t))
}

func TestVMax(t *testing.T) {
	a := Width(VMax(10))
	assert.Equal(t, ` style="width:10vmax"`, testRender(a, t))
}

func TestCh(t *testing.T) {
	a := Width(Ch(10))
	assert.Equal(t, ` style="width:10ch"`, testRender(a, t))
}

func TestEx(t *testing.T) {
	a := Width(Ex(10))
	assert.Equal(t, ` style="width:10ex"`, testRender(a, t))
}
