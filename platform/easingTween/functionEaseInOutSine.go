package easingTween

import "math"

// en: sinusoidal easing in/out - accelerating until halfway, then decelerating
var KEaseInOutSine = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	return -1*delta/2*(math.Cos(math.Pi*currentPercentage)-1) + startValue
}
