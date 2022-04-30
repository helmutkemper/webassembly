//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.webassembly/browser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.webassembly/browser/html"
	"github.com/helmutkemper/iotmaker.webassembly/browser/stage"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryColor"
	"github.com/helmutkemper/iotmaker.webassembly/platform/factoryTween"
	"strconv"
	"time"
)

var img *html.TagImage
var rocketImg *html.TagDiv

func main() {

	done := make(chan struct{}, 0)

	var stage = stage.Stage{}
	stage.Init()

	// Carrega a imagem
	img = factoryBrowser.NewTagImage(
		"spacecraft",
		"./small.png",
		29,
		50,
		true,
	)

	var font html.Font
	font.Family = factoryFontFamily.NewArial()
	font.Size = 20

	factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
		FillStyle(factoryColor.NewYellow()).
		FillRect(50, 50, 250, 100).
		SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
		FillStyle(factoryColor.NewRed()).
		FillRect(50, 50, 250, 100).
		SetTransform(1.0, 0.5, -0.5, 1.0, 30.0, 10.0).
		FillStyle(factoryColor.NewBlue()).
		FillRect(50, 50, 230, 70).
		FillStyle(factoryColor.NewBlackHalfTransparent()).
		FillRect(50, 50, 200, 50).
		AppendById("stage")

	factoryBrowser.NewTagDataList("test_A").
		NewOption("test_A_a", "label a", "value_a", true, false).
		NewOption("test_A_b", "label b", "value_b", false, false).
		NewOption("test_A_c", "label c", "value_c", false, false).
		NewOption("test_A_d", "label d", "value_d", false, true).
		AppendById("stage")

	factoryBrowser.NewTagButton("test_B").
		Value("Ok").
		Name("Ok").
		AppendById("stage")

	factoryBrowser.NewTagInputButton("test_D").
		Value("Button Value").
		Append("stage")

	factoryBrowser.NewTagSelect("test_C").
		NewOption("test_C_X", "label a", "value_a", false, false).
		NewOptionGroup("test_C_op_A", "OP A", true).
		NewOption("test_C_a", "label a", "value_a", true, false).
		NewOption("test_C_b", "label b", "value_b", false, false).
		NewOptionGroup("test_C_op_B", "OP B", false).
		NewOption("test_C_c", "label c", "value_c", false, true).
		NewOption("test_C_d", "label d", "value_d", false, false).
		AppendById("stage")

	var border = 200
	var width = stage.GetWidth() - 29 - border
	var height = stage.GetHeight() - 50 - border

	for a := 0; a != 1; a += 1 {

		var durationX = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond
		var durationY = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond

		var xStart = mathUtil.Float64FomInt(border, width)
		var xEnd = mathUtil.Float64FomInt(border, width)

		var yStart = mathUtil.Float64FomInt(border, height)
		var yEnd = mathUtil.Float64FomInt(border, height)

		var id = "div_" + strconv.FormatInt(int64(a), 10)

		rocketImg = factoryBrowser.NewTagDiv(id).
			Class("animate").
			DragStart().
			AppendToStage()

		factoryTween.NewSelectRandom(durationX, xStart, xEnd, onUpdateX, -1, rocketImg)
		factoryTween.NewSelectRandom(durationY, yStart, yEnd, onUpdateY, -1, rocketImg)
	}

	<-done
}

func onUpdateX(x, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.TagDiv).SetX(int(x))
}

func onUpdateY(y, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.TagDiv).SetY(int(y))
}
