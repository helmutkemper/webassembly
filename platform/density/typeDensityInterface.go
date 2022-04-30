package iotmaker_platform_coordinate

type IDensity interface {
	Set(value float64)
	SetInt(value int)
	Add(value int)
	Sub(value int)
	SetDensityFactor(value interface{})
	Int() int
	Float64() float64
	Float32() float32
	Float() float64
	String() string
}
