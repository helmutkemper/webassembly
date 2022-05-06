//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/algorithm"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryEasingTween"
	"math"
	"strconv"
	"time"
)

var canvas *html.TagCanvas

func main() {

	done := make(chan struct{}, 0)

	var bs = stage.Stage{}
	bs.Init()

	canvas = factoryBrowser.NewTagCanvas("canvas_0", bs.GetWidth(), bs.GetHeight()).
		AppendById("stage")

	var curve algorithm.BezierCurve
	curve.Init()

	border := 50.0
	wight := 400.0
	height := 400.0
	adjust := 15.0

	// E.g.: P0 (1,0) = (1*wight,0*height)
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

	curve.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})
	curve.Add(algorithm.Point{X: 2*wight + border - adjust, Y: 0*height + border + adjust})
	curve.Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border})
	curve.Add(algorithm.Point{X: 2*wight + border - adjust, Y: 2*height + border - adjust})
	curve.Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border})
	curve.Add(algorithm.Point{X: 0*wight + border + adjust, Y: 2*height + border - adjust})
	curve.Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border})
	curve.Add(algorithm.Point{X: 0*wight + border + adjust, Y: 0*height + border + adjust})
	curve.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border})

	curve.Process(0.001)
	curve.AdjustDensity()

	for v, point := range *curve.GetOriginal() {
		AddRedPointer(int(point.X), int(point.Y))
		AddIndex(int(point.X), int(point.Y), v)
	}

	for _, point := range *curve.GetProcessed() {
		AddDot(int(point.X), int(point.Y))
	}

	var div *html.TagDiv
	div = factoryBrowser.NewTagDiv("div_0").
		Class("animate").
		DragStart().
		AddPointsToEasingTween(curve.GetProcessed()).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi / 2).
		AppendToStage()

	factoryEasingTween.NewLinear(
		5*time.Second,
		0,
		10000,
		div.EasingTweenWalkingAndRotateIntoPoints(),
		0,
	).
		SetOnInvertFunc(onInvert).
		SetArgumentsFunc([]interface{}{div}).
		SetDoNotReverseMotion()

	<-done
}

func AddDot(x, y int) {

	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.4, 0, 2*math.Pi, false).
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
