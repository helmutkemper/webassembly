package tween

import "math"

// en: circular easing in/out - acceleration until halfway, then deceleration
var KEaseInOutCircular = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage = currentPercentage * 2
	if currentPercentage < 1 {
		return -1*delta/2*(math.Sqrt(1-math.Pow(currentPercentage, 2.0))-1) + startValue
	}
	currentPercentage -= 2
	return delta/2*(math.Sqrt(1-math.Pow(currentPercentage, 2.0))+1) + startValue
}
