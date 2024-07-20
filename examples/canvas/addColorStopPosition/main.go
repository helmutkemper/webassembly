//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/webassembly/browser/factoryFontStyle"
	"github.com/helmutkemper/webassembly/browser/factoryFontVariant"
	"github.com/helmutkemper/webassembly/browser/html"
	"github.com/helmutkemper/webassembly/platform/factoryColor"
)

var canvas *html.TagCanvas

func main() {

	var fontA html.Font
	fontA.Family = factoryFontFamily.NewArial()
	fontA.Variant = factoryFontVariant.NewSmallCaps()
	fontA.Style = factoryFontStyle.NewItalic()
	fontA.Size = 20

	var fontB html.Font
	fontB.Family = factoryFontFamily.NewVerdana()
	fontB.Size = 35

	canvas = factoryBrowser.NewTagCanvas(800, 600).
		Font(fontA).
		FillText("Hello World!", 10, 50, 300).
		CreateLinearGradient(0, 0, 160, 0).
		AddColorStopPosition(0.0, factoryColor.NewMagenta()).
		AddColorStopPosition(0.5, factoryColor.NewBlue()).
		AddColorStopPosition(1.0, factoryColor.NewRed()).
		FillStyleGradient().
		Font(fontB).
		FillText("Big smile!", 10, 90, 300)

	var stage = factoryBrowser.NewStage()
	stage.Append(canvas)

	done := make(chan struct{}, 0)
	<-done
}
