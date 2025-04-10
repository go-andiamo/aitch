package css

const Important = " !important"

const (
	DisplayBlock       = "block"        // DisplayBlock renders the element as a block
	DisplayInline      = "inline"       // DisplayInline renders the element as an inline element
	DisplayInlineBlock = "inline-block" // DisplayInlineBlock renders as inline-block
	DisplayFlex        = "flex"         // DisplayFlex renders as a flex container
	DisplayGrid        = "grid"         // DisplayGrid renders as a grid container
	DisplayNone        = "none"         // DisplayNone hides the element from layout
)

const (
	PositionStatic   = "static"   // PositionStatic is the default positioning
	PositionRelative = "relative" // PositionRelative positions the element relative to its normal position
	PositionAbsolute = "absolute" // PositionAbsolute removes the element from normal flow and positions it absolutely
	PositionFixed    = "fixed"    // PositionFixed positions relative to the viewport
	PositionSticky   = "sticky"   // PositionSticky toggles between relative and fixed based on scroll
)

const (
	TextLeft    = "left"    // TextLeft aligns text to the left
	TextCenter  = "center"  // TextCenter centers text
	TextRight   = "right"   // TextRight aligns text to the right
	TextJustify = "justify" // TextJustify justifies text to fit the container
)

const (
	BorderNone   = "none"   // BorderNone disables borders
	BorderSolid  = "solid"  // BorderSolid creates a solid border
	BorderDashed = "dashed" // BorderDashed creates a dashed border
	BorderDotted = "dotted" // BorderDotted creates a dotted border
	BorderDouble = "double" // BorderDouble creates a double line border
	BorderGroove = "groove" // BorderGroove creates a grooved border
	BorderRidge  = "ridge"  // BorderRidge creates a ridged border
	BorderInset  = "inset"  // BorderInset makes the border appear embedded
	BorderOutset = "outset" // BorderOutset makes the border appear extruded
)

const (
	FontWeightNormal = "normal"  // FontWeightNormal uses default font weight
	FontWeightBold   = "bold"    // FontWeightBold renders bold text
	FontWeightLight  = "lighter" // FontWeightLight renders lighter text
	FontWeightBolder = "bolder"  // FontWeightBolder renders bolder text than inherited
)

const (
	JustifyStart   = "flex-start"    // JustifyStart aligns items to the start
	JustifyCenter  = "center"        // JustifyCenter centers items
	JustifyEnd     = "flex-end"      // JustifyEnd aligns items to the end
	JustifyBetween = "space-between" // JustifyBetween distributes items with space between
	JustifyAround  = "space-around"  // JustifyAround distributes items with space around
)

const (
	AlignStart    = "start"    // AlignStart aligns items to the start
	AlignCenter   = "center"   // AlignCenter centers items
	AlignEnd      = "end"      // AlignEnd aligns items to the end
	AlignStretch  = "stretch"  // AlignStretch stretches items to fill the container
	AlignBaseline = "baseline" // AlignBaseline aligns items along their text baseline
)

const (
	OverflowVisible = "visible" // OverflowVisible allows content to overflow
	OverflowHidden  = "hidden"  // OverflowHidden clips overflowing content
	OverflowScroll  = "scroll"  // OverflowScroll always shows scrollbars
	OverflowAuto    = "auto"    // OverflowAuto shows scrollbars only when needed
)

const (
	VisibilityVisible  = "visible"  // VisibilityVisible shows the element
	VisibilityHidden   = "hidden"   // VisibilityHidden hides the element (but keeps layout space)
	VisibilityCollapse = "collapse" // VisibilityCollapse hides the element and collapses layout (used in tables)
)

const (
	TextTransformNone       = "none"       // TextTransformNone disables text transformation
	TextTransformCapitalize = "capitalize" // TextTransformCapitalize capitalizes first letter of each word
	TextTransformUppercase  = "uppercase"  // TextTransformUppercase transforms text to all uppercase
	TextTransformLowercase  = "lowercase"  // TextTransformLowercase transforms text to all lowercase
)

