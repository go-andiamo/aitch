package css

import "github.com/go-andiamo/aitch"

// Display sets the `display` style property.
func Display(value ...any) aitch.Node {
	return styleAttribute(propDisplay, value...)
}

// Position sets the `position` style property.
func Position(value ...any) aitch.Node {
	return styleAttribute(propPosition, value...)
}

// Top sets the `top` style property.
func Top(value ...any) aitch.Node {
	return styleAttribute(propTop, value...)
}

// Right sets the `right` style property.
func Right(value ...any) aitch.Node {
	return styleAttribute(propRight, value...)
}

// Bottom sets the `bottom` style property.
func Bottom(value ...any) aitch.Node {
	return styleAttribute(propBottom, value...)
}

// Left sets the `left` style property.
func Left(value ...any) aitch.Node {
	return styleAttribute(propLeft, value...)
}

// ZIndex sets the `z-index` style property.
func ZIndex(value ...any) aitch.Node {
	return styleAttribute(propZIndex, value...)
}

// Float sets the `float` style property.
func Float(value ...any) aitch.Node {
	return styleAttribute(propFloat, value...)
}

// Clear sets the `clear` style property.
func Clear(value ...any) aitch.Node {
	return styleAttribute(propClear, value...)
}

// Flex sets the `flex` style property.
func Flex(value ...any) aitch.Node {
	return styleAttribute(propFlex, value...)
}

// FlexGrow sets the `flex-grow` style property.
func FlexGrow(value ...any) aitch.Node {
	return styleAttribute(propFlexGrow, value...)
}

// FlexShrink sets the `flex-shrink` style property.
func FlexShrink(value ...any) aitch.Node {
	return styleAttribute(propFlexShrink, value...)
}

// FlexBasis sets the `flex-basis` style property.
func FlexBasis(value ...any) aitch.Node {
	return styleAttribute(propFlexBasis, value...)
}

// JustifyContent sets the `justify-content` style property.
func JustifyContent(value ...any) aitch.Node {
	return styleAttribute(propJustifyContent, value...)
}

// AlignItems sets the `align-items` style property.
func AlignItems(value ...any) aitch.Node {
	return styleAttribute(propAlignItems, value...)
}

// AlignContent sets the `align-content` style property.
func AlignContent(value ...any) aitch.Node {
	return styleAttribute(propAlignContent, value...)
}

// AlignSelf sets the `align-self` style property.
func AlignSelf(value ...any) aitch.Node {
	return styleAttribute(propAlignSelf, value...)
}

// Gap sets the `gap` style property.
func Gap(value ...any) aitch.Node {
	return styleAttribute(propGap, value...)
}

// RowGap sets the `row-gap` style property.
func RowGap(value ...any) aitch.Node {
	return styleAttribute(propRowGap, value...)
}

// ColumnGap sets the `column-gap` style property.
func ColumnGap(value ...any) aitch.Node {
	return styleAttribute(propColumnGap, value...)
}

// Width sets the `width` style property.
func Width(value ...any) aitch.Node {
	return styleAttribute(propWidth, value...)
}

// Height sets the `height` style property.
func Height(value ...any) aitch.Node {
	return styleAttribute(propHeight, value...)
}

// MinWidth sets the `min-width` style property.
func MinWidth(value ...any) aitch.Node {
	return styleAttribute(propMinWidth, value...)
}

// MaxWidth sets the `max-width` style property.
func MaxWidth(value ...any) aitch.Node {
	return styleAttribute(propMaxWidth, value...)
}

// MinHeight sets the `min-height` style property.
func MinHeight(value ...any) aitch.Node {
	return styleAttribute(propMinHeight, value...)
}

// MaxHeight sets the `max-height` style property.
func MaxHeight(value ...any) aitch.Node {
	return styleAttribute(propMaxHeight, value...)
}

// Margin sets the `margin` style property.
func Margin(value ...any) aitch.Node {
	return styleAttribute(propMargin, value...)
}

// MarginTop sets the `margin-top` style property.
func MarginTop(value ...any) aitch.Node {
	return styleAttribute(propMarginTop, value...)
}

// MarginRight sets the `margin-right` style property.
func MarginRight(value ...any) aitch.Node {
	return styleAttribute(propMarginRight, value...)
}

