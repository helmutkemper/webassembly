package html

type SvgTransformOrigin string

func (e SvgTransformOrigin) String() string {
	return string(e)
}

const (
	KSvgTransformOriginLeft   SvgTransformOrigin = "left"
	KSvgTransformOriginCenter SvgTransformOrigin = "center"
	KSvgTransformOriginRight  SvgTransformOrigin = "right"
	KSvgTransformOriginTop    SvgTransformOrigin = "top"
	KSvgTransformOriginBottom SvgTransformOrigin = "bottom"
)
