package easingTween

import "math"

var KEaseOutBounce = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {

	if currentPercentage < 0.36363636363 {
		formula := 7.5625 * math.Pow(currentPercentage, 2.0)
		return formula*delta + startValue
	} else if currentPercentage < 0.72727272727 {
		currentPercentage -= 0.54545454545
		formula := 7.5625*math.Pow(currentPercentage, 2.0) + 0.75
		return formula*delta + startValue
	} else if currentPercentage < 0.90909090909 {
		currentPercentage -= 0.81818181818
		formula := 7.5625*math.Pow(currentPercentage, 2.0) + 0.9375
		return formula*delta + startValue
	} else {
		currentPercentage -= 0.95454545454
		formula := 7.5625*math.Pow(currentPercentage, 2.0) + 0.984375
		return formula*delta + startValue
	}
}
