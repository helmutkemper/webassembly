package easingTween

import "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"

var i = -1

func SelectRandom() func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	var list = []func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64{
		KEaseInOutBounce,
		KEaseInBack,
		KEaseInBounce,
		KEaseInElastic,
		KEaseInOutBack,
		KEaseInOutBounce,
		KEaseOutBack,
		KEaseOutBounce,
		KEaseOutElastic,
		KEaseInOutExponential,
		KEaseInOutQuadratic,
		KEaseInOutQuartic,
		KEaseInOutQuintic,
		KEaseInOutSine,
		KEaseInQuadratic,
		KEaseInQuartic,
		KEaseInQuintic,
		KEaseInSine,
		KEaseOutCircular,
		KEaseOutCubic,
		KEaseOutExponential,
		KEaseOutQuadratic,
		KEaseOutQuartic,
		KEaseOutQuintic,
		KEaseOutSine,
		KLinear,
		KEaseInCircular,
		KEaseInCubic,
		KEaseInExponential,
		KEaseInOutCircular,
		KEaseInOutCubic,
	}

	i := mathUtil.Int(0, len(list)-1)

	return list[i]
}
