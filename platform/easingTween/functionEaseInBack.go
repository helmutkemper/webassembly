package easingTween

import "math"

var KEaseInBack = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return math.Pow(currentPercentage, 2.0)*(1.70158+1*currentPercentage-1.70158)*delta + startValue
}