const (
	WhiteSpaceNormal  = "normal"   // WhiteSpaceNormal collapses whitespace and wraps text
	WhiteSpaceNowrap  = "nowrap"   // WhiteSpaceNowrap collapses whitespace and prevents wrapping
	WhiteSpacePre     = "pre"      // WhiteSpacePre preserves whitespace and line breaks
	WhiteSpacePreWrap = "pre-wrap" // WhiteSpacePreWrap preserves whitespace and wraps text
	WhiteSpacePreLine = "pre-line" // WhiteSpacePreLine collapses whitespace but preserves line breaks
)

var (
	stylesDelim = []byte{';', ' '}
	attrStyle   = []byte("style")
	// Layout
	propDisplay  = []byte("display:")
	propPosition = []byte("position:")
	propTop      = []byte("top:")
	propRight    = []byte("right:")
	propBottom   = []byte("bottom:")
	propLeft     = []byte("left:")
	propZIndex   = []byte("z-index:")
	propFloat    = []byte("float:")
	propClear    = []byte("clear:")
	// Flexbox & Grid
	propFlex           = []byte("flex:")
	propFlexGrow       = []byte("flex-grow:")
	propFlexShrink     = []byte("flex-shrink:")
	propFlexBasis      = []byte("flex-basis:")
	propJustifyContent = []byte("justify-content:")
	propAlignItems     = []byte("align-items:")
	propAlignContent   = []byte("align-content:")
	propAlignSelf      = []byte("align-self:")
	propGap            = []byte("gap:")
	propRowGap         = []byte("row-gap:")
	propColumnGap      = []byte("column-gap:")
	// Sizing
	propWidth     = []byte("width:")
	propHeight    = []byte("height:")
	propMinWidth  = []byte("min-width:")
	propMaxWidth  = []byte("max-width:")
	propMinHeight = []byte("min-height:")
	propMaxHeight = []byte("max-height:")
	// Spacing
	propMargin        = []byte("margin:")
	propMarginTop     = []byte("margin-top:")
	propMarginRight   = []byte("margin-right:")
	propMarginBottom  = []byte("margin-bottom:")
	propMarginLeft    = []byte("margin-left:")
	propPadding       = []byte("padding:")
	propPaddingTop    = []byte("padding-top:")
	propPaddingRight  = []byte("padding-right:")
	propPaddingBottom = []byte("padding-bottom:")
	propPaddingLeft   = []byte("padding-left:")
	// Border
	propBorder       = []byte("border:")
	propBorderWidth  = []byte("border-width:")
	propBorderColor  = []byte("border-color:")
	propBorderStyle  = []byte("border-style:")
	propBorderRadius = []byte("border-radius:")
	propBorderTop    = []byte("border-top:")
	propBorderBottom = []byte("border-bottom:")
	propBorderLeft   = []byte("border-left:")
	propBorderRight  = []byte("border-right:")
	// Background
	propBackground         = []byte("background:")
	propBackgroundColor    = []byte("background-color:")
	propBackgroundImage    = []byte("background-image:")
	propBackgroundSize     = []byte("background-size:")
	propBackgroundPosition = []byte("background-position:")
	propBackgroundRepeat   = []byte("background-repeat:")
	// Typography
	propColor          = []byte("color:")
	propFontSize       = []byte("font-size:")
	propFontWeight     = []byte("font-weight:")
	propFontFamily     = []byte("font-family:")
	propLineHeight     = []byte("line-height:")
	propLetterSpacing  = []byte("letter-spacing:")
	propTextAlign      = []byte("text-align:")
	propTextTransform  = []byte("text-transform:")
	propTextDecoration = []byte("text-decoration:")
	propWhiteSpace     = []byte("white-space:")
	propWordBreak      = []byte("word-break:")
	// Visual
	propOpacity    = []byte("opacity:")
	propBoxShadow  = []byte("box-shadow:")
	propVisibility = []byte("visibility:")
	propOverflow   = []byte("overflow:")
	propOverflowX  = []byte("overflow-x:")
	propOverflowY  = []byte("overflow-y:")
	propCursor     = []byte("cursor:")
	// Animation & Transition
	propTransition         = []byte("transition:")
	propTransitionDuration = []byte("transition-duration:")
	propTransitionTiming   = []byte("transition-timing-function:")
	propTransform          = []byte("transform:")
	propAnimation          = []byte("animation:")
	// Misc
	propUserSelect     = []byte("user-select:")
	propPointerEvents  = []byte("pointer-events:")
	propIsolation      = []byte("isolation:")
	propObjectFit      = []byte("object-fit:")
	propObjectPosition = []byte("object-position:")
	// Display & Box
	propBoxSizing         = []byte("box-sizing:")
	propAspectRatio       = []byte("aspect-ratio:")
	propContain           = []byte("contain:")
	propContentVisibility = []byte("content-visibility:")
	propResize            = []byte("resize:")
	// Grid (you’ve got flex already)
	propGridTemplateColumns = []byte("grid-template-columns:")
	propGridTemplateRows    = []byte("grid-template-rows:")
	propGridColumn          = []byte("grid-column:")
	propGridRow             = []byte("grid-row:")
	propGridArea            = []byte("grid-area:")
	propGridAutoFlow        = []byte("grid-auto-flow:")
	propPlaceItems          = []byte("place-items:")
	propPlaceContent        = []byte("place-content:")
	propPlaceSelf           = []byte("place-self:")
	// Typography (extras)
	propFontStyle   = []byte("font-style:")
	propFontVariant = []byte("font-variant:")
	propFontStretch = []byte("font-stretch:")
	propDirection   = []byte("direction:")
	propTabSize     = []byte("tab-size:")
	// Lists
	propListStyle         = []byte("list-style:")
	propListStyleType     = []byte("list-style-type:")
	propListStylePosition = []byte("list-style-position:")
	propListStyleImage    = []byte("list-style-image:")
	// Tables
	propBorderCollapse = []byte("border-collapse:")
	propBorderSpacing  = []byte("border-spacing:")
	propTableLayout    = []byte("table-layout:")
	propEmptyCells     = []byte("empty-cells:")
	// Filters & Effects
	propFilter             = []byte("filter:")
	propBackdropFilter     = []byte("backdrop-filter:")
	propMixBlendMode       = []byte("mix-blend-mode:")
	propBoxDecorationBreak = []byte("box-decoration-break:")
	// Scroll & Overflow
	propScrollBehavior  = []byte("scroll-behavior:")
	propScrollSnapType  = []byte("scroll-snap-type:")
	propScrollSnapAlign = []byte("scroll-snap-align:")
	propScrollMargin    = []byte("scroll-margin:")
	propScrollPadding   = []byte("scroll-padding:")
	// Animations (extras)
	propAnimationName           = []byte("animation-name:")
	propAnimationDuration       = []byte("animation-duration:")
	propAnimationTimingFunction = []byte("animation-timing-function:")
	propAnimationDelay          = []byte("animation-delay:")
	propAnimationIterationCount = []byte("animation-iteration-count:")
	propAnimationDirection      = []byte("animation-direction:")
	propAnimationFillMode       = []byte("animation-fill-mode:")
	propAnimationPlayState      = []byte("animation-play-state:")
	// Transitions (extras)
	propTransitionProperty = []byte("transition-property:")
	propTransitionDelay    = []byte("transition-delay:")
	// Interactivity
	propTouchAction    = []byte("touch-action:")
	propWillChange     = []byte("will-change:")
	propCaretColor     = []byte("caret-color:")
	propScrollSnapStop = []byte("scroll-snap-stop:")
	// Accessibility
	propSpeak = []byte("speak:")
	// Print & Media
	propPageBreakBefore = []byte("page-break-before:")
	propPageBreakAfter  = []byte("page-break-after:")
	propPageBreakInside = []byte("page-break-inside:")
)
