package html

import (
	"github.com/go-andiamo/aitch"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAsync(t *testing.T) {
	a := Async()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "async", a.Name())
	require.Equal(t, ` async`, testRender(a, t))
}

func TestAutoFocus(t *testing.T) {
	a := AutoFocus()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "autofocus", a.Name())
	require.Equal(t, ` autofocus`, testRender(a, t))
}

func TestAutoPlay(t *testing.T) {
	a := AutoPlay()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "autoplay", a.Name())
	require.Equal(t, ` autoplay`, testRender(a, t))
}

func TestChecked(t *testing.T) {
	a := Checked()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "checked", a.Name())
	require.Equal(t, ` checked`, testRender(a, t))
}

func TestControls(t *testing.T) {
	a := Controls()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "controls", a.Name())
	require.Equal(t, ` controls`, testRender(a, t))
}

func TestCrossOrigin(t *testing.T) {
	a := CrossOrigin("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "crossorigin", a.Name())
	require.Equal(t, ` crossorigin="foo"`, testRender(a, t))
}

func TestDateTime(t *testing.T) {
	a := DateTime("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "datetime", a.Name())
	require.Equal(t, ` datetime="foo"`, testRender(a, t))
}

func TestDefer(t *testing.T) {
	a := Defer()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "defer", a.Name())
	require.Equal(t, ` defer`, testRender(a, t))
}

func TestDisabled(t *testing.T) {
	a := Disabled()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "disabled", a.Name())
	require.Equal(t, ` disabled`, testRender(a, t))
}

func TestDownload(t *testing.T) {
	a := Download("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "download", a.Name())
	require.Equal(t, ` download="foo"`, testRender(a, t))
}

func TestDraggable(t *testing.T) {
	a := Draggable("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "draggable", a.Name())
	require.Equal(t, ` draggable="foo"`, testRender(a, t))
}

func TestLoop(t *testing.T) {
	a := Loop()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "loop", a.Name())
	require.Equal(t, ` loop`, testRender(a, t))
}

func TestMultiple(t *testing.T) {
	a := Multiple()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "multiple", a.Name())
	require.Equal(t, ` multiple`, testRender(a, t))
}

func TestMuted(t *testing.T) {
	a := Muted()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "muted", a.Name())
	require.Equal(t, ` muted`, testRender(a, t))
}

func TestPlaysInline(t *testing.T) {
	a := PlaysInline()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "playsinline", a.Name())
	require.Equal(t, ` playsinline`, testRender(a, t))
}

func TestReadOnly(t *testing.T) {
	a := ReadOnly()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "readonly", a.Name())
	require.Equal(t, ` readonly`, testRender(a, t))
}

func TestRequired(t *testing.T) {
	a := Required()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "required", a.Name())
	require.Equal(t, ` required`, testRender(a, t))
}

func TestSelected(t *testing.T) {
	a := Selected()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "selected", a.Name())
	require.Equal(t, ` selected`, testRender(a, t))
}

