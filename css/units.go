package css

import "github.com/go-andiamo/aitch"

var (
	unitPx      = []byte("px")
	unitPt      = []byte("pt")
	unitPc      = []byte("pc")
	unitInch    = []byte("in")
	unitCm      = []byte("cm")
	unitMm      = []byte("mm")
	unitEm      = []byte("em")
	unitRem     = []byte("rem")
	unitPercent = []byte("%")
	unitVw      = []byte("vw")
	unitVh      = []byte("vh")
	unitVmin    = []byte("vmin")
	unitVmax    = []byte("vmax")
	unitCh      = []byte("ch")
	unitEx      = []byte("ex")
)

// Px creates a value with "px" (pixels)
func Px(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitPx)...)
}

// Pt creates a value with "pt" (points)
func Pt(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitPt)...)
}

// Pc creates a value with "pc" (picas)
func Pc(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitPc)...)
}

// In creates a value with "in" (inches)
func In(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitInch)...)
}

// Cm creates a value with "cm" (centimeters)
func Cm(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitCm)...)
}

// Mm creates a value with "mm" (millimeters)
func Mm(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitMm)...)
}

// Em creates a value with "em" (relative to parent font size)
func Em(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitEm)...)
}

// Rem creates a value with "rem" (relative to root font size)
func Rem(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitRem)...)
}

// Percent creates a value with "%" (unitPercent)
func Percent(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitPercent)...)
}

// Vw creates a value with "vw" (viewport width)
func Vw(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitVw)...)
}

// Vh creates a value with "vh" (viewport height)
func Vh(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitVh)...)
}

// VMin creates a value with "vmin" (smaller of vw or vh)
func VMin(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitVmin)...)
}

// VMax creates a value with "vmax" (larger of vw or vh)
func VMax(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitVmax)...)
}

// Ch creates a value with "ch" (width of 0 glyph)
func Ch(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitCh)...)
}

// Ex creates a value with "ex" (x-height of current font)
func Ex(value ...any) aitch.Value {
	return aitch.NewConcatValue(append(value, unitEx)...)
}
