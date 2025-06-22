package rulesDesity

import (
	"strconv"
)

var density Density = 1.0

type Density float64

func (e Density) GetInt() int {
	return int(e * density)
}

func (e Density) GetFloat() float64 {
	return float64(e * density)
}

func (e Density) String() string {
	return strconv.FormatFloat(float64(e*density), 'g', -1, 32)
}

func (e Density) Pixel() string {
	return strconv.FormatFloat(float64(e*density), 'g', -1, 32) + "px"
}
