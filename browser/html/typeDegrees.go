package html

import (
	"math"
	"strconv"
)

type Degrees float64

func (e Degrees) String() string {
	return strconv.FormatFloat(float64(e), 'g', -1, 64) + "deg"
}

func (e *Degrees) SetRad(value float64) {
	*e = Degrees(value * math.Pi * 180.0)
}

func (e Degrees) GetRad() (radians float64) {
	return float64(e * math.Pi / 180.0)
}

func (e Degrees) Get() (degrees float64) {
	return float64(e)
}
