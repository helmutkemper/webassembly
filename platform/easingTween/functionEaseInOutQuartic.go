package easingTween

import "math"

// en: quartic easing in/out - acceleration until halfway, then deceleration
var KEaseInOutQuartic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage = currentPercentage * 2
	if currentPercentage < 1 {
		return delta/2*math.Pow(currentPercentage, 4.0) + startValue
	}
	currentPercentage -= 2
	return -1*delta/2*(math.Pow(currentPercentage, 4.0)-2) + startValue
}
