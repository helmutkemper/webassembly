package html

type SvgAlignmentBaseline string

func (e SvgAlignmentBaseline) String() string {
	return string(e)
}

const (
	KSvgAlignmentBaselineAuto           SvgAlignmentBaseline = "auto"
	KSvgAlignmentBaselineBaseline       SvgAlignmentBaseline = "baseline"
	KSvgAlignmentBaselineBeforeEdge     SvgAlignmentBaseline = "before-edge"
	KSvgAlignmentBaselineTextBeforeEdge SvgAlignmentBaseline = "text-before-edge"
	KSvgAlignmentBaselineMiddle         SvgAlignmentBaseline = "middle"
	KSvgAlignmentBaselineCentral        SvgAlignmentBaseline = "central"
	KSvgAlignmentBaselineAfterEdge      SvgAlignmentBaseline = "after-edge"
	KSvgAlignmentBaselineTextAfterEdge  SvgAlignmentBaseline = "text-after-edge"
	KSvgAlignmentBaselineIdeographic    SvgAlignmentBaseline = "ideographic"
	KSvgAlignmentBaselineAlphabetic     SvgAlignmentBaseline = "alphabetic"
	KSvgAlignmentBaselineHanging        SvgAlignmentBaseline = "hanging"
	KSvgAlignmentBaselineMathematical   SvgAlignmentBaseline = "mathematical"
	KSvgAlignmentBaselineInherit        SvgAlignmentBaseline = "inherit"
)
