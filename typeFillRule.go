package iotmaker_platform_webbrowser

type CanvasFillRule int

var CanvasFillRules = [...]string{
	"",
	"nonzero",
	"evenodd",
}

func (el CanvasFillRule) String() string {
	return CanvasFillRules[el]
}

const (
	// en: In two-dimensional computer graphics, the non-zero winding rule is a means of determining whether a given point
	// falls within an enclosed curve. Unlike the similar even-odd rule, it relies on knowing the direction of stroke for
	// each part of the curve.
	KFillRuleNonZero CanvasFillRule = iota + 1

	// en: The even–odd rule is an algorithm implemented in vector-based graphic software,[1] like the PostScript language
	// and Scalable Vector Graphics (SVG), which determines how a graphical shape with more than one closed outline will
	// be filled. Unlike the nonzero-rule algorithm, this algorithm will alternatively color and leave uncolored shapes
	// defined by nested closed paths irrespective of their winding.
	KFillRuleEvenOdd
)
