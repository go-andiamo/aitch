package htmx

import (
	"bytes"
	"github.com/go-andiamo/aitch"
	"github.com/go-andiamo/aitch/context"
	"github.com/stretchr/testify/require"
	"testing"
)

func testRender(node aitch.Node, t *testing.T) string {
	var w bytes.Buffer
	err := node.Render(&context.Context{Writer: &w})
	require.NoError(t, err)
	return w.String()
}

func TestGet(t *testing.T) {
	a := Get("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-get", a.Name())
	require.Equal(t, ` hx-get="foo"`, testRender(a, t))
}

func TestPost(t *testing.T) {
	a := Post("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-post", a.Name())
	require.Equal(t, ` hx-post="foo"`, testRender(a, t))
}

func TestPut(t *testing.T) {
	a := Put("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-put", a.Name())
	require.Equal(t, ` hx-put="foo"`, testRender(a, t))
}

func TestDelete(t *testing.T) {
	a := Delete("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-delete", a.Name())
	require.Equal(t, ` hx-delete="foo"`, testRender(a, t))
}

func TestPatch(t *testing.T) {
	a := Patch("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-patch", a.Name())
	require.Equal(t, ` hx-patch="foo"`, testRender(a, t))
}

func TestTarget(t *testing.T) {
	a := Target("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-target", a.Name())
	require.Equal(t, ` hx-target="foo"`, testRender(a, t))
}

func TestSwap(t *testing.T) {
	a := Swap("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-swap", a.Name())
	require.Equal(t, ` hx-swap="foo"`, testRender(a, t))
}

func TestSwapOob(t *testing.T) {
	a := SwapOob("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-swap-oob", a.Name())
	require.Equal(t, ` hx-swap-oob="foo"`, testRender(a, t))
}

func TestTrigger(t *testing.T) {
	a := Trigger("click")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-trigger", a.Name())
	require.Equal(t, ` hx-trigger="click"`, testRender(a, t))
}

func TestTrigger_Merge(t *testing.T) {
	e := aitch.Element("div", Trigger("click"), Trigger("submit"))
	require.Equal(t, `<div hx-trigger="click submit"></div>`, testRender(e, t))
}

func TestInclude(t *testing.T) {
	a := Include("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-include", a.Name())
	require.Equal(t, ` hx-include="foo"`, testRender(a, t))
}

func TestSelect(t *testing.T) {
	a := Select("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-select", a.Name())
	require.Equal(t, ` hx-select="foo"`, testRender(a, t))
}

func TestParams(t *testing.T) {
	a := Params("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-params", a.Name())
	require.Equal(t, ` hx-params="foo"`, testRender(a, t))
}

func TestPushURL(t *testing.T) {
	a := PushURL("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-push-url", a.Name())
	require.Equal(t, ` hx-push-url="foo"`, testRender(a, t))
}

func TestReplaceURL(t *testing.T) {
	a := ReplaceURL("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-replace-url", a.Name())
	require.Equal(t, ` hx-replace-url="foo"`, testRender(a, t))
}

func TestConfirm(t *testing.T) {
	a := Confirm("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-confirm", a.Name())
	require.Equal(t, ` hx-confirm="foo"`, testRender(a, t))
}

func TestDisable(t *testing.T) {
	a := Disable("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-disable", a.Name())
	require.Equal(t, ` hx-disable="foo"`, testRender(a, t))
}

func TestDisabledElt(t *testing.T) {
	a := DisabledElt("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-disabled-elt", a.Name())
	require.Equal(t, ` hx-disabled-elt="foo"`, testRender(a, t))
}

func TestBoost(t *testing.T) {
	a := Boost("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-boost", a.Name())
	require.Equal(t, ` hx-boost="foo"`, testRender(a, t))
}

func TestExt(t *testing.T) {
	a := Ext("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-ext", a.Name())
	require.Equal(t, ` hx-ext="foo"`, testRender(a, t))
}

func TestEncoding(t *testing.T) {
	a := Encoding("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-encoding", a.Name())
	require.Equal(t, ` hx-encoding="foo"`, testRender(a, t))
}