// MarginBottom sets the `margin-bottom` style property.
func MarginBottom(value ...any) aitch.Node {
	return styleAttribute(propMarginBottom, value...)
}

// MarginLeft sets the `margin-left` style property.
func MarginLeft(value ...any) aitch.Node {
	return styleAttribute(propMarginLeft, value...)
}

// Padding sets the `padding` style property.
func Padding(value ...any) aitch.Node {
	return styleAttribute(propPadding, value...)
}

// PaddingTop sets the `padding-top` style property.
func PaddingTop(value ...any) aitch.Node {
	return styleAttribute(propPaddingTop, value...)
}

// PaddingRight sets the `padding-right` style property.
func PaddingRight(value ...any) aitch.Node {
	return styleAttribute(propPaddingRight, value...)
}

// PaddingBottom sets the `padding-bottom` style property.
func PaddingBottom(value ...any) aitch.Node {
	return styleAttribute(propPaddingBottom, value...)
}

// PaddingLeft sets the `padding-left` style property.
func PaddingLeft(value ...any) aitch.Node {
	return styleAttribute(propPaddingLeft, value...)
}

// Border sets the `border` style property.
func Border(value ...any) aitch.Node {
	return styleAttribute(propBorder, value...)
}

// BorderWidth sets the `border-width` style property.
func BorderWidth(value ...any) aitch.Node {
	return styleAttribute(propBorderWidth, value...)
}

// BorderColor sets the `border-color` style property.
func BorderColor(value ...any) aitch.Node {
	return styleAttribute(propBorderColor, value...)
}

// BorderStyle sets the `border-style` style property.
func BorderStyle(value ...any) aitch.Node {
	return styleAttribute(propBorderStyle, value...)
}

// BorderRadius sets the `border-radius` style property.
func BorderRadius(value ...any) aitch.Node {
	return styleAttribute(propBorderRadius, value...)
}

// BorderTop sets the `border-top` style property.
func BorderTop(value ...any) aitch.Node {
	return styleAttribute(propBorderTop, value...)
}

// BorderBottom sets the `border-bottom` style property.
func BorderBottom(value ...any) aitch.Node {
	return styleAttribute(propBorderBottom, value...)
}

// BorderLeft sets the `border-left` style property.
func BorderLeft(value ...any) aitch.Node {
	return styleAttribute(propBorderLeft, value...)
}

// BorderRight sets the `border-right` style property.
func BorderRight(value ...any) aitch.Node {
	return styleAttribute(propBorderRight, value...)
}

// Background sets the `background` style property.
func Background(value ...any) aitch.Node {
	return styleAttribute(propBackground, value...)
}

// BackgroundColor sets the `background-color` style property.
func BackgroundColor(value ...any) aitch.Node {
	return styleAttribute(propBackgroundColor, value...)
}

// BackgroundImage sets the `background-image` style property.
func BackgroundImage(value ...any) aitch.Node {
	return styleAttribute(propBackgroundImage, value...)
}

// BackgroundSize sets the `background-size` style property.
func BackgroundSize(value ...any) aitch.Node {
	return styleAttribute(propBackgroundSize, value...)
}

// BackgroundPosition sets the `background-position` style property.
func BackgroundPosition(value ...any) aitch.Node {
	return styleAttribute(propBackgroundPosition, value...)
}

// BackgroundRepeat sets the `background-repeat` style property.
func BackgroundRepeat(value ...any) aitch.Node {
	return styleAttribute(propBackgroundRepeat, value...)
}

// Color sets the `color` style property.
func Color(value ...any) aitch.Node {
	return styleAttribute(propColor, value...)
}

// FontSize sets the `font-size` style property.
func FontSize(value ...any) aitch.Node {
	return styleAttribute(propFontSize, value...)
}

// FontWeight sets the `font-weight` style property.
func FontWeight(value ...any) aitch.Node {
	return styleAttribute(propFontWeight, value...)
}

// FontFamily sets the `font-family` style property.
func FontFamily(value ...any) aitch.Node {
	return styleAttribute(propFontFamily, value...)
}

// LineHeight sets the `line-height` style property.
func LineHeight(value ...any) aitch.Node {
	return styleAttribute(propLineHeight, value...)
}

