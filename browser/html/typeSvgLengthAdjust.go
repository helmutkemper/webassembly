package html

type SvgLengthAdjust string

func (e SvgLengthAdjust) String() string {
	return string(e)
}

const (
	KSvgLengthAdjustSpacing          SvgLengthAdjust = "spacing"
	KSvgLengthAdjustSpacingAndGlyphs SvgLengthAdjust = "spacingAndGlyphs"
)
