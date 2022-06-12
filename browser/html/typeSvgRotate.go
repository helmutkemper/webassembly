package html

type SvgRotate string

func (e SvgRotate) String() string {
	return string(e)
}

const (
	KSvgRotateAuto        SvgRotate = "auto"
	KSvgRotateAutoReverse SvgRotate = "auto-reverse"
)
