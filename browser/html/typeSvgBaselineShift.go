package html

type SvgBaselineShift string

func (e SvgBaselineShift) String() string {
	return string(e)
}

const (
	KSvgBaselineShiftAuto     SvgBaselineShift = "auto"
	KSvgBaselineShiftBaseline SvgBaselineShift = "baseline"
	KSvgBaselineShiftSuper    SvgBaselineShift = "super"
	KSvgBaselineShiftSub      SvgBaselineShift = "sub"
	KSvgBaselineShiftInherit  SvgBaselineShift = "inherit"
)
