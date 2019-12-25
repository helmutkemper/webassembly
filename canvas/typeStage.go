package canvas

import (
	iotmaker_platform_coordinate "github.com/helmutkemper/iotmaker.platform.coordinate"
)

// todo: density
type Stage struct {
	Canvas
	ScratchPad     Canvas
	Density        float64
	originalX      float64
	originalY      float64
	originalHeight float64
	originalWidth  float64
	X              iotmaker_platform_coordinate.Density
	Y              iotmaker_platform_coordinate.Density
	Width          iotmaker_platform_coordinate.Density
	Height         iotmaker_platform_coordinate.Density
}

func (el *Stage) SetDensityFactor(density float64) {
	el.Density = density
	el.Width.SetDensityFactor(density)
	el.Height.SetDensityFactor(density)
}

func (el *Stage) SetX(x float64) {
	el.originalX = x
	el.X.Set(x)
}

func (el *Stage) SetY(y float64) {
	el.originalY = y
	el.Y.Set(y)
}

func (el *Stage) SetWidth(width float64) {
	el.originalWidth = width
	el.Width.Set(width)
}

func (el *Stage) SetHeight(height float64) {
	el.originalHeight = height
	el.Height.Set(height)
}
