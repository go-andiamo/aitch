package css

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDisplay(t *testing.T) {
	a := Display("foo")
	assert.Equal(t, ` style="display:foo"`, testRender(a, t))
}

func TestPosition(t *testing.T) {
	a := Position("foo")
	assert.Equal(t, ` style="position:foo"`, testRender(a, t))
}

func TestTop(t *testing.T) {
	a := Top("foo")
	assert.Equal(t, ` style="top:foo"`, testRender(a, t))
}

func TestRight(t *testing.T) {
	a := Right("foo")
	assert.Equal(t, ` style="right:foo"`, testRender(a, t))
}

func TestBottom(t *testing.T) {
	a := Bottom("foo")
	assert.Equal(t, ` style="bottom:foo"`, testRender(a, t))
}

func TestLeft(t *testing.T) {
	a := Left("foo")
	assert.Equal(t, ` style="left:foo"`, testRender(a, t))
}

func TestZIndex(t *testing.T) {
	a := ZIndex("foo")
	assert.Equal(t, ` style="z-index:foo"`, testRender(a, t))
}

func TestFloat(t *testing.T) {
	a := Float("foo")
	assert.Equal(t, ` style="float:foo"`, testRender(a, t))
}

func TestClear(t *testing.T) {
	a := Clear("foo")
	assert.Equal(t, ` style="clear:foo"`, testRender(a, t))
}

func TestFlex(t *testing.T) {
	a := Flex("foo")
	assert.Equal(t, ` style="flex:foo"`, testRender(a, t))
}

func TestFlexGrow(t *testing.T) {
	a := FlexGrow("foo")
	assert.Equal(t, ` style="flex-grow:foo"`, testRender(a, t))
}

func TestFlexShrink(t *testing.T) {
	a := FlexShrink("foo")
	assert.Equal(t, ` style="flex-shrink:foo"`, testRender(a, t))
}

func TestFlexBasis(t *testing.T) {
	a := FlexBasis("foo")
	assert.Equal(t, ` style="flex-basis:foo"`, testRender(a, t))
}

func TestJustifyContent(t *testing.T) {
	a := JustifyContent("foo")
	assert.Equal(t, ` style="justify-content:foo"`, testRender(a, t))
}

func TestAlignItems(t *testing.T) {
	a := AlignItems("foo")
	assert.Equal(t, ` style="align-items:foo"`, testRender(a, t))
}

func TestAlignContent(t *testing.T) {
	a := AlignContent("foo")
	assert.Equal(t, ` style="align-content:foo"`, testRender(a, t))
}

func TestAlignSelf(t *testing.T) {
	a := AlignSelf("foo")
	assert.Equal(t, ` style="align-self:foo"`, testRender(a, t))
}

func TestGap(t *testing.T) {
	a := Gap("foo")
	assert.Equal(t, ` style="gap:foo"`, testRender(a, t))
}

func TestRowGap(t *testing.T) {
	a := RowGap("foo")
	assert.Equal(t, ` style="row-gap:foo"`, testRender(a, t))
}

func TestColumnGap(t *testing.T) {
	a := ColumnGap("foo")
	assert.Equal(t, ` style="column-gap:foo"`, testRender(a, t))
}

func TestWidth(t *testing.T) {
	a := Width("foo")
	assert.Equal(t, ` style="width:foo"`, testRender(a, t))
}

func TestHeight(t *testing.T) {
	a := Height("foo")
	assert.Equal(t, ` style="height:foo"`, testRender(a, t))
}

func TestMinWidth(t *testing.T) {
	a := MinWidth("foo")
	assert.Equal(t, ` style="min-width:foo"`, testRender(a, t))
}

func TestMaxWidth(t *testing.T) {
	a := MaxWidth("foo")
	assert.Equal(t, ` style="max-width:foo"`, testRender(a, t))
}

func TestMinHeight(t *testing.T) {
	a := MinHeight("foo")
	assert.Equal(t, ` style="min-height:foo"`, testRender(a, t))
}

func TestMaxHeight(t *testing.T) {
	a := MaxHeight("foo")
	assert.Equal(t, ` style="max-height:foo"`, testRender(a, t))
}

func TestMargin(t *testing.T) {
	a := Margin("foo")
	assert.Equal(t, ` style="margin:foo"`, testRender(a, t))
}

func TestMarginTop(t *testing.T) {
	a := MarginTop("foo")
	assert.Equal(t, ` style="margin-top:foo"`, testRender(a, t))
}

