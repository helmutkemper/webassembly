// +build js

//
package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"strconv"
	"time"
)

var (
	density                                   = 1.0
	densityManager coordinateManager.IDensity = &coordinateManager.Density{}
	stage                                     = &canvas.Stage{}
	eng                                       = engine.Engine{}
)

var htm iotmakerPlatformIDraw.IHtml
var browserDocument document.Document
var imgSpace interface{}

func prepareDataBeforeRun() {

	htm = &html.Html{}

	browserDocument = factoryBrowserDocument.NewDocument()
	stage = factoryBrowserStage.NewStage(
		htm,
		&eng,
		browserDocument,
		"stage",
		density,
		densityManager,
	)

	imgSpace = factoryBrowserImage.NewImage(
		htm,
		browserDocument.SelfDocument,
		map[string]interface{}{
			"width": 29,
			"heght": 50,
			"id":    "spacecraft",
			"src":   "./small.png",
		},
		true,
		true,
	)
}

func main() {

	done := make(chan struct{}, 0)
	prepareDataBeforeRun()

	for a := 0; a != 1; a += 1 {
		i := factoryImage.NewImage(
			"id_"+strconv.FormatInt(int64(a), 10),
			stage,
			&stage.Canvas,
			&stage.ScratchPad,
			nil,
			imgSpace,
			-100,
			-100,
			29,
			50,
			density,
			densityManager,
		)
		//i.SetDraggable(true)
		stage.AddToDraw(i)
		factoryTween.NewSelectRandom(
			&eng,
			time.Millisecond*time.Duration(mathUtil.Float64FomInt(500, 3000)),
			mathUtil.Float64FomInt(0, 800),
			mathUtil.Float64FomInt(0, 800),
			nil,
			nil,
			nil,
			func(x float64, ars ...interface{}) {
				i.Dimensions.X = int(x)
				i.OutBoxDimensions.X = int(x)
			},
			nil,
			nil,
			0,
		)
		factoryTween.NewSelectRandom(
			&eng,
			time.Millisecond*time.Duration(mathUtil.Float64FomInt(500, 3000)),
			mathUtil.Float64FomInt(0, 600),
			mathUtil.Float64FomInt(0, 600),
			nil,
			nil,
			nil,
			func(y float64, ars ...interface{}) {
				i.Dimensions.Y = int(y)
				i.OutBoxDimensions.Y = int(y)
			},
			nil,
			nil,
			0,
		)
	}

	<-done
}
