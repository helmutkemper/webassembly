package coordinatesystem

import "strconv"

type Density struct {
	OriginalValue float64
	DensityFactor float64
	DensityValue  float64
}

func (e *Density) adjustDensity() {
	if e.DensityFactor == 0 {
		e.DensityFactor = 1
	}

	e.DensityValue = e.OriginalValue * e.DensityFactor
}

func (e *Density) Set(value float64) {
	e.OriginalValue = value
	e.adjustDensity()
}

func (e *Density) Add(value float64) {
	e.OriginalValue += value
	e.adjustDensity()
}

func (e *Density) Sub(value float64) {
	e.OriginalValue -= value
	e.adjustDensity()
}

func (e *Density) SetDensityFactor(value float64) {
	e.DensityFactor = value
	e.adjustDensity()
}

func (e Density) Get() float64 {
	return e.DensityValue
}

func (e Density) String() string {
	return strconv.FormatFloat(e.OriginalValue, 'g', -1, 64)
}
