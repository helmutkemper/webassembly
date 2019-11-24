package iotmaker_platform_webbrowser

import (
	"errors"
)

func NewStage(id string, x, y, width, height int, density float64) (error, Stage) {

	if density <= 0 {
		return errors.New("density must be greater then 0"), Stage{}
	}

	//ele := NewElement()
	//ele.Create("canvas", id)
	//ele.AppendElementToDocumentBody()

	//ele.Create("p", "id")
	//ele.AppendElementToDocumentBody()

	stage := Stage{}
	stage.Canvas = NewCanvasWith2DContext(id, width, height)
	//stage.SetX( x )
	//stage.SetY( y )
	//stage.SetWidth( width )
	//stage.SetHeight( height )

	stage.Canvas.BeginPath()
	stage.Canvas.MoveTo(0, 0)
	stage.Canvas.LineTo(width, height)
	stage.Canvas.Fill()

	stage.Canvas.AppendElementToDocumentBody()

	return nil, stage
}
