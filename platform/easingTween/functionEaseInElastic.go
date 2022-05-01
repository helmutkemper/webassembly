package easingTween

import "math"

var KEaseInElastic = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {

	if currentPercentage == 0 {
		return startValue
	}

	if currentPercentage == 1.0 {
		return 1.0*delta + startValue
	}

	currentPercentage -= 1

	return -(math.Pow(2, 10*currentPercentage)*math.Sin((currentPercentage-0.075)*20.9435102))*delta + startValue
}