func TestAccept(t *testing.T) {
	a := Accept("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "accept", a.Name())
	require.Equal(t, ` accept="foo"`, testRender(a, t))
}

func TestAction(t *testing.T) {
	a := Action("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "action", a.Name())
	require.Equal(t, ` action="foo"`, testRender(a, t))
}

func TestAlt(t *testing.T) {
	a := Alt("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "alt", a.Name())
	require.Equal(t, ` alt="foo"`, testRender(a, t))
}

func TestAria(t *testing.T) {
	a := Aria("foo", "foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-foo", a.Name())
	require.Equal(t, ` aria-foo="foo"`, testRender(a, t))
}

func TestAs(t *testing.T) {
	a := As("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "as", a.Name())
	require.Equal(t, ` as="foo"`, testRender(a, t))
}

func TestAutoComplete(t *testing.T) {
	a := AutoComplete("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "autocomplete", a.Name())
	require.Equal(t, ` autocomplete="foo"`, testRender(a, t))
}

func TestCharset(t *testing.T) {
	a := Charset("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "charset", a.Name())
	require.Equal(t, ` charset="foo"`, testRender(a, t))
}

func TestCiteAttr(t *testing.T) {
	a := CiteAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "cite", a.Name())
	require.Equal(t, ` cite="foo"`, testRender(a, t))
}

func TestClass(t *testing.T) {
	a := Class("foo", "bar")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "class", a.Name())
	require.Equal(t, ` class="foo bar"`, testRender(a, t))
}

func TestCols(t *testing.T) {
	a := Cols("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "cols", a.Name())
	require.Equal(t, ` cols="foo"`, testRender(a, t))
}

func TestColSpan(t *testing.T) {
	a := ColSpan("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "colspan", a.Name())
	require.Equal(t, ` colspan="foo"`, testRender(a, t))
}

func TestContent(t *testing.T) {
	a := Content("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "content", a.Name())
	require.Equal(t, ` content="foo"`, testRender(a, t))
}

func TestContentEditable(t *testing.T) {
	a := ContentEditable(true)
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "contenteditable", a.Name())
	require.Equal(t, ` contenteditable="true"`, testRender(a, t))
}

func TestData(t *testing.T) {
	a := Data("foo", "foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "data-foo", a.Name())
	require.Equal(t, ` data-foo="foo"`, testRender(a, t))
}

func TestDir(t *testing.T) {
	a := Dir("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "dir", a.Name())
	require.Equal(t, ` dir="foo"`, testRender(a, t))
}

func TestEncType(t *testing.T) {
	a := EncType("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "enctype", a.Name())
	require.Equal(t, ` enctype="foo"`, testRender(a, t))
}

func TestFor(t *testing.T) {
	a := For("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "for", a.Name())
	require.Equal(t, ` for="foo"`, testRender(a, t))
}

func TestFormAttr(t *testing.T) {
	a := FormAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "form", a.Name())
	require.Equal(t, ` form="foo"`, testRender(a, t))
}

func TestNoValidate(t *testing.T) {
	a := NoValidate()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "novalidate", a.Name())
	require.Equal(t, ` novalidate`, testRender(a, t))
}

func TestFormNoValidate(t *testing.T) {
	a := FormNoValidate()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "formnovalidate", a.Name())
	require.Equal(t, ` formnovalidate`, testRender(a, t))
}

func TestHeight(t *testing.T) {
	a := Height("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "height", a.Name())
	require.Equal(t, ` height="foo"`, testRender(a, t))
}

func TestHidden(t *testing.T) {
	a := Hidden()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "hidden", a.Name())
	require.Equal(t, ` hidden`, testRender(a, t))
}

func TestHref(t *testing.T) {
	a := Href("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "href", a.Name())
	require.Equal(t, ` href="foo"`, testRender(a, t))
}

func TestId(t *testing.T) {
	a := Id("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "id", a.Name())
	require.Equal(t, ` id="foo"`, testRender(a, t))
}

func TestIntegrity(t *testing.T) {
	a := Integrity("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "integrity", a.Name())
	require.Equal(t, ` integrity="foo"`, testRender(a, t))
}

func TestLabelAttr(t *testing.T) {
	a := LabelAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "label", a.Name())
	require.Equal(t, ` label="foo"`, testRender(a, t))
}

func TestLang(t *testing.T) {
	a := Lang("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "lang", a.Name())
	require.Equal(t, ` lang="foo"`, testRender(a, t))
}

func TestList(t *testing.T) {
	a := List("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "list", a.Name())
	require.Equal(t, ` list="foo"`, testRender(a, t))
}

func TestLoading(t *testing.T) {
	a := Loading("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "loading", a.Name())
	require.Equal(t, ` loading="foo"`, testRender(a, t))
}

func TestMax(t *testing.T) {
	a := Max("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "max", a.Name())
	require.Equal(t, ` max="foo"`, testRender(a, t))
}

func TestMaxLength(t *testing.T) {
	a := MaxLength("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "maxlength", a.Name())
	require.Equal(t, ` maxlength="foo"`, testRender(a, t))
}

func TestMethod(t *testing.T) {
	a := Method("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "method", a.Name())
	require.Equal(t, ` method="foo"`, testRender(a, t))
}

func TestMin(t *testing.T) {
	a := Min("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "min", a.Name())
	require.Equal(t, ` min="foo"`, testRender(a, t))
}

func TestMinLength(t *testing.T) {
	a := MinLength("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "minlength", a.Name())
	require.Equal(t, ` minlength="foo"`, testRender(a, t))
}

func TestName(t *testing.T) {
	a := Name("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "name", a.Name())
	require.Equal(t, ` name="foo"`, testRender(a, t))
}

func TestPattern(t *testing.T) {
	a := Pattern("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "pattern", a.Name())
	require.Equal(t, ` pattern="foo"`, testRender(a, t))
}

func TestPlaceholder(t *testing.T) {
	a := Placeholder("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "placeholder", a.Name())
	require.Equal(t, ` placeholder="foo"`, testRender(a, t))
}

func TestPopOver(t *testing.T) {
	a := PopOver("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "popover", a.Name())
	require.Equal(t, ` popover="foo"`, testRender(a, t))
}

func TestPopOverTarget(t *testing.T) {
	a := PopOverTarget("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "popovertarget", a.Name())
	require.Equal(t, ` popovertarget="foo"`, testRender(a, t))
}

func TestPopOverTargetAction(t *testing.T) {
	a := PopOverTargetAction("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "popovertargetaction", a.Name())
	require.Equal(t, ` popovertargetaction="foo"`, testRender(a, t))
}

func TestPoster(t *testing.T) {
	a := Poster("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "poster", a.Name())
	require.Equal(t, ` poster="foo"`, testRender(a, t))
}

func TestPreload(t *testing.T) {
	a := Preload("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "preload", a.Name())
	require.Equal(t, ` preload="foo"`, testRender(a, t))
}

func TestReferrerPolicy(t *testing.T) {
	a := ReferrerPolicy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "referrerpolicy", a.Name())
	require.Equal(t, ` referrerpolicy="foo"`, testRender(a, t))
}

func TestRel(t *testing.T) {
	a := Rel("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "rel", a.Name())
	require.Equal(t, ` rel="foo"`, testRender(a, t))
}

func TestRole(t *testing.T) {
	a := Role("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "role", a.Name())
	require.Equal(t, ` role="foo"`, testRender(a, t))
}

func TestRows(t *testing.T) {
	a := Rows("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "rows", a.Name())
	require.Equal(t, ` rows="foo"`, testRender(a, t))
}

func TestRowSpan(t *testing.T) {
	a := RowSpan("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "rowspan", a.Name())
	require.Equal(t, ` rowspan="foo"`, testRender(a, t))
}

func TestSlotAttr(t *testing.T) {
	a := SlotAttr("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "slot", a.Name())
	require.Equal(t, ` slot="foo"`, testRender(a, t))
}

func TestSrc(t *testing.T) {
	a := Src("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "src", a.Name())
	require.Equal(t, ` src="foo"`, testRender(a, t))
}

func TestSrcSet(t *testing.T) {
	a := SrcSet("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "srcset", a.Name())
	require.Equal(t, ` srcset="foo"`, testRender(a, t))
}

func TestStep(t *testing.T) {
	a := Step("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "step", a.Name())
	require.Equal(t, ` step="foo"`, testRender(a, t))
}

func TestStyle(t *testing.T) {
	a := Style("foo", "bar")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "style", a.Name())
	require.Equal(t, ` style="foo; bar"`, testRender(a, t))
}

func TestTabIndex(t *testing.T) {
	a := TabIndex("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "tabindex", a.Name())
	require.Equal(t, ` tabindex="foo"`, testRender(a, t))
}

func TestTarget(t *testing.T) {
	a := Target("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "target", a.Name())
	require.Equal(t, ` target="foo"`, testRender(a, t))
}

func TestTitle(t *testing.T) {
	a := Title("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "title", a.Name())
	require.Equal(t, ` title="foo"`, testRender(a, t))
}

func TestType(t *testing.T) {
	a := Type("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "type", a.Name())
	require.Equal(t, ` type="foo"`, testRender(a, t))
}

func TestValue(t *testing.T) {
	a := Value("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "value", a.Name())
	require.Equal(t, ` value="foo"`, testRender(a, t))
}

func TestWidth(t *testing.T) {
	a := Width("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "width", a.Name())
	require.Equal(t, ` width="foo"`, testRender(a, t))
}

func TestAllowFullScreen(t *testing.T) {
	a := AllowFullScreen()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "allowfullscreen", a.Name())
	require.Equal(t, ` allowfullscreen`, testRender(a, t))
}

func TestDefault(t *testing.T) {
	a := Default()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "default", a.Name())
	require.Equal(t, ` default`, testRender(a, t))
}

func TestInert(t *testing.T) {
	a := Inert()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "inert", a.Name())
	require.Equal(t, ` inert`, testRender(a, t))
}

func TestIsMap(t *testing.T) {
	a := IsMap()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ismap", a.Name())
	require.Equal(t, ` ismap`, testRender(a, t))
}

func TestItemScope(t *testing.T) {
	a := ItemScope()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "itemscope", a.Name())
	require.Equal(t, ` itemscope`, testRender(a, t))
}

func TestNoModule(t *testing.T) {
	a := NoModule()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "nomodule", a.Name())
	require.Equal(t, ` nomodule`, testRender(a, t))
}

func TestOpen(t *testing.T) {
	a := Open()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "open", a.Name())
	require.Equal(t, ` open`, testRender(a, t))
}

func TestReversed(t *testing.T) {
	a := Reversed()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "reversed", a.Name())
	require.Equal(t, ` reversed`, testRender(a, t))
}

func TestSpellcheck(t *testing.T) {
	a := Spellcheck()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "spellcheck", a.Name())
	require.Equal(t, ` spellcheck`, testRender(a, t))
}

func TestTranslate(t *testing.T) {
	a := Translate()
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "translate", a.Name())
	require.Equal(t, ` translate`, testRender(a, t))
}

func TestAriaAtomic(t *testing.T) {
	a := AriaAtomic("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-atomic", a.Name())
	require.Equal(t, ` aria-atomic="foo"`, testRender(a, t))
}

func TestAriaBusy(t *testing.T) {
	a := AriaBusy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-busy", a.Name())
	require.Equal(t, ` aria-busy="foo"`, testRender(a, t))
}

func TestAriaChecked(t *testing.T) {
	a := AriaChecked("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-checked", a.Name())
	require.Equal(t, ` aria-checked="foo"`, testRender(a, t))
}

func TestAriaControls(t *testing.T) {
	a := AriaControls("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-controls", a.Name())
	require.Equal(t, ` aria-controls="foo"`, testRender(a, t))
}

func TestAriaDescribedBy(t *testing.T) {
	a := AriaDescribedBy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-describedby", a.Name())
	require.Equal(t, ` aria-describedby="foo"`, testRender(a, t))
}

func TestAriaDisabled(t *testing.T) {
	a := AriaDisabled("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-disabled", a.Name())
	require.Equal(t, ` aria-disabled="foo"`, testRender(a, t))
}

func TestAriaExpanded(t *testing.T) {
	a := AriaExpanded("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-expanded", a.Name())
	require.Equal(t, ` aria-expanded="foo"`, testRender(a, t))
}

func TestAriaFlowTo(t *testing.T) {
	a := AriaFlowTo("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-flowto", a.Name())
	require.Equal(t, ` aria-flowto="foo"`, testRender(a, t))
}

func TestAriaHidden(t *testing.T) {
	a := AriaHidden("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-hidden", a.Name())
	require.Equal(t, ` aria-hidden="foo"`, testRender(a, t))
}

func TestAriaInvalid(t *testing.T) {
	a := AriaInvalid("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-invalid", a.Name())
	require.Equal(t, ` aria-invalid="foo"`, testRender(a, t))
}

func TestAriaLabel(t *testing.T) {
	a := AriaLabel("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-label", a.Name())
	require.Equal(t, ` aria-label="foo"`, testRender(a, t))
}

func TestAriaLabelledBy(t *testing.T) {
	a := AriaLabelledBy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-labelledby", a.Name())
	require.Equal(t, ` aria-labelledby="foo"`, testRender(a, t))
}

func TestAriaLive(t *testing.T) {
	a := AriaLive("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-live", a.Name())
	require.Equal(t, ` aria-live="foo"`, testRender(a, t))
}

func TestAriaOwns(t *testing.T) {
	a := AriaOwns("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-owns", a.Name())
	require.Equal(t, ` aria-owns="foo"`, testRender(a, t))
}

func TestAriaPlaceholder(t *testing.T) {
	a := AriaPlaceholder("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-placeholder", a.Name())
	require.Equal(t, ` aria-placeholder="foo"`, testRender(a, t))
}

func TestAriaPosInSet(t *testing.T) {
	a := AriaPosInSet("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-posinset", a.Name())
	require.Equal(t, ` aria-posinset="foo"`, testRender(a, t))
}

func TestAriaPressed(t *testing.T) {
	a := AriaPressed("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-pressed", a.Name())
	require.Equal(t, ` aria-pressed="foo"`, testRender(a, t))
}

func TestAriaReadonly(t *testing.T) {
	a := AriaReadonly("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-readonly", a.Name())
	require.Equal(t, ` aria-readonly="foo"`, testRender(a, t))
}

func TestAriaRelevant(t *testing.T) {
	a := AriaRelevant("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-relevant", a.Name())
	require.Equal(t, ` aria-relevant="foo"`, testRender(a, t))
}

func TestAriaRequired(t *testing.T) {
	a := AriaRequired("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-required", a.Name())
	require.Equal(t, ` aria-required="foo"`, testRender(a, t))
}

func TestAriaSelected(t *testing.T) {
	a := AriaSelected("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-selected", a.Name())
	require.Equal(t, ` aria-selected="foo"`, testRender(a, t))
}

func TestAriaSetSize(t *testing.T) {
	a := AriaSetSize("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-setsize", a.Name())
	require.Equal(t, ` aria-setsize="foo"`, testRender(a, t))
}

func TestAriaValueMax(t *testing.T) {
	a := AriaValueMax("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-valuemax", a.Name())
	require.Equal(t, ` aria-valuemax="foo"`, testRender(a, t))
}

func TestAriaValueMin(t *testing.T) {
	a := AriaValueMin("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-valuemin", a.Name())
	require.Equal(t, ` aria-valuemin="foo"`, testRender(a, t))
}

func TestAriaValueNow(t *testing.T) {
	a := AriaValueNow("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-valuenow", a.Name())
	require.Equal(t, ` aria-valuenow="foo"`, testRender(a, t))
}

func TestAriaValueText(t *testing.T) {
	a := AriaValueText("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "aria-valuetext", a.Name())
	require.Equal(t, ` aria-valuetext="foo"`, testRender(a, t))
}

func TestOnAbort(t *testing.T) {
	a := OnAbort("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onabort", a.Name())
	require.Equal(t, ` onabort="foo"`, testRender(a, t))
}

func TestOnBeforeUnload(t *testing.T) {
	a := OnBeforeUnload("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onbeforeunload", a.Name())
	require.Equal(t, ` onbeforeunload="foo"`, testRender(a, t))
}

func TestOnBlur(t *testing.T) {
	a := OnBlur("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onblur", a.Name())
	require.Equal(t, ` onblur="foo"`, testRender(a, t))
}

func TestOnCanPlay(t *testing.T) {
	a := OnCanPlay("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oncanplay", a.Name())
	require.Equal(t, ` oncanplay="foo"`, testRender(a, t))
}

func TestOnCanPlayThrough(t *testing.T) {
	a := OnCanPlayThrough("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oncanplaythrough", a.Name())
	require.Equal(t, ` oncanplaythrough="foo"`, testRender(a, t))
}

func TestOnChange(t *testing.T) {
	a := OnChange("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onchange", a.Name())
	require.Equal(t, ` onchange="foo"`, testRender(a, t))
}

func TestOnClick(t *testing.T) {
	a := OnClick("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onclick", a.Name())
	require.Equal(t, ` onclick="foo"`, testRender(a, t))
}

func TestOnContextMenu(t *testing.T) {
	a := OnContextMenu("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oncontextmenu", a.Name())
	require.Equal(t, ` oncontextmenu="foo"`, testRender(a, t))
}

func TestOnCopy(t *testing.T) {
	a := OnCopy("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oncopy", a.Name())
	require.Equal(t, ` oncopy="foo"`, testRender(a, t))
}

func TestOnCut(t *testing.T) {
	a := OnCut("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oncut", a.Name())
	require.Equal(t, ` oncut="foo"`, testRender(a, t))
}

func TestOnDblClick(t *testing.T) {
	a := OnDblClick("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondblclick", a.Name())
	require.Equal(t, ` ondblclick="foo"`, testRender(a, t))
}

func TestOnDrag(t *testing.T) {
	a := OnDrag("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondrag", a.Name())
	require.Equal(t, ` ondrag="foo"`, testRender(a, t))
}

func TestOnDragEnd(t *testing.T) {
	a := OnDragEnd("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondragend", a.Name())
	require.Equal(t, ` ondragend="foo"`, testRender(a, t))
}

func TestOnDragEnter(t *testing.T) {
	a := OnDragEnter("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondragenter", a.Name())
	require.Equal(t, ` ondragenter="foo"`, testRender(a, t))
}

func TestOnDragLeave(t *testing.T) {
	a := OnDragLeave("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondragleave", a.Name())
	require.Equal(t, ` ondragleave="foo"`, testRender(a, t))
}

func TestOnDragOver(t *testing.T) {
	a := OnDragOver("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondragover", a.Name())
	require.Equal(t, ` ondragover="foo"`, testRender(a, t))
}

func TestOnDragStart(t *testing.T) {
	a := OnDragStart("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondragstart", a.Name())
	require.Equal(t, ` ondragstart="foo"`, testRender(a, t))
}

func TestOnDrop(t *testing.T) {
	a := OnDrop("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondrop", a.Name())
	require.Equal(t, ` ondrop="foo"`, testRender(a, t))
}

func TestOnDurationChange(t *testing.T) {
	a := OnDurationChange("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ondurationchange", a.Name())
	require.Equal(t, ` ondurationchange="foo"`, testRender(a, t))
}

func TestOnEnded(t *testing.T) {
	a := OnEnded("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onended", a.Name())
	require.Equal(t, ` onended="foo"`, testRender(a, t))
}

func TestOnError(t *testing.T) {
	a := OnError("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onerror", a.Name())
	require.Equal(t, ` onerror="foo"`, testRender(a, t))
}

func TestOnFocus(t *testing.T) {
	a := OnFocus("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onfocus", a.Name())
	require.Equal(t, ` onfocus="foo"`, testRender(a, t))
}

func TestOnHashChange(t *testing.T) {
	a := OnHashChange("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onhashchange", a.Name())
	require.Equal(t, ` onhashchange="foo"`, testRender(a, t))
}

func TestOnInput(t *testing.T) {
	a := OnInput("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oninput", a.Name())
	require.Equal(t, ` oninput="foo"`, testRender(a, t))
}

func TestOnInvalid(t *testing.T) {
	a := OnInvalid("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "oninvalid", a.Name())
	require.Equal(t, ` oninvalid="foo"`, testRender(a, t))
}

func TestOnKeyDown(t *testing.T) {
	a := OnKeyDown("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onkeydown", a.Name())
	require.Equal(t, ` onkeydown="foo"`, testRender(a, t))
}

func TestOnKeyPress(t *testing.T) {
	a := OnKeyPress("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onkeypress", a.Name())
	require.Equal(t, ` onkeypress="foo"`, testRender(a, t))
}

func TestOnKeyUp(t *testing.T) {
	a := OnKeyUp("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onkeyup", a.Name())
	require.Equal(t, ` onkeyup="foo"`, testRender(a, t))
}

func TestOnLoad(t *testing.T) {
	a := OnLoad("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onload", a.Name())
	require.Equal(t, ` onload="foo"`, testRender(a, t))
}

func TestOnLoadedData(t *testing.T) {
	a := OnLoadedData("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onloadeddata", a.Name())
	require.Equal(t, ` onloadeddata="foo"`, testRender(a, t))
}

func TestOnLoadedMetadata(t *testing.T) {
	a := OnLoadedMetadata("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onloadedmetadata", a.Name())
	require.Equal(t, ` onloadedmetadata="foo"`, testRender(a, t))
}

func TestOnMouseDown(t *testing.T) {
	a := OnMouseDown("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmousedown", a.Name())
	require.Equal(t, ` onmousedown="foo"`, testRender(a, t))
}

func TestOnMouseEnter(t *testing.T) {
	a := OnMouseEnter("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseenter", a.Name())
	require.Equal(t, ` onmouseenter="foo"`, testRender(a, t))
}

func TestOnMouseLeave(t *testing.T) {
	a := OnMouseLeave("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseleave", a.Name())
	require.Equal(t, ` onmouseleave="foo"`, testRender(a, t))
}

func TestOnMouseMove(t *testing.T) {
	a := OnMouseMove("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmousemove", a.Name())
	require.Equal(t, ` onmousemove="foo"`, testRender(a, t))
}

func TestOnMouseOut(t *testing.T) {
	a := OnMouseOut("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseout", a.Name())
	require.Equal(t, ` onmouseout="foo"`, testRender(a, t))
}

func TestOnMouseOver(t *testing.T) {
	a := OnMouseOver("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseover", a.Name())
	require.Equal(t, ` onmouseover="foo"`, testRender(a, t))
}

func TestOnMouseUp(t *testing.T) {
	a := OnMouseUp("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onmouseup", a.Name())
	require.Equal(t, ` onmouseup="foo"`, testRender(a, t))
}

func TestOnPageHide(t *testing.T) {
	a := OnPageHide("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onpagehide", a.Name())
	require.Equal(t, ` onpagehide="foo"`, testRender(a, t))
}

func TestOnPageShow(t *testing.T) {
	a := OnPageShow("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onpageshow", a.Name())
	require.Equal(t, ` onpageshow="foo"`, testRender(a, t))
}

func TestOnPaste(t *testing.T) {
	a := OnPaste("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onpaste", a.Name())
	require.Equal(t, ` onpaste="foo"`, testRender(a, t))
}

func TestOnPause(t *testing.T) {
	a := OnPause("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onpause", a.Name())
	require.Equal(t, ` onpause="foo"`, testRender(a, t))
}

func TestOnPlay(t *testing.T) {
	a := OnPlay("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onplay", a.Name())
	require.Equal(t, ` onplay="foo"`, testRender(a, t))
}

func TestOnPlaying(t *testing.T) {
	a := OnPlaying("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onplaying", a.Name())
	require.Equal(t, ` onplaying="foo"`, testRender(a, t))
}

func TestOnPopState(t *testing.T) {
	a := OnPopState("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onpopstate", a.Name())
	require.Equal(t, ` onpopstate="foo"`, testRender(a, t))
}

func TestOnProgress(t *testing.T) {
	a := OnProgress("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onprogress", a.Name())
	require.Equal(t, ` onprogress="foo"`, testRender(a, t))
}

func TestOnRateChange(t *testing.T) {
	a := OnRateChange("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onratechange", a.Name())
	require.Equal(t, ` onratechange="foo"`, testRender(a, t))
}

func TestOnReset(t *testing.T) {
	a := OnReset("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onreset", a.Name())
	require.Equal(t, ` onreset="foo"`, testRender(a, t))
}

func TestOnResize(t *testing.T) {
	a := OnResize("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onresize", a.Name())
	require.Equal(t, ` onresize="foo"`, testRender(a, t))
}

func TestOnScroll(t *testing.T) {
	a := OnScroll("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onscroll", a.Name())
	require.Equal(t, ` onscroll="foo"`, testRender(a, t))
}

func TestOnSeeked(t *testing.T) {
	a := OnSeeked("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onseeked", a.Name())
	require.Equal(t, ` onseeked="foo"`, testRender(a, t))
}

func TestOnSeeking(t *testing.T) {
	a := OnSeeking("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onseeking", a.Name())
	require.Equal(t, ` onseeking="foo"`, testRender(a, t))
}

func TestOnSelect(t *testing.T) {
	a := OnSelect("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onselect", a.Name())
	require.Equal(t, ` onselect="foo"`, testRender(a, t))
}

func TestOnStalled(t *testing.T) {
	a := OnStalled("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onstalled", a.Name())
	require.Equal(t, ` onstalled="foo"`, testRender(a, t))
}

func TestOnSubmit(t *testing.T) {
	a := OnSubmit("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onsubmit", a.Name())
	require.Equal(t, ` onsubmit="foo"`, testRender(a, t))
}

func TestOnSuspend(t *testing.T) {
	a := OnSuspend("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onsuspend", a.Name())
	require.Equal(t, ` onsuspend="foo"`, testRender(a, t))
}

func TestOnTimeUpdate(t *testing.T) {
	a := OnTimeUpdate("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ontimeupdate", a.Name())
	require.Equal(t, ` ontimeupdate="foo"`, testRender(a, t))
}

func TestOnToggle(t *testing.T) {
	a := OnToggle("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "ontoggle", a.Name())
	require.Equal(t, ` ontoggle="foo"`, testRender(a, t))
}

func TestOnUnload(t *testing.T) {
	a := OnUnload("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onunload", a.Name())
	require.Equal(t, ` onunload="foo"`, testRender(a, t))
}

func TestOnVolumeChange(t *testing.T) {
	a := OnVolumeChange("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onvolumechange", a.Name())
	require.Equal(t, ` onvolumechange="foo"`, testRender(a, t))
}

func TestOnWaiting(t *testing.T) {
	a := OnWaiting("foo")
	require.Equal(t, aitch.AttributeNode, a.Type())
	require.Equal(t, "onwaiting", a.Name())
	require.Equal(t, ` onwaiting="foo"`, testRender(a, t))
}
