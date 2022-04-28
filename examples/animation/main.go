//go:build js
// +build js

//
package main

import (
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/browserStage"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/browser/mouse"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryBrowser"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/factoryFontFamily"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/html"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform.webbrowser/platform/factoryColor"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/factoryTween"
	"github.com/helmutkemper/iotmaker.santa_isabel_theater.platform/mathUtil"
	"log"
	"strconv"
	"time"
)

var img *html.TagImage
var rocketImg *html.TagDiv

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
			AppendById("stage")

		factoryTween.NewSelectRandom(durationX, xStart, xEnd, onUpdateX, 0, rocketImg)
		factoryTween.NewSelectRandom(durationY, yStart, yEnd, onUpdateY, 0, rocketImg)
	}

	stage.AddListener(mouse.KEventMouseOver, onMouseEvent)
	timer := time.NewTimer(10 * time.Second)
	go func() {
		select {
		case <-timer.C:
			stage.RemoveListener(mouse.KEventMouseOver)
		}
	}()

	//stage.AddListener(mouse.KEventMouseUp, DragStop)
	//stage.AddListener(mouse.KEventMouseMove, Drag)
	//document.AddEventListener(browserMouse.KEventMouseEnter, browserMouse.SetMouseSimpleEventManager(stage.CursorShow))
	<-done
}

func onMouseEvent(event mouse.MouseEvent) {
	isNull, target := event.GetRelatedTarget()
	if isNull == false {
		log.Print("id: ", target.Get("id"))
		log.Print("tagName: ", target.Get("tagName"))
	}
	log.Print(event.GetScreenX())
	log.Print(event.GetScreenY())
}

var drag bool
var difX, difY int

func DragStart(event mouse.MouseEvent) {
	log.Printf("start")
	rocketImg.Mouse(mouse.KCursorCopy)
	drag = true
}
func DragStop(event mouse.MouseEvent) {
	log.Printf("false")
	drag = false
	rocketImg.Mouse(mouse.KCursorAuto)

	difX = 0
	difY = 0
}
func Drag(event mouse.MouseEvent) {
	if drag == false {
		return
	}

	if difX == 0 || difY == 0 {
		var x = rocketImg.GetX()
		var y = rocketImg.GetY()

		var screenX = int(event.GetScreenX())
		var screenY = int(event.GetScreenY())

		difX = screenX - x
		difY = screenY - y
	}

	//isNull, _ := event.GetRelatedTarget()
	//if isNull == true {
	//	return
	//}

	log.Printf("move")

	var x = int(event.GetScreenX()) - difX
	var y = int(event.GetScreenY()) - difY

	rocketImg.SetXY(x, y)
}

func move(event mouse.MouseEvent) {
	return
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