func TestMarginRight(t *testing.T) {
	a := MarginRight("foo")
	assert.Equal(t, ` style="margin-right:foo"`, testRender(a, t))
}

func TestMarginBottom(t *testing.T) {
	a := MarginBottom("foo")
	assert.Equal(t, ` style="margin-bottom:foo"`, testRender(a, t))
}

func TestMarginLeft(t *testing.T) {
	a := MarginLeft("foo")
	assert.Equal(t, ` style="margin-left:foo"`, testRender(a, t))
}

func TestPadding(t *testing.T) {
	a := Padding("foo")
	assert.Equal(t, ` style="padding:foo"`, testRender(a, t))
}

func TestPaddingTop(t *testing.T) {
	a := PaddingTop("foo")
	assert.Equal(t, ` style="padding-top:foo"`, testRender(a, t))
}

func TestPaddingRight(t *testing.T) {
	a := PaddingRight("foo")
	assert.Equal(t, ` style="padding-right:foo"`, testRender(a, t))
}

func TestPaddingBottom(t *testing.T) {
	a := PaddingBottom("foo")
	assert.Equal(t, ` style="padding-bottom:foo"`, testRender(a, t))
}

func TestPaddingLeft(t *testing.T) {
	a := PaddingLeft("foo")
	assert.Equal(t, ` style="padding-left:foo"`, testRender(a, t))
}

func TestBorder(t *testing.T) {
	a := Border("foo")
	assert.Equal(t, ` style="border:foo"`, testRender(a, t))
}

func TestBorderWidth(t *testing.T) {
	a := BorderWidth("foo")
	assert.Equal(t, ` style="border-width:foo"`, testRender(a, t))
}

func TestBorderColor(t *testing.T) {
	a := BorderColor("foo")
	assert.Equal(t, ` style="border-color:foo"`, testRender(a, t))
}

func TestBorderStyle(t *testing.T) {
	a := BorderStyle("foo")
	assert.Equal(t, ` style="border-style:foo"`, testRender(a, t))
}

func TestBorderRadius(t *testing.T) {
	a := BorderRadius("foo")
	assert.Equal(t, ` style="border-radius:foo"`, testRender(a, t))
}

func TestBorderTop(t *testing.T) {
	a := BorderTop("foo")
	assert.Equal(t, ` style="border-top:foo"`, testRender(a, t))
}

func TestBorderBottom(t *testing.T) {
	a := BorderBottom("foo")
	assert.Equal(t, ` style="border-bottom:foo"`, testRender(a, t))
}

func TestBorderLeft(t *testing.T) {
	a := BorderLeft("foo")
	assert.Equal(t, ` style="border-left:foo"`, testRender(a, t))
}

func TestBorderRight(t *testing.T) {
	a := BorderRight("foo")
	assert.Equal(t, ` style="border-right:foo"`, testRender(a, t))
}

func TestBackground(t *testing.T) {
	a := Background("foo")
	assert.Equal(t, ` style="background:foo"`, testRender(a, t))
}

func TestBackgroundColor(t *testing.T) {
	a := BackgroundColor("foo")
	assert.Equal(t, ` style="background-color:foo"`, testRender(a, t))
}

func TestBackgroundImage(t *testing.T) {
	a := BackgroundImage("foo")
	assert.Equal(t, ` style="background-image:foo"`, testRender(a, t))
}

func TestBackgroundSize(t *testing.T) {
	a := BackgroundSize("foo")
	assert.Equal(t, ` style="background-size:foo"`, testRender(a, t))
}

func TestBackgroundPosition(t *testing.T) {
	a := BackgroundPosition("foo")
	assert.Equal(t, ` style="background-position:foo"`, testRender(a, t))
}

func TestBackgroundRepeat(t *testing.T) {
	a := BackgroundRepeat("foo")
	assert.Equal(t, ` style="background-repeat:foo"`, testRender(a, t))
}

func TestColor(t *testing.T) {
	a := Color("foo")
	assert.Equal(t, ` style="color:foo"`, testRender(a, t))
}

func TestFontSize(t *testing.T) {
	a := FontSize("foo")
	assert.Equal(t, ` style="font-size:foo"`, testRender(a, t))
}

func TestFontWeight(t *testing.T) {
	a := FontWeight("foo")
	assert.Equal(t, ` style="font-weight:foo"`, testRender(a, t))
}

func TestFontFamily(t *testing.T) {
	a := FontFamily("foo")
	assert.Equal(t, ` style="font-family:foo"`, testRender(a, t))
}

