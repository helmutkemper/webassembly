package html

type SvgAccumulate string

func (e SvgAccumulate) String() string {
	return string(e)
}

const (

	// KSvgAccumulateSum
	//
	// English:
	//
	//  Specifies that each repeat iteration after the first builds upon the last value of the previous iteration.
	//
	// Português:
	//
	//  Especifica que cada iteração repetida após a primeira se baseia no último valor da iteração anterior.
	KSvgAccumulateSum SvgAccumulate = "sum"

	// KSvgAccumulateNone
	//
	// English:
	//
	//  Specifies that repeat iterations are not cumulative.
	//
	// Português:
	//
	//  Especifica que as iterações repetidas não são cumulativas.
	KSvgAccumulateNone SvgAccumulate = "none"
)
