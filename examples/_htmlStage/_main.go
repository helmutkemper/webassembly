// +build js

package main

import (
  coordinateManager "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.coordinate"
	"github.com/helmutkemper/iotmaker.webassembly/factoryBrowserDocument"
	"github.com/helmutkemper/iotmaker.webassembly/factoryBrowserStage"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/engine"
  "github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factorySimpleBox"
)

func main() {

	var density = 1.0
	var densityManager coordinateManager.IDensity = &coordinateManager.Density{}

	done := make(chan struct{}, 0)

	browserDocument := factoryBrowserDocument.NewDocument()

	var htmlB = &html.Html{}
	var stage = factoryBrowserStage.NewStage(
	  htmlB,
	  &engine.Engine{},
    browserDocument,
		"stage",
		density,
		densityManager,
	)

	factorySimpleBox.NewBoxWithRoundedCorners()


	//var gradientFilter iotmakerPlatformIDraw.IFilterGradientInterface
  //
	//var shadowFilter = shadow.NewShadowFilter(
  //  factoryColorNames.NewBlackHalfTransparent(),
	//	5,
	//	2,
	//	2,
	//	density,
	//	densityManager,
	//)

	//colorWhite := factoryColor.NewColorPosition(factoryColorNames.NewRed(), 0.5)
	//colorBlack := factoryColor.NewColorPosition(factoryColorNames.NewBlack(), 1)
	//colorList := factoryColor.NewColorList(colorBlack, colorWhite)

	//coordinateP0 := factoryPoint.NewPoint(0, 0, density, densityManager)
	//coordinateP1 := factoryPoint.NewPoint(120, 150, density, densityManager)
	//gradientFilter = factoryGradient.NewGradientLinearToFillAndStroke(coordinateP0, coordinateP1, colorList)
  //
	//draw.NewBasicBox(
	//	&stage.Canvas,
	//	&stage.ScratchPad,
	//	"bbox",
	//	20,
	//	50,
	//	100,
	//	100,
	//	10,
	//	2,
	//	shadowFilter,
	//	gradientFilter,
	//	density,
	//	densityManager,
	//)

	<-done
}
