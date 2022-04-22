//go:build js
// +build js

//
package main

import (
	global "github.com/helmutkemper/iotmaker.santa_isabel_theater.globalConfig"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowserImage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"strconv"
	"time"
)

func main() {

	done := make(chan struct{}, 0)

	// Carrega a imagem
	factoryBrowserImage.NewImage(
		29,
		50,
		map[string]interface{}{
			"id":  "spacecraft",
			"src": "./small.png",
		},
		true,
		false,
	)

	factoryBrowser.NewTagDataList("test_A").
		NewOption("test_A_a", "label a", "value_a", true, false).
		NewOption("test_A_b", "label b", "value_b", false, false).
		NewOption("test_A_c", "label c", "value_c", false, false).
		NewOption("test_A_d", "label d", "value_d", false, true).
		AppendById("stage")

	var border = 200
	var width = global.Global.Document.GetDocumentWidth() - 29 - border
	var height = global.Global.Document.GetDocumentHeight() - 50 - border

	for a := 0; a != 20; a += 1 {

		var durationX = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond
		var durationY = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond

		var xStart = mathUtil.Float64FomInt(border, width)
		var xEnd = mathUtil.Float64FomInt(border, width)

		var yStart = mathUtil.Float64FomInt(border, height)
		var yEnd = mathUtil.Float64FomInt(border, height)

		var loopX = mathUtil.Int(-1, 10)
		var loopY = mathUtil.Int(-1, 10)

		var id = "div_" + strconv.FormatInt(int64(a), 10)

		rocketImg := factoryBrowser.NewTagDiv(id).
			SetClass("animate").
			AppendById("stage")

		factoryTween.NewSelectRandom(durationX, xStart, xEnd, onUpdateX, loopX, rocketImg)
		factoryTween.NewSelectRandom(durationY, yStart, yEnd, onUpdateY, loopY, rocketImg)
	}

	<-done
}

func onUpdateX(x, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.GlobalAttributes).SetX(int(x))
}

func onUpdateY(y, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.GlobalAttributes).SetY(int(y))
}