// LetterSpacing sets the `letter-spacing` style property.
func LetterSpacing(value ...any) aitch.Node {
	return styleAttribute(propLetterSpacing, value...)
}

// TextAlign sets the `text-align` style property.
func TextAlign(value ...any) aitch.Node {
	return styleAttribute(propTextAlign, value...)
}

// TextTransform sets the `text-transform` style property.
func TextTransform(value ...any) aitch.Node {
	return styleAttribute(propTextTransform, value...)
}

// TextDecoration sets the `text-decoration` style property.
func TextDecoration(value ...any) aitch.Node {
	return styleAttribute(propTextDecoration, value...)
}

// WhiteSpace sets the `white-space` style property.
func WhiteSpace(value ...any) aitch.Node {
	return styleAttribute(propWhiteSpace, value...)
}

// WordBreak sets the `word-break` style property.
func WordBreak(value ...any) aitch.Node {
	return styleAttribute(propWordBreak, value...)
}

// Opacity sets the `opacity` style property.
func Opacity(value ...any) aitch.Node {
	return styleAttribute(propOpacity, value...)
}

// BoxShadow sets the `box-shadow` style property.
func BoxShadow(value ...any) aitch.Node {
	return styleAttribute(propBoxShadow, value...)
}

// Visibility sets the `visibility` style property.
func Visibility(value ...any) aitch.Node {
	return styleAttribute(propVisibility, value...)
}

// Overflow sets the `overflow` style property.
func Overflow(value ...any) aitch.Node {
	return styleAttribute(propOverflow, value...)
}

// OverflowX sets the `overflow-x` style property.
func OverflowX(value ...any) aitch.Node {
	return styleAttribute(propOverflowX, value...)
}

// OverflowY sets the `overflow-y` style property.
func OverflowY(value ...any) aitch.Node {
	return styleAttribute(propOverflowY, value...)
}

// Cursor sets the `cursor` style property.
func Cursor(value ...any) aitch.Node {
	return styleAttribute(propCursor, value...)
}

// Transition sets the `transition` style property.
func Transition(value ...any) aitch.Node {
	return styleAttribute(propTransition, value...)
}

// TransitionDuration sets the `transition-duration` style property.
func TransitionDuration(value ...any) aitch.Node {
	return styleAttribute(propTransitionDuration, value...)
}

// TransitionTimingFunction sets the `transition-timing-function` style property.
func TransitionTimingFunction(value ...any) aitch.Node {
	return styleAttribute(propTransitionTiming, value...)
}

// Transform sets the `transform` style property.
func Transform(value ...any) aitch.Node {
	return styleAttribute(propTransform, value...)
}

// Animation sets the `animation` style property.
func Animation(value ...any) aitch.Node {
	return styleAttribute(propAnimation, value...)
}

// UserSelect sets the `user-select` style property.
func UserSelect(value ...any) aitch.Node {
	return styleAttribute(propUserSelect, value...)
}

// PointerEvents sets the `pointer-events` style property.
func PointerEvents(value ...any) aitch.Node {
	return styleAttribute(propPointerEvents, value...)
}

// Isolation sets the `isolation` style property.
func Isolation(value ...any) aitch.Node {
	return styleAttribute(propIsolation, value...)
}

// ObjectFit sets the `object-fit` style property.
func ObjectFit(value ...any) aitch.Node {
	return styleAttribute(propObjectFit, value...)
}

// ObjectPosition sets the `object-position` style property.
func ObjectPosition(value ...any) aitch.Node {
	return styleAttribute(propObjectPosition, value...)
}

// BoxSizing sets the `box-sizing` style property.
func BoxSizing(value ...any) aitch.Node {
	return styleAttribute(propBoxSizing, value...)
}

// AspectRatio sets the `aspect-ratio` style property.
func AspectRatio(value ...any) aitch.Node {
	return styleAttribute(propAspectRatio, value...)
}

// Contain sets the `contain` style property.
func Contain(value ...any) aitch.Node {
	return styleAttribute(propContain, value...)
}

// ContentVisibility sets the `content-visibility` style property.
func ContentVisibility(value ...any) aitch.Node {
	return styleAttribute(propContentVisibility, value...)
}

// Resize sets the `resize` style property.
func Resize(value ...any) aitch.Node {
	return styleAttribute(propResize, value...)
}

