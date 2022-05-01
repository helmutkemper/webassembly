package easingTween

import "math"

var KEaseInOutElastic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {

	value := interactionCurrent / interactionTotal

	if value == 0 {
		return startValue
	}

	value /= 0.5
	if value == 2 {
		return endValue
	}

	if value < 1.0 {
		value -= 1
		return -0.5*(math.Pow(2.0, 10.0*value)*math.Sin((value-0.075)*20.9435102))*delta + startValue
	}

	value -= 1
	return math.Pow(2.0, -10.0*value)*math.Sin((value-0.075)*20.9435102)*0.5*delta + endValue
}
