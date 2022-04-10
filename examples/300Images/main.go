//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"strconv"
	"time"
)

var imgSpace html.Image

func main() {

	done := make(chan struct{}, 0)
	stage := global.Global.Stage

	imgSpace = factoryBrowserImage.NewImage(
		29,
		50,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		true,
	)

	for a := 0; a != 100; a += 1 {
		i := factoryImage.NewImage(
			"id_"+strconv.FormatInt(int64(a), 10),
			global.Global.Stage,
			global.Global.Canvas,
			global.Global.ScratchPad,
			nil,
			imgSpace.Get(),
			-100,
			-100,
			29,
			50,
			global.Global.Density,
			global.Global.DensityManager,
		)
		i.SetDraggableToDesktop()
		stage.AddToDraw(i)

		factoryTween.NewLinear(
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, 1000),
			mathUtil.Float64FomInt(0, 1000),
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
			func(x, p float64, ars ...interface{}) {
				//log.Printf("x: %v", x)
				//i.Dimensions.X = int(x)
				//i.OutBoxDimensions.X = int(x)
				//i.MoveY(int(x), int(100))
				i.MoveX(int(x))
				//i.Draw()
			},
			-1,
		)

		factoryTween.NewLinear(
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentHeight()-50),
			mathUtil.Float64FomInt(0, global.Global.Document.GetDocumentHeight()-50),
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
			func(y, p float64, ars ...interface{}) {
				//log.Printf("x: %v", x)
				//i.Dimensions.X = int(x)
				//i.OutBoxDimensions.X = int(x)
				//i.MoveY(int(x), int(100))
				i.MoveY(int(y))
				//i.Draw()
			},
			-1,
		)

	}

	<-done
}
