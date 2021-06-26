package global

import (
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/canvas"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/document"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
)

type Config struct {
	Canvas         *canvas.Canvas
	ScratchPad     *canvas.Canvas
	Stage          *canvas.Stage
	Engine         *engine.Engine
	Html           *html.Html
	Document       document.Document
	Density        float64
	DensityManager *coordinateManager.Density
}

var Global Config

func init() {
	Global.Document = factoryBrowserDocument.NewDocument()
	Global.Html = &html.Html{}
	Global.Engine = &engine.Engine{}
	Global.Density = 1.0
	Global.DensityManager = &coordinateManager.Density{}
	Global.Stage = factoryBrowserStage.NewStage(
		Global.Html,
		Global.Engine,
		Global.Document,
		"stage-auto-build",
		Global.Density,
		Global.DensityManager,
	)

	Global.Canvas = &Global.Stage.Canvas
	Global.ScratchPad = &Global.Stage.ScratchPad
}
