package easingTween

var KEaseInBounce = func(interactionCurrent, interactionTotal, currentPercentage, startValue, endValue, delta float64) float64 {
	//fixme: melhorar
	/*
			end -= start;
		  float d = 1f;
		  return end - EaseOutBounce(0, end, d - value) + start;
	*/
	value := interactionCurrent / interactionTotal
	endValue -= startValue
	return endValue - EaseOutBounce(0, endValue, value) + startValue
}

func EaseOutBounce(start, end, value float64) float64 {
	end -= start
	if value < 1.0/2.75 {
		return end*(7.5625*value*value) + start
	} else if value < 2.0/2.75 {
		value -= 1.5 / 2.75
		return end*(7.5625*value*value+0.75) + start
	} else if value < 2.5/2.75 {
		value -= 2.25 / 2.75
		return end*(7.5625*value*value+0.9375) + start
	} else {
		value -= 2.625 / 2.75
		return end*(7.5625*value*value+0.984375) + start
	}
}
