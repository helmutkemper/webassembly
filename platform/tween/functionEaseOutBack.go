package tween

import "math"

var KEaseOutBack = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage = currentPercentage - 1
	return (math.Pow(currentPercentage, 2.0)*(2.70158*currentPercentage+1.70158)+1)*delta + startValue
}
