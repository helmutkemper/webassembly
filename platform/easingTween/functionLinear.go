package easingTween

// en: simple linear tweening - no easing, no acceleration
var KLinear = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return delta*currentPercentage + startValue
}
