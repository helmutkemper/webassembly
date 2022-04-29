package tween

import "math"

var KEaseOutElastic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	value := interactionCurrent / interactionTotal

	if value == 0 {
		return startValue
	}

	if value == 1.0 {
		return 1.0*delta + startValue
	}

	return ((math.Pow(2.0, -10.0*value)*math.Sin((value-0.075)*20.9435102))+1.0)*delta + startValue
}
