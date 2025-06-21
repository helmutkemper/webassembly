package rulesDesity

import (
	"strconv"
)

var density = 1.0

type Density struct {
	value float64
}

func (e *Density) GetInt() int {
	return int(e.value * density)
}

func (e *Density) GetFloat() float64 {
	return e.value * density
}

func (e *Density) SetInt(value int) {
	e.value = float64(value)
}

func (e *Density) SetFloat(value float64) {
	e.value = value
}

func (e *Density) String() string {
	return strconv.FormatFloat(e.value*density, 'g', -1, 32)
}

func (e *Density) Pixel() string {
	return strconv.FormatFloat(e.value*density, 'g', -1, 32) + "px"
}