func TestHistoryElt(t *testing.T) {
	a := HistoryElt("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-history-elt", a.Name())
	require.Equal(t, ` hx-history-elt="foo"`, testRender(a, t))
}

func TestIndicator(t *testing.T) {
	a := Indicator("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-indicator", a.Name())
	require.Equal(t, ` hx-indicator="foo"`, testRender(a, t))
}

func TestRequest(t *testing.T) {
	a := Request("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-request", a.Name())
	require.Equal(t, ` hx-request="foo"`, testRender(a, t))
}

func TestOn(t *testing.T) {
	a := On("foo", "bar")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:foo", a.Name())
	require.Equal(t, ` hx-on:foo="bar"`, testRender(a, t))
}

func TestOnBeforeRequest(t *testing.T) {
	a := OnBeforeRequest("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:before-request", a.Name())
	require.Equal(t, ` hx-on:before-request="foo"`, testRender(a, t))
}

func TestOnSend(t *testing.T) {
	a := OnSend("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:send", a.Name())
	require.Equal(t, ` hx-on:send="foo"`, testRender(a, t))
}

func TestOnAfterRequest(t *testing.T) {
	a := OnAfterRequest("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:after-request", a.Name())
	require.Equal(t, ` hx-on:after-request="foo"`, testRender(a, t))
}

func TestOnBeforeSwap(t *testing.T) {
	a := OnBeforeSwap("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:before-swap", a.Name())
	require.Equal(t, ` hx-on:before-swap="foo"`, testRender(a, t))
}

func TestOnAfterSwap(t *testing.T) {
	a := OnAfterSwap("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:after-swap", a.Name())
	require.Equal(t, ` hx-on:after-swap="foo"`, testRender(a, t))
}

func TestOnResponseError(t *testing.T) {
	a := OnResponseError("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:response-error", a.Name())
	require.Equal(t, ` hx-on:response-error="foo"`, testRender(a, t))
}

func TestOnSendError(t *testing.T) {
	a := OnSendError("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-on:send-error", a.Name())
	require.Equal(t, ` hx-on:send-error="foo"`, testRender(a, t))
}

func TestVals(t *testing.T) {
	a := Vals("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-vals", a.Name())
	require.Equal(t, ` hx-vals="foo"`, testRender(a, t))
}

func TestVars(t *testing.T) {
	a := Vars("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-vars", a.Name())
	require.Equal(t, ` hx-vars="foo"`, testRender(a, t))
}

func TestHeaders(t *testing.T) {
	a := Headers("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-headers", a.Name())
	require.Equal(t, ` hx-headers="foo"`, testRender(a, t))
}

func TestInherit(t *testing.T) {
	a := Inherit("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-inherit", a.Name())
	require.Equal(t, ` hx-inherit="foo"`, testRender(a, t))
}

func TestInherit_Merge(t *testing.T) {
	e := aitch.Element("div", Inherit(AttrSync), Inherit(AttrTrigger))
	require.Equal(t, `<div hx-inherit="hx-sync hx-trigger"></div>`, testRender(e, t))
}

func TestDisinherit(t *testing.T) {
	a := Disinherit("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-disinherit", a.Name())
	require.Equal(t, ` hx-disinherit="foo"`, testRender(a, t))
}

func TestDisinherit_Merge(t *testing.T) {
	e := aitch.Element("div", Disinherit(AttrSync), Disinherit(AttrTrigger))
	require.Equal(t, `<div hx-disinherit="hx-sync hx-trigger"></div>`, testRender(e, t))
}

func TestSync(t *testing.T) {
	a := Sync("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-sync", a.Name())
	require.Equal(t, ` hx-sync="foo"`, testRender(a, t))
}

func TestPreserve(t *testing.T) {
	a := Preserve("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-preserve", a.Name())
	require.Equal(t, ` hx-preserve="foo"`, testRender(a, t))
}

func TestPrompt(t *testing.T) {
	a := Prompt("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-prompt", a.Name())
	require.Equal(t, ` hx-prompt="foo"`, testRender(a, t))
}

func TestValidate(t *testing.T) {
	a := Validate("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hx-validate", a.Name())
	require.Equal(t, ` hx-validate="foo"`, testRender(a, t))
}
