package iotmaker_platform_webbrowser

import iotmaker_types "github.com/helmutkemper/iotmaker.types"

type Stage struct {
	Canvas
	Density        float64
	originalX      int
	originalY      int
	originalHeight int
	originalWidth  int
	X              iotmaker_types.Coordinate
	Y              iotmaker_types.Coordinate
	Width          iotmaker_types.Coordinate
	Height         iotmaker_types.Coordinate
}

func (el *Stage) SetDensityFactor(density float64) {
	el.Density = density
	el.Width.SetDensityFactor(density)
	el.Height.SetDensityFactor(density)
}

func (el *Stage) SetX(x int) {
	el.originalX = x
	el.X.Set(x)
}

func (el *Stage) SetY(y int) {
	el.originalY = y
	el.Y.Set(y)
}

func (el *Stage) SetWidth(width int) {
	el.originalWidth = width
	el.Width.Set(width)
}

func (el *Stage) SetHeight(height int) {
	el.originalHeight = height
	el.Height.Set(height)
}
