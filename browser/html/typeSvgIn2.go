package html

type SvgIn2 string

func (e SvgIn2) String() string {
	return string(e)
}

const (
	KSvgIn2SourceGraphic   SvgIn2 = "SourceGraphic"
	KSvgIn2SourceAlpha     SvgIn2 = "SourceAlpha"
	KSvgIn2BackgroundImage SvgIn2 = "BackgroundImage"
	KSvgIn2BackgroundAlpha SvgIn2 = "BackgroundAlpha"
	KSvgIn2FillPaint       SvgIn2 = "FillPaint"
	KSvgIn2StrokePaint     SvgIn2 = "StrokePaint"
)
