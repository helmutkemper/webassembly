package iotmaker_platform_webbrowser

import (
	"errors"
)

// fixme: density compatibility
// fixme: return must be an interface
func NewStage(id string, width, height int, density float64) (error, Stage) {

	if density <= 0 {
		return errors.New("density must be greater then 0"), Stage{}
	}

	stage := Stage{}
	stage.Canvas = NewCanvasWith2DContext(id, width, height)

	stage.Canvas.AppendElementToDocumentBody()

	return nil, stage
}
