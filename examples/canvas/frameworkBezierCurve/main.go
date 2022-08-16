//go:build js

package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryAlgorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryEasingTween"
	"math"
	"strconv"
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
	adjust := 15.0

	// E.g.: P0 (1,0) = (1*wight,0*height)
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
	bezier.Add(algorithm.Point{X: 2*wight + border - adjust, Y: 0*height + border + adjust})
	bezier.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 2*wight + border - adjust, Y: 2*height + border - adjust})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border + adjust, Y: 2*height + border - adjust})
	bezier.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
	bezier.Add(algorithm.Point{X: 0*wight + border + adjust, Y: 0*height + border + adjust})
	bezier.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	bezier.Process(1000)

	var decimatesCurve = factoryAlgorithm.NewBezierCurve()
	decimatesCurve.Copy(bezier)
	decimatesCurve.SetNumberOfSegments(3)
	for _, point := range *decimatesCurve.GetProcessed() {
		AddDotYellow(int(point.X), int(point.Y))
	}

	var density = factoryAlgorithm.NewBezierCurve()
	density.Copy(decimatesCurve)
	density.IncreaseDensityBetweenPoints(300)
	for _, point := range *density.GetProcessed() {
		AddDotGreen(int(point.X), int(point.Y))
	}

	for v, point := range *bezier.GetOriginal() {
		AddRedPointer(int(point.X), int(point.Y))
		AddIndex(int(point.X), int(point.Y), v)
	}

	for _, point := range *bezier.GetProcessed() {
		AddDotBlue(int(point.X), int(point.Y))
	}

	var ripple = factoryAlgorithm.NewBezierCurve()
	ripple.Copy(bezier)
	ripple.GenerateRipple(20, 30)
	for _, point := range *ripple.GetProcessed() {
		AddDotPalegoldenrod(int(point.X), int(point.Y))
	}

	div := factoryBrowser.NewTagDiv().
		Class("animate").
		AddPointsToEasingTween(ripple).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi / 2).
		AppendToStage()

	factoryEasingTween.NewLinear(
		30*time.Second,
		0,
		10000,
		div.EasingTweenWalkingAndRotateIntoPoints,
		-1,
	).
		SetOnInvertFunc(onInvert).
		SetArgumentsFunc([]interface{}{div}).
		SetDoNotReverseMotion()

	done := make(chan struct{}, 0)
	<-done
}

func AddDotBlue(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.6, 0, 2*math.Pi, false).
		Fill()
}

func AddDotPalegoldenrod(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewPalegoldenrod()).
		Arc(x, y, 1.0, 0, 2*math.Pi, false).
		Fill()
}

func AddDotYellow(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewYellow()).
		Arc(x, y, 10.0, 0, 2*math.Pi, false).
		Fill()
}

func AddDotGreen(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewGreen()).
		Arc(x, y, 0.5, 0, 2*math.Pi, false).
		Fill()
}

func AddRedPointer(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewRedHalfTransparent()).
		Arc(x, y, 3, 0, 2*math.Pi, false).
		Fill()
}

func AddIndex(x, y, i int) {
	xStr := strconv.FormatInt(int64(x), 10)
	yStr := strconv.FormatInt(int64(y), 10)
	iStr := strconv.FormatInt(int64(i), 10)

	if i == 8 {
		y += 16
	}

	x += 5
	y += 20
	var font html.Font
	font.Family = factoryFontFamily.NewArial()
	font.Size = 17

	canvas.BeginPath().
		Font(font).
		FillStyle(factoryColor.NewRed()).
		FillText(
			"#"+iStr,
			x,
			y,
			300,
		)

	font.Size = 12
	canvas.BeginPath().
		Font(font).
		FillStyle(factoryColor.NewRed()).
		FillText(
			"("+xStr+", "+yStr+")",
			x+20,
			y,
			300,
		)
}

func onInvert(_ float64, args interface{}) {
	this := args.([]interface{})[0].(*html.TagDiv)
	delta := this.GetRotateDelta()
	if delta > 0 {
		this.RotateDelta(-math.Pi / 2)
	} else {
		this.RotateDelta(math.Pi / 2)
	}
}
