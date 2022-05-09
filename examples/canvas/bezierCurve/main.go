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

	// browser stage
	var bs = stage.Stage{}
	bs.Init()

	canvas = factoryBrowser.NewTagCanvas("canvas_0", bs.GetWidth(), bs.GetHeight()).
		AppendById("stage")

	var curve = &algorithm.BezierCurve{}
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

	curve.Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border}).
		Add(algorithm.Point{X: 2*wight + border - adjust, Y: 0*height + border + adjust}).
		Add(algorithm.Point{X: 2*wight + border, Y: 1*height + border}).
		Add(algorithm.Point{X: 2*wight + border - adjust, Y: 2*height + border - adjust}).
		Add(algorithm.Point{X: 1*wight + border, Y: 2*height + border}).
		Add(algorithm.Point{X: 0*wight + border + adjust, Y: 2*height + border - adjust}).
		Add(algorithm.Point{X: 0*wight + border, Y: 1*height + border}).
		Add(algorithm.Point{X: 0*wight + border + adjust, Y: 0*height + border + adjust}).
		Add(algorithm.Point{X: 1*wight + border, Y: 0*height + border}).
		Process(0.001)

	var decimatesCurve = &algorithm.Rdp{}
	decimatesCurve.Init()

	for _, point := range *curve.GetProcessed() {
		decimatesCurve.Add(point)
	}
	decimatesCurve.Process(12.0)

	var density = &algorithm.Density{}
	density.Init()

	for _, point := range *decimatesCurve.GetProcessed() {
		density.Add(point)
		AddDotYellow(int(point.X), int(point.Y))
	}
	density.IncreaseDensityBetweenPoints(60)

	for _, point := range *density.GetProcessed() {
		AddDotGreen(int(point.X), int(point.Y))
	}

	//log.Printf("len: %v", len(*decimatesCurve.GetProcessed()))
	//l := len(*decimatesCurve.GetProcessed()) - 1
	//for k, point := range *decimatesCurve.GetProcessed() {
	//	if k != l {
	//		p1 := (*decimatesCurve.GetProcessed())[k]
	//		p2 := (*decimatesCurve.GetProcessed())[k+1]
	//
	//		d := math.Sqrt(math.Pow(p1.X-p2.X, 2.0)+math.Pow(p1.Y-p2.Y, 2.0)) / 2.0
	//		a := math.Atan2(p1.Y-p2.Y, p1.X-p2.X)
	//		x := p1.X - math.Cos(a)*d
	//		y := p1.Y - math.Sin(a)*d
	//		p3 := algorithm.Point{X: x, Y: y}
	//		AddDotGreen(int(p3.X), int(p3.Y))
	//	}
	//
	//	AddDotYellow(int(point.X), int(point.Y))
	//}

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
		AddPointsToEasingTween(density).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi / 2).
		AppendToStage()

	factoryEasingTween.NewInOutBounce(
		15*time.Second,
		0,
		10000,
		div.EasingTweenWalkingAndRotateIntoPoints,
		-1,
	).
		SetOnInvertFunc(onInvert).
		SetArgumentsFunc([]interface{}{div}).
		SetDoNotReverseMotion()

	<-done
}

func AddDot(x, y int) {
	return
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.4, 0, 2*math.Pi, false).
		Fill()
}

func AddDotYellow(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewYellow()).
		Arc(x, y, 5.0, 0, 2*math.Pi, false).
		Fill()
}

func AddDotGreen(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewGreen()).
		Arc(x, y, 2.0, 0, 2*math.Pi, false).
		Fill()
}

func AddRedPointer(x, y int) {
	return
	canvas.BeginPath().
		FillStyle(factoryColor.NewRedHalfTransparent()).
		Arc(x, y, 3, 0, 2*math.Pi, false).
		Fill()
}

func AddIndex(x, y, i int) {
	return
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
