// English:
//
// This example was taken from https://developer.mozilla.org/en-US/docs/Web/SVG/Element/switch
//
// Português:
//
// Este exemplo foi retirado do site https://developer.mozilla.org/en-US/docs/Web/SVG/Element/switch
//
//  <svg viewBox="0 -20 100 50">
//    <switch>
//      <text systemLanguage="ar">مرحبا</text>
//      <text systemLanguage="de,nl">Hallo!</text>
//      <text systemLanguage="en-us">Howdy!</text>
//      <text systemLanguage="en-gb">Wotcha!</text>
//      <text systemLanguage="en-au">G'day!</text>
//      <text systemLanguage="en">Hello!</text>
//      <text systemLanguage="es">Hola!</text>
//      <text systemLanguage="fr">Bonjour!</text>
//      <text systemLanguage="ja">こんにちは</text>
//      <text systemLanguage="ru">Привет!</text>
//      <text>☺</text>
//    </switch>
//  </svg>

//go:build js

package main

import (
	"github.com/helmutkemper/webassembly/browser/factoryBrowser"
)

func main() {

	done := make(chan struct{}, 0)

	stage := factoryBrowser.NewStage()

	s1 := factoryBrowser.NewTagSvg().ViewBox([]float64{0, -20, 100, 50}).Append(
		factoryBrowser.NewTagSvgSwitch().Append(
			factoryBrowser.NewTagSvgText().SystemLanguage("ar").Text("مرحبا"),
			factoryBrowser.NewTagSvgText().SystemLanguage("de,nl").Text("Hallo!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("en-us").Text("Howdy!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("en-gb").Text("Wotcha!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("en-au").Text("G'day!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("en").Text("Hello!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("es").Text("Hola!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("fr").Text("Bonjour!"),
			factoryBrowser.NewTagSvgText().SystemLanguage("ja").Text("こんにちは"),
			factoryBrowser.NewTagSvgText().SystemLanguage("ru").Text("Привет!"),
			factoryBrowser.NewTagSvgText().Text("☺"),
		),
	)

	stage.Append(s1)

	<-done
}
