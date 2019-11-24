package iotmaker_platform_webbrowser

type ICoordinate interface {
	Set(value int)
	Add(value int)
	Sub(value int)
	SetDensityFactor(value float64)
	Int() int
	Float64() float64
	Float32() float32
	Float() float64
	String() string
}
