package easingTween

import "math"

// en: Quadratic easing in/out - acceleration until halfway, then deceleration
var KEaseInOutQuadratic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	interactionCurrent /= interactionTotal / 2
	if interactionCurrent < 1 {
		return delta/2*math.Pow(interactionCurrent, 2.0) + startValue
	}
	interactionCurrent--
	return -1*delta/2*(interactionCurrent*(interactionCurrent-2)-1) + startValue
}
