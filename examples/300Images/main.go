//go:build js
// +build js

//
package main

import (
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontStyle"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/basic"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
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

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
)

var imgSpace html.Image

func main() {

	done := make(chan struct{}, 0)

	var browserDocument = factoryBrowserDocument.NewDocument()
	var stage = &canvas.Stage{}
	var eng = &engine.Engine{}
	var hml = &html.Html{}

	stage = factoryBrowserStage.NewStage(
		hml,
		eng,
		browserDocument,
		"stage",
		density,
		densityManager,
	)

	fontText := factoryFont.NewFont(45, factoryFontFamily.NewArial(), factoryFontStyle.NewNotSet(), density, densityManager)
	inkText := factoryInk.NewInk(ink.Ink{}, 1, factoryColorNames.NewRed(), nil, nil, density, densityManager)

	text := factoryText.NewText(
		"text",
		stage,
		&stage.Canvas,
		&stage.ScratchPad,
		&inkText,
		fontText,
		"Olá Mundo!",
		125,
		20,
		density,
		densityManager,
	)
	text.SetDragMode(basic.KDragModeDesktop)
	stage.AddToDraw(text)

	imgSpace = factoryBrowserImage.NewImage(
		hml,
		browserDocument.SelfDocument,
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
			stage,
			&stage.Canvas,
			&stage.ScratchPad,
			nil,
			imgSpace.Get(),
			-100,
			-100,
			29,
			50,
			density,
			densityManager,
		)
		i.SetDraggableToDesktop()
		stage.AddToDraw(i)

		factoryTween.NewSelectRandom(
			eng,
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, 1000),
			mathUtil.Float64FomInt(0, 1000),
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

		factoryTween.NewSelectRandom(
			eng,
			time.Duration(mathUtil.Int(500, 3000))*time.Millisecond,
			mathUtil.Float64FomInt(0, 800),
			mathUtil.Float64FomInt(0, 900),
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
