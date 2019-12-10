// +build js

package main

import (
	iotmakerPlatformIDraw "github.com/helmutkemper/iotmaker.platform.IDraw"
	coordinateManager "github.com/helmutkemper/iotmaker.platform.coordinate"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.platform.webbrowser/factoryBrowserStage"
	"github.com/helmutkemper/iotmaker.platform/abstractType/basicBox"
	"github.com/helmutkemper/iotmaker.platform/abstractType/colornames"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryColor"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryGradient"
	"github.com/helmutkemper/iotmaker.platform/abstractType/factoryPoint"
	"github.com/helmutkemper/iotmaker.platform/abstractType/shadow"
)

func main() {

	var density = 1.0
	var densityManager coordinateManager.IDensity = &coordinateManager.Density{}

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()

	var stage = factoryBrowserStage.NewStage(
		browserDocument,
		"stage",
		300,
		300,
		density,
		densityManager,
	)

	var gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface

	var shadowFilter = shadow.NewShadowFilter(
		colornames.BlackHalfTransparent,
		5,
		2,
		2,
		density,
		densityManager,
	)

	colorWhite := factoryColor.NewColorPosition(colornames.Red, 0.5)
	colorBlack := factoryColor.NewColorPosition(colornames.Black, 1)
	colorList := factoryColor.NewColorList(colorBlack, colorWhite)

	coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)

	basicBox.NewBasicBox(
		&stage.Canvas,
		&stage.ScratchPad,
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