func TestLineHeight(t *testing.T) {
	a := LineHeight("foo")
	assert.Equal(t, ` style="line-height:foo"`, testRender(a, t))
}

func TestLetterSpacing(t *testing.T) {
	a := LetterSpacing("foo")
	assert.Equal(t, ` style="letter-spacing:foo"`, testRender(a, t))
}

func TestTextAlign(t *testing.T) {
	a := TextAlign("foo")
	assert.Equal(t, ` style="text-align:foo"`, testRender(a, t))
}

func TestTextTransform(t *testing.T) {
	a := TextTransform("foo")
	assert.Equal(t, ` style="text-transform:foo"`, testRender(a, t))
}

func TestTextDecoration(t *testing.T) {
	a := TextDecoration("foo")
	assert.Equal(t, ` style="text-decoration:foo"`, testRender(a, t))
}

func TestWhiteSpace(t *testing.T) {
	a := WhiteSpace("foo")
	assert.Equal(t, ` style="white-space:foo"`, testRender(a, t))
}

func TestWordBreak(t *testing.T) {
	a := WordBreak("foo")
	assert.Equal(t, ` style="word-break:foo"`, testRender(a, t))
}

func TestOpacity(t *testing.T) {
	a := Opacity("foo")
	assert.Equal(t, ` style="opacity:foo"`, testRender(a, t))
}

func TestBoxShadow(t *testing.T) {
	a := BoxShadow("foo")
	assert.Equal(t, ` style="box-shadow:foo"`, testRender(a, t))
}

func TestVisibility(t *testing.T) {
	a := Visibility("foo")
	assert.Equal(t, ` style="visibility:foo"`, testRender(a, t))
}

func TestOverflow(t *testing.T) {
	a := Overflow("foo")
	assert.Equal(t, ` style="overflow:foo"`, testRender(a, t))
}

func TestOverflowX(t *testing.T) {
	a := OverflowX("foo")
	assert.Equal(t, ` style="overflow-x:foo"`, testRender(a, t))
}

func TestOverflowY(t *testing.T) {
	a := OverflowY("foo")
	assert.Equal(t, ` style="overflow-y:foo"`, testRender(a, t))
}

func TestCursor(t *testing.T) {
	a := Cursor("foo")
	assert.Equal(t, ` style="cursor:foo"`, testRender(a, t))
}

func TestTransition(t *testing.T) {
	a := Transition("foo")
	assert.Equal(t, ` style="transition:foo"`, testRender(a, t))
}

func TestTransitionDuration(t *testing.T) {
	a := TransitionDuration("foo")
	assert.Equal(t, ` style="transition-duration:foo"`, testRender(a, t))
}

func TestTransitionTimingFunction(t *testing.T) {
	a := TransitionTimingFunction("foo")
	assert.Equal(t, ` style="transition-timing-function:foo"`, testRender(a, t))
}

func TestTransform(t *testing.T) {
	a := Transform("foo")
	assert.Equal(t, ` style="transform:foo"`, testRender(a, t))
}

func TestAnimation(t *testing.T) {
	a := Animation("foo")
	assert.Equal(t, ` style="animation:foo"`, testRender(a, t))
}

func TestUserSelect(t *testing.T) {
	a := UserSelect("foo")
	assert.Equal(t, ` style="user-select:foo"`, testRender(a, t))
}

func TestPointerEvents(t *testing.T) {
	a := PointerEvents("foo")
	assert.Equal(t, ` style="pointer-events:foo"`, testRender(a, t))
}

func TestIsolation(t *testing.T) {
	a := Isolation("foo")
	assert.Equal(t, ` style="isolation:foo"`, testRender(a, t))
}

func TestObjectFit(t *testing.T) {
	a := ObjectFit("foo")
	assert.Equal(t, ` style="object-fit:foo"`, testRender(a, t))
}

func TestObjectPosition(t *testing.T) {
	a := ObjectPosition("foo")
	assert.Equal(t, ` style="object-position:foo"`, testRender(a, t))
}

func TestBoxSizing(t *testing.T) {
	a := BoxSizing("foo")
	assert.Equal(t, ` style="box-sizing:foo"`, testRender(a, t))
}

func TestAspectRatio(t *testing.T) {
	a := AspectRatio("foo")
	assert.Equal(t, ` style="aspect-ratio:foo"`, testRender(a, t))
}

func TestContain(t *testing.T) {
	a := Contain("foo")
	assert.Equal(t, ` style="contain:foo"`, testRender(a, t))
}

