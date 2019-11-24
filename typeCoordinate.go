package iotmaker_platform_webbrowser

import "strconv"

type Coordinate struct {
	OriginalValue int
	DensityFactor float64
	DensityValue  float64
}

func (el *Coordinate) adjustDensity() {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.DensityValue = float64(el.OriginalValue) * el.DensityFactor
}

func (el *Coordinate) Set(value int) {
	el.OriginalValue = value
	el.adjustDensity()
}

func (el *Coordinate) Add(value int) {
	el.OriginalValue += value
	el.adjustDensity()
}

func (el *Coordinate) Sub(value int) {
	el.OriginalValue -= value
	el.adjustDensity()
}

func (el *Coordinate) SetDensityFactor(value float64) {
	el.DensityFactor = value
	el.adjustDensity()
}

func (el Coordinate) Int() int {
	return int(el.DensityValue)
}

func (el Coordinate) Float64() float64 {
	return el.DensityValue
}

func (el Coordinate) Float32() float32 {
	return float32(el.DensityValue)
}

func (el Coordinate) Float() float64 {
	return el.DensityValue
}

func (el Coordinate) String() string {
	return strconv.FormatInt(int64(el.OriginalValue), 10)
}
