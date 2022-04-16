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

	var border = 200

	for a := 0; a != 20; a += 1 {

		var intervalX = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond
		var intervalY = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond
		var width = global.Global.Document.GetDocumentWidth() - 29 - border
		var height = global.Global.Document.GetDocumentHeight() - 50 - border
		var positionInicialX = mathUtil.Float64FomInt(border, width)
		var positionFinalX = mathUtil.Float64FomInt(border, width)
		var positionInicialY = mathUtil.Float64FomInt(border, height)
		var positionFinalY = mathUtil.Float64FomInt(border, height)

		var id = "div_" + strconv.FormatInt(int64(a), 10)

		rocket := factoryBrowser.NewTagDiv(id).
			SetClass("animate").
			AppendById("stage")

		factoryTween.NewSelectRandom(
			intervalX,
			positionInicialX,
			positionFinalX,
			updateX,
			-1,
			rocket,
		)

		factoryTween.NewSelectRandom(
			intervalY,
			positionInicialY,
			positionFinalY,
			updateY,
			-1,
			rocket,
		)
	}

	<-done
}

func updateX(x, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.GlobalAttributes).SetX(int(x))
}

func updateY(y, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.GlobalAttributes).SetY(int(y))
}
