package iotmaker_platform_webbrowser

import (
	"errors"
)

func NewStage(id string, width, height int, density float64, coordinateSystem ICoordinate) (error, Stage) {

	if density <= 0 {
		return errors.New("density must be greater then 0"), Stage{}
	}

	stage := Stage{
		X:      coordinateSystem,
		Y:      coordinateSystem,
		Width:  coordinateSystem,
		Height: coordinateSystem,
	}
	stage.Canvas = NewCanvasWith2DContext(id, width, height)

	stage.Canvas.AppendElementToDocumentBody()

	return nil, stage
}
