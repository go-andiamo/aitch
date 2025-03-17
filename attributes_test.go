package aitch

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAsync(t *testing.T) {
	a := Async()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "async", a.Name())
}

func TestAutoFocus(t *testing.T) {
	a := AutoFocus()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "autofocus", a.Name())
}

func TestAutoPlay(t *testing.T) {
	a := AutoPlay()
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "autoplay", a.Name())
}

func TestChecked(t *testing.T) {
	a := Checked()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "checked", a.Name())
}

func TestControls(t *testing.T) {
	a := Controls()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "controls", a.Name())
}

func TestCrossOrigin(t *testing.T) {
	a := CrossOrigin("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "crossorigin", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestDateTime(t *testing.T) {
	a := DateTime("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "datetime", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestDefer(t *testing.T) {
	a := Defer()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "defer", a.Name())
}

func TestDisabled(t *testing.T) {
	a := Disabled()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "disabled", a.Name())
}

func TestDownload(t *testing.T) {
	a := Download("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "download", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestDraggable(t *testing.T) {
	a := Draggable("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "draggable", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestLoop(t *testing.T) {
	a := Loop()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "loop", a.Name())
}

func TestMultiple(t *testing.T) {
	a := Multiple()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "multiple", a.Name())
}

func TestMuted(t *testing.T) {
	a := Muted()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "muted", a.Name())
}

func TestPlaysInline(t *testing.T) {
	a := PlaysInline()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "playsinline", a.Name())
}

func TestReadOnly(t *testing.T) {
	a := ReadOnly()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "readonly", a.Name())
}

func TestRequired(t *testing.T) {
	a := Required()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "required", a.Name())
}

func TestSelected(t *testing.T) {
	a := Selected()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "selected", a.Name())
}

