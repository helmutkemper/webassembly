package easingTween

import "math"

// en: quartic easing in - accelerating from zero velocity
var KEaseInQuartic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	interactionCurrent /= interactionTotal
	return delta*math.Pow(interactionCurrent, 4.0) + startValue
}