// GridTemplateColumns sets the `grid-template-columns` style property.
func GridTemplateColumns(value ...any) aitch.Node {
	return styleAttribute(propGridTemplateColumns, value...)
}

// GridTemplateRows sets the `grid-template-rows` style property.
func GridTemplateRows(value ...any) aitch.Node {
	return styleAttribute(propGridTemplateRows, value...)
}

// GridColumn sets the `grid-column` style property.
func GridColumn(value ...any) aitch.Node {
	return styleAttribute(propGridColumn, value...)
}

// GridRow sets the `grid-row` style property.
func GridRow(value ...any) aitch.Node {
	return styleAttribute(propGridRow, value...)
}

// GridArea sets the `grid-area` style property.
func GridArea(value ...any) aitch.Node {
	return styleAttribute(propGridArea, value...)
}

// GridAutoFlow sets the `grid-auto-flow` style property.
func GridAutoFlow(value ...any) aitch.Node {
	return styleAttribute(propGridAutoFlow, value...)
}

// PlaceItems sets the `place-items` style property.
func PlaceItems(value ...any) aitch.Node {
	return styleAttribute(propPlaceItems, value...)
}

// PlaceContent sets the `place-content` style property.
func PlaceContent(value ...any) aitch.Node {
	return styleAttribute(propPlaceContent, value...)
}

// PlaceSelf sets the `place-self` style property.
func PlaceSelf(value ...any) aitch.Node {
	return styleAttribute(propPlaceSelf, value...)
}

// FontStyle sets the `font-style` style property.
func FontStyle(value ...any) aitch.Node {
	return styleAttribute(propFontStyle, value...)
}

// FontVariant sets the `font-variant` style property.
func FontVariant(value ...any) aitch.Node {
	return styleAttribute(propFontVariant, value...)
}

// FontStretch sets the `font-stretch` style property.
func FontStretch(value ...any) aitch.Node {
	return styleAttribute(propFontStretch, value...)
}

// Direction sets the `direction` style property.
func Direction(value ...any) aitch.Node {
	return styleAttribute(propDirection, value...)
}

// TabSize sets the `tab-size` style property.
func TabSize(value ...any) aitch.Node {
	return styleAttribute(propTabSize, value...)
}

// ListStyle sets the `list-style` style property.
func ListStyle(value ...any) aitch.Node {
	return styleAttribute(propListStyle, value...)
}

// ListStyleType sets the `list-style-type` style property.
func ListStyleType(value ...any) aitch.Node {
	return styleAttribute(propListStyleType, value...)
}

// ListStylePosition sets the `list-style-position` style property.
func ListStylePosition(value ...any) aitch.Node {
	return styleAttribute(propListStylePosition, value...)
}

// ListStyleImage sets the `list-style-image` style property.
func ListStyleImage(value ...any) aitch.Node {
	return styleAttribute(propListStyleImage, value...)
}

// BorderCollapse sets the `border-collapse` style property.
func BorderCollapse(value ...any) aitch.Node {
	return styleAttribute(propBorderCollapse, value...)
}

// BorderSpacing sets the `border-spacing` style property.
func BorderSpacing(value ...any) aitch.Node {
	return styleAttribute(propBorderSpacing, value...)
}

// TableLayout sets the `table-layout` style property.
func TableLayout(value ...any) aitch.Node {
	return styleAttribute(propTableLayout, value...)
}

// EmptyCells sets the `empty-cells` style property.
func EmptyCells(value ...any) aitch.Node {
	return styleAttribute(propEmptyCells, value...)
}

// Filter sets the `filter` style property.
func Filter(value ...any) aitch.Node {
	return styleAttribute(propFilter, value...)
}

// BackdropFilter sets the `backdrop-filter` style property.
func BackdropFilter(value ...any) aitch.Node {
	return styleAttribute(propBackdropFilter, value...)
}

// MixBlendMode sets the `mix-blend-mode` style property.
func MixBlendMode(value ...any) aitch.Node {
	return styleAttribute(propMixBlendMode, value...)
}

// BoxDecorationBreak sets the `box-decoration-break` style property.
func BoxDecorationBreak(value ...any) aitch.Node {
	return styleAttribute(propBoxDecorationBreak, value...)
}