func TestAccept(t *testing.T) {
	a := Accept("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "accept", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAction(t *testing.T) {
	a := Action("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "action", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAlt(t *testing.T) {
	a := Alt("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "alt", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAria(t *testing.T) {
	a := Aria("foo", "foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-foo", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAs(t *testing.T) {
	a := As("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "as", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAutoComplete(t *testing.T) {
	a := AutoComplete("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "autocomplete", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestCharset(t *testing.T) {
	a := Charset("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "charset", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestCiteAttr(t *testing.T) {
	a := CiteAttr("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "cite", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestClass(t *testing.T) {
	a := Class("foo")
	require.IsType(t, &delimitedAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "class", a.Name())
	require.Equal(t, 1, len(a.(*delimitedAttribute).values))
}

func TestCols(t *testing.T) {
	a := Cols("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "cols", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestColSpan(t *testing.T) {
	a := ColSpan("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "colspan", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestContent(t *testing.T) {
	a := Content("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "content", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestContentEditable(t *testing.T) {
	a := ContentEditable(true)
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "contenteditable", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestData(t *testing.T) {
	a := Data("foo", "foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "data-foo", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestDir(t *testing.T) {
	a := Dir("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "dir", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestEncType(t *testing.T) {
	a := EncType("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "enctype", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestFor(t *testing.T) {
	a := For("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "for", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestFormAttr(t *testing.T) {
	a := FormAttr("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "form", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestNoValidate(t *testing.T) {
	a := NoValidate()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "novalidate", a.Name())
}

func TestFormNoValidate(t *testing.T) {
	a := FormNoValidate()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "formnovalidate", a.Name())
}

func TestHeight(t *testing.T) {
	a := Height("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "height", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestHidden(t *testing.T) {
	a := Hidden()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "hidden", a.Name())
}

func TestHref(t *testing.T) {
	a := Href("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "href", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestId(t *testing.T) {
	a := Id("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "id", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestIntegrity(t *testing.T) {
	a := Integrity("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "integrity", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestLabelAttr(t *testing.T) {
	a := LabelAttr("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "label", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestLang(t *testing.T) {
	a := Lang("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "lang", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestList(t *testing.T) {
	a := List("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "list", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestLoading(t *testing.T) {
	a := Loading("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "loading", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestMax(t *testing.T) {
	a := Max("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "max", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestMaxLength(t *testing.T) {
	a := MaxLength("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "maxlength", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestMethod(t *testing.T) {
	a := Method("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "method", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestMin(t *testing.T) {
	a := Min("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "min", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestMinLength(t *testing.T) {
	a := MinLength("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "minlength", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestName(t *testing.T) {
	a := Name("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "name", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPattern(t *testing.T) {
	a := Pattern("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "pattern", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPlaceholder(t *testing.T) {
	a := Placeholder("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "placeholder", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPopover(t *testing.T) {
	a := Popover("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "popover", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPopoverTarget(t *testing.T) {
	a := PopoverTarget("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "popovertarget", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPopoverTargetAction(t *testing.T) {
	a := PopoverTargetAction("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "popovertargetaction", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPoster(t *testing.T) {
	a := Poster("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "poster", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestPreload(t *testing.T) {
	a := Preload("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "preload", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestReferrerPolicy(t *testing.T) {
	a := ReferrerPolicy("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "referrerpolicy", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestRel(t *testing.T) {
	a := Rel("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "rel", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestRole(t *testing.T) {
	a := Role("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "role", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestRows(t *testing.T) {
	a := Rows("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "rows", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestRowSpan(t *testing.T) {
	a := RowSpan("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "rowspan", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestSlotAttr(t *testing.T) {
	a := SlotAttr("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "slot", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestSrc(t *testing.T) {
	a := Src("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "src", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestSrcSet(t *testing.T) {
	a := SrcSet("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "srcset", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestStep(t *testing.T) {
	a := Step("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "step", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestStyle(t *testing.T) {
	a := Style("foo")
	require.IsType(t, &delimitedAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "style", a.Name())
	require.Equal(t, 1, len(a.(*delimitedAttribute).values))
}

func TestTabIndex(t *testing.T) {
	a := TabIndex("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "tabindex", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestTarget(t *testing.T) {
	a := Target("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "target", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestTitle(t *testing.T) {
	a := Title("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "title", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestType(t *testing.T) {
	a := Type("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "type", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestValue(t *testing.T) {
	a := Value("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "value", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestWidth(t *testing.T) {
	a := Width("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "width", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAllowFullScreen(t *testing.T) {
	a := AllowFullScreen()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "allowfullscreen", a.Name())
}

func TestDefault(t *testing.T) {
	a := Default()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "default", a.Name())
}

func TestInert(t *testing.T) {
	a := Inert()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "inert", a.Name())
}

func TestIsMap(t *testing.T) {
	a := IsMap()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ismap", a.Name())
}

func TestItemScope(t *testing.T) {
	a := ItemScope()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "itemscope", a.Name())
}

func TestNoModule(t *testing.T) {
	a := NoModule()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "nomodule", a.Name())
}

func TestOpen(t *testing.T) {
	a := Open()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "open", a.Name())
}

func TestReversed(t *testing.T) {
	a := Reversed()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "reversed", a.Name())
}

func TestSpellcheck(t *testing.T) {
	a := Spellcheck()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "spellcheck", a.Name())
}

func TestTranslate(t *testing.T) {
	a := Translate()
	require.IsType(t, &booleanAttribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "translate", a.Name())
}

func TestAriaAtomic(t *testing.T) {
	a := AriaAtomic("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-atomic", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaBusy(t *testing.T) {
	a := AriaBusy("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-busy", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaChecked(t *testing.T) {
	a := AriaChecked("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-checked", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaControls(t *testing.T) {
	a := AriaControls("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-controls", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaDescribedBy(t *testing.T) {
	a := AriaDescribedBy("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-describedby", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaDisabled(t *testing.T) {
	a := AriaDisabled("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-disabled", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaExpanded(t *testing.T) {
	a := AriaExpanded("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-expanded", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaFlowTo(t *testing.T) {
	a := AriaFlowTo("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-flowto", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaHidden(t *testing.T) {
	a := AriaHidden("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-hidden", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaInvalid(t *testing.T) {
	a := AriaInvalid("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-invalid", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaLabel(t *testing.T) {
	a := AriaLabel("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-label", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaLabelledBy(t *testing.T) {
	a := AriaLabelledBy("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-labelledby", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaLive(t *testing.T) {
	a := AriaLive("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-live", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaOwns(t *testing.T) {
	a := AriaOwns("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-owns", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaPlaceholder(t *testing.T) {
	a := AriaPlaceholder("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-placeholder", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaPosInSet(t *testing.T) {
	a := AriaPosInSet("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-posinset", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaPressed(t *testing.T) {
	a := AriaPressed("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-pressed", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaReadonly(t *testing.T) {
	a := AriaReadonly("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-readonly", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaRelevant(t *testing.T) {
	a := AriaRelevant("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-relevant", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaRequired(t *testing.T) {
	a := AriaRequired("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-required", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaSelected(t *testing.T) {
	a := AriaSelected("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-selected", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaSetSize(t *testing.T) {
	a := AriaSetSize("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-setsize", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaValueMax(t *testing.T) {
	a := AriaValueMax("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-valuemax", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaValueMin(t *testing.T) {
	a := AriaValueMin("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-valuemin", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaValueNow(t *testing.T) {
	a := AriaValueNow("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-valuenow", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestAriaValueText(t *testing.T) {
	a := AriaValueText("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "aria-valuetext", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnAbort(t *testing.T) {
	a := OnAbort("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onabort", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnBeforeUnload(t *testing.T) {
	a := OnBeforeUnload("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onbeforeunload", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnBlur(t *testing.T) {
	a := OnBlur("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onblur", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnCanPlay(t *testing.T) {
	a := OnCanPlay("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oncanplay", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnCanPlayThrough(t *testing.T) {
	a := OnCanPlayThrough("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oncanplaythrough", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnChange(t *testing.T) {
	a := OnChange("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onchange", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnClick(t *testing.T) {
	a := OnClick("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onclick", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnContextMenu(t *testing.T) {
	a := OnContextMenu("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oncontextmenu", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnCopy(t *testing.T) {
	a := OnCopy("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oncopy", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnCut(t *testing.T) {
	a := OnCut("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oncut", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDblClick(t *testing.T) {
	a := OnDblClick("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondblclick", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDrag(t *testing.T) {
	a := OnDrag("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondrag", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDragEnd(t *testing.T) {
	a := OnDragEnd("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondragend", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDragEnter(t *testing.T) {
	a := OnDragEnter("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondragenter", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDragLeave(t *testing.T) {
	a := OnDragLeave("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondragleave", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDragOver(t *testing.T) {
	a := OnDragOver("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondragover", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDragStart(t *testing.T) {
	a := OnDragStart("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondragstart", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDrop(t *testing.T) {
	a := OnDrop("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondrop", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnDurationChange(t *testing.T) {
	a := OnDurationChange("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ondurationchange", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnEnded(t *testing.T) {
	a := OnEnded("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onended", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnError(t *testing.T) {
	a := OnError("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onerror", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnFocus(t *testing.T) {
	a := OnFocus("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onfocus", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnHashChange(t *testing.T) {
	a := OnHashChange("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onhashchange", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnInput(t *testing.T) {
	a := OnInput("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oninput", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnInvalid(t *testing.T) {
	a := OnInvalid("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "oninvalid", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnKeyDown(t *testing.T) {
	a := OnKeyDown("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onkeydown", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnKeyPress(t *testing.T) {
	a := OnKeyPress("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onkeypress", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnKeyUp(t *testing.T) {
	a := OnKeyUp("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onkeyup", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnLoad(t *testing.T) {
	a := OnLoad("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onload", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnLoadedData(t *testing.T) {
	a := OnLoadedData("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onloadeddata", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnLoadedMetadata(t *testing.T) {
	a := OnLoadedMetadata("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onloadedmetadata", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseDown(t *testing.T) {
	a := OnMouseDown("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmousedown", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseEnter(t *testing.T) {
	a := OnMouseEnter("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmouseenter", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseLeave(t *testing.T) {
	a := OnMouseLeave("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmouseleave", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseMove(t *testing.T) {
	a := OnMouseMove("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmousemove", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseOut(t *testing.T) {
	a := OnMouseOut("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmouseout", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseOver(t *testing.T) {
	a := OnMouseOver("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmouseover", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnMouseUp(t *testing.T) {
	a := OnMouseUp("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onmouseup", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPageHide(t *testing.T) {
	a := OnPageHide("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onpagehide", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPageShow(t *testing.T) {
	a := OnPageShow("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onpageshow", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPaste(t *testing.T) {
	a := OnPaste("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onpaste", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPause(t *testing.T) {
	a := OnPause("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onpause", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPlay(t *testing.T) {
	a := OnPlay("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onplay", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPlaying(t *testing.T) {
	a := OnPlaying("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onplaying", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnPopState(t *testing.T) {
	a := OnPopState("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onpopstate", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnProgress(t *testing.T) {
	a := OnProgress("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onprogress", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnRateChange(t *testing.T) {
	a := OnRateChange("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onratechange", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnReset(t *testing.T) {
	a := OnReset("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onreset", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnResize(t *testing.T) {
	a := OnResize("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onresize", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnScroll(t *testing.T) {
	a := OnScroll("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onscroll", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnSeeked(t *testing.T) {
	a := OnSeeked("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onseeked", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnSeeking(t *testing.T) {
	a := OnSeeking("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onseeking", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnSelect(t *testing.T) {
	a := OnSelect("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onselect", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnStalled(t *testing.T) {
	a := OnStalled("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onstalled", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnSubmit(t *testing.T) {
	a := OnSubmit("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onsubmit", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnSuspend(t *testing.T) {
	a := OnSuspend("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onsuspend", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnTimeUpdate(t *testing.T) {
	a := OnTimeUpdate("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ontimeupdate", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnToggle(t *testing.T) {
	a := OnToggle("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "ontoggle", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnUnload(t *testing.T) {
	a := OnUnload("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onunload", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnVolumeChange(t *testing.T) {
	a := OnVolumeChange("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onvolumechange", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}

func TestOnWaiting(t *testing.T) {
	a := OnWaiting("foo")
	require.IsType(t, &attribute{}, a)
	require.Equal(t, AttributeNode, a.Type())
	require.Equal(t, "onwaiting", a.Name())
	require.Equal(t, 1, len(a.(*attribute).values))
}