func TestContentVisibility(t *testing.T) {
	a := ContentVisibility("foo")
	assert.Equal(t, ` style="content-visibility:foo"`, testRender(a, t))
}

func TestResize(t *testing.T) {
	a := Resize("foo")
	assert.Equal(t, ` style="resize:foo"`, testRender(a, t))
}

func TestGridTemplateColumns(t *testing.T) {
	a := GridTemplateColumns("foo")
	assert.Equal(t, ` style="grid-template-columns:foo"`, testRender(a, t))
}

func TestGridTemplateRows(t *testing.T) {
	a := GridTemplateRows("foo")
	assert.Equal(t, ` style="grid-template-rows:foo"`, testRender(a, t))
}

func TestGridColumn(t *testing.T) {
	a := GridColumn("foo")
	assert.Equal(t, ` style="grid-column:foo"`, testRender(a, t))
}

func TestGridRow(t *testing.T) {
	a := GridRow("foo")
	assert.Equal(t, ` style="grid-row:foo"`, testRender(a, t))
}

func TestGridArea(t *testing.T) {
	a := GridArea("foo")
	assert.Equal(t, ` style="grid-area:foo"`, testRender(a, t))
}

func TestGridAutoFlow(t *testing.T) {
	a := GridAutoFlow("foo")
	assert.Equal(t, ` style="grid-auto-flow:foo"`, testRender(a, t))
}

func TestPlaceItems(t *testing.T) {
	a := PlaceItems("foo")
	assert.Equal(t, ` style="place-items:foo"`, testRender(a, t))
}

func TestPlaceContent(t *testing.T) {
	a := PlaceContent("foo")
	assert.Equal(t, ` style="place-content:foo"`, testRender(a, t))
}

func TestPlaceSelf(t *testing.T) {
	a := PlaceSelf("foo")
	assert.Equal(t, ` style="place-self:foo"`, testRender(a, t))
}

func TestFontStyle(t *testing.T) {
	a := FontStyle("foo")
	assert.Equal(t, ` style="font-style:foo"`, testRender(a, t))
}

func TestFontVariant(t *testing.T) {
	a := FontVariant("foo")
	assert.Equal(t, ` style="font-variant:foo"`, testRender(a, t))
}

func TestFontStretch(t *testing.T) {
	a := FontStretch("foo")
	assert.Equal(t, ` style="font-stretch:foo"`, testRender(a, t))
}

func TestDirection(t *testing.T) {
	a := Direction("foo")
	assert.Equal(t, ` style="direction:foo"`, testRender(a, t))
}

func TestTabSize(t *testing.T) {
	a := TabSize("foo")
	assert.Equal(t, ` style="tab-size:foo"`, testRender(a, t))
}

func TestListStyle(t *testing.T) {
	a := ListStyle("foo")
	assert.Equal(t, ` style="list-style:foo"`, testRender(a, t))
}

func TestListStyleType(t *testing.T) {
	a := ListStyleType("foo")
	assert.Equal(t, ` style="list-style-type:foo"`, testRender(a, t))
}

func TestListStylePosition(t *testing.T) {
	a := ListStylePosition("foo")
	assert.Equal(t, ` style="list-style-position:foo"`, testRender(a, t))
}

func TestListStyleImage(t *testing.T) {
	a := ListStyleImage("foo")
	assert.Equal(t, ` style="list-style-image:foo"`, testRender(a, t))
}

func TestBorderCollapse(t *testing.T) {
	a := BorderCollapse("foo")
	assert.Equal(t, ` style="border-collapse:foo"`, testRender(a, t))
}

func TestBorderSpacing(t *testing.T) {
	a := BorderSpacing("foo")
	assert.Equal(t, ` style="border-spacing:foo"`, testRender(a, t))
}

func TestTableLayout(t *testing.T) {
	a := TableLayout("foo")
	assert.Equal(t, ` style="table-layout:foo"`, testRender(a, t))
}

func TestEmptyCells(t *testing.T) {
	a := EmptyCells("foo")
	assert.Equal(t, ` style="empty-cells:foo"`, testRender(a, t))
}

func TestFilter(t *testing.T) {
	a := Filter("foo")
	assert.Equal(t, ` style="filter:foo"`, testRender(a, t))
}

func TestBackdropFilter(t *testing.T) {
	a := BackdropFilter("foo")
	assert.Equal(t, ` style="backdrop-filter:foo"`, testRender(a, t))
}

