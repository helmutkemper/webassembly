package html

type SvgTypeFeFunc string

func (e SvgTypeFeFunc) String() string {
	return string(e)
}

const (
	KSvgTypeFeFuncIdentity SvgTypeFeFunc = "identity"
	KSvgTypeFeFuncTable    SvgTypeFeFunc = "table"
	KSvgTypeFeFuncDiscrete SvgTypeFeFunc = "discrete"
	KSvgTypeFeFuncLinear   SvgTypeFeFunc = "linear"
	KSvgTypeFeFuncGamma    SvgTypeFeFunc = "gamma"
)
