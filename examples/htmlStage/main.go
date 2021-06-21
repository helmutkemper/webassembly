package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorGradient"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryColorNames"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryDraw"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryGradient"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryPoint"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryShadow"
)

func main() {
	var density = 1.0
	var densityManager coordinateManager.IDensity = &coordinateManager.Density{}

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()

	var eng = engine.Engine{}
	var hml = &html.Html{}
	var stage = factoryBrowserStage.NewStage(
		hml,
		&eng,
		browserDocument,
		"stage",
		density,
		densityManager,
	)

	var gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface

	var shadowFilter = factoryShadow.NewShadow(
		factoryColorNames.NewBlackHalfTransparent(),
		5,
		2,
		2,
		density,
		densityManager,
	)

	colorRed := factoryColorGradient.NewColorPosition(factoryColorNames.NewRed(), 0.5)
	colorBlack := factoryColorGradient.NewColorPosition(factoryColorNames.NewBlack(), 1.0)
	colorList := factoryColorGradient.NewColorList(colorRed, colorBlack)

	coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)

	factoryDraw.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
		factoryColorNames.NewGray(),
		"bbox",
		20,
		50,
		100,
		100,
		10,
		2,
		shadowFilter,
		gradientFilter,
		density,
		densityManager,
	)

	<-done
}
