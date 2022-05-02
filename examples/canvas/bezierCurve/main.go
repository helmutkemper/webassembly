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
	"math"
	"strconv"
	"time"
)

var canvas *html.TagCanvas

func main() {

	done := make(chan struct{}, 0)

	var stage = stage.Stage{}
	stage.Init()

	canvas = factoryBrowser.NewTagCanvas("canvas_0", stage.GetWidth(), stage.GetHeight()).
		AppendById("stage")

	var curve algorithm.BezierCurve
	curve.Init()

	border := 50.0
	wight := 400.0
	height := 400.0
	adjust := -15.0

	//    0,0    1,0    2,0
	//     7------0------1
	//     |             |
	// 0,1 6             2 2,1
	//     |             |
	//     5------4------3
	//    0,2    1,2    2,2

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
	div = factoryBrowser.NewTagDiv("div_0")
	div.Class("animate").
		AddPoints(curve.GetProcessed()).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi/2).
		//NewEasingTweenLinear("line", 5*time.Second, 0, 10000, div.WalkingIntoPoints, -1).
		NewEasingTweenLinear("line", 5*time.Second, 0, 10000, div.WalkingAndRotateIntoPoints, -1).
		EasingTweenOnInvertFunc("line", onInvert).
		EasingTweenDoNotReverseMotion("line").
		AppendToStage()

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

	x += 5
	y += 20
	var font html.Font
	font.Family = factoryFontFamily.NewArial()
	font.Size = 15

	canvas.BeginPath().
		Font(font).
		FillStyle(factoryColor.NewRedHalfTransparent()).
		FillText(
			strconv.FormatInt(int64(i), 10),
			x,
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