func TestMixBlendMode(t *testing.T) {
	a := MixBlendMode("foo")
	assert.Equal(t, ` style="mix-blend-mode:foo"`, testRender(a, t))
}

func TestBoxDecorationBreak(t *testing.T) {
	a := BoxDecorationBreak("foo")
	assert.Equal(t, ` style="box-decoration-break:foo"`, testRender(a, t))
}

func TestScrollBehavior(t *testing.T) {
	a := ScrollBehavior("foo")
	assert.Equal(t, ` style="scroll-behavior:foo"`, testRender(a, t))
}

func TestScrollSnapType(t *testing.T) {
	a := ScrollSnapType("foo")
	assert.Equal(t, ` style="scroll-snap-type:foo"`, testRender(a, t))
}

func TestScrollSnapAlign(t *testing.T) {
	a := ScrollSnapAlign("foo")
	assert.Equal(t, ` style="scroll-snap-align:foo"`, testRender(a, t))
}

func TestScrollMargin(t *testing.T) {
	a := ScrollMargin("foo")
	assert.Equal(t, ` style="scroll-margin:foo"`, testRender(a, t))
}

func TestScrollPadding(t *testing.T) {
	a := ScrollPadding("foo")
	assert.Equal(t, ` style="scroll-padding:foo"`, testRender(a, t))
}

func TestAnimationName(t *testing.T) {
	a := AnimationName("foo")
	assert.Equal(t, ` style="animation-name:foo"`, testRender(a, t))
}

func TestAnimationDuration(t *testing.T) {
	a := AnimationDuration("foo")
	assert.Equal(t, ` style="animation-duration:foo"`, testRender(a, t))
}

func TestAnimationTimingFunction(t *testing.T) {
	a := AnimationTimingFunction("foo")
	assert.Equal(t, ` style="animation-timing-function:foo"`, testRender(a, t))
}

func TestAnimationDelay(t *testing.T) {
	a := AnimationDelay("foo")
	assert.Equal(t, ` style="animation-delay:foo"`, testRender(a, t))
}

func TestAnimationIterationCount(t *testing.T) {
	a := AnimationIterationCount("foo")
	assert.Equal(t, ` style="animation-iteration-count:foo"`, testRender(a, t))
}

func TestAnimationDirection(t *testing.T) {
	a := AnimationDirection("foo")
	assert.Equal(t, ` style="animation-direction:foo"`, testRender(a, t))
}

func TestAnimationFillMode(t *testing.T) {
	a := AnimationFillMode("foo")
	assert.Equal(t, ` style="animation-fill-mode:foo"`, testRender(a, t))
}

func TestAnimationPlayState(t *testing.T) {
	a := AnimationPlayState("foo")
	assert.Equal(t, ` style="animation-play-state:foo"`, testRender(a, t))
}

func TestTransitionProperty(t *testing.T) {
	a := TransitionProperty("foo")
	assert.Equal(t, ` style="transition-property:foo"`, testRender(a, t))
}

func TestTransitionDelay(t *testing.T) {
	a := TransitionDelay("foo")
	assert.Equal(t, ` style="transition-delay:foo"`, testRender(a, t))
}

func TestTouchAction(t *testing.T) {
	a := TouchAction("foo")
	assert.Equal(t, ` style="touch-action:foo"`, testRender(a, t))
}

func TestWillChange(t *testing.T) {
	a := WillChange("foo")
	assert.Equal(t, ` style="will-change:foo"`, testRender(a, t))
}

func TestCaretColor(t *testing.T) {
	a := CaretColor("foo")
	assert.Equal(t, ` style="caret-color:foo"`, testRender(a, t))
}

func TestScrollSnapStop(t *testing.T) {
	a := ScrollSnapStop("foo")
	assert.Equal(t, ` style="scroll-snap-stop:foo"`, testRender(a, t))
}

func TestSpeak(t *testing.T) {
	a := Speak("foo")
	assert.Equal(t, ` style="speak:foo"`, testRender(a, t))
}

func TestPageBreakBefore(t *testing.T) {
	a := PageBreakBefore("foo")
	assert.Equal(t, ` style="page-break-before:foo"`, testRender(a, t))
}

func TestPageBreakAfter(t *testing.T) {
	a := PageBreakAfter("foo")
	assert.Equal(t, ` style="page-break-after:foo"`, testRender(a, t))
}

func TestPageBreakInside(t *testing.T) {
	a := PageBreakInside("foo")
	assert.Equal(t, ` style="page-break-inside:foo"`, testRender(a, t))
}
