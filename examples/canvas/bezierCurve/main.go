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

	var b algorithm.BezierCurve
	b.Init()

	//var y int
	//var delta = 0
	//var change bool
	//for i := 0; i != 21; i += 1 {
	//	if i%2 == 0 {
	//		y = 300
	//	} else if change == false {
	//		y = 400
	//		change = !change
	//	} else {
	//		y = 200
	//		change = !change
	//	}
	//
	//	y += delta
	//	delta += 0
	//
	//	b.Add(algorithm.Point{X: float64(i * 50), Y: float64(y)})
	//}
	//
	//// 1000, 300
	//b.Add(algorithm.Point{X: 1050, Y: 500})
	//b.Add(algorithm.Point{X: 700, Y: 600})
	//b.Add(algorithm.Point{X: 400, Y: 700})
	//b.Add(algorithm.Point{X: 300, Y: 400})
	//b.Add(algorithm.Point{X: 300, Y: 300})
	//b.Add(algorithm.Point{X: 400, Y: 200})
	//b.Add(algorithm.Point{X: 600, Y: 100})
	//b.Add(algorithm.Point{X: 700, Y: 200})
	//b.Add(algorithm.Point{X: 800, Y: 300})
	//b.Add(algorithm.Point{X: 700, Y: 400})
	//b.Add(algorithm.Point{X: 600, Y: 500})
	//b.Add(algorithm.Point{X: 500, Y: 400})
	//b.Add(algorithm.Point{X: 400, Y: 300})
	//b.Add(algorithm.Point{X: 500, Y: 200})
	//b.Add(algorithm.Point{X: 600, Y: 100})
	//b.Add(algorithm.Point{X: 700, Y: 200})

	b.Add(algorithm.Point{X: 50, Y: 400}) //0
	b.Add(algorithm.Point{X: 50, Y: 50})
	b.Add(algorithm.Point{X: 600, Y: 50})
	b.Add(algorithm.Point{X: 1200, Y: 50})
	b.Add(algorithm.Point{X: 1200, Y: 400})
	b.Add(algorithm.Point{X: 1200, Y: 800}) //5
	b.Add(algorithm.Point{X: 600, Y: 800})
	b.Add(algorithm.Point{X: 150, Y: 800})
	b.Add(algorithm.Point{X: 150, Y: 400})
	b.Add(algorithm.Point{X: 150, Y: 150})
	b.Add(algorithm.Point{X: 600, Y: 150}) //10
	b.Add(algorithm.Point{X: 1100, Y: 150})
	b.Add(algorithm.Point{X: 1100, Y: 400})
	b.Add(algorithm.Point{X: 1100, Y: 700})
	b.Add(algorithm.Point{X: 600, Y: 700})
	b.Add(algorithm.Point{X: 250, Y: 700}) //15
	b.Add(algorithm.Point{X: 250, Y: 400})
	b.Add(algorithm.Point{X: 250, Y: 250})
	b.Add(algorithm.Point{X: 600, Y: 250})
	b.Add(algorithm.Point{X: 1000, Y: 250})
	b.Add(algorithm.Point{X: 970, Y: 400}) //20
	b.Add(algorithm.Point{X: 930, Y: 600})
	b.Add(algorithm.Point{X: 800, Y: 350})
	b.Add(algorithm.Point{X: 700, Y: 200})
	b.Add(algorithm.Point{X: 600, Y: 400})
	b.Add(algorithm.Point{X: 500, Y: 600})
	b.Add(algorithm.Point{X: 400, Y: 400})

	b.Process(0.005)

	for v, point := range *b.GetOriginal() {
		AddRedPointer(int(point.X), int(point.Y))
		AddIndex(int(point.X), int(point.Y), v)
	}

	for _, point := range *b.GetProcessed() {
		AddDot(int(point.X), int(point.Y))
	}

	var div *html.TagDiv
	div = factoryBrowser.NewTagDiv("div_0")
	div.Class("animate").
		AddPoints(b.GetProcessed()).
		SetDeltaX(-15).
		SetDeltaY(-25).
		RotateDelta(-math.Pi/2).
		//NewEasingTweenRandom("line", 5*time.Second, 0, 10000, div.WalkingIntoPoints, -1).
		NewEasingTweenLinear("line", 30*time.Second, 0, 10000, div.WalkingAndRotateIntoPoints, -1).
		EasingTweenOnInvertFunc("line", onInvert).
		//EasingTweenDoNotReverseMotion("line").
		AppendToStage()

	<-done
}

func AddDot(x, y int) {
	canvas.BeginPath().
		FillStyle(factoryColor.NewBlueHalfTransparent()).
		Arc(x, y, 0.3, 0, 2*math.Pi, false).
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
