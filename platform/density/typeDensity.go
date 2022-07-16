package coordinatesystem

import "strconv"

type Density struct {
	OriginalValue float64
	DensityFactor float64
	DensityValue  float64
}

func (el *Density) adjustDensity() {
	if el.DensityFactor == 0 {
		el.DensityFactor = 1
	}

	el.DensityValue = el.OriginalValue * el.DensityFactor
}

func (el *Density) Set(value float64) {
	el.OriginalValue = value
	el.adjustDensity()
}

func (el *Density) Add(value float64) {
	el.OriginalValue += value
	el.adjustDensity()
}

func (el *Density) Sub(value float64) {
	el.OriginalValue -= value
	el.adjustDensity()
}

func (el *Density) SetDensityFactor(value float64) {
	el.DensityFactor = value
	el.adjustDensity()
}

func (el Density) Get() float64 {
	return el.DensityValue
}

func (el Density) String() string {
	return strconv.FormatFloat(el.OriginalValue, 'g', -1, 64)
}
