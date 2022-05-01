package easingTween

import "math"

// en: sinusoidal easing in - accelerating from zero velocity
var KEaseInSine = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return -1*delta*math.Cos(currentPercentage*1.570796) + delta + startValue
}
