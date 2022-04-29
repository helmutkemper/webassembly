package tween

import "math"

var KEaseInOutBack = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	currentPercentage = currentPercentage / 0.5

	if currentPercentage < 1 {
		return 0.5*(math.Pow(currentPercentage, 2.0)*(3.5949095*currentPercentage-2.5949095))*delta + startValue
	}
	currentPercentage -= 2
	return 0.5*(math.Pow(currentPercentage, 2.0)*(3.5949095*currentPercentage+2.5949095)+2)*delta + startValue
}
