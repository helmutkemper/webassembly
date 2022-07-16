package coordinatesystem

type IDensity interface {
	Set(value float64)
	Add(value float64)
	Sub(value float64)
	SetDensityFactor(value float64)
	Get() float64
	String() string
}