// ScrollBehavior sets the `scroll-behavior` style property.
func ScrollBehavior(value ...any) aitch.Node {
	return styleAttribute(propScrollBehavior, value...)
}

// ScrollSnapType sets the `scroll-snap-type` style property.
func ScrollSnapType(value ...any) aitch.Node {
	return styleAttribute(propScrollSnapType, value...)
}

// ScrollSnapAlign sets the `scroll-snap-align` style property.
func ScrollSnapAlign(value ...any) aitch.Node {
	return styleAttribute(propScrollSnapAlign, value...)
}

// ScrollMargin sets the `scroll-margin` style property.
func ScrollMargin(value ...any) aitch.Node {
	return styleAttribute(propScrollMargin, value...)
}

// ScrollPadding sets the `scroll-padding` style property.
func ScrollPadding(value ...any) aitch.Node {
	return styleAttribute(propScrollPadding, value...)
}

// AnimationName sets the `animation-name` style property.
func AnimationName(value ...any) aitch.Node {
	return styleAttribute(propAnimationName, value...)
}

// AnimationDuration sets the `animation-duration` style property.
func AnimationDuration(value ...any) aitch.Node {
	return styleAttribute(propAnimationDuration, value...)
}

// AnimationTimingFunction sets the `animation-timing-function` style property.
func AnimationTimingFunction(value ...any) aitch.Node {
	return styleAttribute(propAnimationTimingFunction, value...)
}

// AnimationDelay sets the `animation-delay` style property.
func AnimationDelay(value ...any) aitch.Node {
	return styleAttribute(propAnimationDelay, value...)
}

// AnimationIterationCount sets the `animation-iteration-count` style property.
func AnimationIterationCount(value ...any) aitch.Node {
	return styleAttribute(propAnimationIterationCount, value...)
}

// AnimationDirection sets the `animation-direction` style property.
func AnimationDirection(value ...any) aitch.Node {
	return styleAttribute(propAnimationDirection, value...)
}

// AnimationFillMode sets the `animation-fill-mode` style property.
func AnimationFillMode(value ...any) aitch.Node {
	return styleAttribute(propAnimationFillMode, value...)
}

// AnimationPlayState sets the `animation-play-state` style property.
func AnimationPlayState(value ...any) aitch.Node {
	return styleAttribute(propAnimationPlayState, value...)
}

// TransitionProperty sets the `transition-property` style property.
func TransitionProperty(value ...any) aitch.Node {
	return styleAttribute(propTransitionProperty, value...)
}

// TransitionDelay sets the `transition-delay` style property.
func TransitionDelay(value ...any) aitch.Node {
	return styleAttribute(propTransitionDelay, value...)
}

// TouchAction sets the `touch-action` style property.
func TouchAction(value ...any) aitch.Node {
	return styleAttribute(propTouchAction, value...)
}

// WillChange sets the `will-change` style property.
func WillChange(value ...any) aitch.Node {
	return styleAttribute(propWillChange, value...)
}

// CaretColor sets the `caret-color` style property.
func CaretColor(value ...any) aitch.Node {
	return styleAttribute(propCaretColor, value...)
}

// ScrollSnapStop sets the `scroll-snap-stop` style property.
func ScrollSnapStop(value ...any) aitch.Node {
	return styleAttribute(propScrollSnapStop, value...)
}

// Speak sets the `speak` style property.
func Speak(value ...any) aitch.Node {
	return styleAttribute(propSpeak, value...)
}

// PageBreakBefore sets the `page-break-before` style property.
func PageBreakBefore(value ...any) aitch.Node {
	return styleAttribute(propPageBreakBefore, value...)
}

// PageBreakAfter sets the `page-break-after` style property.
func PageBreakAfter(value ...any) aitch.Node {
	return styleAttribute(propPageBreakAfter, value...)
}

// PageBreakInside sets the `page-break-inside` style property.
func PageBreakInside(value ...any) aitch.Node {
	return styleAttribute(propPageBreakInside, value...)
}

func styleAttribute(name []byte, value ...any) aitch.Node {
	return aitch.NewDelimitedAttribute(attrStyle, stylesDelim, aitch.NewConcatValue(append([]any{name}, value...)...))
}
