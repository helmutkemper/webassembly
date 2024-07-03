package easingTween

var KEaseInBounce = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	x := currentPercentage
	p := 1 - bounceOut(1-x)
	return delta*p + startValue
}
