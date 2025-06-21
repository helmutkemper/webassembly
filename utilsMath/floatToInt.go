package utilsMath

import "math"

func FloatToInt(value float64) int {
	var roundOn = 0.5
	var places = 0.0

	var round float64
	pow := math.Pow(10, places)
	digit := pow * value
	_, div := math.Modf(digit)

	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}

	return int(round / pow)
}
