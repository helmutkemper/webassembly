//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/algorithm"
	"github.com/helmutkemper/webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
	"github.com/helmutkemper/webassembly/platform/factoryEasingTween"
	"math"
	"time"
)

var canvas *html.TagCanvas

func main() {

	var stage = factoryBrowser.NewStage()

	canvas = factoryBrowser.NewTagCanvas(stage.GetWidth(), stage.GetHeight())
	stage.Append(canvas)

	var bezier = factoryAlgorithm.NewBezierCurve()

	border := 50.0
	wight := 400.0
	height := 400.0

	// E.g.: P0 (1,0) = (1*wight,0*height)
	// E.g.: P1 (2,0) = (2*wight,0*height)
	// E.g.: P2 (2,1) = (2*wight,1*height)
	//
	//     (0,0)            (1,0)            (2,0)
	//       +----------------+----------------+
	//       | P7            P0             P1 |
	//       |                                 |
	//       |                                 |
	//       |                                 |
	// (0,1) + P6                           P2 + (2,1)
	//       |                                 |
	//       |                                 |
	//       |                                 |
	//       | P5            P4             P3 |
	//       +----------------+----------------+
	//     (0,2)            (1,2)            (2,2)

	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 0*height + border})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Process()
	bezier.SetNumberOfSegments(2000)

	//bezier.GenerateRipple(20, 30)
	for _, point := range *bezier.GetProcessed() {
		AddDotBlue(int(point.X), int(point.Y))
	}

	div := factoryBrowser.NewTagDiv().
		Class("animate").
		AddPointsToEasingTween(bezier).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi / 2)
	stage.Append(div)

	wasm := factoryBrowser.NewTagDiv().
		Style("font-size:40px;color:#555555").
		SetXY(200, 360).
		Html("golang = wasm = fast javascript<br>" +
			"<br>" +
			"&nbsp;&nbsp;This rocket is the background<br>" +
			"&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;of a simple html div")
	stage.Append(wasm)

	factoryEasingTween.NewInOutBounce(
		20*time.Second,
		0,
		10000,
		div.EasingTweenWalkingAndRotateIntoPoints,
		-1,
	).
		SetArgumentsFunc(any(div)).
		SetDoNotReverseMotion()

	done := make(chan struct{}, 0)
	<-done
}

func AddDotBlue(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.5, 0, 2*math.Pi, false).
		Fill()
}
