package iotmaker_platform_coordinate

import "strconv"

type Density struct {
	OriginalValue int
	DensityFactor int
	DensityValue  int
}

func (el *Density) adjustDensity() {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.DensityValue = el.OriginalValue * el.DensityFactor
}

func (el *Density) Set(value float64) {
	el.OriginalValue = int(value)
	el.adjustDensity()
}

func (el *Density) SetInt(value int) {
	el.OriginalValue = value
	el.adjustDensity()
}

func (el *Density) Add(value int) {
	el.OriginalValue += value
	el.adjustDensity()
}

func (el *Density) Sub(value int) {
	el.OriginalValue -= value
	el.adjustDensity()
}

func (el *Density) SetDensityFactor(value interface{}) {

	switch converted := value.(type) {
	case float64:
		el.DensityFactor = int(converted)
	case int:
		el.DensityFactor = converted
	}

	el.adjustDensity()
}

func (el Density) Int() int {
	return int(el.DensityValue)
}

func (el Density) Float64() float64 {
	return float64(el.DensityValue)
}

func (el Density) Float32() float32 {
	return float32(el.DensityValue)
}

func (el Density) Float() float64 {
	return float64(el.DensityValue)
}

func (el Density) String() string {
	return strconv.FormatInt(int64(el.OriginalValue), 10)
}
