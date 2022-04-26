//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/browserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browserMouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"log"
	"strconv"
	"time"
)

var img *html.TagImage

func main() {

	done := make(chan struct{}, 0)

	var stage = browserStage.Stage{}
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

	//     ctx.font = "30px Arial";
	//     var txt = "Hello World"
	//     ctx.fillText("width:" + ctx.measureText(txt).width, 10, 50)
	//     ctx.fillText(txt, 10, 100);

	canvas := factoryBrowser.NewTagCanvas("canvas_0", 800, 600).
		Font(font)

	w := canvas.MeasureText("Hello Word!")
	wText := strconv.FormatInt(int64(w), 10)

	canvas.FillText("width:"+wText, 10, 50, 300).
		AppendById("stage")

	//FillText("width: ").
	//BeginPath().
	//LineWidth(5).
	//LineJoin(html.KJoinRuleRound).
	//MoveTo(20, 20).
	//LineTo(100, 50).
	//LineTo(20, 100).
	//Stroke().
	//AppendById("stage")

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

	for a := 0; a != 20; a += 1 {

		var durationX = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond
		var durationY = time.Duration(mathUtil.Int(1000, 3000)) * time.Millisecond

		var xStart = mathUtil.Float64FomInt(border, width)
		var xEnd = mathUtil.Float64FomInt(border, width)

		var yStart = mathUtil.Float64FomInt(border, height)
		var yEnd = mathUtil.Float64FomInt(border, height)

		var id = "div_" + strconv.FormatInt(int64(a), 10)

		rocketImg := factoryBrowser.NewTagDiv(id).
			Class("animate").
			AppendById("stage")

		factoryTween.NewSelectRandom(durationX, xStart, xEnd, onUpdateX, -1, rocketImg)
		factoryTween.NewSelectRandom(durationY, yStart, yEnd, onUpdateY, -1, rocketImg)
	}

	stage.AddListener(browserMouse.KEventMouseOver, move)
	//document.AddEventListener(browserMouse.KEventMouseEnter, browserMouse.SetMouseSimpleEventManager(stage.CursorShow))
	<-done
}

func move(event browserMouse.MouseEvent) {
	isNull, target := event.GetRelatedTarget()
	if isNull == false {
		log.Print("id: ", target.Get("id"))
		log.Print("tagName: ", target.Get("tagName"))
	}
	log.Print(event.GetScreenX())
	log.Print(event.GetScreenY())
}

func onUpdateX(x, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.TagDiv).SetX(int(x))
}

func onUpdateY(y, p float64, args ...interface{}) {
	args[0].([]interface{})[0].(*html.TagDiv).SetY(int(y))
}
