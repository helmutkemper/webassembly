//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontStyle"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorNames"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryFont"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryInk"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryText"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/ink"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"strconv"
	"time"
)

var imgSpace html.Image

func main() {

	done := make(chan struct{}, 0)
	stage := global.Global.Stage

	fontText := factoryFont.NewFont(45, factoryFontFamily.NewArial(), factoryFontStyle.NewNotSet(), global.Global.Density, global.Global.DensityManager)
	inkText := factoryInk.NewInk(ink.Ink{}, 1, factoryColorNames.NewRed(), nil, nil, global.Global.Density, global.Global.DensityManager)

	text := factoryText.NewText(
		"text",
		global.Global.Stage,
		global.Global.Canvas,
		global.Global.ScratchPad,
		&inkText,
		fontText,
		"Olá Mundo!",
		125,
		20,
		global.Global.Density,
		global.Global.DensityManager,
	)
	text.SetDragMode(basic.KDragModeDesktop)
	stage.AddToDraw(text)

	imgSpace = factoryBrowserImage.NewImage(
		global.Global.Html,
		global.Global.Document.SelfDocument,
		map[string]interface{}{
			"width": 29,
			"heght": 50,
			"id":    "spacecraft",
			"src":   "./small.png",
		},
		true,
		false,
	)

	for a := 0; a != 10; a += 1 {
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
			mathUtil.Float64FomInt(0, 800),
			mathUtil.Float64FomInt(0, 900),
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
